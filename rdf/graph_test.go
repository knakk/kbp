package rdf

import "testing"

func TestGraphIsomorphism(t *testing.T) {
	tests := []struct {
		a  string
		b  string
		eq bool
	}{
		{
			"",
			"",
			true,
		},
		{
			`<a> <b> "c" .`,
			`<a> <b> "c" .`,
			true,
		},
		{
			`<a> <b> "c" .`,
			`<a> <b> "c"^^<http://www.w3.org/2001/XMLSchema#string> .`,
			true,
		},
		{
			`<a> <b> "c" .`,
			`<a> <b> "C" .`,
			false,
		},
		{
			`<a> <b> "c" .
			 <a> <b> "xyz" .`,
			`<a> <b> "c" .`,
			false,
		},
		{
			`_:a <p> "1" .`,
			`_:a <p> "1" .`,
			true,
		},
		{
			`_:a <p> "1" .`,
			`_:b <p> "1" .`,
			true,
		},
		{
			`<a> <b> _:a .`,
			`<a> <b> _:a .`,
			true,
		},
		{
			`<a> <b> _:a .`,
			`<a> <b> _:b .`,
			true,
		},
	}

	for _, test := range tests {
		a := mustDecode(test.a)
		b := mustDecode(test.b)
		if got := a.Eq(b); got != test.eq {
			t.Errorf("\n%v  Eq\n%v  = %v; want %v", mustEncode(a), mustEncode(b), got, test.eq)
		}
	}
}
