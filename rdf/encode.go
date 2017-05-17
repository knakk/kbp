package rdf

import "io"

type Format int

const (
	NTtriples Format = iota
	Turtle
)

type Encoder struct {
	w io.Writer
}

func NewNTriplesEncoder(w io.Writer) *Encoder {
	return &Encoder{w: w}
}

func (e *Encoder) Encode(tr Triple) error {
	_, err := e.w.Write([]byte(tr.String()))
	return err
}
