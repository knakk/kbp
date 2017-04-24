package rdf

import (
	"errors"
	"fmt"
	"io"

	"github.com/knakk/kbp/rdf/internal/scanner"
)

// Decoder decodes RDF triples i N-Triples format.
type Decoder struct {
	s              *scanner.Scanner
	ParseVariables bool
}

// NewDecoder returns a new Decoder on the given stream.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{s: scanner.NewStreamingScanner(r)}
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

func (d *Decoder) oneOf(types ...nodeType) (node, bool, error) {
	node, parsedDot, err := d.parseNode()
	if err != nil {
		return nil, false, err
	}
	for _, t := range types {
		if t == node.nodeType() {
			return node, parsedDot, nil
		}
	}
	return nil, false, fmt.Errorf("unexpected node type: %v", node.nodeType())
}

var errEOL = errors.New("EOL")

func (d *Decoder) parseNode() (n node, parsedDot bool, err error) {
	tok := d.s.Scan()
	switch tok.Type {
	case scanner.TokenLiteral:
		n = Literal{val: tok.Text, dt: NamedNode{name: xsdString}}
	case scanner.TokenIllegal:
		return n, false, errors.New(d.s.Error + ": " + tok.Text)
	case scanner.TokenEOF:
		return n, false, io.EOF
	case scanner.TokenEOL:
		return n, false, errEOL
	case scanner.TokenURI:
		return NamedNode{name: tok.Text}, false, nil
	case scanner.TokenBNode:
		return BlankNode{id: tok.Text}, false, nil
	case scanner.TokenVariable:
		return Variable{name: tok.Text}, false, nil
	default:
		return n, false, fmt.Errorf("unexpected %v", tok.Type)
	}

	// Check if we have datatype or language-tag for literal
	tok = d.s.Scan()
	switch tok.Type {
	case scanner.TokenIllegal:
		err = errors.New(d.s.Error + ": " + tok.Text)
	case scanner.TokenEOF:
		err = io.EOF
	case scanner.TokenEOL:
		err = errEOL
	case scanner.TokenDot:
		parsedDot = true
	case scanner.TokenLangTag:
		n = Literal{
			val:  n.(Literal).val,
			dt:   NamedNode{name: rdfLangString},
			lang: tok.Text,
		}
	case scanner.TokenTypeMarker:
		tok = d.s.Scan()
		if tok.Type == scanner.TokenURI {
			n = Literal{
				val: n.(Literal).val,
				dt:  NamedNode{name: tok.Text},
			}
			break
		}
		fallthrough
	default:
		err = fmt.Errorf("unexpected %v", tok.Type)
	}

	return n, parsedDot, err
}

// Decode returns the next valid triple in the the stream, or an error.
func (d *Decoder) Decode() (Triple, error) {
	var tr Triple
	subj, _, err := d.oneOf(typeNamedNode, typeBlankNode)
	if err == errEOL {
		return d.Decode() // Continue looking for new triple after empty line
	} else if err == io.EOF {
		return tr, io.EOF
	} else if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing subject: %q", d.s.Row, d.s.Col, err)
	}
	tr.Subject = subj.(SubjectNode)

	pred, _, err := d.oneOf(typeNamedNode)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing preicate: %q", d.s.Row, d.s.Col, err)
	}
	tr.Predicate = pred.(NamedNode)

	obj, parsedDot, err := d.oneOf(typeNamedNode, typeBlankNode, typeLiteral)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing object: %q", d.s.Row, d.s.Col, err)
	}
	tr.Object = obj.(Node)

	if !parsedDot {
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
	subj, _, err := d.oneOf(typeNamedNode, typeBlankNode, typeVariable)
	if err == errEOL {
		return d.DecodePattern() // Continue looking for new triple after empty line
	} else if err == io.EOF {
		return tr, io.EOF
	} else if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing subject: %q", d.s.Row, d.s.Col, err)
	}
	tr.Subject = subj.(subject)

	pred, _, err := d.oneOf(typeNamedNode, typeVariable)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing preicate: %q", d.s.Row, d.s.Col, err)
	}
	tr.Predicate = pred.(predicate)

	obj, parsedDot, err := d.oneOf(typeNamedNode, typeBlankNode, typeLiteral, typeVariable)
	if err != nil {
		d.ignoreLine()
		return tr, fmt.Errorf("%d:%d: error parsing object: %q", d.s.Row, d.s.Col, err)
	}
	tr.Object = obj.(object)

	if !parsedDot {
		if err = d.parseEnd(); err != nil {
			return tr, err
		}
	}
	d.ignoreLine()
	return tr, nil
}

// TODO implement this properly
func MustParseNode(node string) Node {
	d := Decoder{
		s: scanner.NewScanner(node),
	}
	n, _, err := d.parseNode()
	if err != nil && err != io.EOF {
		panic(err)
	}
	return n.(Node)
}
