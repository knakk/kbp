package rdf

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

// Decoder is a decodes RDF triples i N-Triples format.
type Decoder struct {
	s *scanner
}

// NewDecoder returns a new Decoder on the given stream.
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{s: newScanner(r)}
}

func (d *Decoder) ignoreLine() {
	for tok := d.s.Scan(); tok.Type != tokenEOL &&
		tok.Type != tokenEOF; tok = d.s.Scan() {
	}
}

func (d *Decoder) parseSubject() (token, error) {
	var tok token
	for tok = d.s.Scan(); tok.Type == tokenEOL; tok = d.s.Scan() {
	}
	switch tok.Type {
	case tokenIllegal:
		return token{}, errors.New(tok.Text)
	case tokenEOF:
		return token{}, io.EOF
	case tokenBNode, tokenURI:
		break
	default:
		return token{}, fmt.Errorf("expected URI/BNode as subject, got %s", tok.Type.String())
	}

	return tok, nil
}

func (d *Decoder) parsePredicate() (token, error) {
	tok := d.s.Scan()
	switch tok.Type {
	case tokenIllegal:
		return token{}, errors.New(d.s.Error + ": " + tok.Text)
	case tokenEOF:
		return token{}, io.EOF
	case tokenURI:
		break
	default:
		return token{}, fmt.Errorf("expected URI as predicate, got %s", tok.Type.String())
	}
	return tok, nil
}

func (d *Decoder) parseObject() (token, error) {
	tok := d.s.Scan()
	switch tok.Type {
	case tokenIllegal:
		return token{}, errors.New(d.s.Error + ": " + tok.Text)
	case tokenEOF:
		return token{}, io.EOF
	case tokenURI, tokenLiteral, tokenBNode:
		break
	default:
		return token{}, fmt.Errorf("expected URI/BNode/Literal as object, got %s", tok.Type.String())
	}

	return tok, nil
}

func (d *Decoder) parseEnd() error {
	// Each statement must end in a dot (.)
	tok := d.s.Scan()
	switch tok.Type {
	case tokenIllegal:
		return errors.New(d.s.Error + ": " + tok.Text)
	case tokenEOF:
		return io.EOF
	case tokenDot:
		break
	default:
		return fmt.Errorf("expected dot, got %s", tok.Type.String())
	}

	// Any tokens after dot until end of line are ignored
	d.ignoreLine()

	return nil
}

// Decode returns the next valid triple in the the stream, or an error.
func (d *Decoder) Decode() (Triple, error) {
	var tr Triple
	// subject
	tok, err := d.parseSubject()
	if err == io.EOF {
		return tr, io.EOF
	}
	if err != nil {
		d.ignoreLine()
		return Triple{}, fmt.Errorf("%d:%d: error parsing subject: %q", d.s.Row, d.s.Col, err)
	}
	if tok.Type == tokenBNode {
		tr.Subj = BlankNode{id: tok.Text}
	} else {
		tr.Subj = URI{val: tok.Text}
	}

	// predicate
	tok, err = d.parsePredicate()
	if err != nil {
		d.ignoreLine()
		return Triple{}, fmt.Errorf("%d:%d: error parsing predicate: %q", d.s.Row, d.s.Col, err)
	}
	tr.Pred = URI{val: tok.Text}

	// object
	tok, err = d.parseObject()
	if err != nil {
		d.ignoreLine()
		return Triple{}, fmt.Errorf("%d:%d: error parsing object: %q", d.s.Row, d.s.Col, err)
	}
	if tok.Type == tokenBNode {
		tr.Obj = BlankNode{id: tok.Text}
	} else if tok.Type == tokenURI {
		tr.Obj = URI{val: tok.Text}
	} else {
		// literal
		obj := tok.Text
		next := d.s.Scan()

		switch next.Type {
		case tokenDot:
			// plain literal xsd:String
			tr.Obj = Literal{val: obj, dt: URI{val: "http://www.w3.org/2001/XMLSchema#String"}}
			d.ignoreLine()
			return tr, nil
		case tokenLangTag:
			// rdf:langString
			tr.Obj = Literal{val: obj, lang: next.Text, dt: URI{val: "http://www.w3.org/1999/02/22-rdf-syntax-ns#langString"}}
		case tokenTypeMarker:
			// typed literal
			next = d.s.Scan()
			if next.Type != tokenURI {
				d.ignoreLine()
				return Triple{}, fmt.Errorf("%d: expected URI as literal datatype, got %v: %q", d.s.line, tok.Type, next.Text)
			}
			tr.Obj = Literal{val: tok.Text, dt: URI{val: next.Text}}
		case tokenIllegal:
			d.ignoreLine()
			return Triple{}, fmt.Errorf("%d:%d: error parsing object: %s: %v", d.s.Row, d.s.Col, d.s.Error, next.Text)
		default:
			return Triple{}, fmt.Errorf("unexpected %s: %q after object", next.Type, next.Text)
		}
	}

	// dot+newline/eof
	if err := d.parseEnd(); err != nil {
		d.ignoreLine()
		return Triple{}, err
	}

	return tr, nil
}

// DecodeGraph consumes stream until the end and decodes all triples into a graph.
// If an error occurs, it will return it along with the graph parsed so far.
func (d *Decoder) DecodeGraph() (*Graph, error) {
	g := NewGraph()
	for tr, err := d.Decode(); err != io.EOF; tr, err = d.Decode() {
		if err != nil {
			return g, err
		}
		g.Insert(tr)
	}
	return g, nil
}

type Encoder struct {
	w *bufio.Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: bufio.NewWriter(w)}
}

func (e *Encoder) EncodeGraph(g *Graph) error {
	for s, po := range g.nodes {
		for p, objs := range po {
			for _, o := range objs {
				if _, err := e.w.WriteString(Triple{Subj: s, Pred: p, Obj: o}.String()); err != nil {
					return err
				}
			}
		}
	}
	return e.w.Flush()
}
