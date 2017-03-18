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

// Node represents a node in an RDF graph.
type Node interface {
	Eq(Node) bool // TODO is this that usefull?
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

func (b BlankNode) validAsNode()    {}
func (b BlankNode) validAsSubject() {}

// String returns a string representation of a Blank Node in N-Triples format.
func (b BlankNode) String() string { return "_:" + b.id }

// Eq tests for equality against another RDF term.
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
func (u NamedNode) validAsNode()      {}
func (u NamedNode) validAsPredicate() {}
func (u NamedNode) validAsSubject()   {}
func (u NamedNode) validAsObject()    {}

// String returns a string representation of an URI in N-Triples format.
func (u NamedNode) String() string { return "<" + u.val + ">" }

// Eq tests for equality against another RDF term.
func (u NamedNode) Eq(other Node) bool {
	switch t := other.(type) {
	case NamedNode:
		return u.val == t.val
	default:
		return false
	}
}

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

func (l Literal) validAsNode()   {}
func (l Literal) validAsObject() {}

// DataType returns the Datatype of a Literal.
func (l Literal) DataType() NamedNode { return l.dt }

// Lang returns the language tag of a Literal, or an empty
// string if it is not a rdf:langString.
func (l Literal) Lang() string { return l.lang }

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

// Eq tests for equality against another RDF term.
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

// Variable represents a variable which can be bound to RDF nodes in a query.
type Variable struct {
	id string
}

func (v Variable) validAsSubject()   {}
func (v Variable) validAsPredicate() {}
func (v Variable) validAsObj()       {}

type subject interface {
	validAsSubject()
}

type predicate interface {
	validAsPredicate()
}

type object interface {
	validAsObject()
}

// TriplePattern represents a pattern which can be used to match against a graph.
type TriplePattern struct {
	Subj subject
	Pred predicate
	Obj  object
}
