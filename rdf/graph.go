package rdf

import (
	"fmt"
	"sort"
	"strings"
)

// Graph is an in-memory representation of an RDF graph:
// a colllection of triples which can be queried for subgraphs
// and mutated by inserting and deleting triples.
//
// Graph is not thread-safe, so concurrent writes/reads must be
// synchronized.
type Graph struct {
	// mappings between node and integer id
	node2id map[Node]int
	id2node map[int]Node

	// indexes
	spo index
	osp index
	pos index
}

type index map[int]map[int][]int

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		node2id: make(map[Node]int),
		id2node: make(map[int]Node),
		spo:     make(index),
		osp:     make(index),
		pos:     make(index),
	}
}

// Stats returns statistics about the graph.
func (g *Graph) Stats() Stats {
	// Count node types
	var bnode, literal, named int
	for node := range g.node2id {
		switch node.(type) {
		case NamedNode:
			named++
		case BlankNode:
			bnode++
		case Literal:
			literal++
		}
	}

	var triples int
	predicates := make(map[string]int)
	for _, po := range g.spo {
		for p, objs := range po {
			pred := g.id2node[p].(NamedNode).val
			triples += len(objs)
			predicates[pred] = predicates[pred] + len(objs)
		}
	}

	return Stats{
		NumNodes:      len(g.node2id),
		NumSubjects:   len(g.spo),
		NumPredicates: len(g.pos),
		NumObjects:    len(g.osp),
		NumBlankNodes: bnode,
		NumLiterals:   literal,
		NumNamedNodes: named,
		NumTriples:    triples,
		Predicates:    predicates,
	}
}

// Stats represents statistics about a graph.
type Stats struct {
	// Counts
	NumNodes      int
	NumNamedNodes int
	NumLiterals   int
	NumBlankNodes int
	NumSubjects   int
	NumPredicates int
	NumObjects    int
	NumTriples    int
	// Distributions
	Predicates map[string]int
}

// Size returns the number of triples in the Graph.
func (g *Graph) Size() int {
	return len(g.node2id)
}

// Triples returns all the triples in the Graph.
func (g *Graph) Triples() []Triple {
	var res []Triple
	for s, po := range g.spo {
		for p, objs := range po {
			for _, o := range objs {
				res = append(res, Triple{
					Subject:   g.id2node[s].(SubjectNode),
					Predicate: g.id2node[p].(NamedNode),
					Object:    g.id2node[o],
				})
			}
		}
	}
	return res
}

// Insert adds one or more triples to the Graph. It returns the number
// of triples inserted which where not already present.
//
// Blank nodes are assumed to be disjoint from the blank nodes already
// present in the graph, and will be inserted with "fresh" node IDs.
// However, any blank nodes in with identical IDs will be inserted as identical.
func (g *Graph) Insert(trs ...Triple) int {
	n := 0
	bnodes := make(map[BlankNode]int)
	for _, tr := range trs {

		var sid, pid, oid int

		// get subject id
		if bnode, ok := tr.Subject.(BlankNode); ok {
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
		if bnode, ok := tr.Object.(BlankNode); ok {
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
	return n
}

// addNode adds a node to the node dictonaries, and returns its new ID, or
// existing ID if already present.
func (g *Graph) addNode(node Node) int {
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

func (g *Graph) ids(tr Triple) (sid, pid, oid int, found bool) {
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
func (g *Graph) Delete(trs ...Triple) int {
	n := 0

	for _, tr := range trs {
		// Skip triples with blank nodes.
		if _, ok := tr.Subject.(BlankNode); ok {
			continue
		}
		if _, ok := tr.Object.(BlankNode); ok {
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
	return n
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

// func (g *Graph) DeleteWhere(trs...TriplePattern) int

// TriplePattern represents a pattern which can be used to match against a graph.
type TriplePattern struct {
	Subject   subject
	Predicate predicate
	Object    object
}

// Eq tests the euality between two TriplePatterns.
func (p TriplePattern) Eq(other TriplePattern) bool {
	return p.Subject == other.Subject &&
		p.Predicate == other.Predicate &&
		p.Object == other.Object
}

func (p TriplePattern) variables() (res []Variable) {
	if v, ok := p.Subject.(Variable); ok {
		res = append(res, v)
	}
	if v, ok := p.Predicate.(Variable); ok {
		res = append(res, v)
	}
	if v, ok := p.Object.(Variable); ok {
		res = append(res, v)
	}
	return res
}

// selectivity returns a selectivity score, determined by the number
// of and position of variables. This score can be used to select the
// execution order of graph patterns. This idea is proposed in the paper:
// Tsialiamanis, Petros, et al. "Heuristics-based query optimisation for SPARQL."
// Proceedings of the 15th International Conference on Extending Database Technology. ACM, 2012.
func (p TriplePattern) selectivity() int {
	// The pattern score from lowest (most selective) to highest (least selective) is
	// using the following order:
	//
	//   {s,p,o} < {s,?,o} < {?,p,o} < {s,p,?} < {?,?,o} < {s,?,?} < {?,p,?} < {?,?,?}
	//
	// In addition, patterns where the node in object position is a literal are scored
	// lower than if it is a named node/blank node, since a literal cannot have outgoing edges.
	//
	// TODO rewrite using node.selectivity() score

	vars := [3]bool{}
	objIsLiteral := 0
	if _, ok := p.Subject.(Variable); ok {
		vars[0] = true
	}
	if _, ok := p.Predicate.(Variable); ok {
		vars[1] = true
	}
	switch p.Object.(type) {
	case Variable:
		vars[2] = true
	case Literal:
		objIsLiteral = 1
	}

	switch vars {
	case [3]bool{false, false, false}:
		return 1 - objIsLiteral
	case [3]bool{false, true, false}:
		return 2 - objIsLiteral
	case [3]bool{true, false, false}:
		return 3 - objIsLiteral
	case [3]bool{false, false, true}:
		return 4
	case [3]bool{true, true, false}:
		return 5 - objIsLiteral
	case [3]bool{false, true, true}:
		return 6
	case [3]bool{true, false, true}:
		return 7
	case [3]bool{true, true, true}:
		return 8
	default:
		panic("BUG: TriplePattern.selectivity: unhandeled pattern")
	}
}

func (g *Graph) encodePattern(p TriplePattern, vars map[Variable]int) (ep encPattern) {
	// positive integer = nodeId (NamedNode, Literal)
	// 0 = node missing
	// negative integer = variable id

	for i, node := range []node{p.Subject, p.Predicate, p.Object} {
		if v, ok := node.(Variable); ok {
			if vid, ok := vars[v]; ok {
				ep[i] = vid
			} else {
				vars[v] = (len(vars) + 1) * -1
				ep[i] = vars[v]
			}
		} else {
			if id, ok := g.node2id[node.(Node)]; ok {
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

// encPattern is a TriplePattern encoded for query processing, where index
// 0-2 represents s,p,o nodes and the last represents the estimated cardinality.
// If a node equals=0, it means it is not stored in graph, meaning that the
// pattern (and all patterns sharing variables with it) will not match any.
// If a nodes < 0, it's a variable ID.
type encPattern [4]int

func (e encPattern) s() int              { return e[0] }
func (e encPattern) p() int              { return e[1] }
func (e encPattern) o() int              { return e[2] }
func (e encPattern) estCardinality() int { return e[3] }

// Where returns a new graph with the triples matching the given patterns.
// It corresponds to a SPARQL CONSTRUCT WHERE query, i.e where both template
// and patterns are identical.
func (g *Graph) Where(patterns ...TriplePattern) *Graph {
	if len(patterns) == 0 {
		return NewGraph()
	}

	encPatterns := make([]encPattern, len(patterns))
	vars := make(map[Variable]int)
	for i, pattern := range patterns {
		encPatterns[i] = g.encodePattern(pattern, vars)
	}

	res := NewGraph()
	bound := make([][]int, 0, len(vars))
	groups := groupPatternsByVariable(encPatterns)
	for n := range groups {
		// sort each group by estimated cardinality for the graph pattern
		sort.Slice(groups[n], func(i, j int) bool {
			return groups[n][i].estCardinality() < groups[n][j].estCardinality()
		})
	}

	for _, group := range groups {
		var matches []Triple
		for len(group) > 0 {
			var triples []Triple
			triples, bound = g.triplesMatching(group[0], bound)
			matches = append(matches, triples...)
			group = group[1:]
			reorderGroup(group, bound)
		}

		res.Insert(matches...)
	}

	return res
}

func (g *Graph) Select(vars []Variable, patterns ...TriplePattern) (res [][]Node) {

	encPatterns := make([]encPattern, len(patterns))
	varstmp := make(map[Variable]int, len(vars))
	for i, pattern := range patterns {
		encPatterns[i] = g.encodePattern(pattern, varstmp)
	}
	encVars := make([]int, len(vars))
	for i, v := range vars {
		if ev, ok := varstmp[v]; ok {
			encVars[i] = ev
		} // else: if variable does not occur in patterns, we ignore it.
	}

	bound := make([][]int, 0, len(vars))
	groups := groupPatternsByVariable(encPatterns)
	for n := range groups {
		// sort each group by estimated cardinality for the graph pattern
		sort.Slice(groups[n], func(i, j int) bool {
			return groups[n][i].estCardinality() < groups[n][j].estCardinality()
		})
	}

	var solutions [][]Node
	for _, group := range groups {
		for len(group) > 0 {
			pattern := group[0]
			solutions, bound = g.solutionsMatching(encVars, pattern, bound)
			res = append(res, solutions...)
			group = group[1:]
			reorderGroup(group, bound)
		}
	}
	// merge rows
	for i, row := range res {
		for j, node := range row {
			if node == nil {
				for y := range res {
					if res[y][j] != nil {
						res[i][j] = res[y][j]
						res[y][j] = nil
						break
					}
				}
			}
		}
	}
	// prune rows with nils
	for i := 0; i < len(res); i++ {
	checkNil:
		for j := 0; j < len(res[i]); j++ {
			if res[i][j] == nil {
				res[i] = res[len(res)-1]
				res = res[:len(res)-1]
				if i > 0 {
					i--
				}
				continue checkNil
			}
		}
	}

	return res
}

func indexVariables(ep encPattern, n int, v2g map[int]int) {
	for _, nid := range ep[:3] {
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

// reorderGroup moves pattern with bound variables to the top
func reorderGroup(patterns []encPattern, bound [][]int) {
	if len(patterns) <= 1 {
		return
	}
	for i, pattern := range patterns {
		for _, v := range pattern[:3] {
			if v < 0 {
				if v*-1 <= len(bound) && len(bound[(v*-1)-1]) > 0 {
					if i > 0 {
						patterns[i], patterns[i-1] = patterns[i-1], patterns[i]
					}
					return
				}
			}
		}
	}
}

func (g *Graph) solutionsMatching(vars []int, pattern encPattern, bound [][]int) ([][]Node, [][]int) {
	var solutions [][]Node
	var triples []Triple
	triples, bound = g.triplesMatching(pattern, bound)
	for _, triple := range triples {
		row := make([]Node, len(vars))
		match := false
		for i, v := range vars {
			if pattern[0] == v {
				row[i] = triple.Subject
				match = true
				continue
			}
			if pattern[1] == v {
				row[i] = triple.Predicate
				match = true
				continue
			}
			if pattern[2] == v {
				row[i] = triple.Object
				match = true
			}
		}
		if match {
			solutions = append(solutions, row)
		}
	}
	return solutions, bound
}

func (g *Graph) triplesMatching(ep encPattern, bound [][]int) ([]Triple, [][]int) {
	if ep[0] > 0 && ep[1] > 0 && ep[2] > 0 {
		// It's a concrete triple, we now it exists because patterns with
		// missing nodes has been pruned. Return it
		return []Triple{
			Triple{
				g.id2node[ep[0]].(SubjectNode),
				g.id2node[ep[1]].(NamedNode),
				g.id2node[ep[2]].(Node),
			},
		}, bound
	}
	if ep[2] > 0 {
		// Matching {any,any,n} using OSP index
		return g.triplesMatchingO(ep, bound)
	}
	if ep[1] > 0 {
		// Matching {?,n,n} using POS index
		return g.triplesMatchingP(ep, bound)
	}
	// Matching {any,any,any} using SPO index
	return g.triplesMatchingAny(ep, bound)
}

func (g *Graph) triplesMatchingO(ep encPattern, bound [][]int) ([]Triple, [][]int) {
	// {any,any,literal/uri}

	var solutions [][3]int

	for sid, preds := range g.osp[ep[2]] {
		if ep[0] < 0 {
			if !freeOrBound(ep[0], sid, bound) {
				continue
			}
		} else if ep[0] != sid {
			continue
		}
		for _, pid := range preds {
			if ep[1] < 0 {
				if !freeOrBound(ep[1], pid, bound) {
					continue
				}
			} else if ep[1] != pid {
				continue
			}
			solutions = append(solutions, [3]int{sid, pid, ep[2]})
		}
	}

	for _, match := range solutions {
		bound = updateBound(ep, match, bound)
	}
	var res []Triple
	for _, sol := range solutions {
		res = append(res, Triple{
			g.id2node[sol[0]].(SubjectNode),
			g.id2node[sol[1]].(NamedNode),
			g.id2node[ep[2]],
		})
	}
	return res, bound
}

func (g *Graph) triplesMatchingP(ep encPattern, bound [][]int) ([]Triple, [][]int) {
	// {var,uri,var}

	var solutions [][3]int

	for oid, subjs := range g.pos[ep[1]] {
		if ep[2] < 0 {
			if !freeOrBound(ep[2], oid, bound) {
				continue
			}
		} else if ep[2] != oid {
			continue
		}
		for _, sid := range subjs {
			if ep[0] < 0 {
				if !freeOrBound(ep[0], sid, bound) {
					continue
				}
			} else if ep[0] != sid {
				continue
			}
			solutions = append(solutions, [3]int{sid, ep[1], oid})
		}
	}

	for _, match := range solutions {
		bound = updateBound(ep, match, bound)
	}
	var res []Triple
	for _, sol := range solutions {
		res = append(res, Triple{
			g.id2node[sol[0]].(SubjectNode),
			g.id2node[ep[1]].(NamedNode),
			g.id2node[sol[2]],
		})
	}
	return res, bound
}

func (g *Graph) triplesMatchingAny(ep encPattern, bound [][]int) ([]Triple, [][]int) {
	// {any,any,any}
	var (
		solutions  [][3]int
		subjects   []int
		predicates []int
		objects    []int
	)

	/*
		if ep[0] < 0 {
			// subject is variable
			if len(bound) > ep[0]*-1 {
				// subject is already bound
				for _, sid := range bound[(ep[0]*-1)-1] {
					// iterate over bound subjects
					if ep[1] < 0 {
						// predicate is variable
					} else {
						// predicate is a node
					}
				}
			} else {
				// iterate over all subjects
			}
		} else {
			// subject is a node
			for pid, objs := range g.spo[ep[0]] {
				// iterate over subjects predicate->objects
			}
		}*/

	if ep[0] < 0 && ep[0]*-1 <= len(bound) && len(bound[(ep[0]*-1)-1]) > 0 {
		subjects = bound[(ep[0]*-1)-1]
	} else {
		subjects = g.subjects()
	}

	for _, sid := range subjects {
		if ep[0] > 0 && ep[0] != sid {
			continue
		}

		if ep[1] < 0 && ep[1]*-1 <= len(bound) && len(bound[(ep[1]*-1)-1]) > 0 {
			predicates = bound[(ep[1]*-1)-1]
		} else {
			predicates = g.predicatesForSubj(sid)
		}

		for _, pid := range predicates {
			if ep[1] > 0 && ep[1] != pid {
				continue
			}

			if ep[2] < 0 && ep[2]*-1 <= len(bound) && len(bound[(ep[2]*-1)-1]) > 0 {
				objects = bound[(ep[2]*-1)-1]
			} else {
				objects = g.spo[sid][pid]
			}
			for _, oid := range objects {
				if !includeNode(g.spo[sid][pid], oid) {
					continue
				}
				if ep[2] > 0 && ep[2] != oid {
					continue
				}

				solutions = append(solutions, [3]int{sid, pid, oid})
			}
		}
	}

	for _, match := range solutions {
		bound = updateBound(ep, match, bound)
	}
	var res []Triple
	for _, sol := range solutions {
		res = append(res, Triple{
			g.id2node[sol[0]].(SubjectNode),
			g.id2node[sol[1]].(NamedNode),
			g.id2node[sol[2]],
		})
	}
	return res, bound
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

func updateBound(ep encPattern, tr [3]int, bound [][]int) [][]int {
	for i, epid := range ep[:3] {
		if epid < 0 {
			if diff := epid*-1 - len(bound); diff != 0 {
				for j := 0; j < diff; j++ {
					bound = append(bound, []int{})
				}
				bound[(epid*-1)-1] = append(bound[(epid*-1)-1], tr[i])
			} else if !includeNode(bound[(epid*-1)-1], tr[i]) {
				bound[(epid*-1)-1] = append(bound[(epid*-1)-1], tr[i])
			}
		}
	}
	return bound
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
func freeOrBound(v int, id int, bound [][]int) bool {
	if len(bound) < v*-1 || len(bound[(v*-1)-1]) == 0 {
		return true
	}

	for _, node := range bound[(v*-1)-1] {
		if node == id {
			return true
		}
	}
	return false
}

// Eq checks if two graphs are equal (isomorphic).
func (g *Graph) Eq(other *Graph) bool {
	if g.Size() != other.Size() {
		return false
	}

	set := make(map[BlankNode]bool)

	// Check if all triples without blank-nodes are present in other.
	// Collect all blank nodes while iterating the graph.
	for s, po := range g.spo {
		if bnode, ok := g.id2node[s].(BlankNode); ok {
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
				if bnode, ok := g.id2node[o].(BlankNode); ok {
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

	var aBlankNodes []BlankNode
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

func (g *Graph) bnodes() []BlankNode {
	set := make(map[BlankNode]bool)
	for s, po := range g.spo {
		if bnode, ok := g.id2node[s].(BlankNode); ok {
			set[bnode] = true
			continue
		}
		for _, objs := range po {
			for _, o := range objs {
				if bnode, ok := g.id2node[o].(BlankNode); ok {
					set[bnode] = true
				}
			}
		}
	}
	var res []BlankNode
	for bnode := range set {
		res = append(res, bnode)
	}
	return res
}

func (g *Graph) signature(bnode BlankNode) string {
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
