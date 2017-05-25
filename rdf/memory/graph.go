package memory

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/knakk/kbp/rdf"
)

// Graph is an in-memory implementation of rdf.Graph.
//
// Graph is not thread-safe, so concurrent writes/reads must be
// synchronized.
type Graph struct {
	// mappings between node and integer id
	node2id map[rdf.Node]int
	id2node map[int]rdf.Node

	// indexes
	spo index
	osp index
	pos index
}

type index map[int]map[int][]int

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		node2id: make(map[rdf.Node]int),
		id2node: make(map[int]rdf.Node),
		spo:     make(index),
		osp:     make(index),
		pos:     make(index),
	}
}

// Stats returns statistics about the graph.
func (g *Graph) Stats() (rdf.Stats, error) {
	// Count node types
	var bnode, literal, named int
	for node := range g.node2id {
		switch node.(type) {
		case rdf.NamedNode:
			named++
		case rdf.BlankNode:
			bnode++
		case rdf.Literal:
			literal++
		}
	}

	var triples int
	predicates := make(map[string]int)
	for _, po := range g.spo {
		for p, objs := range po {
			pred := g.id2node[p].(rdf.NamedNode).Name()
			triples += len(objs)
			predicates[pred] = predicates[pred] + len(objs)
		}
	}

	return rdf.Stats{
		NumNodes:      len(g.node2id),
		NumSubjects:   len(g.spo),
		NumPredicates: len(g.pos),
		NumObjects:    len(g.osp),
		NumBlankNodes: bnode,
		NumLiterals:   literal,
		NumNamedNodes: named,
		NumTriples:    triples,
		Predicates:    predicates,
	}, nil
}

// Triples returns all the triples in the Graph.
func (g *Graph) Triples() []rdf.Triple {
	var res []rdf.Triple
	for s, po := range g.spo {
		for p, objs := range po {
			for _, o := range objs {
				res = append(res, rdf.Triple{
					Subject:   g.id2node[s].(rdf.SubjectNode),
					Predicate: g.id2node[p].(rdf.NamedNode),
					Object:    g.id2node[o],
				})
			}
		}
	}
	return res
}

// Decode decodes values from the graph into a struct, starting at the given node
// and traversing the graph following the path specified in "rdf" struct tags.
// Each path fragment (separated by ;) represents a predicate and a direction:
//
//   -> or >> outgoing direction from node
//   <- or << incomming direction from node
//
// The double arrow variants signify that you want to collect all matches, as
// opposed to just one match otherwise.
//
// In addition, a single 'id' as last fragment in a chain states that you want
// to collect the URI of the resource.
func (g *Graph) Decode(v interface{}, startNode, base rdf.NamedNode) error {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Ptr {
		fmt.Errorf("Graph.Decode of non-pointer %s", reflect.TypeOf(v))
	}
	if rv.IsNil() {
		fmt.Errorf("Graph.Decode of nil %s", reflect.TypeOf(v))
	}

	nodeID, ok := g.node2id[startNode]
	if !ok {
		return nil
	}

	return g.decodeStruct(rv, nodeID, base.Name())
}

// traverse traverses the graph starting at the given node, using the predicates and directions encoded
// in paths. Collect and return the nodes at the end of the traversal. If path includes '>>' or '<<'
// return all mathcing nodes, otherwise at most 1 node.
func (g *Graph) traverse(startID int, base string, paths []string) (res []int, err error) {
	ids := []int{startID}
	for len(paths) > 0 {
		pred := paths[0]
		paths = paths[1:]
		var tempNodes []int
		for _, nID := range ids {

			if pred == "id" {
				return []int{nID}, nil
			}

			if len(pred) < 3 {
				return nil, errors.New("rdf path fragment too short")
			}
			pID, ok := g.node2id[rdf.NewNamedNode(base+pred[2:])]
			if !ok {
				return
			}

			var (
				found = false
				nodes []int
			)
			switch pred[:2] {
			case "->":
				nodes, found = g.spo[nID][pID]
				if found {
					nodes = nodes[:1] // keep only one match
				}
			case "<-":
				nodes, found = g.pos[pID][nID]
				if found {
					nodes = nodes[:1] // keep only one match
				}
			case ">>":
				nodes, found = g.spo[nID][pID]
			case "<<":
				nodes, found = g.pos[pID][nID]
			default:
				return nil, fmt.Errorf("unknown path fragment prefix: %q", pred[:2])
			}

			if !found {
				continue
			}

			if len(paths) == 0 {
				// traversal finished, collect nodes
				res = append(res, nodes...)

			}

			tempNodes = append(tempNodes, nodes...)
		}
		ids = tempNodes
	}

	return res, nil
}

func (g *Graph) decodeStruct(rv reflect.Value, nodeID int, base string) error {
	s := rv.Elem()
	t := s.Type()

structFields:
	for i := 0; i < s.NumField(); i++ {
		tag := t.Field(i).Tag.Get("rdf")
		if tag == "" {
			continue
		}
		predicates := strings.Split(tag, ";")
		nodes, err := g.traverse(nodeID, base, predicates)
		if err != nil {
			return err
		}
		if len(nodes) == 0 {
			continue structFields
		}

		switch s.Field(i).Kind() {
		case reflect.Slice:
			if err := g.decodeSlice(s.Field(i), nodes, base); err != nil {
				return err
			}
		case reflect.Struct:
			if err := g.decodeStruct(s.Field(i).Addr(), nodes[0], base); err != nil {
				return err
			}
		default:
			if err := g.decodePrimitive(s.Field(i), g.id2node[nodes[0]]); err != nil {
				return err
			}
		}

	}
	return nil
}

func (g *Graph) decodePrimitive(v reflect.Value, node rdf.Node) error {
	switch v.Kind() {
	case reflect.String:
		switch n := node.(type) {
		case rdf.Literal:
			v.SetString(n.ValueAsString())
		case rdf.NamedNode:
			v.SetString(n.Name())
		default:
			return errors.New("only named node or literal can be set to struct.string field")
		}
	case reflect.Int:
		lit, ok := node.(rdf.Literal)
		if !ok {
			return errors.New("only literal can be sto struct.int field")
		}
		litInt, ok := lit.Value().(int)
		if !ok {
			return fmt.Errorf("struct field is int, but RDF Literal incompatible: %v", lit.DataType())
		}
		v.SetInt(int64(litInt))
	default:
		panic("decodePrimitive TODO")
	}
	return nil
}

func (g *Graph) decodeSlice(rv reflect.Value, nodes []int, base string) error {
	slice := reflect.MakeSlice(rv.Type(), 0, len(nodes))
	switch slice.Type().Elem().Kind() {
	case reflect.String:
		vals := make([]string, 0, len(nodes))
		for _, id := range nodes {
			switch node := g.id2node[id].(type) {
			case rdf.Literal:
				vals = append(vals, node.ValueAsString())
			case rdf.NamedNode:
				vals = append(vals, node.Name())
			default:
				return fmt.Errorf("only named node or literal can be set to struct.[]string field, got %T %v", node, node)
			}
		}
		rv.Set(reflect.AppendSlice(slice, reflect.ValueOf(vals)))
	case reflect.Struct:
		for _, id := range nodes {
			elem := reflect.New(slice.Type().Elem())
			if err := g.decodeStruct(elem, id, base); err != nil {
				return err
			}
			slice = reflect.Append(slice, reflect.Indirect(elem))
		}
		rv.Set(slice)
	default:
		panic("decodeSlice TODO")
	}
	return nil
}

func (g *Graph) Describe(mode rdf.DescribeMode, nodes ...rdf.NamedNode) (rdf.Graph, error) {
	res := NewGraph()
	described := make(map[int]struct{})
	for _, node := range nodes {
		g.describe(mode, node, res, described)
	}
	return res, nil
}

func (g *Graph) describe(mode rdf.DescribeMode, node rdf.Node, res *Graph, described map[int]struct{}) {
	var trs []rdf.Triple
	if s, found := g.node2id[node]; found {
		if _, ok := described[s]; ok {
			return
		}
		for p, objs := range g.spo[s] {
			for _, o := range objs {
				trs = append(trs, rdf.Triple{
					Subject:   node.(rdf.SubjectNode),
					Predicate: g.id2node[p].(rdf.NamedNode),
					Object:    g.id2node[o],
				})
				switch mode {
				case rdf.DescForward:
					if _, isBnode := g.id2node[o].(rdf.BlankNode); isBnode {
						g.describe(mode, g.id2node[o], res, described)
					}
				case rdf.DescForwardRecursive, rdf.DescSymmetricRecursive:
					if _, isLiteral := g.id2node[o].(rdf.Literal); !isLiteral {
						// Node must be Blank Node or Named Node
						g.describe(rdf.DescForwardRecursive, g.id2node[o], res, described)
					}
				}
			}
		}
		res.Insert(trs...)
		described[s] = struct{}{}

		if mode == rdf.DescSymmetricRecursive {
			for sub, _ := range g.osp[s] {
				g.describe(mode, g.id2node[sub], res, described)
			}
		}
	}
}

func (g *Graph) describeW(enc *rdf.Encoder, mode rdf.DescribeMode, node rdf.Node, described map[int]struct{}) error {
	if s, found := g.node2id[node]; found {
		if _, ok := described[s]; ok {
			return nil
		}
		for p, objs := range g.spo[s] {
			for _, o := range objs {
				if err := enc.Encode(rdf.Triple{
					Subject:   node.(rdf.SubjectNode),
					Predicate: g.id2node[p].(rdf.NamedNode),
					Object:    g.id2node[o],
				}); err != nil {
					return err
				}
				switch mode {
				case rdf.DescForward:
					if _, isBnode := g.id2node[o].(rdf.BlankNode); isBnode {
						g.describeW(enc, mode, g.id2node[o], described)
					}
				case rdf.DescForwardRecursive, rdf.DescSymmetricRecursive:
					if _, isLiteral := g.id2node[o].(rdf.Literal); !isLiteral {
						// Node must be Blank Node or Named Node
						if err := g.describeW(enc, rdf.DescForwardRecursive, g.id2node[o], described); err != nil {
							return err
						}
					}
				}
			}
		}
		described[s] = struct{}{}
		if mode == rdf.DescSymmetricRecursive {
			for sub, _ := range g.osp[s] {
				if err := g.describeW(enc, mode, g.id2node[sub], described); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (g *Graph) DescribeW(enc *rdf.Encoder, mode rdf.DescribeMode, nodes ...rdf.NamedNode) error {
	described := make(map[int]struct{})
	for _, node := range nodes {
		if err := g.describeW(enc, mode, node, described); err != nil {
			return err
		}
	}
	return nil
}

func (g *Graph) Update(del []rdf.TriplePattern, ins []rdf.TriplePattern, where []rdf.TriplePattern) (delN, insN int, err error) {
	if where != nil {
		// Encode patterns
		encPatterns := make([]encPattern, len(where))
		vars := make(map[rdf.Variable]int)
		for i, pattern := range where {
			encPatterns[i] = g.encodePattern(pattern, vars)
			//fmt.Printf("\n%v => %v\n", pattern, encPatterns[i][:3])
		}

		// Evaluate query
		_, solutions := g.where(encPatterns)

		// Substitute variables with solutions
		if len(solutions.Rows) > 0 {
			for _, pattern := range ins {
				n := g.insertFor(pattern, solutions, vars)
				insN += n
			}

			for _, pattern := range del {
				n := g.deleteFor(pattern, solutions, vars)
				delN += n
			}
		}

		return delN, insN, nil
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

	return delN, insN, err
}

// Insert adds one or more triples to the Graph. It returns the number
// of triples inserted which where not already present.
//
// Blank nodes are assumed to be disjoint from the blank nodes already
// present in the graph, and will be inserted with "fresh" node IDs.
// However, blank nodes in with identical IDs will be given the same, new ID.
func (g *Graph) Insert(trs ...rdf.Triple) (int, error) {
	n := 0
	bnodes := make(map[rdf.BlankNode]int)
	for _, tr := range trs {

		var sid, pid, oid int

		// get subject id
		if bnode, ok := tr.Subject.(rdf.BlankNode); ok {
			if _, ok := bnodes[bnode]; !ok {
				bnodes[bnode] = g.addNode(bnode)
			}
			sid = bnodes[bnode]
		} else {
			sid = g.addNode(tr.Subject)
		}

		// get predicate id
		pid = g.addNode(tr.Predicate)

		// get object id
		if bnode, ok := tr.Object.(rdf.BlankNode); ok {
			if _, ok := bnodes[bnode]; !ok {
				bnodes[bnode] = g.addNode(bnode)
			}
			oid = bnodes[bnode]
		} else {
			oid = g.addNode(tr.Object)
		}

		if g.index(sid, pid, oid) {
			n++
		}
	}
	return n, nil
}

func (g *Graph) deleteFor(p rdf.TriplePattern, s encSolutions, vars map[rdf.Variable]int) int {
	var encP [3]int
	for i, node := range []rdf.TriplePatternNode{p.Subject, p.Predicate, p.Object} {
		if v, ok := node.(rdf.Variable); ok {
			encP[i] = vars[v]
			if vars[v] == 0 {
				panic("BUG: deleteFor got variable not in cache")
			}
			continue
		}

		id, ok := g.node2id[node.(rdf.Node)]
		if !ok {
			// Node is not stored, so we have no triple to delete
			return 0
		}
		encP[i] = id
	}

	if p.IsConcrete() {
		if stored := g.unindex(encP[0], encP[1], encP[2]); !stored {
			return 0
		}
		return 1
	}

	var n int
	tr := [3]int{encP[0], encP[1], encP[2]}
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
		stored := g.unindex(tr[0], tr[1], tr[2])
		if stored {
			n++
		}
	}

	return n
}

func (g *Graph) insertFor(p rdf.TriplePattern, s encSolutions, vars map[rdf.Variable]int) int {
	var encP [3]int
	for i, node := range []rdf.TriplePatternNode{p.Subject, p.Predicate, p.Object} {
		if v, ok := node.(rdf.Variable); ok {
			encP[i] = vars[v]
			if vars[v] == 0 {
				panic("BUG: insertFor got variable not in cache")
			}
			continue
		}
		id := g.addNode(node.(rdf.Node))
		encP[i] = id
	}

	if p.IsConcrete() {
		stored := g.index(encP[0], encP[1], encP[2])
		if !stored {
			return 0
		}
		return 1
	}

	var n int
	tr := [3]int{encP[0], encP[1], encP[2]}
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
		stored := g.index(tr[0], tr[1], tr[2])
		if stored {
			n++
		}
	}

	return n
}

// addNode adds a node to the node dictonaries, and returns its new ID, or
// existing ID if already present.
func (g *Graph) addNode(node rdf.Node) int {
	if id, ok := g.node2id[node]; ok {
		return id
	}
	id := len(g.node2id) + 1
	g.node2id[node] = id
	g.id2node[id] = node
	return id
}

func (g *Graph) removeNode(id int) bool {
	node, ok := g.id2node[id]
	if !ok {
		return false
	}
	delete(g.node2id, node)
	delete(g.id2node, id)
	return true
}

// index indexes a triple in all three indexes. It returns false if the
// triple was already stored.
func (g *Graph) index(sid, pid, oid int) bool {
	// Add to spo index
	if _, ok := g.spo[sid]; !ok {
		// Subject is a new node.
		g.spo[sid] = make(map[int][]int)
	}
	if _, ok := g.spo[sid][pid]; !ok {
		// New predicate for subject.
		g.spo[sid][pid] = []int{oid}
	} else {
		for _, o := range g.spo[sid][pid] {
			if o == oid {
				return false
			}
		}
		// Triple not in graph, add it
		g.spo[sid][pid] = append(g.spo[sid][pid], oid)
	}

	// Add to osp index
	if _, ok := g.osp[oid]; !ok {
		g.osp[oid] = make(map[int][]int)
	}
	if _, ok := g.osp[oid][sid]; !ok {
		g.osp[oid][sid] = []int{pid}
	} else {
		g.osp[oid][sid] = append(g.osp[oid][sid], pid)
	}

	// Add to pos index
	if _, ok := g.pos[pid]; !ok {
		g.pos[pid] = make(map[int][]int)
	}
	if _, ok := g.pos[pid][oid]; !ok {
		g.pos[pid][oid] = []int{sid}
	} else {
		g.pos[pid][oid] = append(g.pos[pid][oid], sid)
	}

	return true
}

// unindex removes a triple from all three indexes. It returns false if the
// triple was not stored.
func (g *Graph) unindex(sid, pid, oid int) bool {
	// Remove from spo index
	if _, ok := g.spo[sid]; !ok {
		return false
	}
	if _, ok := g.spo[sid][pid]; !ok {
		return false
	}
	found := false
	for i, o := range g.spo[sid][pid] {
		if o == oid {
			found = true
			if len(g.spo[sid][pid]) == 1 {
				// It's the only object left for given SP
				delete(g.spo[sid], pid)
				break
			}
			l := len(g.spo[sid][pid]) - 1
			g.spo[sid][pid][i] = g.spo[sid][pid][l]
			g.spo[sid][pid] = g.spo[sid][pid][:l]
			break
		}
	}
	if !found {
		return false
	}

	// Remove from osp index
	for i, p := range g.osp[oid][sid] {
		if p == pid {
			if len(g.osp[oid][sid]) == 1 {
				// It's the only predicate left for given OS
				delete(g.osp[oid], sid)
				break
			}
			l := len(g.osp[oid][sid]) - 1
			g.osp[oid][sid][i] = g.osp[oid][sid][l]
			g.osp[oid][sid] = g.osp[oid][sid][:l]
			break
		}
	}

	// Remove from pos index
	for i, s := range g.pos[pid][oid] {
		if s == sid {
			if len(g.pos[pid][oid]) == 1 {
				// It's the last subject left for given PO
				delete(g.pos[pid], oid)
				break
			}
			l := len(g.pos[pid][oid]) - 1
			g.pos[pid][oid][i] = g.pos[pid][oid][l]
			g.pos[pid][oid] = g.pos[pid][oid][:l]
			break
		}
	}

	g.removeOrphanNodes(sid, pid, oid)

	return true
}

func (g *Graph) removeOrphanNodes(sid, pid, oid int) {
	if len(g.spo[sid]) == 0 && len(g.pos[sid]) == 0 && len(g.osp[sid]) == 0 {
		g.removeNode(sid)
	}
	if len(g.pos[pid]) == 0 && len(g.spo[pid]) == 0 && len(g.osp[pid]) == 0 {
		g.removeNode(pid)
	}
	if len(g.osp[oid]) == 0 && len(g.spo[oid]) == 0 && len(g.pos[oid]) == 0 {
		g.removeNode(oid)
	}
}

func (g *Graph) ids(tr rdf.Triple) (sid, pid, oid int, found bool) {
	sid, found = g.node2id[tr.Subject]
	if !found {
		return
	}
	pid, found = g.node2id[tr.Predicate]
	if !found {
		return
	}
	oid, found = g.node2id[tr.Object]
	return
}

// Delete removes one or more triples to the Graph. It returns the number
// of triples deleted. The delete operation only supports deleting triples with
// concrete data, that means, without blank nodes; use the DeleteWhere method for that.
func (g *Graph) Delete(trs ...rdf.Triple) (int, error) {
	n := 0

	for _, tr := range trs {
		// Skip triples with blank nodes.
		if _, ok := tr.Subject.(rdf.BlankNode); ok {
			continue
		}
		if _, ok := tr.Object.(rdf.BlankNode); ok {
			continue
		}

		sid, pid, oid, found := g.ids(tr)
		if !found {
			continue
		}

		if g.unindex(sid, pid, oid) {
			n++
		}

	}
	return n, nil
}

func (g *Graph) estCardinality(ep encPattern) int {
	if ep.s() == 0 || ep.p() == 0 || ep.o() == 0 {
		// triple is not in graph
		return 0 // exact
	}
	if ep.s() > 0 && ep.p() > 0 && ep.o() > 0 {
		// {n,n,n}
		return 1 // exact
	}
	if ep.s() < 0 && ep.p() < 0 && ep.o() < 0 {
		// {v,v,v}
		return len(g.spo) * len(g.osp) // estimate
	}
	if ep.s() < 0 && ep.p() > 0 && ep.o() > 0 {
		// {v,n,n}
		return len(g.pos[ep.p()][ep.o()]) // exact
	}
	if ep.s() < 0 && ep.p() < 0 && ep.o() > 0 {
		// {v,v,n}
		for _, preds := range g.osp[ep.o()] {
			return len(g.osp[ep.o()]) * len(preds) // estimate
		}
	}
	if ep.s() > 0 && ep.p() > 0 && ep.o() < 0 {
		// {n,n,v}
		return len(g.spo[ep.s()][ep.p()]) // exact
	}
	if ep.s() > 0 && ep.p() < 0 && ep.o() > 0 {
		// {n,v,n}
		return len(g.osp[ep.o()][ep.s()]) // exact
	}
	if ep.s() > 0 && ep.p() < 0 && ep.o() < 0 {
		// {n,v,v}
		for _, objs := range g.spo[ep.s()] {
			return len(g.spo[ep.s()]) * len(objs) // estimate
		}
	}
	if ep.s() < 0 && ep.p() > 0 && ep.o() < 0 {
		// {v,n,v}
		for _, subjs := range g.pos[ep.p()] {
			return len(g.pos[ep.p()]) * len(subjs) // estimate
		}
	}
	panic(fmt.Sprintf("BUG: Graph.estCardinality: unhandled pattern: %v", ep))
}

func (g *Graph) encodePattern(p rdf.TriplePattern, vars map[rdf.Variable]int) (ep encPattern) {
	// positive integer = nodeId (rdf.NamedNode, Literal)
	// 0 = node missing
	// negative integer = variable id

	for i, node := range []rdf.TriplePatternNode{p.Subject, p.Predicate, p.Object} {
		if v, ok := node.(rdf.Variable); ok {
			if vid, ok := vars[v]; ok {
				ep[i] = vid
			} else {
				vars[v] = (len(vars) + 1) * -1
				ep[i] = vars[v]
			}
		} else {
			if id, ok := g.node2id[node.(rdf.Node)]; ok {
				ep[i] = id
			} else {
				// node is missing
				ep[i] = 0
			}
		}
	}

	ep[3] = g.estCardinality(ep)

	return ep
}

// encPattern is a rdf.TriplePattern encoded for query processing, where index
// 0-2 represents s,p,o nodes and the last represents the estimated cardinality.
// If a node equals=0, it means it is not stored in graph, meaning that the
// pattern (and all patterns sharing variables with it) will not match any.
// If a nodes < 0, it's a variable ID.
type encPattern [4]int

func (e encPattern) s() int              { return e[0] }
func (e encPattern) p() int              { return e[1] }
func (e encPattern) o() int              { return e[2] }
func (e encPattern) estCardinality() int { return e[3] }

type encSolutions struct {
	//SortedBy []int // vars
	Vars []int
	Rows [][]int
}

func compatible(pairs [][2]int, rowa, rowb []int) bool {
	for _, p := range pairs {
		if rowa[p[0]] != rowb[p[1]] {
			return false
		}
	}
	return true
}

func (s encSolutions) join(other encSolutions, bound map[int][]int) encSolutions {
	if len(s.Vars) == 0 {
		// We're joining with the empty solution set
		return other
	}

	if len(other.Vars) == 0 {
		return s
	}

	var out [][]int
	var pairs [][2]int // the indices of shared variables in a & b
	for i, va := range s.Vars {
		for j, vb := range other.Vars {
			if va == vb {
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}
	var addVars []int
	for i, v := range other.Vars {
		if !includeNode(s.Vars, v) {
			s.Vars = append(s.Vars, v)
			addVars = append(addVars, i)
		}
	}

	for _, a := range s.Rows {
		for _, b := range other.Rows {
			if compatible(pairs, a, b) {
				row := make([]int, 0, len(a)+len(addVars))
				row = append(row, a...)
				for _, v := range addVars {
					row = append(row, b[v])
				}
				out = append(out, row)
			} else {
				// TODO remove from bound
			}
		}
	}

	s.Rows = out
	return s
}

func (s encSolutions) project(vars []int) encSolutions {
	rows := make([][]int, 0, len(s.Rows))
	var svars []int
	for _, v := range vars {
		for i, vs := range s.Vars {
			if v == vs {
				svars = append(svars, i)
			}
		}
	}
	for _, row := range s.Rows {
		newrow := make([]int, 0, len(svars))
		for _, i := range svars {
			newrow = append(newrow, row[i])
		}
		rows = append(rows, newrow)
	}
	return encSolutions{
		Vars: vars,
		Rows: rows,
	}
}

func (g *Graph) evalBGP(p encPattern, bound map[int][]int) encSolutions {
	if p[0] > 0 && p[1] > 0 && p[2] > 0 {
		// It's a concrete triple
		return encSolutions{}
	}
	if p[2] > 0 {
		// Matching {any,any,n} using OSP index
		return g.scanOSP(p, bound)
	}
	if p[1] > 0 {
		// Matching {?,n,n} using POS index
		return g.scanPOS(p, bound)
	}
	// Matching {any,any,any} using SPO index
	return g.scanSPO(p, bound)
}

func (g *Graph) scanSPO(p encPattern, bound map[int][]int) encSolutions {
	// {any,any,any}
	var (
		solutions  encSolutions
		subjects   []int
		predicates []int
		objects    []int
	)

	if p[0] < 0 {
		solutions.Vars = []int{p[0]}
	}
	if p[1] < 0 {
		solutions.Vars = append(solutions.Vars, p[1])
	}
	if p[2] < 0 {
		solutions.Vars = append(solutions.Vars, p[2])
	}

	if subjs, ok := bound[p[0]]; ok {
		subjects = subjs
	} else {
		subjects = g.subjects()
	}

	for _, sid := range subjects {
		if p[0] > 0 && p[0] != sid {
			continue
		}

		if preds, ok := bound[p[1]]; ok {
			predicates = preds
		} else {
			predicates = g.predicatesForSubj(sid)
		}

		for _, pid := range predicates {
			if p[1] > 0 && p[1] != pid {
				continue
			}

			if objs, ok := bound[p[2]]; ok {
				objects = objs
			} else {
				objects = g.spo[sid][pid]
			}
			for _, oid := range objects {
				if !includeNode(g.spo[sid][pid], oid) {
					continue
				}
				if p[2] > 0 && p[2] != oid {
					continue
				}

				row := make([]int, 0, len(solutions.Vars))
				if p[0] < 0 {
					row = append(row, sid)
				}
				if p[1] < 0 {
					row = append(row, pid)
				}
				if p[2] < 0 {
					row = append(row, oid)
				}
				solutions.Rows = append(solutions.Rows, row)
			}
		}
	}

	updateBound(solutions, bound)
	return solutions
}

func (g *Graph) scanOSP(p encPattern, bound map[int][]int) encSolutions {
	// {any,any,literal/uri}

	var solutions encSolutions
	if p[0] < 0 {
		solutions.Vars = []int{p[0]}
	}
	if p[1] < 0 {
		solutions.Vars = append(solutions.Vars, p[1])
	}
	// p[2] is concrete

	for sid, preds := range g.osp[p[2]] {
		if p[0] < 0 {
			if !freeOrBound(p[0], sid, bound) {
				continue
			}
		} else if p[0] != sid {
			continue
		}
		for _, pid := range preds {
			if p[1] < 0 {
				if !freeOrBound(p[1], pid, bound) {
					continue
				}
			} else if p[1] != pid {
				continue
			}
			row := make([]int, 0, len(solutions.Vars))
			if p[0] < 0 {
				row = append(row, sid)
			}
			if p[1] < 0 {
				row = append(row, pid)
			}
			solutions.Rows = append(solutions.Rows, row)
		}
	}

	updateBound(solutions, bound)
	return solutions
}

func (g *Graph) scanPOS(p encPattern, bound map[int][]int) encSolutions {
	// {var,uri,var}

	var solutions encSolutions
	if p[0] < 0 {
		solutions.Vars = []int{p[0]}
	}
	// p[1] is concrete
	if p[2] < 0 {
		solutions.Vars = append(solutions.Vars, p[2])
	}

	for oid, subjs := range g.pos[p[1]] {
		if p[2] < 0 {
			if !freeOrBound(p[2], oid, bound) {
				continue
			}
		} else if p[2] != oid {
			continue
		}
		for _, sid := range subjs {
			if p[0] < 0 {
				if !freeOrBound(p[0], sid, bound) {
					continue
				}
			} else if p[0] != sid {
				continue
			}

			row := make([]int, 0, len(solutions.Vars))
			if p[0] < 0 {
				row = append(row, sid)
			}
			if p[2] < 0 {
				row = append(row, oid)
			}
			solutions.Rows = append(solutions.Rows, row)
		}
	}

	updateBound(solutions, bound)
	return solutions
}

func (g *Graph) substitute(p encPattern, s encSolutions) (res []rdf.Triple) {
	for _, row := range s.Rows {
		var tr rdf.Triple
		for i, v := range s.Vars {
			if v == p.s() {
				tr.Subject = g.id2node[row[i]].(rdf.SubjectNode)
			} else if p.s() > 0 {
				tr.Subject = g.id2node[p.s()].(rdf.SubjectNode)
			}

			if v == p.p() {
				tr.Predicate = g.id2node[row[i]].(rdf.NamedNode)
			} else if p.p() > 0 {
				tr.Predicate = g.id2node[p.p()].(rdf.NamedNode)
			}

			if v == p.o() {
				tr.Object = g.id2node[row[i]]
			} else if p.o() > 0 {
				tr.Object = g.id2node[p.o()]
			}
		}

		res = append(res, tr)
	}
	return res
}

// Where returns a new graph with the triples matching the given patterns.
// It corresponds to a SPARQL CONSTRUCT WHERE query, i.e where both template
// and patterns are identical.
func (g *Graph) Where(patterns ...rdf.TriplePattern) (rdf.Graph, error) {

	// Fast path for no patterns
	if len(patterns) == 0 {
		return NewGraph(), nil
	}

	// Encode patterns
	encPatterns := make([]encPattern, len(patterns))
	vars := make(map[rdf.Variable]int)
	for i, pattern := range patterns {
		encPatterns[i] = g.encodePattern(pattern, vars)
	}

	groups, solutions := g.where(encPatterns)

	res := NewGraph()
	if len(solutions.Rows) > 0 {
		for _, group := range groups {
			for _, pattern := range group {
				triples := g.substitute(pattern, solutions)
				res.Insert(triples...)
			}
		}
	}

	return res, nil
}

func (g *Graph) where(patterns []encPattern) ([][]encPattern, encSolutions) {

	// Group disjoint pattern groups and sort them by estimated cardinality
	groups := groupPatternsByVariable(patterns)
	for n := range groups {
		sort.Slice(groups[n], func(i, j int) bool {
			return groups[n][i].estCardinality() < groups[n][j].estCardinality()
		})
	}

	bound := make(map[int][]int)
	var left, right encSolutions
	// Evaluate each BGP sequentially, each BGP using any bound values from previous BGPs
	for _, group := range groups {
		for _, pattern := range group {
			right = g.evalBGP(pattern, bound)
			left = left.join(right, bound)
		}
	}

	return groups, left
}

func (g *Graph) Select(vars []rdf.Variable, patterns ...rdf.TriplePattern) ([][]rdf.Node, error) {

	var res [][]rdf.Node

	// Fast path for no patterns
	if len(patterns) == 0 {
		return res, nil
	}

	// Encode patterns
	encPatterns := make([]encPattern, len(patterns))
	varstmp := make(map[rdf.Variable]int)
	for i, pattern := range patterns {
		encPatterns[i] = g.encodePattern(pattern, varstmp)
	}
	// Encode variables using the mappings from above
	encVars := make([]int, len(vars))
	for i, v := range vars {
		if ev, ok := varstmp[v]; ok {
			encVars[i] = ev
		} // else: if variable does not occur in patterns, we ignore it.
	}

	_, solutions := g.where(encPatterns)

	if len(solutions.Rows) == 0 {
		return res, nil
	}

	solutions = solutions.project(encVars)

	for _, row := range solutions.Rows {
		nodes := make([]rdf.Node, 0, len(encVars))
		for _, id := range row {
			nodes = append(nodes, g.id2node[id])
		}

		res = append(res, nodes)
	}

	return res, nil
}

func indexVariables(p encPattern, n int, v2g map[int]int) {
	for _, nid := range p[:3] {
		if nid < 0 {
			v2g[nid] = n
		}
	}
}

// groupPatternsByVariable groups encoded patterns by shared variables, either direct or transitive.
// Disjoint groups can be processed separately (eg. in paralell) and its solutions merged by union.
// It also prunes any patterns with missing nodes, and any other patterns sharing variable with those,
// as they will not match any triples and should be discarded as early as possible.
func groupPatternsByVariable(patterns []encPattern) [][]encPattern {
	groups := make([][]encPattern, 0, len(patterns))
	if len(patterns) == 0 {
		return groups
	}

	// variable to group mapping
	v2g := make(map[int]int)
	// checked all patterns for variable
	explored := make(map[int]bool)
	// variables shared with pattern with missing node, should be ignore
	ignore := make(map[int]bool)

	for i := 0; i < len(patterns); i++ {
		if patterns[i].estCardinality() == 0 {
			for _, node := range patterns[i][:3] {
				if node < 0 {
					ignore[node] = true
				}
			}
			patterns = append(patterns[:i], patterns[i+1:]...)
			i--
		}
	}

	for k := range ignore {
		for i := 0; i < len(patterns); i++ {
			if patterns[i].s() == k || patterns[i].p() == k || patterns[i].o() == k {
				patterns = append(patterns[:i], patterns[i+1:]...)
				i--
			}
		}
	}
	// TODO if len(ignore) > 0 loop over patterns and remove those which should be pruned

	for len(patterns) > 0 {
		// assign first pattern to first group
		groups = append(groups, []encPattern{patterns[0]})
		indexVariables(patterns[0], len(groups)-1, v2g)
		patterns = patterns[1:]
		if len(v2g) > 0 {
			// we have a variable
			break
		}
	}

again:
	for v, n := range v2g {
		if explored[v] {
			continue
		}
		for i := 0; i < len(patterns); i++ {
			for _, node := range patterns[i][:3] {
				if node > 0 {
					continue // not a variable
				}
				if node == v {
					groups[n] = append(groups[n], patterns[i])
					indexVariables(patterns[i], n, v2g)
					patterns = append(patterns[:i], patterns[i+1:]...)
					i--
					break
				}
			}
		}
		explored[v] = true
		if len(patterns) > 0 {
			goto again
		}
	}
	for i := 0; i < len(patterns); i++ {
		for _, node := range patterns[i][:3] {
			if node > 0 {
				continue // not a variable
			}
			if n, ok := v2g[node]; ok {
				groups[n] = append(groups[n], patterns[i])
				indexVariables(patterns[i], n, v2g)
				patterns = append(patterns[:i], patterns[i+1:]...)
				i--
				break
			}
		}
	}
	for len(patterns) > 0 {
		// assign first pattern to new group
		groups = append(groups, []encPattern{patterns[0]})
		indexVariables(patterns[0], len(groups)-1, v2g)
		patterns = patterns[1:]
		if len(v2g) > 0 {
			// we have a variable
			break
		}
	}
	if len(patterns) > 0 {
		goto again
	}

	return groups
}

func (g *Graph) subjects() (res []int) {
	for sid := range g.spo {
		res = append(res, sid)
	}
	return res
}

func (g *Graph) predicatesForSubj(sid int) (res []int) {
	for pid := range g.spo[sid] {
		res = append(res, pid)
	}
	return res
}

func updateBound(s encSolutions, bound map[int][]int) {
	for _, v := range s.Vars {
		if _, ok := bound[v]; !ok {
			bound[v] = []int{}
		}
	}
	for _, row := range s.Rows {
		for i, v := range s.Vars {
			if !includeNode(bound[v], row[i]) {
				bound[v] = append(bound[v], row[i])
			}
		}
	}
}

func includeNode(nodes []int, find int) bool {
	for _, node := range nodes {
		if node == find {
			return true
		}
	}
	return false
}

// freeOrBound returns true if there are no bound nodes for the given variable,
// or if the given id is included in the already bound nodes.
// TODO find a better name
func freeOrBound(v int, id int, bound map[int][]int) bool {
	if vs, ok := bound[v]; ok {
		return includeNode(vs, id)
	}
	return true
}

// Eq checks if two graphs are equal (isomorphic).
func (g *Graph) Eq(other *Graph) bool {
	if len(g.node2id) != len(other.node2id) {
		return false
	}

	set := make(map[rdf.BlankNode]bool)

	// Check if all triples without blank-nodes are present in other.
	// Collect all blank nodes while iterating the graph.
	for s, po := range g.spo {
		if bnode, ok := g.id2node[s].(rdf.BlankNode); ok {
			set[bnode] = true
			continue
		}
		others, ok := other.node2id[g.id2node[s]]
		if !ok {
			return false
		}
		for p, objs := range po {
			otherp, ok := other.node2id[g.id2node[p]]
			if !ok {
				return false
			}
			if _, ok := other.spo[others][otherp]; !ok {
				return false
			}
			for _, o := range objs {
				if bnode, ok := g.id2node[o].(rdf.BlankNode); ok {
					set[bnode] = true
					continue
				}
				othero, ok := other.node2id[g.id2node[o]]
				if !ok {
					return false
				}
				if !includeNode(other.spo[others][otherp], othero) {
					return false
				}
			}
		}
	}

	// Graphs excluding blank nodes are equal.

	var aBlankNodes []rdf.BlankNode
	for bnode := range set {
		aBlankNodes = append(aBlankNodes, bnode)
	}
	bBlankNodes := other.bnodes()
	if len(bBlankNodes) != len(aBlankNodes) {
		return false
	}
	var (
		aSign []string
		bSign []string
	)

	for _, bnode := range aBlankNodes {
		aSign = append(aSign, g.signature(bnode))
	}

	for _, bnode := range bBlankNodes {
		bSign = append(bSign, other.signature(bnode))
	}

	sort.Strings(aSign)
	sort.Strings(bSign)
	for i, s := range aSign {
		if s != bSign[i] {
			return false
		}
	}

	return true
}

func (g *Graph) bnodes() []rdf.BlankNode {
	set := make(map[rdf.BlankNode]bool)
	for s, po := range g.spo {
		if bnode, ok := g.id2node[s].(rdf.BlankNode); ok {
			set[bnode] = true
			continue
		}
		for _, objs := range po {
			for _, o := range objs {
				if bnode, ok := g.id2node[o].(rdf.BlankNode); ok {
					set[bnode] = true
				}
			}
		}
	}
	var res []rdf.BlankNode
	for bnode := range set {
		res = append(res, bnode)
	}
	return res
}

func (g *Graph) signature(bnode rdf.BlankNode) string {
	// TODO function shoud take nodeID int

	// We keep track of visited blank nodes, as not to trigger infinite
	// recursion if there is a circular relationship between nodes.
	//visited := make(map[BlankNode]bool)
	// TODO

	var (
		incoming []string
		outgoing []string
	)

	// incoming relations
	for s, preds := range g.osp[g.node2id[bnode]] {
		for _, p := range preds {
			outgoing = append(outgoing, g.id2node[s].String()+g.id2node[p].String())
		}
	}
	sort.Strings(incoming)

	// outgoing relations
	for p, objs := range g.spo[g.node2id[bnode]] {
		for _, o := range objs {
			outgoing = append(outgoing, g.id2node[p].String()+g.id2node[o].String())
		}
	}
	sort.Strings(outgoing)

	return strings.Join(incoming, "") + strings.Join(outgoing, "")
}

func (g *Graph) EncodeNTriples(w io.Writer) error {
	bw := bufio.NewWriter(w)
	for _, tr := range g.Triples() {
		if _, err := bw.WriteString(tr.String()); err != nil {
			return err
		}

	}
	return bw.Flush()
}
