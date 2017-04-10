package rdf

import (
	"sort"
	"strings"
	"testing"

	"github.com/knakk/kbp/rdf"
	"github.com/knakk/kbp/rdf/memory"
)

func wantGraph(t *testing.T, got *memory.Graph, wantGraph string) {
	want := mustDecode(wantGraph)
	if eq, _ := want.Eq(got); !eq {
		t.Fatalf("\ngot:\n%v\nwant:\n%v", mustEncode(got), mustEncode(want))
	}
}

func mustTriples(s string) []rdf.Triple {
	return mustDecode(s).Triples()
}

func mustParseVariables(s string) []rdf.Variable {
	var res []rdf.Variable
	for _, v := range strings.Split(s, "?") {
		if strings.TrimSpace(v) == "" {
			continue
		}
		res = append(res, rdf.NewVariable(strings.TrimSpace(v)))
	}
	return res
}

func mustParseSolutions(s string) (res [][]rdf.Node) {
	rows := strings.Split(s, "\n")[1:] // discard first line (column header listing variables)
	for _, row := range rows {
		var solution []rdf.Node
		for _, node := range strings.Split(row, "\t") {
			solution = append(solution, rdf.MustParseNode(node))
		}
		res = append(res, solution)
	}
	return res
}

func nodesToString(nodes []rdf.Node) string {
	var s string
	for _, n := range nodes {
		s += n.String()
	}
	return s
}

func solutionsEq(a, b [][]rdf.Node) bool {
	if len(a) != len(b) {
		return false
	}

	sort.Slice(a, func(i, j int) bool {
		return nodesToString(a[i]) < nodesToString(a[j])
	})

	sort.Slice(b, func(i, j int) bool {
		return nodesToString(b[i]) < nodesToString(b[j])
	})

	for i, row := range a {
		for j, node := range row {
			if b[i][j] != node {
				return false
			}
		}
	}

	return true
}

func TestGraphMutateOperations(t *testing.T) {
	g := memory.NewGraph()

	rdf.Insert(g, mustTriples(`
		<h1> <name> "A" .
		<h1> <knows> <h2> .
		<h2> <name> "B" .
		<h2> <knows> <h1> .`)...,
	)
	wantGraph(t, g, `
		<h1> <name> "A" .
		<h1> <knows> <h2> .
		<h2> <name> "B" .
		<h2> <knows> <h1> .
		`,
	)

	rdf.Insert(g, mustTriples(`
		<b1> <title> "book" .
		<b1> <contributor> _:c1 .
		_:c1 <role> <author> .
		_:c1 <agent> <h1> .
		<b1> <contributor> _:c2 .
		_:c2 <role> <illustrator> .
		_:c2 <agent> <h2> .
		`)...,
	)
	wantGraph(t, g, `
		<h1> <name> "A" .
		<h1> <knows> <h2> .
		<h2> <name> "B" .
		<h2> <knows> <h1> .
		<b1> <title> "book" .
		<b1> <contributor> _:123 .
		_:123 <role> <author> .
		_:123 <agent> <h1> .
		<b1> <contributor> _:456 .
		_:456 <role> <illustrator> .
		_:456 <agent> <h2> .`,
	)

	rdf.Delete(g, mustTriples(`
		<b1> <contributor> _:a .
		_:a <role> <illustrator> .
		_:a <agent> <h2> `)...,
	)

	wantGraph(t, g, `
		<h1> <name> "A" .
		<h1> <knows> <h2> .
		<h2> <name> "B" .
		<h2> <knows> <h1> .
		<b1> <title> "book" .
		<b1> <contributor> _:1 .
		_:1 <role> <author> .
		_:1 <agent> <h1> .
		<b1> <contributor> _:2 .
		_:2 <role> <illustrator> .
		_:2 <agent> <h2> .`,
	)

	rdf.Delete(g, mustTriples(`
		<h1> <knows> <h2> .
		<h2> <knows> <h1> .
		`)...,
	)

	wantGraph(t, g, `
		<h1> <name> "A" .
		<h2> <name> "B" .
		<b1> <title> "book" .
		<b1> <contributor> _:1 .
		_:1 <role> <author> .
		_:1 <agent> <h1> .
		<b1> <contributor> _:2 .
		_:2 <role> <illustrator> .
		_:2 <agent> <h2> .
		`,
	)
}

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
			`<s> <p> "o" .`,
			`<s> <p> "o" .`,
			true,
		},
		{
			`<s> <p> "o" .`,
			`<s> <p> "o"^^<http://www.w3.org/2001/XMLSchema#string> .`,
			true,
		},
		{
			`<s> <p> "o" .`,
			`<s> <p> "o2" .`,
			false,
		},
		{
			`<s> <p> "o" .
			 <s> <p> "xyz" .`,
			`<s> <p> "o" .`,
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
			`<s> <p> _:a .`,
			`<s> <p> _:a .`,
			true,
		},
		{
			`<s> <p> _:a .`,
			`<s> <p> _:b .`,
			true,
		},
		{
			`<s> <p> _:a .
			 _:a <p> "xyz" . `,
			`<s> <p> _:b .
			 _:b <p> "xyz" . `,
			true,
		},
		{
			`<s> <p> _:a .
			 _:a <p> "xyz" .
			 _:a <p> "æøå" . `,
			`<s> <p> _:a .
			_:a <p> "xyz" . `,
			false,
		},
		{
			`<s> <p> _:a .
			 _:a <p> "xyz" .
			 _:a <p> "æøå" .`,
			`<s> <p> _:b .
			 _:b <p> "xyz" .
			 _:b <p> "æøå" .`,
			true,
		},
		{
			`<s> <p> _:a .
			 _:a <p> "xyz" . `,
			`<s> <p> _:a .
			 _:a <p> "XYZ" . `,
			false,
		},
		{
			`<s> <p> _:a .
			 _:a <p> "xyz" .
			 _:a <p> "æøå" .`,
			`<s> <p> _:b .
			 _:b <p> "xyz" .
			 _:b <p> "æøå!!!" .`,
			false,
		},
		{
			`<s> <p> _:a .
			 _:a <p> "xyz" .
			 <s> <p> _:a2 .
			 _:a2 <p> "123" .`,
			`<s> <p> _:a .
			 _:a <p> "xyz" .
			 <s> <p> _:a2 .
			 _:a2 <p> "123" .`,
			true,
		},
	}

	for _, test := range tests {
		a := mustDecode(test.a)
		b := mustDecode(test.b)
		if got, _ := a.Eq(b); got != test.eq {
			t.Fatalf("\n%v  Eq\n%v  = %v; want %v", mustEncode(a), mustEncode(b), got, test.eq)
		}
	}
}

var testGraph = mustDecode(`
		<a1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Person> .
		<a1> <hasName> "Italo Calvino" .
		<a1> <hasBirthYear> "1923"^^<http://www.w3.org/2001/XMLSchema#gYear> .
		<a1> <hasDeathYear> "1985"^^<http://www.w3.org/2001/XMLSchema#gYear> .

		<w1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Work> .
		<w1> <hasMainTitle> "Le Cosmicomiche" .
		<w1> <hasPublicationYear> "1965"^^<http://www.w3.org/2001/XMLSchema#gYear> .
		<w1> <hasContributor> _:c1 .

		_:c1 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
		_:c1 <hasRole> <author> .
		_:c1 <hasAgent> <a1> .

		<p1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Publication> .
		<p1> <isPublicationOf> <w1> .
		<p1> <hasMainTitle> "The Complete Cosmicomics" .
		<p1> <hasISBN> "9781846141652" .
		<p1> <hasPublishYear> "2009"^^<http://www.w3.org/2001/XMLSchema#gYear> .
		<p1> <isPublishedBy> <c1> .
		<p1> <hasContributor> _:c2 .
		<p1> <hasContributor> _:c3 .
		<p1> <hasContributor> _:c4 .
		<p1> <hasAbstract> "The definitive edition of Calvino’s cosmicomics, bringing together all of these enchanting stories—including some never before translated — in one volume for the first time" .

		_:c2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
		_:c2 <hasRole> <translator> .
		_:c2 <hasAgent> <a2> .

		_:c3 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
		_:c3 <hasRole> <translator> .
		_:c3 <hasAgent> <a3> .

		_:c4 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
		_:c4 <hasRole> <translator> .
		_:c4 <hasAgent> <a4> .

		<a2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Person> .
		<a2> <hasName> "Martin L. McLaughlin" .

		<a3> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Person> .
		<a3> <hasName> "Tim Parks" .

		<a4> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Person> .
		<a4> <hasName> "William Weaver" .

		<c1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Corporation> .
		<c1> <hasName> "Penguin" .

		<w2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Work> .
		<w2> <hasMainTitle> "Il barone rampante" .
		<w2> <hasPublicationYear> "1957"^^<http://www.w3.org/2001/XMLSchema#gYear> .
		<w2> <hasContributor> _:c5 .

		_:c5 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
		_:c5 <hasRole> <author> .
		_:c5 <hasAgent> <a1> .

		<p2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Publication> .
		<p2> <isPublicationOf> <w2> .
		<p2> <hasMainTitle> "Klatrebaronen" .
		<p2> <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .
		<p2> <isPublishedBy> <c2> .
		<p2> <hasContributor> _:c6 .

		_:c6 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
		_:c6 <hasRole> <translator> .
		_:c6 <hasAgent> <a4> .

		<a5> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Person> .
		<a5> <hasName> "Ingeborg Hagemann" .

		<c1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Corporation> .
		<c1> <hasName> "Gyldendal" .
		`)

func TestGraphWhere(t *testing.T) {
	tests := []struct {
		patterns []rdf.TriplePattern
		want     string
	}{
		{
			mustParsePatterns(""),
			"",
		},
		{
			mustParsePatterns("?s ?p ?o ."),
			mustEncode(testGraph),
		},
		{
			mustParsePatterns("?pub <hasMainTitle> ?title ."),
			`
			<w1> <hasMainTitle> "Le Cosmicomiche" .
			<p1> <hasMainTitle> "The Complete Cosmicomics" .
			<w2> <hasMainTitle> "Il barone rampante" .
			<p2> <hasMainTitle> "Klatrebaronen" .`,
		},
		{
			mustParsePatterns(`?p2 <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .`),
			`<p2> <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .`,
		},
		{
			mustParsePatterns("<p2> ?p ?o ."),
			`
			<p2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Publication> .
			<p2> <isPublicationOf> <w2> .
			<p2> <hasMainTitle> "Klatrebaronen" .
			<p2> <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .
			<p2> <isPublishedBy> <c2> .
			<p2> <hasContributor> _:1 .`,
		},
		{
			mustParsePatterns(`?s ?p "Penguin" .`),
			`<c1> <hasName> "Penguin" .`,
		},
		{
			mustParsePatterns(`?s ?p "Penguin"@en .`),
			"",
		},
		{
			mustParsePatterns(`
				?w <hasMainTitle> "Le Cosmicomiche" .
				?p <isPublicationOf> ?w .
				?p <hasMainTitle> ?pubTitle .`,
			),
			`
			<w1> <hasMainTitle> "Le Cosmicomiche" .
			<p1> <isPublicationOf> <w1> .
			<p1> <hasMainTitle> "The Complete Cosmicomics" .
			`,
		},
		{
			mustParsePatterns(`
				?c <hasRole> <translator> .
				?c <hasAgent> ?agent .
				?agent <hasName> ?name .`,
			),
			`
			_:1 <hasAgent> <a4> .
			_:1 <hasRole> <translator> .
			_:2 <hasAgent> <a2> .
			_:2 <hasRole> <translator> .
			_:3 <hasAgent> <a3> .
			_:3 <hasRole> <translator> .
			_:4 <hasAgent> <a4> .
			_:4 <hasRole> <translator> .
			<a2> <hasName> "Martin L. McLaughlin" .
			<a4> <hasName> "William Weaver" .
			<a3> <hasName> "Tim Parks" .
			`,
		},
		{
			mustParsePatterns(`
				?agent <hasName> "William Weaver" .
				?c <hasAgent> ?agent .
				?c <hasRole> <translator> .
				`,
			),
			`
			_:1 <hasAgent> <a4> .
			_:1 <hasRole> <translator> .
			_:4 <hasAgent> <a4> .
			_:4 <hasRole> <translator> .
			<a4> <hasName> "William Weaver" .
			`,
		},
		{
			mustParsePatterns(`
				?c <hasRole> <translator> .
				?c <hasAgent> ?agent .
				?agent <hasName> "William Weaver" .
				`,
			),
			`
			_:1 <hasAgent> <a4> .
			_:1 <hasRole> <translator> .
			_:4 <hasAgent> <a4> .
			_:4 <hasRole> <translator> .
			<a4> <hasName> "William Weaver" .
			`,
		},
		{
			mustParsePatterns(`
				?work <hasMainTitle> "Le Cosmicomiche" .
				?book <isPublicationOf> ?work.
				?book <hasContributor> ?contrib .
				?contrib <hasRole> <translator> .
				?contrib <hasAgent> ?person .
				?person <hasName> ?translator .`,
			),
			`
			<w1> <hasMainTitle> "Le Cosmicomiche" .
			<p1> <isPublicationOf> <w1> .
			<p1> <hasContributor> _:c2 .
			<p1> <hasContributor> _:c3 .
			<p1> <hasContributor> _:c4 .
			_:c2 <hasRole> <translator> .
			_:c2 <hasAgent> <a2> .
			_:c3 <hasRole> <translator> .
			_:c3 <hasAgent> <a3> .
			_:c4 <hasRole> <translator> .
			_:c4 <hasAgent> <a4> .
			<a2> <hasName> "Martin L. McLaughlin" .
			<a3> <hasName> "Tim Parks" .
			<a4> <hasName> "William Weaver" .
			`,
		},
	}

	for i, test := range tests {
		got, _ := testGraph.Where(test.patterns...)
		if eq, _ := got.Eq(mustDecode(test.want)); !eq {
			t.Fatalf("%d:got:\n%v\nwant:\n%v", i, mustEncode(got.(*memory.Graph)), test.want)
		}
	}
}

func TestGraphSelect(t *testing.T) {
	tests := []struct {
		vars     []rdf.Variable
		patterns []rdf.TriplePattern
		want     string
	}{
		{
			mustParseVariables(""),
			mustParsePatterns(""),
			"",
		},
		{
			mustParseVariables("?name"),
			mustParsePatterns("?person <hasName> ?name ."),
			`name
             "Italo Calvino"
             "Martin L. McLaughlin"
             "Tim Parks"
             "William Weaver"
             "Penguin"
             "Ingeborg Hagemann"
             "Gyldendal"`,
		},
		{
			mustParseVariables("?name"),
			mustParsePatterns(`?c <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Corporation> .
							   ?c <hasName> ?name .`),
			`name
             "Penguin"
             "Gyldendal"`,
		},
		{
			mustParseVariables("?book ?translator"),
			mustParsePatterns(`?book <hasContributor> ?contrib .
							   ?book <isPublicationOf> ?work.
							   ?work <hasMainTitle> "Le Cosmicomiche" .
							   ?contrib <hasRole> <translator> .
							   ?contrib <hasAgent> ?person .
							   ?person <hasName> ?translator .`),
			`book	translator
             <p1>	"Martin L. McLaughlin"
             <p1>	"Tim Parks"
             <p1>	"William Weaver"`,
		},
	}

	for _, test := range tests {
		want := mustParseSolutions(test.want)
		got, _ := testGraph.Select(test.vars, test.patterns...)
		if !solutionsEq(got, want) {
			t.Fatalf("got:\n%v\nwant:\n%v", got, want)
		}
	}
}
