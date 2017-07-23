package rdf

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/knakk/kbp/rdf/internal/scanner"
)

type context [2]Node // [0] = subject, [1] = predicate

// Decoder decodes RDF triples i Turtle/N-Triples format.
type Decoder struct {
	s              *scanner.Scanner
	stack          []context
	bnodeN         int
	prefixes       map[string]NamedNode
	ParseVariables bool
}

// NewDecoder returns a new Decoder on the given stream.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{
		s:        scanner.NewStreamingScanner(r),
		prefixes: make(map[string]NamedNode),
	}
}

func (d *Decoder) ignoreLine() {
	for tok := d.s.Scan(); tok.Type != scanner.TokenEOL &&
		tok.Type != scanner.TokenEOF; tok = d.s.Scan() {
	}
}

func (d *Decoder) parseEnd() error {
	// Each statement must end in a dot (.)
	tok := d.s.Scan()
	switch tok.Type {
	case scanner.TokenIllegal:
		return errors.New(d.s.Error + ": " + tok.Text)
	case scanner.TokenEOF:
		return io.EOF
	case scanner.TokenDot:
	default:
		return fmt.Errorf("expected dot, got %s", tok.Type.String())
	}

	return nil
}

func (d *Decoder) oneOf(types ...nodeType) (TriplePatternNode, error) {
	node, err := d.parseNode()
	if err != nil {
		return nil, err
	}
	for _, t := range types {
		if t == node.nodeType() {
			return node, nil
		}
	}
	return nil, fmt.Errorf("unexpected node type: %v", node.nodeType())
}

var errEOL = errors.New("EOL")

func (d *Decoder) parseNode() (n TriplePatternNode, err error) {
	tok := d.s.Scan()
	switch tok.Type {
	case scanner.TokenPrefixDirective:
		tok = d.s.Scan()
		if tok.Type != scanner.TokenPrefix {
			return n, fmt.Errorf("unexpected %v", tok.Type)
		}
		prefix := tok.Text
		tok = d.s.Scan()
		if tok.Type != scanner.TokenURI {
			return n, fmt.Errorf("unexpected %v", tok.Type)
		}
		d.prefixes[prefix] = NewNamedNode(tok.Text)
		tok = d.s.Scan()
		if tok.Type != scanner.TokenDot {
			return n, fmt.Errorf("unexpected %v", tok.Type)
		}
		return d.parseNode()
	case scanner.TokenPrefix:
		prefix := tok.Text
		tok = d.s.Scan()
		if tok.Type != scanner.TokenSuffix {
			return n, fmt.Errorf("unexpected %v", tok.Type)
		}
		if _, ok := d.prefixes[prefix]; !ok {
			return n, fmt.Errorf("unknown prefix %q", prefix)
		}
		return d.prefixes[prefix].Resolve(tok.Text), nil
	case scanner.TokenLiteral:
		n = Literal{val: tok.Text, dt: XSDstring}
	case scanner.TokenIllegal:
		return n, errors.New(d.s.Error + ": " + tok.Text)
	case scanner.TokenEOF:
		return n, io.EOF
	case scanner.TokenEOL:
		return n, errEOL
	case scanner.TokenURI:
		return NamedNode{name: tok.Text}, nil
	case scanner.TokenBNode:
		return BlankNode{id: tok.Text}, nil
	case scanner.TokenVariable:
		return Variable{name: tok.Text}, nil
	case scanner.TokenBracketStart:
		d.bnodeN++
		return BlankNode{id: "b" + strconv.Itoa(d.bnodeN)}, nil
	default:
		return n, fmt.Errorf("unexpected %v", tok.Type)
	}

	// Check if we have datatype or language-tag for literal

	switch d.s.Peek().Type {
	case scanner.TokenLangTag:
		tok = d.s.Scan()
		n = Literal{
			val:  n.(Literal).val,
			dt:   RDFlangString,
			lang: tok.Text,
		}
	case scanner.TokenTypeMarker:
		d.s.Scan() // consume ^^
		tok = d.s.Scan()
		if tok.Type == scanner.TokenURI {
			n = Literal{
				val: n.(Literal).val,
				dt:  NamedNode{name: tok.Text},
			}
			break
		}
		err = fmt.Errorf("unexpected %v", tok.Type)
	}

	return n, err
}

// Decode returns the next valid triple in the the stream, or an error.
func (d *Decoder) Decode() (Triple, error) {
	var tr Triple
	var c context

	if len(d.stack) > 0 {
		c = d.stack[len(d.stack)-1]
		d.stack = d.stack[:len(d.stack)-1]
	}

	bnodeN := d.bnodeN
	if c[0] != nil {
		tr.Subject = c[0].(SubjectNode)
	} else {
		subj, err := d.oneOf(typeNamedNode, typeBlankNode)
		if err == errEOL {
			return d.Decode() // Continue looking for new triple after empty line
		} else if err == io.EOF {
			return tr, io.EOF
		} else if err != nil {
			d.ignoreLine()
			return tr, fmt.Errorf("%d:%d: error parsing subject: %q", d.s.Row, d.s.Col, err)
		}
		tr.Subject = subj.(SubjectNode)
		if bnodeN != d.bnodeN {
			d.stack = append(d.stack, context{subj.(Node)})
		}
	}

	if c[1] != nil {
		tr.Predicate = c[1].(NamedNode)
	} else {
		pred, err := d.oneOf(typeNamedNode)
		if err == errEOL {
			if c[0] != nil {
				d.stack = append(d.stack, c)
				return d.Decode()
			}
		} else if err != nil {
			d.ignoreLine()
			return tr, fmt.Errorf("%d:%d: error parsing preicate: %q", d.s.Row, d.s.Col, err)
		}
		tr.Predicate = pred.(NamedNode)
	}

	bnodeN = d.bnodeN
	obj, err := d.oneOf(typeNamedNode, typeBlankNode, typeLiteral)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing object: %q", d.s.Row, d.s.Col, err)
	}
	tr.Object = obj.(Node)
	if bnodeN != d.bnodeN {
		d.stack = append(d.stack, context{obj.(Node)})
	}

	switch d.s.Peek().Type {
	case scanner.TokenSemicolon:
		d.s.Scan() // consume ;
		d.stack = append(d.stack, context{tr.Subject.(Node)})
		return tr, nil
	case scanner.TokenComma:
		d.s.Scan() // consume ,
		d.stack = append(d.stack, context{tr.Subject.(Node), Node(tr.Predicate)})
		return tr, nil
	case scanner.TokenBracketEnd:
		d.s.Scan() // consume ]
		if d.s.Peek().Type == scanner.TokenDot {
			d.s.Scan()
		}
		switch d.s.Peek().Type {
		case scanner.TokenSemicolon:
			d.s.Scan() // consume ;
			d.stack = append(d.stack, context{tr.Subject.(Node)})
			return tr, nil
		case scanner.TokenComma:
			d.s.Scan() // consume ,
			d.stack = append(d.stack, context{tr.Subject.(Node), Node(tr.Predicate)})
			return tr, nil
		}
		return tr, nil
	}
	if len(d.stack) == 0 {
		if err = d.parseEnd(); err != nil {
			return tr, err
		}
	}
	d.ignoreLine()
	return tr, nil
}

// DecodePattern decodes a TriplePattern.
func (d *Decoder) DecodePattern() (TriplePattern, error) {
	var tr TriplePattern
	subj, err := d.oneOf(typeNamedNode, typeBlankNode, typeVariable)
	if err == errEOL {
		return d.DecodePattern() // Continue looking for new triple after empty line
	} else if err == io.EOF {
		return tr, io.EOF
	} else if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing subject: %q", d.s.Row, d.s.Col, err)
	}
	tr.Subject = subj.(subject)

	pred, err := d.oneOf(typeNamedNode, typeVariable)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing preicate: %q", d.s.Row, d.s.Col, err)
	}
	tr.Predicate = pred.(predicate)

	obj, err := d.oneOf(typeNamedNode, typeBlankNode, typeLiteral, typeVariable)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing object: %q", d.s.Row, d.s.Col, err)
	}
	tr.Object = obj.(object)

	if err = d.parseEnd(); err != nil {
		return tr, err
	}
	d.ignoreLine()
	return tr, nil
}

func ParseUpdateQuery(s string) (del, ins, where []TriplePattern, err error) {
	d := Decoder{
		s: scanner.NewScanner(s),
	}
	var tr TriplePattern
	for {
		switch d.s.Peek().Type {
		case scanner.TokenEOL:
			d.s.Scan() // consume
		case scanner.TokenPlus:
			d.s.Scan() // consume +
			tr, err = d.DecodePattern()
			if err != nil {
				return
			}
			ins = append(ins, tr)
		case scanner.TokenMinus:
			d.s.Scan() // consume -
			tr, err = d.DecodePattern()
			if err != nil {
				return
			}
			del = append(del, tr)
		case scanner.TokenEOF:
			return
		default:
			tr, err = d.DecodePattern()
			if err != nil {
				return
			}
			where = append(where, tr)
		}
	}
	return
}

// TODO implement this properly
func MustParseNode(node string) Node {
	d := Decoder{
		s: scanner.NewScanner(node),
	}
	n, err := d.parseNode()
	if err != nil && err != io.EOF {
		panic(err)
	}
	if n == nil {
		return nil
	}
	return n.(Node)
}
