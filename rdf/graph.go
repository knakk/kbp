package rdf

import (
	"sort"
	"strconv"
)

// Graph is an in-memory representation of an RDF graph:
// a colllection of triples which can be queried for subgraphs
// and mutated with inserting and deleting triples.
//
// A Graph is not thread-safe; concurrent writes/reads must be
// synchronized.
type Graph struct {
	// The graph is stored internally as a map of subjects to a map of
	// predicates to a slice of objects.
	nodes map[SubjectNode]map[NamedNode][]Node

	// nextID is the next unique integer to be used as Blank Node ID.
	nextID int
}

// NewGraph returns a new Graph.
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[SubjectNode]map[NamedNode][]Node),
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
					Subject:   s,
					Predicate: p,
					Object:    o,
				})
			}
		}
	}
	return res
}

// Contains returns true if the given Triple is present in Graph, otherwise false.
// TODO handle Blank Node subj/obj.
func (g *Graph) Contains(tr Triple) bool {
	if _, ok := g.nodes[tr.Subject]; !ok {
		return false
	}
	if _, ok := g.nodes[tr.Subject][tr.Predicate]; !ok {
		return false
	}

	for _, o := range g.nodes[tr.Subject][tr.Predicate] {
		if o.Eq(tr.Object) {
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
		switch t := tr.Subject.(type) {
		case BlankNode:
			if newID, ok := bnodes[t.id]; ok {
				tr.Subject = BlankNode{id: newID}
			} else {
				g.nextID++
				bnodes[t.id] = strconv.Itoa(g.nextID)
				tr.Subject = BlankNode{id: bnodes[t.id]}
			}
		}
		switch t := tr.Object.(type) {
		case BlankNode:
			if newID, ok := bnodes[t.id]; ok {
				tr.Object = BlankNode{id: newID}
			} else {
				g.nextID++
				bnodes[t.id] = strconv.Itoa(g.nextID)
				tr.Object = BlankNode{id: bnodes[t.id]}
			}
		}
		if _, ok := g.nodes[tr.Subject]; !ok {
			// new subject
			g.nodes[tr.Subject] = make(map[NamedNode][]Node)
		}
		if _, ok := g.nodes[tr.Subject][tr.Predicate]; !ok {
			// new predicate
			g.nodes[tr.Subject][tr.Predicate] = make([]Node, 0, 1)
		}
		for _, o := range g.nodes[tr.Subject][tr.Predicate] {
			if o == tr.Object {
				// triple already in graph
				continue outer
			}
		}
		// add object
		g.nodes[tr.Subject][tr.Predicate] = append(g.nodes[tr.Subject][tr.Predicate], tr.Object)
		n++
	}
	return n
}

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

// Where returns a new graph with the triples matching the given patterns.
// It corresponds to a SPARQL CONSTRUCT WHERE query, i.e where both template
// and patterns are identical.
func (g *Graph) Where(patterns []TriplePattern) *Graph {
	if len(patterns) == 0 {
		return NewGraph()
	}

	if len(patterns) == 1 {
		res := NewGraph()
		res.Insert(g.solutionsFor(patterns[0])...)
		return res
	}

	res := NewGraph()

	// Group patterns by shared variables, either direct or transitive.
	// Disjoint groups can be processed separately and its solutions merged by union.
	for _, group := range groupPatternsByVariable(patterns) {
		var matches []Triple
		for _, pattern := range group {
			matches = append(matches, g.solutionsFor(pattern)...)
		}
		res.Insert(matches...)
	}

	return res
}

func (g *Graph) solutionsFor(pattern TriplePattern) []Triple {
	var solutions []Triple
	for subj, po := range g.nodes {
		if !matchSubj(pattern.Subject, subj) {
			continue
		}

		for p, objs := range po {
			if !matchPred(pattern.Predicate, p) {
				continue
			}
			for _, o := range objs {
				if !matchObj(pattern.Object, o) {
					continue
				}

				solutions = append(solutions, Triple{subj, p, o})
			}
		}
	}
	return solutions
}

func matchSubj(s subject, other SubjectNode) bool {
	switch s.(type) {
	case Variable:
		return true
	default:
		return s.(SubjectNode).Eq(other)
	}
}

func matchPred(p predicate, other NamedNode) bool {
	switch p.(type) {
	case Variable:
		return true
	default:
		return p.(NamedNode).Eq(other)
	}
}

func matchObj(o object, other Node) bool {
	switch o.(type) {
	case Variable:
		return true
	default:
		return o.(Node).Eq(other)
	}
}

func matchPattern(p TriplePattern, t Triple) bool {
	return matchSubj(p.Subject, t.Subject) &&
		matchPred(p.Predicate, t.Predicate) &&
		matchObj(p.Object, t.Object)
}

func commonVar(p, other TriplePattern) bool {
	if a, ok := p.Subject.(Variable); ok {
		if b, ok := other.Subject.(Variable); ok {
			if a == b {
				return true
			}
		}
		if b, ok := other.Object.(Variable); ok {
			if a == b {
				return true
			}
			if c, ok := p.Object.(Variable); ok {
				if c == b {
					return true
				}
			}
		}
	}

	return false
}

func groupPatternsByVariable(patterns []TriplePattern) [][]TriplePattern {
	variables := make(map[Variable]int) // variable to group number
	groups := make([][]TriplePattern, 0)
	n := 0

	// TODO first sort patterns by variable name?

	for _, pattern := range patterns {
		assignedTo := 0
		for _, v := range pattern.variables() {
			if g, ok := variables[v]; ok {
				assignedTo = g
			}
		}
		if assignedTo == 0 {
			n++
			assignedTo = n
		}
		for _, v := range pattern.variables() {
			variables[v] = assignedTo
		}
		if len(groups) < n {
			groups = append(groups, []TriplePattern{pattern})
		} else {
			groups[n-1] = append(groups[n-1], pattern)
		}
	}

	return groups
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
						Triple{Subject: subj, Predicate: pred, Object: obj})
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
						Triple{Subject: subj, Predicate: pred, Object: obj})
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

	// Sort triples with Bnode in obj position here,
	// as not to risk sort them multiple times in isMatch function.
	for _, objs := range aBNodesAsObj {
		sort.Slice(objs, func(i, j int) bool {
			return objs[i].String() < objs[j].String()
		})
	}
	for _, objs := range bBNodesAsObj {
		sort.Slice(objs, func(i, j int) bool {
			return objs[i].String() < objs[j].String()
		})
	}

	matchA := make(map[BlankNode]bool)
	matchB := make(map[BlankNode]bool)
outer:
	for _, aNode := range aBNodes {
		for _, bNode := range bBNodes {
			if matchB[bNode] {
				continue
			}
			if len(g.nodes[aNode]) == len(other.nodes[bNode]) && len(aBNodesAsObj[aNode]) == len(bBNodesAsObj[bNode]) {
				if isMatch(g, other, aNode, bNode, aBNodesAsObj[aNode], bBNodesAsObj[bNode]) {
					matchA[aNode] = true
					matchB[bNode] = true
					continue outer
				}
			}
		}
		// no match was found for aNode in range bBNodes loop
		return false
	}

	return len(matchA) == len(aBNodes)
}

func isMatch(a, b *Graph, aNode, bNode BlankNode, aAsObj, bAsObj []Triple) bool {

	// Check for match in obj position first, as it is likely
	// less triples to compare, so we can return early if it's no match.
	for i, tr := range aAsObj {
		// Triples are allready sorted
		if !tr.Eq(bAsObj[i]) {
			return false
		}
	}

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

//func (g *Graph) Describe(u URI, depth int) *Graph {}
//func (g *Graph) Describe(u URI, depth int) *Graph {}
