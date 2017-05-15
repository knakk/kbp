package disk

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/RoaringBitmap/roaring"
	"github.com/boltdb/bolt"
	"github.com/knakk/kbp/rdf"
	"github.com/knakk/kbp/rdf/internal/query"
	"github.com/knakk/kbp/rdf/memory"
)

// Exported errors
var (
	// ErrNotFound is an error signifying that the resource is not present in the database.
	ErrNotFound = errors.New("not found")

	// ErrDBFull is returned when the database cannot store more nodes.
	ErrDBFull = errors.New("database is full")
)

const (
	// MaxNodes is the maximum number of unique RDF nodes that can be stored.
	MaxNodes = 4294967295
)

// Graph is an implementation of rdf.Graph which persists to disk.
type Graph struct {
	// Graph is persisted in Bolt - a transactional key-value store backed
	// by a memory-mapped file.
	kv *bolt.DB

	// The majority of the URIs in a RDF graph are typically made up
	// using the same base URI, so we optimize storage for those by
	// making it the default case, and only store the absolute part
	// of an URI when it's not the default one.
	//
	// The base should include the scheme and hostname, ex: http://example.org/
	//
	// It must not be changed as long as the database is open, but may
	// be set in the call to Open() when opening a database.
	base rdf.NamedNode
}

// Open opens a persisted Graph at the given path. If the file does
// not exist, it will be created.
// Only one process will have access to the Graph at a time.
func Open(path string, base string) (*Graph, error) {
	kv, err := bolt.Open(path, 0644, nil)
	if err != nil {
		return nil, err
	}
	g := &Graph{
		kv:   kv,
		base: rdf.NewNamedNode(base),
	}
	return g.setup()
}

// Close closes the Graph, relasing the lock on the underlying file.
func (g *Graph) Close() error { return g.kv.Close() }

// Buckets in the key-value store:
var (
	//bucketMeta = []byte("meta")

	// Dictionaries
	bucketId2Node = []byte("id2node") // uint32 -> node
	bucketNode2Id = []byte("node2id") // node -> uint32

	// Indexes
	bucketSPO = []byte("spo") // subect + predicate -> object bitmap
	bucketOSP = []byte("osp") // object + subject   -> predicate bitmap
	bucketPOS = []byte("pos") // predicate + object -> subject bitmap
)

// setup makes sure the Graph is ready to be used.
func (g *Graph) setup() (*Graph, error) {

	// Create the required buckets in the key-value store if the database is new.
	err := g.kv.Update(func(tx *bolt.Tx) error {
		for _, b := range [][]byte{bucketId2Node, bucketNode2Id, bucketSPO, bucketOSP, bucketPOS} {
			_, err := tx.CreateBucketIfNotExists(b)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return g, err
}

func (g *Graph) spo(tx *bolt.Tx, s, p uint32) (*roaring.Bitmap, error) {
	return g.getBitmap(tx, bucketSPO, s, p)
}

func (g *Graph) pos(tx *bolt.Tx, p, o uint32) (*roaring.Bitmap, error) {
	return g.getBitmap(tx, bucketPOS, p, o)
}

func (g *Graph) osp(tx *bolt.Tx, o, s uint32) (*roaring.Bitmap, error) {
	return g.getBitmap(tx, bucketOSP, o, s)
}

func (g *Graph) getBitmap(tx *bolt.Tx, bkt []byte, a, b uint32) (*roaring.Bitmap, error) {
	idx := make([]byte, 8)
	copy(idx, u32tob(a))
	copy(idx[4:], u32tob(b))

	bb := tx.Bucket(bkt).Get(idx)
	if bb == nil {
		return nil, ErrNotFound
	}

	bitmap := roaring.New()
	_, err := bitmap.ReadFrom(bytes.NewReader(bb))
	if err != nil {
		return nil, err
	}

	return bitmap, nil
}

func (g *Graph) estCardinality(tx *bolt.Tx, p *query.EncTriplePattern) error {

	// triple is not in graph
	if p[0] == 0 || p[1] == 0 || p[2] == 0 {
		p[3] = 0 // exact
		return nil
	}

	// {n,n,n}
	// TODO what if n is blank node?
	if p[0] > query.MaxVariables && p[1] > query.MaxVariables && p[2] > query.MaxVariables {
		p[3] = 1 // exact
		return nil
	}

	// {?,?,?}
	if p[0] < query.MaxVariables && p[1] < query.MaxVariables && p[2] < query.MaxVariables {
		p[3] = uint32(tx.Bucket(bucketSPO).Stats().KeyN * tx.Bucket(bucketOSP).Stats().KeyN) // estimate
		return nil
	}

	// {?,n,n}
	if p[0] < query.MaxVariables && p[1] > query.MaxVariables && p[2] > query.MaxVariables {
		pos, err := g.pos(tx, p[1], p[2])
		if err == ErrNotFound {
			p[3] = 0 // exact
			return nil
		}
		if err != nil {
			return err
		}
		p[3] = uint32(pos.GetCardinality()) // exact
		return nil
	}

	// {n,n,?}
	if p[0] > query.MaxVariables && p[1] > query.MaxVariables && p[2] < query.MaxVariables {
		spo, err := g.spo(tx, p[0], p[1])
		if err == ErrNotFound {
			p[3] = 0 // exact
			return nil
		}
		if err != nil {
			return err
		}
		p[3] = uint32(spo.GetCardinality()) // exact
		return nil
	}

	// {n,?,n}
	if p[0] > query.MaxVariables && p[1] < query.MaxVariables && p[2] > query.MaxVariables {
		osp, err := g.osp(tx, p[2], p[0])
		if err == ErrNotFound {
			p[3] = 0 // exact
			return nil
		}
		if err != nil {
			return err
		}
		p[3] = uint32(osp.GetCardinality()) // exact
		return nil
	}

	// {?,?,n}
	if p[0] < query.MaxVariables && p[1] < query.MaxVariables && p[2] > query.MaxVariables {
		// TODO find a way to approach this
		p[3] = 10000 // estimate
		return nil
	}

	// {n,?,?}
	if p[0] > query.MaxVariables && p[1] < query.MaxVariables && p[2] < query.MaxVariables {
		// TODO find a way to approach this
		p[3] = 10000 // estimate
		return nil
	}

	// {?,n,?}
	if p[0] < query.MaxVariables && p[1] > query.MaxVariables && p[2] < query.MaxVariables {
		// TODO find a way to approach this
		p[3] = 10000 // estimate
		return nil
	}

	panic(fmt.Sprintf("BUG: Graph.estCardinality: unhandled pattern: %v", p[0:3]))
}

func (g *Graph) encodePattern(tx *bolt.Tx, p rdf.TriplePattern, cache map[rdf.TriplePatternNode]uint32, numVars *int) (query.EncTriplePattern, error) {
	var ep query.EncTriplePattern
	for i, node := range []rdf.TriplePatternNode{p.Subject, p.Predicate, p.Object} {
		if id, ok := cache[node]; ok {
			ep[i] = id
			continue
		}
		switch t := node.(type) {
		case rdf.Variable, rdf.BlankNode:
			*numVars++
			cache[node] = uint32(*numVars)
		default:
			id, err := g.getID(tx, t.(rdf.Node))
			if err == ErrNotFound {
				ep[i] = 0
			} else if err != nil {
				return ep, err
			}
			cache[node] = id
		}
		ep[i] = cache[node]

	}

	err := g.estCardinality(tx, &ep)

	return ep, err
}

// Scanning functions TODO
// - Move to separate file scan.go ?
// - A lot of code duplication right now, refactor when evalBGP is assumed correct (needs more tests)
// - If only one pattern, or last pattern in group, updating bound is unessecary

func (g *Graph) scanOSPwBoundSP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanOSPwBoundSP(%v) bound %v\n", p[:3], bound)
	// O must be a concrete, exisiting node.

	// SP are bound, so they must be variables
	sol := query.EncSolutions{
		Vars: []uint32{p[0], p[1]},
	}

	boundS := bound[p[0]]
	boundP := bound[p[1]]
	iterS := boundS.Iterator()
	missS := roaring.New()

	for iterS.HasNext() {
		sID := iterS.Next()
		preds, err := g.osp(tx, p[2], sID)
		if err == ErrNotFound {
			missS.Add(sID)
			continue
		} else if err != nil {
			return sol, err
		}
		boundP.And(preds)
		iterP := boundP.Iterator()
		for iterP.HasNext() {
			sol.Rows = append(sol.Rows, []uint32{sID, iterP.Next()})
		}
	}

	boundS.AndNot(missS)

	return sol, nil
}

func (g *Graph) scanOSPwBoundS(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanOSPwBoundS(%v) bound %v\n", p[:3], bound)
	// O must be a concrete, exisiting node.

	// S is bound, so must be a variable
	sol := query.EncSolutions{
		Vars: []uint32{p[0]},
	}
	if p[1] < query.MaxVariables {
		sol.Vars = append(sol.Vars, p[1])
		bound[p[1]] = roaring.New()
	}

	boundS := bound[p[0]]
	iterS := boundS.Iterator()
	missS := roaring.New()

	for iterS.HasNext() {
		sID := iterS.Next()
		preds, err := g.osp(tx, p[2], sID)
		if err == ErrNotFound {
			missS.Add(sID)
			continue
		} else if err != nil {
			return sol, err
		}

		if p[1] < query.MaxVariables {
			bound[p[1]].And(preds)
			iterP := preds.Iterator()
			for iterP.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{sID, iterP.Next()})
			}
		} else {
			sol.Rows = append(sol.Rows, []uint32{sID})
		}
	}

	boundS.AndNot(missS)

	return sol, nil
}

func (g *Graph) scanPOSwBoundP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// O must be a concrete, exisiting node.

	sol := query.EncSolutions{}

	if p[0] < query.MaxVariables {
		sol.Vars = []uint32{p[0], p[1]}
		bound[p[0]] = roaring.New()
	} else {
		// P is bound, so must be a variable
		sol.Vars = []uint32{p[1]}
	}

	boundP := bound[p[1]]
	iterP := boundP.Iterator()
	missP := roaring.New()

	for iterP.HasNext() {
		pID := iterP.Next()
		subjs, err := g.pos(tx, pID, p[2])
		if err == ErrNotFound {
			missP.Add(pID)
			continue
		} else if err != nil {
			return sol, err
		}

		if p[0] < query.MaxVariables {
			bound[p[0]].And(subjs)
			iterS := subjs.Iterator()
			for iterS.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{iterS.Next(), pID})
			}
		} else {
			sol.Rows = append(sol.Rows, []uint32{pID})
		}
	}

	boundP.AndNot(missP)

	return sol, nil
}

func (g *Graph) scanOSP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanOSP(%v) bound %v\n", p[:3], bound)
	// O must be a concrete, exisiting node.
	// At leas one, or both of SP must be variable.

	sol := query.EncSolutions{}

	if p[0] < query.MaxVariables {
		if p[1] < query.MaxVariables {
			// Both S & P are variable
			sol.Vars = []uint32{p[0], p[1]}
			bound[p[0]] = roaring.New()
			bound[p[1]] = roaring.New()

			bo := u32tob(p[2])
			cur := tx.Bucket(bucketOSP).Cursor()
		outerOSP:
			for k, _ := cur.Seek(bo); k != nil; k, _ = cur.Next() {
				switch bytes.Compare(k[:4], bo) {
				case 0:
					sID := btou32(k[4:])
					preds, err := g.osp(tx, p[2], sID)
					if err != nil {
						// err cannot be ErrNotFound
						return sol, err
					}
					bound[p[0]].Add(sID)
					iterP := preds.Iterator()
					for iterP.HasNext() {
						sol.Rows = append(sol.Rows, []uint32{sID, iterP.Next()})
					}
					bound[p[1]].Or(preds)
				case 1:
					break outerOSP
				}
			}
			return sol, nil
		}
		// S is variable, P is concrete
		sol.Vars = []uint32{p[0]}
		bound[p[0]] = roaring.New()
		subjs, err := g.pos(tx, p[1], p[2])
		if err == ErrNotFound {
			return sol, nil
		}
		if err != nil {
			return sol, err
		}
		bound[p[0]].Or(subjs)
		iterS := subjs.Iterator()
		for iterS.HasNext() {
			sol.Rows = append(sol.Rows, []uint32{iterS.Next()})
		}
		return sol, nil
	}

	// P is variable, S is concrete
	sol.Vars = []uint32{p[1]}
	bound[p[1]] = roaring.New()
	preds, err := g.osp(tx, p[2], p[0])
	if err == ErrNotFound {
		return sol, nil
	}
	if err != nil {
		return sol, err
	}
	iterP := preds.Iterator()
	for iterP.HasNext() {
		sol.Rows = append(sol.Rows, []uint32{iterP.Next()})
	}
	bound[p[1]].Or(preds)

	return sol, nil
}

func (g *Graph) scanPOSwBoundSO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanPOSwBoundSO(%v) bound %v\n", p[:3], bound)
	// P must be a concrete, exisiting node.

	// SO are bound, so they must be variables
	sol := query.EncSolutions{
		Vars: []uint32{p[0], p[2]},
	}

	boundS := bound[p[0]]
	boundO := bound[p[2]]
	iterO := boundO.Iterator()
	missO := roaring.New()

	for iterO.HasNext() {
		oID := iterO.Next()
		subjs, err := g.pos(tx, p[1], oID)
		if err == ErrNotFound {
			missO.Add(oID)
			boundS.AndNot(subjs)
			continue
		}
		if err != nil {
			return sol, err
		}
		iterS := subjs.Iterator()
		for iterS.HasNext() {
			sol.Rows = append(sol.Rows, []uint32{iterS.Next(), oID})
		}
		boundS.And(subjs)
	}

	boundO.AndNot(missO)

	return sol, nil
}

func (g *Graph) scanSPOwBoundS(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanSPOwBoundS(%v) bound %v\n", p[:3], bound)
	// P must be a concrete, exisiting node.

	// S is bound, so must be a variable
	sol := query.EncSolutions{
		Vars: []uint32{p[0]},
	}
	if p[2] < query.MaxVariables {
		sol.Vars = append(sol.Vars, p[2])
		bound[p[2]] = roaring.New()
	}

	boundS := bound[p[0]]
	iterS := boundS.Iterator()
	missS := roaring.New()

	for iterS.HasNext() {
		sID := iterS.Next()
		objs, err := g.spo(tx, sID, p[1])
		if err == ErrNotFound {
			missS.Add(sID)
			continue
		} else if err != nil {
			return sol, err
		}

		if p[2] < query.MaxVariables {
			iterO := objs.Iterator()
			for iterO.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{sID, iterO.Next()})
			}
			bound[p[2]].Or(objs)
		} else {
			sol.Rows = append(sol.Rows, []uint32{sID})
		}
	}

	boundS.AndNot(missS)

	return sol, nil
}

func (g *Graph) scanPOSwBoundO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanPOSwBoundO(%v) bound %v\n", p[:3], bound)
	// P must be a concrete, exisiting node.

	sol := query.EncSolutions{}

	if p[0] < query.MaxVariables {
		sol.Vars = []uint32{p[0], p[2]}
		bound[p[0]] = roaring.New()
	} else {
		// O is bound, so must be a variable
		sol.Vars = []uint32{p[2]}
	}

	boundO := bound[p[2]]
	iterO := boundO.Iterator()
	missO := roaring.New()

	for iterO.HasNext() {
		oID := iterO.Next()
		subjs, err := g.pos(tx, p[1], oID)
		if err == ErrNotFound {
			missO.Add(oID)
			if p[0] < query.MaxVariables {
				bound[p[0]].AndNot(subjs)
			}
			continue
		}
		if err != nil {
			return sol, err
		}

		if p[0] < query.MaxVariables {
			iterS := subjs.Iterator()
			for iterS.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{iterS.Next(), oID})
			}
			bound[p[0]].Or(subjs)
		} else {
			sol.Rows = append(sol.Rows, []uint32{oID})
		}
	}

	boundO.AndNot(missO)

	return sol, nil
}

func (g *Graph) scanPOS(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanPOS(%v) bound %v\n", p[:3], bound)
	// P must be a concrete, exisiting node.
	// At least one, or both of S & O must be variable.

	sol := query.EncSolutions{}

	if p[0] < query.MaxVariables {
		if p[2] < query.MaxVariables {
			// S and O are both variables
			sol.Vars = []uint32{p[0], p[2]}
			bound[p[0]] = roaring.New()
			bound[p[2]] = roaring.New()

			bp := u32tob(p[1])
			cur := tx.Bucket(bucketPOS).Cursor()
		outerPOS:
			for k, _ := cur.Seek(bp); k != nil; k, _ = cur.Next() {
				switch bytes.Compare(k[:4], bp) {
				case 0:
					oID := btou32(k[4:])
					subjs, err := g.pos(tx, p[1], oID)
					if err != nil {
						// err cannot be ErrNotFound
						return sol, err
					}
					bound[p[2]].Add(oID)
					bound[p[0]].Or(subjs)
					itr := subjs.Iterator()
					for itr.HasNext() {
						sol.Rows = append(sol.Rows, []uint32{itr.Next(), oID})
					}
				case 1:
					break outerPOS
				}
			}
			return sol, nil
		}

		// S is variable, O is concrete
		sol.Vars = []uint32{p[0]}
		bound[p[0]] = roaring.New()

		subjs, err := g.pos(tx, p[1], p[2])
		if err == ErrNotFound {
			return sol, nil
		}
		if err != nil {
			return sol, err
		}
		bound[p[0]].Or(subjs)
		iterS := subjs.Iterator()
		for iterS.HasNext() {
			sol.Rows = append(sol.Rows, []uint32{iterS.Next()})
		}
		return sol, nil
	}

	// O is variable, S is concrete
	sol.Vars = []uint32{p[2]}
	bound[p[2]] = roaring.New()
	objs, err := g.spo(tx, p[0], p[1])
	if err == ErrNotFound {
		return sol, nil
	}
	if err != nil {
		return sol, err
	}
	itr := objs.Iterator()
	for itr.HasNext() {
		sol.Rows = append(sol.Rows, []uint32{itr.Next()})
	}
	bound[p[2]].Or(objs)

	return sol, nil
}

func (g *Graph) scanSPOwBoundPO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// S must be a concrete, exisiting node.

	// PO are bound, so they must be variables
	sol := query.EncSolutions{
		Vars: []uint32{p[1], p[2]},
	}

	boundP := bound[p[1]]
	boundO := bound[p[2]]
	iterP := boundP.Iterator()
	missP := roaring.New()

	for iterP.HasNext() {
		pID := iterP.Next()
		objs, err := g.spo(tx, p[0], pID)
		if err == ErrNotFound {
			missP.Add(pID)
			continue
		} else if err != nil {
			return sol, err
		}
		iterO := objs.Iterator()
		for iterO.HasNext() {
			sol.Rows = append(sol.Rows, []uint32{pID, iterO.Next()})
		}
		boundO.And(objs)
	}

	boundP.AndNot(missP)

	return sol, nil
}

func (g *Graph) scanSPOwBoundP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// S must be a concrete, exisiting node.

	sol := query.EncSolutions{}

	if p[2] < query.MaxVariables {
		sol.Vars = []uint32{p[1], p[2]}
		bound[p[2]] = roaring.New()
	} else {
		// P is bound, so must be a variable
		sol.Vars = []uint32{p[1]}
	}

	boundP := bound[p[1]]
	iterP := boundP.Iterator()
	missP := roaring.New()

	for iterP.HasNext() {
		pID := iterP.Next()
		objs, err := g.spo(tx, p[0], pID)
		if err == ErrNotFound {
			missP.Add(pID)
			continue
		} else if err != nil {
			return sol, err
		}

		if p[2] < query.MaxVariables {
			bound[p[2]].And(objs)
			iterO := objs.Iterator()
			for iterO.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{pID, iterO.Next()})
			}
		} else {
			sol.Rows = append(sol.Rows, []uint32{pID})
		}
	}

	boundP.AndNot(missP)

	return sol, nil
}

func (g *Graph) scanOSPwBoundO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanOSPwBoundO(%v) bound %v\n", p[:3], bound)
	// S must be a concrete, exisiting node.

	sol := query.EncSolutions{}

	if p[1] < query.MaxVariables {
		sol.Vars = []uint32{p[1], p[2]}
		bound[p[1]] = roaring.New()
	} else {
		// O is bound, so must be a variable
		sol.Vars = []uint32{p[2]}
	}

	boundO := bound[p[1]]
	iterO := boundO.Iterator()
	missO := roaring.New()

	for iterO.HasNext() {
		oID := iterO.Next()
		preds, err := g.osp(tx, oID, p[0])
		if err == ErrNotFound {
			missO.Add(oID)
			continue
		} else if err != nil {
			return sol, err
		}

		if p[1] < query.MaxVariables {
			iterP := preds.Iterator()
			for iterP.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{iterP.Next(), oID})
			}
			bound[p[1]].And(preds)
		} else {
			sol.Rows = append(sol.Rows, []uint32{oID})
		}
	}

	boundO.AndNot(missO)

	return sol, nil
}

func (g *Graph) scanSPO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanSPO(%v) bound %v\n", p[:3], bound)
	// S must be a concrete, exisiting node.
	// At least one, or both of P & O must be variable.

	sol := query.EncSolutions{}

	// P is variable
	if p[1] < query.MaxVariables {
		if p[2] < query.MaxVariables {
			// Both P & O are variable
			sol.Vars = []uint32{p[1], p[2]}
			bound[p[1]] = roaring.New()
			bound[p[2]] = roaring.New()

			bs := u32tob(p[0])
			cur := tx.Bucket(bucketSPO).Cursor()
		outerSPO:
			for k, _ := cur.Seek(bs); k != nil; k, _ = cur.Next() {
				switch bytes.Compare(k[:4], bs) {
				case 0:
					pID := btou32(k[4:])
					objs, err := g.spo(tx, p[0], pID)
					if err != nil {
						// err cannot be ErrNotFound
						return sol, err
					}
					bound[p[1]].Add(pID)
					iterO := objs.Iterator()
					for iterO.HasNext() {
						sol.Rows = append(sol.Rows, []uint32{pID, iterO.Next()})
					}
					bound[p[2]].Or(objs)
				case 1:
					break outerSPO
				}
			}
			return sol, nil
		}
		// P is variable, O is concrete
		sol.Vars = []uint32{p[1]}
		bound[p[1]] = roaring.New()

		bs := u32tob(p[0])
		cur := tx.Bucket(bucketSPO).Cursor()
	outerSPO2:
		for k, _ := cur.Seek(bs); k != nil; k, _ = cur.Next() {
			switch bytes.Compare(k[:4], bs) {
			case 0:
				pID := btou32(k[4:])
				objs, err := g.spo(tx, p[0], pID)
				if err != nil {
					// err cannot be ErrNotFound
					return sol, err
				}
				if objs.Contains(p[2]) {
					bound[p[1]].Add(pID)
					sol.Rows = append(sol.Rows, []uint32{pID})
				}
			case 1:
				break outerSPO2
			}
		}
		return sol, nil
	}

	// O is variable, P is concrete
	sol.Vars = []uint32{p[2]}
	bound[p[2]] = roaring.New()
	objs, err := g.spo(tx, p[0], p[1])
	if err == ErrNotFound {
		return sol, nil
	}
	if err != nil {
		return sol, err
	}
	itr := objs.Iterator()
	for itr.HasNext() {
		oID := itr.Next()
		bound[p[2]].Add(oID)
		sol.Rows = append(sol.Rows, []uint32{oID})
	}

	return sol, nil
}

func (g *Graph) scanAllwBoundSPO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// SPO are all variables, and bound
	sol := query.EncSolutions{
		Vars: []uint32{p[0], p[1], p[2]},
	}

	boundS := bound[p[0]]
	boundP := bound[p[1]]
	boundO := bound[p[2]]
	iterS := boundS.Iterator()
	iterP := boundP.Iterator()
	missS := roaring.New()
	//missP := roaring.New()

	for iterS.HasNext() {
		sID := iterS.Next()
		for iterP.HasNext() {
			pID := iterP.Next()
			objs, err := g.spo(tx, sID, pID)
			if err == ErrNotFound {
				missS.Add(sID)
				//missP.Add(pID) TODO ?
				continue
			} else if err != nil {
				return sol, err
			}
			boundO.And(objs)
			iterO := boundO.Iterator()
			for iterO.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{sID, pID, iterO.Next()})
			}

		}
	}

	boundS.AndNot(missS)
	//boundP.AndNot(missP) TODO ?

	return sol, nil
}

func (g *Graph) scanAllwBoundSP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// SPO are all variables

	sol := query.EncSolutions{}

	sol.Vars = []uint32{p[0], p[1], p[2]}
	bound[p[2]] = roaring.New()

	boundS := bound[p[0]]
	boundP := bound[p[1]]
	iterS := boundS.Iterator()
	iterP := boundP.Iterator()
	missS := roaring.New()
	//missP := roaring.New()

	for iterS.HasNext() {
		sID := iterS.Next()
		for iterP.HasNext() {
			pID := iterP.Next()
			objs, err := g.spo(tx, sID, pID)
			if err == ErrNotFound {
				missS.Add(sID)
				//missP.Add(pID) TODO ?
				continue
			} else if err != nil {
				return sol, err
			}
			iterO := objs.Iterator()
			for iterO.HasNext() {
				oID := iterO.Next()
				sol.Rows = append(sol.Rows, []uint32{sID, pID, oID})
			}
			bound[p[2]].Or(objs)
		}
	}

	boundS.AndNot(missS)
	//boundP.AndNot(missP) TODO ?

	return sol, nil
}

func (g *Graph) scanAllwBoundSO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// SPO are all variables

	sol := query.EncSolutions{}

	sol.Vars = []uint32{p[0], p[1], p[2]}
	bound[p[1]] = roaring.New()

	boundS := bound[p[0]]
	boundO := bound[p[2]]
	iterS := boundS.Iterator()
	for iterS.HasNext() {
		sID := iterS.Next()
		bs := u32tob(sID)
		cur := tx.Bucket(bucketSPO).Cursor()
	outerSPO:
		for k, _ := cur.Seek(bs); k != nil; k, _ = cur.Next() {
			switch bytes.Compare(k[:4], bs) {
			case 0:
				pID := btou32(k[4:])
				objs, err := g.osp(tx, sID, pID)
				if err != nil {
					// err cannot be ErrNotFound
					return sol, err
				}
				bound[p[1]].Add(pID)
				boundO.And(objs)
				iterO := boundO.Iterator()
				for iterO.HasNext() {
					sol.Rows = append(sol.Rows, []uint32{sID, pID, iterO.Next()})
				}
			case 1:
				break outerSPO
			}
		}
	}

	return sol, nil
}

func (g *Graph) scanAllwBoundS(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanAllwBoundS(%v) bound %v\n", p[:3], bound)
	// SPO are all variables

	sol := query.EncSolutions{}

	sol.Vars = []uint32{p[0], p[1], p[2]}
	bound[p[1]] = roaring.New()
	bound[p[2]] = roaring.New()

	iterS := bound[p[0]].Iterator()
	for iterS.HasNext() {
		sID := iterS.Next()
		bs := u32tob(sID)
		cur := tx.Bucket(bucketSPO).Cursor()
	outerSPO:
		for k, _ := cur.Seek(bs); k != nil; k, _ = cur.Next() {
			switch bytes.Compare(k[:4], bs) {
			case 0:
				pID := btou32(k[4:])
				objs, err := g.spo(tx, sID, pID)
				if err != nil {
					// err cannot be ErrNotFound
					return sol, err
				}
				bound[p[1]].Add(pID)
				iterO := objs.Iterator()
				for iterO.HasNext() {
					sol.Rows = append(sol.Rows, []uint32{sID, pID, iterO.Next()})
				}
				bound[p[2]].Or(objs)
			case 1:
				break outerSPO
			}
		}
	}

	return sol, nil
}

func (g *Graph) scanAllwBoundPO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanAll(%v) bound %v\n", p[:3], bound)
	// SPO are all variables

	sol := query.EncSolutions{}

	sol.Vars = []uint32{p[0], p[1], p[2]}
	bound[p[0]] = roaring.New()

	boundP := bound[p[1]]
	boundO := bound[p[2]]
	iterP := boundP.Iterator()
	iterO := boundO.Iterator()
	missP := roaring.New()
	//missO := roaring.New()

	for iterP.HasNext() {
		pID := iterP.Next()
		for iterO.HasNext() {
			oID := iterO.Next()
			subjs, err := g.pos(tx, pID, oID)
			if err == ErrNotFound {
				missP.Add(pID)
				//missS.Add(sID) TODO ?
				continue
			} else if err != nil {
				return sol, err
			}
			iterS := subjs.Iterator()
			for iterS.HasNext() {
				sol.Rows = append(sol.Rows, []uint32{iterS.Next(), pID, oID})
			}
			bound[p[0]].Or(subjs)
		}
	}

	//boundO.AndNot(missO) // TODO ?
	boundP.AndNot(missP)

	return sol, nil
}

func (g *Graph) scanAllwBoundP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// SPO are all variables

	sol := query.EncSolutions{}

	sol.Vars = []uint32{p[0], p[1], p[2]}
	bound[p[0]] = roaring.New()
	bound[p[2]] = roaring.New()

	iterP := bound[p[1]].Iterator()
	for iterP.HasNext() {
		pID := iterP.Next()
		bp := u32tob(pID)
		cur := tx.Bucket(bucketPOS).Cursor()
	outerSPO:
		for k, _ := cur.Seek(bp); k != nil; k, _ = cur.Next() {
			switch bytes.Compare(k[:4], bp) {
			case 0:
				oID := btou32(k[4:])
				subjs, err := g.pos(tx, pID, oID)
				if err != nil {
					// err cannot be ErrNotFound
					return sol, err
				}
				bound[p[1]].Add(pID)
				iterS := subjs.Iterator()
				for iterS.HasNext() {
					sol.Rows = append(sol.Rows, []uint32{iterS.Next(), pID, oID})
				}
				bound[p[0]].Or(subjs)
			case 1:
				break outerSPO
			}
		}
	}

	return sol, nil
}

func (g *Graph) scanAllwBoundO(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	//fmt.Printf("scanAllwBoundO(%v) bound %v\n", p[:3], bound)
	// SPO are all variables

	sol := query.EncSolutions{}

	sol.Vars = []uint32{p[0], p[1], p[2]}
	bound[p[0]] = roaring.New()
	bound[p[1]] = roaring.New()

	iterO := bound[p[2]].Iterator()
	for iterO.HasNext() {
		oID := iterO.Next()
		bo := u32tob(oID)
		cur := tx.Bucket(bucketOSP).Cursor()
	outerSPO:
		for k, _ := cur.Seek(bo); k != nil; k, _ = cur.Next() {
			switch bytes.Compare(k[:4], bo) {
			case 0:
				sID := btou32(k[4:])
				preds, err := g.osp(tx, oID, sID)
				if err != nil {
					// err cannot be ErrNotFound
					return sol, err
				}
				bound[p[0]].Add(sID)
				iterP := preds.Iterator()
				for iterP.HasNext() {
					sol.Rows = append(sol.Rows, []uint32{sID, iterP.Next(), oID})
				}
				bound[p[1]].Or(preds)
			case 1:
				break outerSPO
			}
		}
	}

	return sol, nil
}

func (g *Graph) scanAll(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {
	// SPO are all variables

	// We don't bother to update bound, since this should always be
	// the last pattern evaluated, since it asks for all the triples in the graph.

	sol := query.EncSolutions{
		Vars: []uint32{p[0], p[1], p[2]},
	}

	cur := tx.Bucket(bucketSPO).Cursor()
	for k, _ := cur.First(); k != nil; k, _ = cur.Next() {
		sID, pID := btou32(k[:4]), btou32(k[4:])
		objs, err := g.spo(tx, sID, pID)
		if err != nil {
			return sol, err
		}
		iterO := objs.Iterator()
		for iterO.HasNext() {
			sol.Rows = append(sol.Rows, []uint32{sID, pID, iterO.Next()})
		}
	}

	return sol, nil
}

func (g *Graph) evalBGP(tx *bolt.Tx, p query.EncTriplePattern, bound map[uint32]*roaring.Bitmap) (query.EncSolutions, error) {

	if p.IsConcrete() {
		// It's a concrete triple - no variables to be bound.
		//fmt.Printf("Concrete(%v)\n", p[:3])
		return query.EncSolutions{}, nil
	}

	_, sIsBound := bound[p[0]]
	_, pIsBound := bound[p[1]]
	_, oIsBound := bound[p[2]]

	// O is concrete
	if p[2] > query.MaxVariables {
		if sIsBound {
			if pIsBound {
				// SP bound
				return g.scanOSPwBoundSP(tx, p, bound)
			}
			// S bound
			return g.scanOSPwBoundS(tx, p, bound)
		}
		if pIsBound {
			// P bound
			return g.scanPOSwBoundP(tx, p, bound)
		}
		// None are bound
		return g.scanOSP(tx, p, bound)
	}

	// P is concrete
	if p[1] > query.MaxVariables {
		if sIsBound {
			if oIsBound {
				// SO bound
				return g.scanPOSwBoundSO(tx, p, bound)
			}
			// S bound
			return g.scanSPOwBoundS(tx, p, bound)
		}
		if oIsBound {
			// O bound
			return g.scanPOSwBoundO(tx, p, bound)
		}
		// None are bound
		return g.scanPOS(tx, p, bound)
	}

	// S is concrete
	if p[0] > query.MaxVariables {
		if pIsBound {
			if oIsBound {
				// PO bound
				return g.scanSPOwBoundPO(tx, p, bound)
			}
			// P bound
			return g.scanSPOwBoundP(tx, p, bound)
		}
		if oIsBound {
			// O bound
			return g.scanOSPwBoundO(tx, p, bound)
		}
		// None are bound
		return g.scanSPO(tx, p, bound)
	}

	// SPO are all variables
	if sIsBound {
		if pIsBound {
			if oIsBound {
				// SPO bound
				return g.scanAllwBoundSPO(tx, p, bound)
			}
			// SP bound
			return g.scanAllwBoundSP(tx, p, bound)
		}
		if oIsBound {
			// SO bound
			return g.scanAllwBoundSO(tx, p, bound)
		}
		// S bound
		return g.scanAllwBoundS(tx, p, bound)
	}
	if pIsBound {
		if oIsBound {
			// PO bound
			return g.scanAllwBoundPO(tx, p, bound)
		}
		// P bound
		return g.scanAllwBoundP(tx, p, bound)
	}
	if oIsBound {
		// O bound
		return g.scanAllwBoundO(tx, p, bound)
	}
	// None are bound
	return g.scanAll(tx, p, bound)
}

func (g *Graph) where(tx *bolt.Tx, patterns []query.EncTriplePattern) ([][]query.EncTriplePattern, query.EncSolutions, error) {
	// Group disjoint pattern groups and sort them by estimated cardinality
	groups := query.GroupPatterns(patterns)
	for n := range groups {
		sort.Slice(groups[n], func(i, j int) bool {
			return groups[n][i][3] < groups[n][j][3]
		})
	}

	bound := make(map[uint32]*roaring.Bitmap)
	var left, right query.EncSolutions
	var err error

	// Evaluate each BGP sequentially, each BGP using any bound values from previous BGPs
	for _, group := range groups {
		for _, pattern := range group {
			right, err = g.evalBGP(tx, pattern, bound)
			if err != nil {
				return nil, right, err
			}
			//fmt.Printf("joining: %v %v:\n", left, right)
			left = left.Join(right, bound)
			//fmt.Printf("%v\n", left)
		}
	}

	return groups, left, nil
}

func (g *Graph) substitute(tx *bolt.Tx, p query.EncTriplePattern, s query.EncSolutions) (res []rdf.Triple, err error) {
	for _, row := range s.Rows {
		var tr rdf.Triple
		for i, v := range s.Vars {
			if v == p[0] {
				node, err := g.getNode(tx, row[i])
				if err != nil {
					return nil, err
				}
				tr.Subject = node.(rdf.SubjectNode)
			} else if p[0] > query.MaxVariables {
				node, err := g.getNode(tx, p[0])
				if err != nil {
					return nil, err
				}
				tr.Subject = node.(rdf.SubjectNode)
			}

			if v == p[1] {
				node, err := g.getNode(tx, row[i])
				if err != nil {
					return nil, err
				}
				tr.Predicate = node.(rdf.NamedNode)
			} else if p[1] > query.MaxVariables {
				node, err := g.getNode(tx, p[1])
				if err != nil {
					return nil, err
				}
				tr.Predicate = node.(rdf.NamedNode)
			}

			if v == p[2] {
				node, err := g.getNode(tx, row[i])
				if err != nil {
					return nil, err
				}
				tr.Object = node
			} else if p[2] > query.MaxVariables {
				node, err := g.getNode(tx, p[2])
				if err != nil {
					return nil, err
				}
				tr.Object = node
			}
		}

		res = append(res, tr)
	}
	return res, nil
}

func (g *Graph) deleteFor(tx *bolt.Tx, p rdf.TriplePattern, s query.EncSolutions, cache map[rdf.TriplePatternNode]uint32) (int, error) {
	var encP [3]uint32
	for i, node := range []rdf.TriplePatternNode{p.Subject, p.Predicate, p.Object} {
		id, ok := cache[node]
		if ok {
			encP[i] = id
			continue
		}
		if _, ok := node.(rdf.Variable); ok {
			panic("BUG: deleteFor got variable not in cache")
		}

		id, err := g.getID(tx, node.(rdf.Node))
		if err == ErrNotFound {
			// Node is not stored, so we have no triple
			return 0, nil
		}
		if err != nil {
			return 0, err
		}
		encP[i] = id
		cache[node] = id
	}

	if p.IsConcrete() {
		if err := g.removeTriple(tx, encP[0], encP[1], encP[2]); err != nil {
			return 0, err
		}
		return 1, nil
	}

	var n int
	tr := [3]uint32{encP[0], encP[1], encP[2]}
	for _, row := range s.Rows {
		for i, v := range s.Vars {
			if v == encP[0] {
				tr[0] = row[i]
			}
			if v == encP[1] {
				tr[1] = row[i]
			}
			if v == encP[2] {
				tr[2] = row[i]
			}
		}
		err := g.removeTriple(tx, tr[0], tr[1], tr[2])
		if err == ErrNotFound {
			continue
		}
		if err != nil {
			// TODO, how can removeTriple be called with non-existing triple?
			// Maybe It means one of the solution rows is wrong
			continue //return n, err
		}
		n++
	}

	return n, nil
}

func (g *Graph) insertFor(tx *bolt.Tx, p rdf.TriplePattern, s query.EncSolutions, cache map[rdf.TriplePatternNode]uint32) (int, error) {
	var encP [3]uint32
	for i, node := range []rdf.TriplePatternNode{p.Subject, p.Predicate, p.Object} {
		id, ok := cache[node]
		if ok {
			encP[i] = id
			continue
		}
		if _, ok := node.(rdf.Variable); ok {
			panic("BUG: insertFor got variable not in cache")
		}

		id, err := g.addNode(tx, node.(rdf.Node))
		if err != nil {
			return 0, nil
		}
		encP[i] = id
		cache[node] = id
	}

	if p.IsConcrete() {
		stored, err := g.storeTriple(tx, encP[0], encP[1], encP[2])
		if err != nil || !stored {
			return 0, err
		}
		return 1, nil
	}

	var n int
	tr := [3]uint32{encP[0], encP[1], encP[2]}
	for _, row := range s.Rows {
		for i, v := range s.Vars {
			if v == encP[0] {
				tr[0] = row[i]
			}
			if v == encP[1] {
				tr[1] = row[i]
			}
			if v == encP[2] {
				tr[2] = row[i]
			}
		}
		stored, err := g.storeTriple(tx, tr[0], tr[1], tr[2])
		if err != nil {
			return n, err
		}
		if stored {
			n++
		}
	}

	return n, nil
}

func (g *Graph) Where(patterns ...rdf.TriplePattern) (rdf.Graph, error) {
	// Fast path for no patterns
	if len(patterns) == 0 {
		return memory.NewGraph(), nil
	}

	res := memory.NewGraph()

	if err := g.kv.View(func(tx *bolt.Tx) error {
		// Encode patterns
		encPatterns := make([]query.EncTriplePattern, len(patterns))
		cache := make(map[rdf.TriplePatternNode]uint32)
		numVars := 0
		for i, pattern := range patterns {
			var err error
			encPatterns[i], err = g.encodePattern(tx, pattern, cache, &numVars)
			//fmt.Printf("%v -> %v\n", pattern, encPatterns[i][:3])
			if err != nil {
				return err
			}
		}

		// Evaluate query
		groups, solutions, err := g.where(tx, encPatterns)
		if err != nil {
			return err
		}

		// Substitute variables with solutions
		if len(solutions.Rows) > 0 {
			for _, group := range groups {
				for _, pattern := range group {
					triples, err := g.substitute(tx, pattern, solutions)
					if err != nil {
						return err
					}
					res.Insert(triples...)
				}
			}
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return res, nil
}

func (g *Graph) getID(tx *bolt.Tx, node rdf.Node) (id uint32, err error) {
	bt, err := g.encode(tx, node)
	if err != nil {
		return 0, err
	}
	b := tx.Bucket(bucketNode2Id).Get(bt)
	if b == nil {
		err = ErrNotFound
	} else {
		id = btou32(b)
	}
	return id, err
}

func (g *Graph) getIDb(tx *bolt.Tx, node []byte) (id uint32, err error) {
	b := tx.Bucket(bucketNode2Id).Get(node)
	if b == nil {
		err = ErrNotFound
	} else {
		id = btou32(b)
	}
	return id, err
}

// getNode returns the node for a given ID.
func (g *Graph) getNode(tx *bolt.Tx, id uint32) (rdf.Node, error) {
	b := tx.Bucket(bucketId2Node).Get(u32tob(id))
	if b == nil {
		return nil, ErrNotFound
	}
	return g.decode(b)
}

func (g *Graph) addNode(tx *bolt.Tx, node rdf.Node) (id uint32, err error) {
	b, err := g.encode(tx, node)
	if err != nil {
		return 0, err
	}

	if id, err = g.getIDb(tx, b); err == nil {
		// Term is allready in database
		return id, nil
	} else if err != ErrNotFound {
		// Some other IO error occured
		return 0, err
	}

	// get a new ID
	bkt := tx.Bucket(bucketId2Node)
	n, err := bkt.NextSequence()
	if err != nil {
		return 0, err
	}
	n += query.MaxVariables
	if n > MaxNodes {
		return 0, ErrDBFull
	}

	id = uint32(n)
	idb := u32tob(uint32(n))

	// store node in dictionaries
	err = bkt.Put(idb, b)
	if err != nil {
		return 0, err
	}
	err = tx.Bucket(bucketNode2Id).Put(b, idb)

	return id, err
}

func (g *Graph) storeTriple(tx *bolt.Tx, s, p, o uint32) (bool, error) {
	indices := []struct {
		k1 uint32
		k2 uint32
		v  uint32
		bk []byte
	}{
		{s, p, o, bucketSPO},
		{o, s, p, bucketOSP},
		{p, o, s, bucketPOS},
	}

	key := make([]byte, 8)

	for _, i := range indices {
		bkt := tx.Bucket(i.bk)
		copy(key, u32tob(i.k1))
		copy(key[4:], u32tob(i.k2))
		bitmap := roaring.NewBitmap()

		bo := bkt.Get(key)
		if bo != nil {
			_, err := bitmap.ReadFrom(bytes.NewReader(bo))
			if err != nil {
				return false, err
			}
		}

		newTriple := bitmap.CheckedAdd(i.v)
		if !newTriple {
			// Triple is allready stored
			return false, nil
		}
		b, err := bitmap.ToBytes()
		if err != nil {
			return false, err
		}
		err = bkt.Put(key, b)
		if err != nil {
			return false, err
		}
	}

	return true, nil
}

func (g *Graph) removeTriple(tx *bolt.Tx, s, p, o uint32) error {
	indices := []struct {
		k1 uint32
		k2 uint32
		v  uint32
		bk []byte
	}{
		{s, p, o, bucketSPO},
		{o, s, p, bucketOSP},
		{p, o, s, bucketPOS},
	}

	key := make([]byte, 8)
	found := false
	for _, i := range indices {
		bkt := tx.Bucket(i.bk)
		copy(key, u32tob(i.k1))
		copy(key[4:], u32tob(i.k2))

		bitmap := roaring.New()
		_, err := bitmap.ReadFrom(bytes.NewReader(bkt.Get(key)))
		if err != nil {
			return err
		}
		hasTriple := bitmap.CheckedRemove(i.v)
		if !hasTriple {
			if found {
				panic("BUG: triple stored in one, but not all indexes")
			}
			return ErrNotFound
		}
		found = true
		// Remove from index if bitmap is empty
		if bitmap.GetCardinality() == 0 {
			err = bkt.Delete(key)
			if err != nil {
				return err
			}
		} else {
			b, err := bitmap.ToBytes()
			if err != nil {
				return err
			}
			err = bkt.Put(key, b)
			if err != nil {
				return err
			}
		}
	}

	return nil
	// TODO return db.removeOrphanedTerms(tx, s, p, o)
}

func (g *Graph) tripleNodeIDs(tx *bolt.Tx, tr rdf.Triple, cache map[rdf.Node]uint32) ([3]uint32, error) {
	ids := [3]uint32{}

	if s, ok := cache[tr.Subject]; ok {
		ids[0] = s
	} else {
		s, err := g.getID(tx, tr.Subject)
		if err != nil {
			return ids, err
		}
		ids[0] = s
	}

	if p, ok := cache[tr.Predicate]; ok {
		ids[1] = p
	} else {
		p, err := g.getID(tx, tr.Predicate)
		if err != nil {
			return ids, err
		}
		ids[1] = p
	}

	if o, ok := cache[tr.Object]; ok {
		ids[2] = o
	} else {
		o, err := g.getID(tx, tr.Object)
		if err != nil {
			return ids, err
		}
		ids[2] = o
	}

	return ids, nil
}

func (g *Graph) Insert(trs ...rdf.Triple) (int, error) {

	cache := make(map[rdf.Node]uint32)
	n := 0

	txErr := g.kv.Update(func(tx *bolt.Tx) error {
		for _, tr := range trs {
			var err error
			sID, ok := cache[tr.Subject]
			if !ok {
				sID, err = g.addNode(tx, tr.Subject)
				if err != nil {
					return err
				}
				cache[tr.Subject] = sID
			}

			pID, ok := cache[tr.Predicate]
			if !ok {
				pID, err = g.addNode(tx, tr.Predicate)
				if err != nil {
					return err
				}
				cache[tr.Predicate] = pID
			}

			oID, ok := cache[tr.Object]
			if !ok {
				oID, err = g.addNode(tx, tr.Object)
				if err != nil {
					return err
				}
				cache[tr.Object] = oID
			}

			stored, err := g.storeTriple(tx, sID, pID, oID)
			if err != nil {
				return err
			}
			if stored {
				n++
			}
		}
		return nil
	})
	return n, txErr
}

func (g *Graph) Delete(trs ...rdf.Triple) (int, error) {

	cache := make(map[rdf.Node]uint32)
	n := 0

	txErr := g.kv.Update(func(tx *bolt.Tx) error {
		for _, tr := range trs {
			// Skip triples with blank nodes.
			// TODO or return error?
			if _, ok := tr.Subject.(rdf.BlankNode); ok {
				continue
			}
			if _, ok := tr.Object.(rdf.BlankNode); ok {
				continue
			}

			ids, err := g.tripleNodeIDs(tx, tr, cache)
			if err == ErrNotFound {
				continue
			}
			if err != nil {
				return err
			}

			err = g.removeTriple(tx, ids[0], ids[1], ids[2])
			if err == ErrNotFound {
				continue
			}
			if err != nil {
				return err
			}
			n++
		}
		return nil
	})

	return n, txErr
}

func (g *Graph) Update(del []rdf.TriplePattern, ins []rdf.TriplePattern, where []rdf.TriplePattern) (delN, insN int, err error) {
	if where != nil {
		err = g.kv.Update(func(tx *bolt.Tx) error {
			// Encode patterns
			encWhere := make([]query.EncTriplePattern, len(where))
			cache := make(map[rdf.TriplePatternNode]uint32)
			numVars := 0
			for i, pattern := range where {
				var er error
				encWhere[i], er = g.encodePattern(tx, pattern, cache, &numVars)
				if er != nil {
					return err
				}
				//fmt.Printf("\n%v => %v\n", pattern, encWhere[i][:3])
			}

			// Evaluate query
			_, solutions, er := g.where(tx, encWhere)
			if er != nil {
				return er
			}

			// Substitute variables with solutions
			if len(solutions.Rows) > 0 {
				for _, pattern := range ins {
					n, er := g.insertFor(tx, pattern, solutions, cache)
					if er != nil {
						return er
					}
					insN += n
				}

				for _, pattern := range del {
					n, er := g.deleteFor(tx, pattern, solutions, cache)
					if er != nil {
						return er
					}
					delN += n
				}

			}

			return nil
		})
		return delN, insN, err
	}

	// TODO delete and insert operations must be part of same transaction
	if del != nil {
		trs := make([]rdf.Triple, len(del))
		for i, p := range del {
			if !p.IsConcrete() {
				trs = trs[:len(trs)-1]
				continue
			}
			trs[i] = rdf.Triple{
				Subject:   p.Subject.(rdf.SubjectNode),
				Predicate: p.Predicate.(rdf.NamedNode),
				Object:    p.Object.(rdf.Node),
			}
		}
		delN, err = g.Delete(trs...)
		if err != nil {
			return
		}
	}
	if ins != nil {
		trs := make([]rdf.Triple, len(ins))
		for i, p := range ins {
			if !p.IsConcrete() {
				trs = trs[:len(trs)-1]
				continue
			}
			trs[i] = rdf.Triple{
				Subject:   p.Subject.(rdf.SubjectNode),
				Predicate: p.Predicate.(rdf.NamedNode),
				Object:    p.Object.(rdf.Node),
			}
		}
		insN, err = g.Insert(trs...)
		if err != nil {
			return
		}
	}

	return delN, insN, err
}

func (g *Graph) Select(vars []rdf.Variable, patterns ...rdf.TriplePattern) ([][]rdf.Node, error) {
	var res [][]rdf.Node

	// Fast path for no patterns
	if len(patterns) == 0 {
		return res, nil
	}

	txErr := g.kv.View(func(tx *bolt.Tx) error {
		// Encode patterns
		encPatterns := make([]query.EncTriplePattern, len(patterns))
		cache := make(map[rdf.TriplePatternNode]uint32)
		numVars := 0
		var err error
		for i, pattern := range patterns {
			encPatterns[i], err = g.encodePattern(tx, pattern, cache, &numVars)
			if err != nil {
				return err
			}
		}
		// Encode variables using the mappings from above
		encVars := make([]uint32, len(vars))
		for i, v := range vars {
			if ev, ok := cache[v]; ok {
				encVars[i] = ev
			} // else: if variable does not occur in patterns, we ignore it.
		}

		_, solutions, err := g.where(tx, encPatterns)
		if err != nil {
			return err
		}

		if len(solutions.Rows) == 0 {
			return nil
		}

		solutions = solutions.Project(encVars)

		for _, row := range solutions.Rows {
			nodes := make([]rdf.Node, 0, len(encVars))
			for _, id := range row {
				node, err := g.getNode(tx, id)
				if err != nil {
					return err
				}
				nodes = append(nodes, node)
			}

			res = append(res, nodes)
		}
		return nil
	})

	return res, txErr
}

func (g *Graph) Stats() (rdf.Stats, error) {
	return rdf.Stats{}, nil
}

func (g *Graph) Eq(other rdf.Graph) (bool, error) {
	panic("Graph.Eq: TODO")
}

func (g *Graph) encode(tx *bolt.Tx, node rdf.Node) ([]byte, error) {
	switch t := node.(type) {
	case rdf.NamedNode:
		if strings.HasPrefix(t.Name(), g.base.Name()) {
			l := len(g.base.Name())
			b := make([]byte, len(t.Name())-l+1)
			copy(b[1:], t.Name()[l:])
			return b, nil
		}
		b := make([]byte, len(t.Name())+1)
		b[0] = 0x01
		copy(b[1:], t.Name())
		return b, nil
	case rdf.BlankNode:
		b := make([]byte, 5)
		b[0] = 0x02
		n, err := tx.Bucket(bucketId2Node).NextSequence()
		if err != nil {
			return nil, err
		}
		n += query.MaxVariables
		// TODO err if n > uint32 max
		copy(b[1:], u32tob(uint32(n)))
		return b, nil
	case rdf.Literal:
		var dt byte
		switch t.DataType() {
		case rdf.XSDstring:
			dt = 0x03
		case rdf.RDFlangString:
			ll := len(t.Lang())
			b := make([]byte, len(t.ValueAsString())+ll+2)
			b[0] = 0x04
			b[1] = uint8(ll)
			copy(b[2:], []byte(t.Lang()))
			copy(b[2+ll:], []byte(t.ValueAsString()))
			return b, nil
		case rdf.XSDboolean:
			dt = 0x05
		case rdf.XSDbyte:
			dt = 0x06
		case rdf.XSDint:
			dt = 0x07
		case rdf.XSDshort:
			dt = 0x08
		case rdf.XSDlong:
			dt = 0x09
		case rdf.XSDinteger:
			dt = 0x0A
		case rdf.XSDunsignedShort:
			dt = 0x0B
		case rdf.XSDunsignedInt:
			dt = 0x0C
		case rdf.XSDunsignedLong:
			dt = 0x0D
		case rdf.XSDunsignedByte:
			dt = 0x0E
		case rdf.XSDfloat:
			dt = 0x0F
		case rdf.XSDdouble:
			dt = 0x10
		case rdf.XSDdateTimeStamp:
			dt = 0x11
		default:
			ll := len(t.DataType().Name())
			b := make([]byte, len(t.ValueAsString())+ll+2)
			b[0] = 0xFF
			b[1] = uint8(ll)
			copy(b[2:], []byte(t.DataType().Name()))
			copy(b[2+ll:], []byte(t.ValueAsString()))
			return b, nil

		}
		b := make([]byte, len(t.ValueAsString())+1)
		b[0] = dt
		copy(b[1:], t.ValueAsString())
		return b, nil
	default:
		panic("Graph.encode: unreachable")
	}
}

func (g *Graph) decode(b []byte) (rdf.Node, error) {
	// We control the encoding, so if this method fails, it's either a bug
	// or corruption in the database file.

	if len(b) == 0 {
		return nil, errors.New("Graph.decode: empty byte slice")
	}

	var dt rdf.NamedNode
	switch b[0] {
	case 0x00:
		return rdf.NewNamedNode(g.base.Name() + string(b[1:])), nil
	case 0x01:
		return rdf.NewNamedNode(string(b[1:])), nil
	case 0x02:
		return rdf.NewBlankNode(strconv.Itoa(int(btou32(b[1:])))), nil
	case 0x03:
		return rdf.NewStrLiteral(string(b[1:])), nil
	case 0x04:
		if len(b) < 2 {
			return nil, fmt.Errorf("Graph.decode: corrupt rdf:langString: %v", b)
		}
		ll := int(b[1])
		if len(b) < ll+2 {
			return nil, fmt.Errorf("Graph.decode: corrupt rdf:langString: %v", b)
		}
		return rdf.NewLangLiteral(string(b[ll+2:]), string(b[2:2+ll])), nil
	case 0x05:
		dt = rdf.XSDboolean
	case 0x06:
		dt = rdf.XSDbyte
	case 0x07:
		dt = rdf.XSDint
	case 0x08:
		dt = rdf.XSDshort
	case 0x09:
		dt = rdf.XSDlong
	case 0x0A:
		dt = rdf.XSDinteger
	case 0x0B:
		dt = rdf.XSDunsignedShort
	case 0x0C:
		dt = rdf.XSDunsignedInt
	case 0x0D:
		dt = rdf.XSDunsignedLong
	case 0x0E:
		dt = rdf.XSDunsignedByte
	case 0x0F:
		dt = rdf.XSDfloat
	case 0x10:
		dt = rdf.XSDdouble
	case 0x11:
		dt = rdf.XSDdateTimeStamp
	case 0xFF:
		if len(b) < 2 {
			return nil, fmt.Errorf("Graph.deocde: corrupt literal: %v", b)
		}
		ll := int(b[1])
		if len(b) < ll {
			return nil, fmt.Errorf("Graph.decode: corrupt literal: %v", b)
		}
		return rdf.NewTypedLiteral(string(b[ll+2:]), rdf.NewNamedNode(string(b[2:2+ll]))), nil
	default:
		return nil, fmt.Errorf("Graph.decode: unknown literal datatype: %v", b)
	}

	return rdf.NewTypedLiteral(string(b[1:]), dt), nil
}

// u32tob converts a uint32 into a 4-byte slice.
func u32tob(v uint32) []byte {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, v)
	return b
}

// btou32 converts a 4-byte slice into an uint32.
func btou32(b []byte) uint32 {
	return binary.BigEndian.Uint32(b)
}
