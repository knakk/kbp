package rdf

import (
	"bytes"
	"testing"
)

func mustDecode(s string) *Graph {
	dec := NewDecoder(bytes.NewBufferString(s))
	g, err := dec.DecodeGraph()
	if err != nil {
		panic("mustDecode: " + err.Error())
	}
	return g
}

func mustEncode(g *Graph) string {
	var out bytes.Buffer
	enc := NewEncoder(&out)
	if err := enc.EncodeGraph(g); err != nil {
		panic("mustEncode: " + err.Error())
	}
	return out.String()
}

func TestDecode(t *testing.T) {
	// Testcase lifted from W3C testsuite at: https://www.w3.org/2013/N-TriplesTests/
	const input = `
<http://example.org/resource1> <http://example.org/property> <http://example.org/resource2> .
_:anon <http://example.org/property> <http://example.org/resource2> .
<http://example.org/resource2> <http://example.org/property> _:anon .
# spaces and tabs throughout:
 	 <http://example.org/resource3> 	 <http://example.org/property>	 <http://example.org/resource2> 	.

# line ending with CR NL (ASCII 13, ASCII 10)
<http://example.org/resource4> <http://example.org/property> <http://example.org/resource2> .

# 2 statement lines separated by single CR (ASCII 10)
<http://example.org/resource5> <http://example.org/property> <http://example.org/resource2> .
<http://example.org/resource6> <http://example.org/property> <http://example.org/resource2> .


# All literal escapes
<http://example.org/resource7> <http://example.org/property> "simple literal" .
<http://example.org/resource8> <http://example.org/property> "backslash:\\" .
<http://example.org/resource9> <http://example.org/property> "dquote:\"" .
<http://example.org/resource10> <http://example.org/property> "newline:\n" .
<http://example.org/resource11> <http://example.org/property> "return\r" .
<http://example.org/resource12> <http://example.org/property> "tab:\t" .

# Space is optional before final .
<http://example.org/resource13> <http://example.org/property> <http://example.org/resource2>.
<http://example.org/resource14> <http://example.org/property> "x".
<http://example.org/resource15> <http://example.org/property> _:anon.

# \u and \U escapes
# latin small letter e with acute symbol \u00E9 - 3 UTF-8 bytes #xC3 #A9
<http://example.org/resource16> <http://example.org/property> "\u00E9" .
# Euro symbol \u20ac  - 3 UTF-8 bytes #xE2 #x82 #xAC
<http://example.org/resource17> <http://example.org/property> "\u20AC" .

# XML Literals as Datatyped Literals
<http://example.org/resource21> <http://example.org/property> ""^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource22> <http://example.org/property> " "^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource23> <http://example.org/property> "x"^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource23> <http://example.org/property> "\""^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource24> <http://example.org/property> "<a></a>"^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource25> <http://example.org/property> "a <b></b>"^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource26> <http://example.org/property> "a <b></b> c"^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource26> <http://example.org/property> "a\n<b></b>\nc"^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .
<http://example.org/resource27> <http://example.org/property> "chat"^^<http://www.w3.org/2000/01/rdf-schema#XMLLiteral> .

# Plain literals with languages
<http://example.org/resource30> <http://example.org/property> "chat"@fr .
<http://example.org/resource31> <http://example.org/property> "chat"@en .

# Typed Literals
<http://example.org/resource32> <http://example.org/property> "abc"^^<http://example.org/datatype1> .
`
	g := mustDecode(input)
	nt := mustEncode(g)
	g2 := mustDecode(nt)

	if !g.Eq(g2) {
		t.Fatal("decode-encode-decode roundtrip failed")
	}
}
