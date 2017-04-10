package rdf

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	xsdString     = "http://www.w3.org/2001/XMLSchema#string"
	rdfLangString = "http://www.w3.org/1999/02/22-rdf-syntax-ns#langString"
)

// A Triple is the basic unit of information in an RDF graph.
// By itself, a triple is a statement describing a fact about a resource (subject),
// stating that it has a property (predicate) with a specific value (object).
//
// In the context of a RDF graph, a triple can be seen as two nodes with
// an edge between them.
type Triple struct {

	// The subject denotes the node which the triple statement is about.
	// It can be either named or blank.
	Subject SubjectNode

	// The predicate denotes the property of the subject. It is always named.
	Predicate NamedNode

	// The object denotes the value of the property. It can be any kind of node.
	Object Node
}

// String returns a string representation of a triple in N-Triples format.
func (t Triple) String() string {
	return fmt.Sprintf("%s %s %s .\n", t.Subject, t.Predicate, t.Object)
}

// Eq tests if two triples are equal.
func (t Triple) Eq(other Triple) bool {
	return t.Subject.Eq(other.Subject) && t.Predicate.Eq(other.Predicate) && t.Object.Eq(other.Object)
}

// ToTriplePattern returns the Triple as a TriplePattern.
func (t Triple) ToTriplePattern() TriplePattern {
	return TriplePattern{
		Subject:   t.Subject,
		Predicate: t.Predicate,
		Object:    t.Object.(object),
	}
}

// Node represents a node in an RDF graph.
type Node interface {
	Eq(Node) bool // TODO is this that useful?
	String() string

	validAsNode()
}

// SubjectNode represents the subject of a triple, which can be either
// a named node or a blank node.
type SubjectNode interface {
	Node
	subject
}

// BlankNode represents a blank node; that is, an unnamed RDF node.
type BlankNode struct {
	id string
}

// String returns a string representation of a Blank Node in N-Triples format.
func (b BlankNode) String() string { return "_:" + b.id }

// Eq tests for equality against another RDF node.
//
// Comparing a blank node with another will always return true. A blank
// really node only carries information in relation to it's position in a graph.
func (b BlankNode) Eq(other Node) bool {
	switch other.(type) {
	case BlankNode:
		return true
	default:
		return false
	}
}

func (b BlankNode) validAsNode()       {}
func (b BlankNode) validAsSubject()    {}
func (b BlankNode) validAsObject()     {}
func (b BlankNode) selectivity() int   { return 1 }
func (b BlankNode) nodeType() nodeType { return typeBlankNode }

// NamedNode represent an named node; an RDF node identified by an URI.
type NamedNode struct {
	val string
}

// NewNamedNode creates and returns an URI from the given string, along with an error if it's not valid.
func NewNamedNode(uri string) (NamedNode, error) {
	// A valid URI cannot be empty or contain any of disallowed characters.
	// http://www.ietf.org/rfc/rfc3987.txt
	if len(uri) == 0 {
		return NamedNode{}, errors.New("URI cannot be empty")
	}
	for _, r := range uri {
		if r >= '\x00' && r <= '\x20' {
			return NamedNode{}, fmt.Errorf("disallowed character in URI: %q", r)
		}
		switch r {
		case '<', '>', '"', '{', '}', '|', '^', '`', '\\':
			return NamedNode{}, fmt.Errorf("disallowed character in URI: %q", r)
		}
	}
	return NamedNode{val: uri}, nil
}

// String returns a string representation of an URI in N-Triples format.
func (u NamedNode) String() string { return "<" + u.val + ">" }

// Eq tests for equality against another RDF node.
func (u NamedNode) Eq(other Node) bool {
	switch t := other.(type) {
	case NamedNode:
		return u.val == t.val
	default:
		return false
	}
}

func (u NamedNode) validAsNode()       {}
func (u NamedNode) validAsPredicate()  {}
func (u NamedNode) validAsSubject()    {}
func (u NamedNode) validAsObject()     {}
func (u NamedNode) selectivity() int   { return 2 }
func (u NamedNode) nodeType() nodeType { return typeNamedNode }

// Literal represents an RDF Literal.
type Literal struct {
	val  string
	lang string
	dt   NamedNode
}

// NewStrLiteral returns a new string literal.
func NewStrLiteral(s string) Literal {
	return Literal{
		val: s,
		dt:  NamedNode{val: xsdString},
	}
}

// NewLangLiteral returns a new language-tagged literal.
func NewLangLiteral(s string, lang string) Literal {
	return Literal{
		val:  s,
		lang: lang,
		dt:   NamedNode{val: rdfLangString},
	}
}

// NewTypedLiteral returns a new literal with the give datatype.
func NewTypedLiteral(s string, dt NamedNode) Literal {
	return Literal{
		val: s,
		dt:  dt,
	}
}

// String returns a string representation of a Literal in N-Triples format.
func (l Literal) String() string {
	if (l.lang) != "" {
		return fmt.Sprintf("%q@%s", l.val, l.lang)
	}
	if (l.dt.val) == xsdString {
		return strconv.Quote(l.val)
	}
	return fmt.Sprintf("%q^^%s", l.val, l.dt)
}

// Eq tests for equality against another RDF node.
func (l Literal) Eq(other Node) bool {
	switch t := other.(type) {
	case Literal:
		return l.val == t.val &&
			l.lang == t.lang &&
			l.dt.val == t.dt.val
	default:
		return false
	}
}

// DataType returns the Datatype of a Literal.
func (l Literal) DataType() NamedNode { return l.dt }

// Lang returns the language tag of a Literal, or an empty
// string if it is not a rdf:langString.
func (l Literal) Lang() string { return l.lang }

func (l Literal) validAsNode()       {}
func (l Literal) validAsObject()     {}
func (l Literal) selectivity() int   { return 0 }
func (l Literal) nodeType() nodeType { return typeLiteral }

// Variable represents a variable which can be bound to RDF nodes in a query.
type Variable struct {
	name string
}

// NewVariable returns a new variable with the given name.
func NewVariable(name string) Variable {
	return Variable{name: name}
}

func (v Variable) validAsSubject()    {}
func (v Variable) validAsPredicate()  {}
func (v Variable) validAsObject()     {}
func (v Variable) selectivity() int   { return 3 }
func (v Variable) nodeType() nodeType { return typeVariable }

type node interface {
	// selectivity returns a selectivity score from lowest (most selective) to highest (least selective):
	//
	// 0 literal		Literal is most selective. It cannot have outgoing edges, like a uri/bnode
	// 1 bnode			More selective than uri because it usually only have one incoming edge
	// 2 uri
	// 3 variable
	selectivity() int
	nodeType() nodeType
}

type nodeType int

const (
	typeNamedNode = iota
	typeBlankNode
	typeLiteral
	typeVariable
)

func (t nodeType) String() string {
	switch t {
	case typeNamedNode:
		return "named node"
	case typeBlankNode:
		return "blank node"
	case typeLiteral:
		return "literal"
	case typeVariable:
		return "variable"
	default:
		panic("BUG: nodeType incomplete String()")
	}
}

type subject interface {
	node
	validAsSubject()
}

type predicate interface {
	node
	validAsPredicate()
}

type object interface {
	node
	validAsObject()
}

// TriplePattern represents a pattern which can be used to match against a graph. It differs
// from a Triple by that each of the subject, predicate and object may be variable.
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
