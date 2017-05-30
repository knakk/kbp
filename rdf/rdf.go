package rdf

import (
	"errors"
	"fmt"
	"html/template"
	"strconv"
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
	return t.Subject == other.Subject && t.Predicate == other.Predicate && t.Object == other.Object
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
	String() string

	node()
}

// SubjectNode represents the subject of a triple, which can be either
// a named node or a blank node.
type SubjectNode interface {
	Node
	subject
}

func (BlankNode) node() {}
func (NamedNode) node() {}
func (Literal) node()   {}

// subject represents a Node which can be in Subject position in a TriplePattern.
type subject interface {
	TriplePatternNode
	subject()
}

func (BlankNode) subject() {}
func (NamedNode) subject() {}
func (Variable) subject()  {}

// predicate represents a Node which can be in Predicate position in a TriplePattern.
type predicate interface {
	TriplePatternNode
	predicate()
}

func (NamedNode) predicate() {}
func (Variable) predicate()  {}

// object represents a Node which can be in Object position in a TriplePattern.
type object interface {
	TriplePatternNode
	object()
}

func (BlankNode) object() {}
func (NamedNode) object() {}
func (Literal) object()   {}
func (Variable) object()  {}

// BlankNode represents a blank node; that is, an unnamed RDF node.
type BlankNode struct {
	id string
}

// NewBlankNode returns a BlankNode with the given ID.
func NewBlankNode(id string) BlankNode {
	return BlankNode{id: id}
}

// String returns a string representation of a Blank Node in N-Triples format.
func (b BlankNode) String() string { return "_:" + b.id }

// ID returns the ID of a BlankNode.
func (b BlankNode) ID() string { return b.id }

// NamedNode represent an named node; an RDF node identified by an URI.
type NamedNode struct {
	name string
}

// NewNamedNode creates a new NamedNode.
func NewNamedNode(name string) NamedNode {
	return NamedNode{name: name}
}

// NewURI creates an URI as NamedNode from the given string,
// or an error if it's not valid.
func NewURI(uri string) (NamedNode, error) {
	// A valid URI cannot be empty or contain any disallowed characters.
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
	return NamedNode{name: uri}, nil
}

// String returns a string representation of NamedNode in N-Triples format.
func (n NamedNode) String() string { return "<" + n.name + ">" }

// Name returns the Name of the NamedNode.
func (n NamedNode) Name() string { return n.name }

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
		dt:  XSDstring,
	}
}

// NewLangLiteral returns a new language-tagged literal.
func NewLangLiteral(s string, lang string) Literal {
	return Literal{
		val:  s,
		lang: lang,
		dt:   RDFlangString,
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
	if l.lang != "" {
		return fmt.Sprintf("%q@%s", l.val, l.lang)
	}
	if l.dt == XSDstring {
		return strconv.Quote(l.val)
	}
	return fmt.Sprintf("%q^^%s", l.val, l.dt)
}

// DataType returns the Datatype of a Literal.
func (l Literal) DataType() NamedNode { return l.dt }

// Lang returns the language tag of a Literal, or an empty
// string if it is not a rdf:langString.
func (l Literal) Lang() string { return l.lang }

// Value returns the typed value of a Literal.
func (l Literal) Value() interface{} {
	switch l.DataType() {
	case XSDgYear, XSDinteger, XSDint: // int
		v, err := strconv.Atoi(l.val)
		if err != nil {
			return l.val
		}
		return v
	case XSDunsignedInt:
		v, err := strconv.Atoi(l.val)
		if err != nil {
			return l.val
		}
		return uint(v)
	case RDFHTML:
		v, err := template.New("").Parse(l.val)
		if err != nil {
			return l.val
		}
		return v
	default:
		panic("rdf.Literal.Value() TODO: " + l.DataType().String())
	}
}

// ValueAsString returns the value of a Literal as a string.
func (l Literal) ValueAsString() string { return l.val }

// Variable represents a variable which can be bound to RDF nodes in a query.
type Variable struct {
	name string
}

// NewVariable returns a new variable with the given name.
func NewVariable(name string) Variable {
	return Variable{name: name}
}

// Name returns the name of a variable.
func (v Variable) Name() string { return v.name }

// String returns a string-representation of a Variable.
func (v Variable) String() string { return "?" + v.name }

type nodeType int

const (
	typeNamedNode = iota
	typeBlankNode
	typeLiteral
	typeVariable
)

// TriplePatternNode represent a Node that can be part of a TriplePattern.
// TODO rename BGP?
type TriplePatternNode interface {
	// TODO reconsider this method - it's only used by the decoder.oneOf function
	nodeType() nodeType
}

func (Literal) nodeType() nodeType   { return typeLiteral }
func (NamedNode) nodeType() nodeType { return typeNamedNode }
func (BlankNode) nodeType() nodeType { return typeBlankNode }
func (Variable) nodeType() nodeType  { return typeVariable }

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

// TriplePattern represents a pattern which can be used to match against a graph. It differs
// from a Triple by that each of the subject, predicate and object may be variable.
type TriplePattern struct {
	Subject   subject
	Predicate predicate
	Object    object
}

// IsConcrete tests if the triple pattern represent a concrete triple,
// i.e. contains no variables.
func (p TriplePattern) IsConcrete() bool {
	return p.Subject.nodeType() != typeVariable &&
		p.Predicate.nodeType() != typeVariable &&
		p.Object.nodeType() != typeVariable
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
