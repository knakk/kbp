package rdf

import (
	"sort"
	"strconv"
)

// Graph is an in-memory representation of an RDF graph:
// a colllection of triples which can be queried for subgraphs
// and mutated with inserting and deleting triples.
type Graph struct {
	// The graph is stored internally as a map of subjects to a map of
	// predicates to a slice of objects.
	nodes map[Subject]map[URI][]Term

	// nextID is the next unique integer to be used as Blank Node ID.
	nextID int
}

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[Subject]map[URI][]Term),
	}
}

// Size returns the number of triples in the Graph.
func (g *Graph) Size() int {
	n := 0
	for _, po := range g.nodes {
		for _, objs := range po {
			n += len(objs)
		}
	}
	return n
}

// Triples returns all the triples in the Graph.
func (g *Graph) Triples() []Triple {
	var res []Triple
	for s, po := range g.nodes {
		for p, objs := range po {
			for _, o := range objs {
				res = append(res, Triple{
					Subj: s,
					Pred: p,
					Obj:  o,
				})
			}
		}
	}
	return res
}

// Contains returns true if the given Triple is present in Graph, otherwise false.
func (g *Graph) Contains(tr Triple) bool {
	if _, ok := g.nodes[tr.Subj]; !ok {
		return false
	}
	if _, ok := g.nodes[tr.Subj][tr.Pred]; !ok {
		return false
	}

	for _, o := range g.nodes[tr.Subj][tr.Pred] {
		if o.Eq(tr.Obj) {
			return true
		}
	}

	return false
}

// Insert adds one or more triples to the Graph. It returns the number
// of triples inserted which where not allready present.
func (g *Graph) Insert(trs ...Triple) int {
	n := 0
	bnodes := make(map[string]string)
outer:
	for _, tr := range trs {
		switch t := tr.Subj.(type) {
		case BlankNode:
			if newID, ok := bnodes[t.id]; ok {
				tr.Subj = BlankNode{id: newID}
			} else {
				g.nextID++
				bnodes[t.id] = strconv.Itoa(g.nextID)
				tr.Subj = BlankNode{id: bnodes[t.id]}
			}
		}
		switch t := tr.Obj.(type) {
		case BlankNode:
			if newID, ok := bnodes[t.id]; ok {
				tr.Obj = BlankNode{id: newID}
			} else {
				g.nextID++
				bnodes[t.id] = strconv.Itoa(g.nextID)
				tr.Obj = BlankNode{id: bnodes[t.id]}
			}
		}
		if _, ok := g.nodes[tr.Subj]; !ok {
			// new subject
			g.nodes[tr.Subj] = make(map[URI][]Term)
		}
		if _, ok := g.nodes[tr.Subj][tr.Pred]; !ok {
			// new predicate
			g.nodes[tr.Subj][tr.Pred] = make([]Term, 0, 1)
		}
		for _, o := range g.nodes[tr.Subj][tr.Pred] {
			if o == tr.Obj {
				// triple already in graph
				continue outer
			}
		}
		// add object
		g.nodes[tr.Subj][tr.Pred] = append(g.nodes[tr.Subj][tr.Pred], tr.Obj)
		n++
	}
	return n
}

// Eq checks if two graphs are equal (isomorphic).
//
// The algorithm for checking for isomporphism is rather naive and
// inefficient, and will be slow for large graphs with lots of
// blank nodes. For graphs with few blank nodes it will be fast enough.
func (g *Graph) Eq(other *Graph) bool {
	if g.Size() != other.Size() {
		return false
	}

	var (
		aBNodes []BlankNode
		bBNodes []BlankNode
	)
	aBNodesAsObj := make(map[BlankNode][]Triple)
	bBNodesAsObj := make(map[BlankNode][]Triple)

	for subj, po := range g.nodes {
		switch t := subj.(type) {
		case BlankNode:
			aBNodes = append(aBNodes, t)
			continue
		}
		if _, ok := other.nodes[subj]; !ok {
			return false
		}
		for pred, objs := range po {
			if _, ok := other.nodes[subj][pred]; !ok {
				return false
			}
			if len(objs) != len(other.nodes[subj][pred]) {
				return false
			}
			sort.Slice(objs, func(i, j int) bool {
				return objs[i].String() < objs[j].String()
			})
			sort.Slice(other.nodes[subj][pred], func(i, j int) bool {
				return other.nodes[subj][pred][i].String() < other.nodes[subj][pred][j].String()
			})
			for i, obj := range objs {
				switch t := obj.(type) {
				case BlankNode:
					aBNodesAsObj[t] = append(aBNodesAsObj[t],
						Triple{Subj: subj, Pred: pred, Obj: obj})
					continue
				}
				if obj != other.nodes[subj][pred][i] {
					return false
				}
			}
		}
	}

	// collect bnodes from other graph
	for subj, po := range other.nodes {
		switch t := subj.(type) {
		case BlankNode:
			bBNodes = append(bBNodes, t)
			continue
		}
		for pred, objs := range po {
			for _, obj := range objs {
				switch t := obj.(type) {
				case BlankNode:
					bBNodesAsObj[t] = append(bBNodesAsObj[t],
						Triple{Subj: subj, Pred: pred, Obj: obj})
					continue
				}
			}
		}
	}

	if len(aBNodes) != len(bBNodes) {
		return false
	}

	if len(aBNodesAsObj) != len(bBNodesAsObj) {
		return false
	}

	nAEmptyBNodes := 0
	nBEmptyBNodes := 0
	for bnode := range aBNodesAsObj {
		if _, ok := g.nodes[bnode]; !ok {
			nAEmptyBNodes++
		}
	}
	for bnode := range bBNodesAsObj {
		if _, ok := other.nodes[bnode]; !ok {
			nBEmptyBNodes++
		}
	}
	if nAEmptyBNodes != nBEmptyBNodes {
		return false
	}

	matchA := make(map[BlankNode]bool)
	matchB := make(map[BlankNode]bool)
outer:
	for _, aNode := range aBNodes {
		for _, bNode := range bBNodes {
			if len(other.nodes[bNode]) == len(g.nodes[aNode]) && len(aBNodesAsObj[aNode]) == len(bBNodesAsObj[bNode]) {
				if isMatch(g, other, aNode, bNode, aBNodesAsObj[aNode], bBNodesAsObj[bNode]) {
					matchA[aNode] = true
					matchB[bNode] = true
					continue outer
				}
			}
			// no match was found for aNode in range bBNodes loop
			return false
		}
	}

	return len(matchA) == len(aBNodes)
}

func isMatch(a, b *Graph, aNode, bNode BlankNode, aAsObj, bAsObj []Triple) bool {
	for pred, objs := range a.nodes[aNode] {
		if _, ok := b.nodes[bNode][pred]; !ok {
			return false
		}
		if len(objs) != len(b.nodes[bNode][pred]) {
			return false
		}
		// objs are assumed to allready be sorted
		for i, obj := range objs {
			switch obj.(type) {
			case BlankNode:
				panic("TODO handle blank node pointing to another blank node")
			}
			if !obj.Eq(b.nodes[bNode][pred][i]) {
				return false
			}
		}
	}
	return true
}

//func (g *Graph) Union(other *Graph) *Graph {}
//func (g *Graph) Delete(tr ...Triple) {}
//func (g *Graph) Merge(other *Graph)  {}
