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
			`<a> <b> "c"^^<http://www.w3.org/2001/XMLSchema#String> .`,
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
	}

	for _, test := range tests {
		a := mustDecode(test.a)
		b := mustDecode(test.b)
		if got := a.Eq(b); got != test.eq {
			t.Errorf("%v \nEq\n %v = %v; want %v", mustEncode(a), mustEncode(b), got, test.eq)
		}
	}
}
