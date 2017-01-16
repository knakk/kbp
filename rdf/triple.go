package rdf

import (
	"errors"
	"fmt"
	"strconv"
)

// Triple represent an RDF Triple, also known as an RDF statement.
type Triple struct {
	Subj Subject
	Pred URI
	Obj  Term
}

// String returns a string representation of a Triple in N-Triples format.
func (t Triple) String() string {
	return fmt.Sprintf("%s %s %s .\n", t.Subj, t.Pred, t.Obj)
}

// Term is the interface for an RDF Term, which can be either an URI,
// Literal or Blank Node.
type Term interface {
	Eq(Term) bool
	String() string

	validAsTerm()
}

// Subject represents the subject of a Triple, which can be either
// an URI or a Blank Node.
type Subject interface {
	Term
	subject
}

// BlankNode represents a Blank Node; that is, an unnamed RDF node.
type BlankNode struct {
	id string
}

func (b BlankNode) validAsTerm()    {}
func (b BlankNode) validAsSubject() {}

// String returns a string representation of a Blank Node in N-Triples format.
func (b BlankNode) String() string { return "_:" + b.id }

// Eq checks for equality against another RDF term.
//
// A BNode is always equal to another BNode, as it does not
// make sense to compare without the graph it is part of.
func (b BlankNode) Eq(other Term) bool {
	switch other.(type) {
	case BlankNode:
		return true
	default:
		return false
	}
}

// URI represent an URI; a globally unique named RDF node.
type URI struct {
	val string
}

// NewURI creates and returns an URI from the given string, along with an error if it's not valid.
func NewURI(uri string) (URI, error) {
	// A valid URI cannot be empty or contain any of disallowed characters.
	// http://www.ietf.org/rfc/rfc3987.txt
	if len(uri) == 0 {
		return URI{}, errors.New("URI cannot be empty")
	}
	for _, r := range uri {
		if r >= '\x00' && r <= '\x20' {
			return URI{}, fmt.Errorf("disallowed character in URI: %q", r)
		}
		switch r {
		case '<', '>', '"', '{', '}', '|', '^', '`', '\\':
			return URI{}, fmt.Errorf("disallowed character in URI: %q", r)
		}
	}
	return URI{val: uri}, nil
}
func (u URI) validAsTerm()      {}
func (u URI) validAsPredicate() {}
func (u URI) validAsSubject()   {}
func (u URI) validAsObject()    {}

// String returns a string representation of an URI in N-Triples format.
func (u URI) String() string { return "<" + u.val + ">" }

// Eq checks for euality against another RDF term.
func (u URI) Eq(other Term) bool {
	switch t := other.(type) {
	case URI:
		return u.val == t.val
	default:
		return false
	}
}

// Literal represents an RDF Literal.
type Literal struct {
	val  string
	lang string
	dt   URI
}

func NewStrLiteral(s string) Literal {
	return Literal{
		val: s,
		dt:  URI{val: "http://www.w3.org/2001/XMLSchema#string"},
	}
}

func NewLangLiteral(s string, lang string) Literal {
	return Literal{
		val:  s,
		lang: lang,
		dt:   URI{val: "http://www.w3.org/1999/02/22-rdf-syntax-ns#langString"},
	}
}

func (l Literal) validAsTerm()   {}
func (l Literal) validAsObject() {}

// DataType returns the Datatype of a Literal.
func (l Literal) DataType() URI { return l.dt }

// Lang returns the language tag of a Literal, or an empty
// string if it is not a rdf:langString.
func (l Literal) Lang() string { return l.lang }

// String returns a string representation of a Literal in N-Triples format.
func (l Literal) String() string {
	if (l.lang) != "" {
		return fmt.Sprintf("%q@%s", l.val, l.lang)
	}
	if (l.dt.val) == "http://www.w3.org/2001/XMLSchema#string" {
		return strconv.Quote(l.val)
	}
	return fmt.Sprintf("%q^^%s", l.val, l.dt)
}

// Eq checks for euality against another RDF term.
func (l Literal) Eq(other Term) bool {
	switch t := other.(type) {
	case Literal:
		return l.val == t.val &&
			l.lang == t.lang &&
			l.dt.val == t.dt.val
	default:
		return false
	}
}

// func (l Literal) LangTag() (language.Tag, error) { return language.Parse(l.l) }

// func (l Literal) Int() (int, bool)
// func (l Literal) String() (string, bool)
// func (l Literal) Time() time.Time

// func (l Literal) MustInt() int
// func (l Literal) MustString() string
// func (l Literal) MustTime() time.Time

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

// Eq checks if two triples are equal.
func (tr Triple) Eq(other Triple) bool {
	return tr.Subj.Eq(other.Subj) && tr.Pred.Eq(other.Pred) && tr.Obj.Eq(other.Obj)
}
