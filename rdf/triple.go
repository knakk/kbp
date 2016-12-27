package rdf

import (
	"fmt"
	"strconv"
)

type Triple struct {
	Subj Subject
	Pred URI
	Obj  Term
}

func (t Triple) String() string {
	return fmt.Sprintf("%s %s %s .\n", t.Subj, t.Pred, t.Obj)
}

type Term interface {
	Eq(Term) bool
	String() string

	validAsTerm()
}

type Subject interface {
	Term
	subject
}

type BlankNode struct {
	id string
}

func (b BlankNode) validAsTerm()    {}
func (b BlankNode) validAsSubject() {}

func (b BlankNode) String() string { return "_:" + b.id }

func (b BlankNode) Eq(other Term) bool {
	switch t := other.(type) {
	case BlankNode:
		return b.id == t.id
	default:
		return false
	}
}

type URI struct {
	val string
}

func (u URI) validAsTerm()      {}
func (u URI) validAsPredicate() {}
func (u URI) validAsSubject()   {}
func (u URI) validAsObject()    {}

func (u URI) String() string { return "<" + u.val + ">" }

func (u URI) Eq(other Term) bool {
	switch t := other.(type) {
	case URI:
		return u.val == t.val
	default:
		return false
	}
}

type Literal struct {
	val  string
	lang string
	dt   URI
}

func (l Literal) validAsTerm()   {}
func (l Literal) validAsObject() {}
func (l Literal) DataType() URI  { return l.dt }
func (l Literal) Lang() string   { return l.lang }

func (l Literal) String() string {
	if (l.lang) != "" {
		return fmt.Sprintf("%q@%s", l.val, l.lang)
	}
	if (l.dt.val) == "http://www.w3.org/2001/XMLSchema#string" {
		return strconv.Quote(l.val)
	}
	return fmt.Sprintf("%q^^%s", l.val, l.dt)
}

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

type TriplePattern struct {
	Subj subject
	Pred predicate
	Obj  object
}
