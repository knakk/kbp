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

// specificity returns a specificity score, determined by the number
// of and position of variables.
func (p TriplePattern) specificity() int {
	// {s,p,o} < {s,?,o} < {?,p,o} < {s,p,?} < {?,?,o} < {s,?,?} < {?,p,?} < {?,?,?}

	vars := [3]bool{}
	objLiteral := 0
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
		// A Literal in object position is more specific than an URI
		objLiteral = 1
	}

	switch vars {
	case [3]bool{false, false, false}:
		return 1 - objLiteral
	case [3]bool{false, true, false}:
		return 2 - objLiteral
	case [3]bool{true, false, false}:
		return 3 - objLiteral
	case [3]bool{false, false, true}:
		return 4
	case [3]bool{true, true, false}:
		return 5 - objLiteral
	case [3]bool{false, true, true}:
		return 6
	case [3]bool{true, false, true}:
		return 7
	case [3]bool{true, true, true}:
		return 8
	default:
		panic("BUG: TriplePattern.specificity: unhandeled pattern")
	}
}

// Where returns a new graph with the triples matching the given patterns.
// It corresponds to a SPARQL CONSTRUCT WHERE query, i.e where both template
// and patterns are identical.
func (g *Graph) Where(patterns []TriplePattern) *Graph {
	if len(patterns) == 0 {
		return NewGraph()
	}

	bound := make(map[Variable][]Node)
	res := NewGraph()

	if len(patterns) == 1 {
		res.Insert(g.solutionsFor(patterns[0], bound)...)
		return res
	}

	// Group patterns by shared variables, either direct or transitive.
	// Disjoint groups can be processed separately and its solutions merged by union.
	for _, group := range groupPatternsByVariable(patterns) {
		var matches []Triple
		for len(group) > 0 {
			pattern := group[0]
			solutions := g.solutionsFor(pattern, bound)
			matches = append(matches, solutions...)
			group = group[1:]
			reorderGroup(group, bound)
		}

		res.Insert(matches...)
	}

	return res
}

// reorderGroup moves pattern with bound variables to the top
func reorderGroup(patterns []TriplePattern, bound map[Variable][]Node) {
	if len(patterns) <= 1 {
		return
	}
	for i, pattern := range patterns {
		for _, v := range pattern.variables() {
			if _, ok := bound[v]; ok {
				if i > 0 {
					patterns[i], patterns[i-1] = patterns[i-1], patterns[i]
				}
				return
			}
		}
	}
}

func (g *Graph) solutionsFor(pattern TriplePattern, bound map[Variable][]Node) []Triple {
	var (
		solutions  []Triple
		subjects   []SubjectNode
		predicates []NamedNode
		objects    []Node
	)

	if nodes, ok := boundSubject(pattern.Subject, bound); ok {
		subjects = nodes
	} else {
		subjects = g.subjects()
	}

	for _, subj := range subjects {
		if !matchSubj(pattern.Subject, subj) {
			continue
		}

		if nodes, ok := boundPredicate(pattern.Predicate, bound); ok {
			predicates = nodes
		} else {
			predicates = g.predicatesForSubj(subj)
		}

		for _, p := range predicates {
			if !matchPred(pattern.Predicate, p) {
				continue
			}

			if nodes, ok := boundObject(pattern.Object, bound); ok {
				objects = nodes
			} else {
				objects = g.nodes[subj][p]
			}
			for _, o := range objects {
				if !includeNode(g.nodes[subj][p], o) {
					continue
				}
				if !matchObj(pattern.Object, o) {
					continue
				}

				match := Triple{subj, p, o}
				solutions = append(solutions, match)
			}
		}
	}

	for _, match := range solutions {
		updateBound(pattern, match, bound)
	}
	return solutions
}

func (g *Graph) subjects() (res []SubjectNode) {
	for subj := range g.nodes {
		res = append(res, subj)
	}
	return res
}

func (g *Graph) predicatesForSubj(s SubjectNode) (res []NamedNode) {
	for pred := range g.nodes[s] {
		res = append(res, pred)
	}
	return res
}

func updateBound(p TriplePattern, match Triple, bound map[Variable][]Node) {
	if v, ok := p.Subject.(Variable); ok {
		if !includeNode(bound[v], match.Subject) {
			bound[v] = append(bound[v], match.Subject)
		}
	}
	if v, ok := p.Predicate.(Variable); ok {
		if !includeNode(bound[v], match.Predicate) {
			bound[v] = append(bound[v], match.Predicate)
		}
	}
	if v, ok := p.Object.(Variable); ok {
		if !includeNode(bound[v], match.Object) {
			bound[v] = append(bound[v], match.Object)
		}
	}
}

func includeNode(nodes []Node, find Node) bool {
	for _, node := range nodes {
		if node == find {
			return true
		}
	}
	return false
}

func boundSubject(s subject, variables map[Variable][]Node) ([]SubjectNode, bool) {
	if v, ok := s.(Variable); ok {
		if nodes, bound := variables[v]; bound {
			var res []SubjectNode
			for _, n := range nodes {
				if subj, ok := n.(SubjectNode); ok {
					res = append(res, subj)
				}
			}
			return res, true
		}
	}
	return nil, false
}

func boundPredicate(p predicate, variables map[Variable][]Node) ([]NamedNode, bool) {
	if v, ok := p.(Variable); ok {
		if nodes, bound := variables[v]; bound {
			var res []NamedNode
			for _, n := range nodes {
				if subj, ok := n.(NamedNode); ok {
					res = append(res, subj)
				}
			}
			return res, true
		}
	}
	return nil, false
}

func boundObject(o object, variables map[Variable][]Node) ([]Node, bool) {
	if v, ok := o.(Variable); ok {
		if nodes, bound := variables[v]; bound {
			return nodes, true
		}
	}
	return nil, false
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

func groupPatternsByVariable(patterns []TriplePattern) [][]TriplePattern {
	variables := make(map[Variable]int) // variable to group number
	groups := make([][]TriplePattern, 0)
	n := 0

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

	// Sort the patterns in each group by specificity
	for i := range groups {
		sortBySpecificity(groups[i])
	}

	return groups
}

func sortBySpecificity(patterns []TriplePattern) {
	sort.Slice(patterns, func(i, j int) bool {
		return patterns[i].specificity() < patterns[j].specificity()
	})
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
