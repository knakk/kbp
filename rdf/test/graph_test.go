package rdf

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/knakk/kbp/rdf"
	"github.com/knakk/kbp/rdf/disk"
	"github.com/knakk/kbp/rdf/memory"
)

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

func mustParseUpdate(s string) (del, ins, where []rdf.TriplePattern) {
	del, ins, where, err := rdf.ParseUpdateQuery(s)
	if err != nil {
		panic(err)
	}
	return
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

// tempfile returns a temporary file path.
func tempfile() string {
	f, _ := ioutil.TempFile("", "kbprdfgraphdisktemp")
	f.Close()
	os.Remove(f.Name())
	return f.Name()
}

func newDiskGraph(f string) *disk.Graph {
	db, err := disk.Open(f, "http://test.org/")
	if err != nil {
		panic("cannot open db: " + err.Error())
	}
	return db
}

type graphImpl struct {
	name    string
	graph   rdf.Graph
	closeFn func()
}

func newGraphImplementations() []graphImpl {
	f := tempfile()
	dg := newDiskGraph(f)
	impls := []graphImpl{
		{
			"memory",
			memory.NewGraph(),
			func() {},
		},
		{
			"disk",
			dg,
			func() {
				defer os.Remove(f)
				dg.Close()
			},
		},
	}

	return impls
}

func TestGraphUpdate(t *testing.T) {

	tests := []struct {
		input  string
		update string
		want   string
	}{
		{ // 1
			``,
			``,
			``,
		},
		{ // 2
			``,

			`+ <s> <p> <o> .`,

			`<s> <p> <o> .`,
		},
		{ // 3
			`<s> <p> <o> .`,

			`+ <s> <p> <o2> .`,

			`<s> <p> <o> .
			 <s> <p> <o2> .`,
		},
		{ // 4
			`<s> <p> <o> .`,

			`- <s> <p> <o> .`,

			``,
		},
		{ // 5
			`<s> <p> "a" .`,

			`- <s> <p> "a" .
			 + <s> <p> "b" .`,

			`<s> <p> "b" .`,
		},
		{ // 6
			`<s> <p> "a" .
			 <s2> <p> "a" .`,

			`- ?s <p> "a" .
			 ?s <p> "a" .`,

			``,
		},
		{ // 7
			`<s> <p> "a" .
			 <s2> <p> "a" .`,

			`- ?s <p> "a" .
			 + ?s <p> "b" .
			 ?s <p> "a" .`,

			`<s> <p> "b" .
			 <s2> <p> "b" .`,
		},
		{ // 8
			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .`,

			`- ?c <hasRole> <author> .
			 + ?c <hasRole> <editor> .
			 <book1> <hasContribution> ?c .
			 ?c <hasAgent> <h1> .`,

			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <editor> .`,
		},
		{ // 9
			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .
			 <book2> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <author> .`,

			`- ?c <hasRole> <author> .
			 + ?c <hasRole> <editor> .
			 ?book <hasContribution> ?c .
			 ?c <hasRole> <author> .`,

			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <editor> .
			 <book2> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <editor> .`,
		},
		{ // 10
			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .
			 <book2> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <author> .`,

			`- ?c <hasRole> <author> .
			 + ?c <hasRole> <editor> .
			 ?book <hasContribution> ?c .
			 ?c <hasRole> <author> .
			 ?c <hasAgent> <h2> .`,

			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .
			 <book2> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <editor> .`,
		},
		{ // 11
			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .
			 <book2> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <author> .`,

			`- ?book <hasContribution> ?c .
			 - ?c ?p ?o .
			 ?book <hasContribution> ?c .
			 ?c <hasAgent> <h2> .
			 ?c ?p ?o .`,

			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .`,
		},
		{ // 12
			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .
			 <book2> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <author> .`,

			`- <book2> <hasContribution> ?c .
			 + <book1> <hasContribution> ?c .
			 <book2> <hasContribution> ?c .`,

			`<book1> <hasContribution> _:c1 .
			 _:c1 <hasAgent> <h1> .
			 _:c1 <hasRole> <author> .
			 <book1> <hasContribution> _:c2 .
			 _:c2 <hasAgent> <h2> .
			 _:c2 <hasRole> <author> .`,
		},
		{
			// 13
			`<person> <hasName> "Name" .
			 <person> <hasBirthYear> "1988" .
			 <book> <hasTitle> "title" .
			 <book> <hasContribution> _:c .
			 _:c <hasAgent> <person> .
			 _:c <hasRole> <author> .`,

			`- <person> <hasBirthYear> "1988" .
			 + <person> <hasBirthYear> "1888" .
			 + <person> <hasDeathYear> "1958" .`,

			`<person> <hasName> "Name" .
			 <person> <hasBirthYear> "1888" .
			 <person> <hasDeathYear> "1958" .
			 <book> <hasTitle> "title" .
			 <book> <hasContribution> _:c .
			 _:c <hasAgent> <person> .
			 _:c <hasRole> <author> ..`,
		},
	}

	impls := newGraphImplementations()
	matchAll := mustParsePatterns("?s ?p ?o .")

	for _, impl := range impls {
		t.Run(fmt.Sprintf("%s implementation", impl.name), func(t *testing.T) {
			defer impl.closeFn()
			for i, test := range tests {

				// clear graph
				_, _, err := impl.graph.Update(matchAll, nil, matchAll)
				if err != nil {
					t.Fatalf("delete {?s ?p ?o}: %v", err)
				}

				// insert
				_, err = impl.graph.Insert(mustTriples(test.input)...)
				if err != nil {
					t.Fatal("insert: %v", err)
				}

				// update
				del, ins, where := mustParseUpdate(test.update)
				_, _, err = impl.graph.Update(del, ins, where)
				if err != nil {
					t.Fatalf("update: %v", err)
				}

				// verify
				got, err := impl.graph.Where(matchAll...)
				if err != nil {
					t.Fatalf("where: %v", err)
				}
				want := mustDecode(test.want)
				if !got.(*memory.Graph).Eq(want) {
					t.Fatalf("#%d\ngot:\n%v\nwant:\n%v\n", i+1, mustEncode(got.(*memory.Graph)), test.want)
				}
			}
		})
	}
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
		if got := a.Eq(b); got != test.eq {
			t.Fatalf("\n%v  Eq\n%v  = %v; want %v", mustEncode(a), mustEncode(b), got, test.eq)
		}
	}
}

var testTriples = `
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

		<c2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Corporation> .
		<c2> <hasName> "Gyldendal" .
		`
var testGraph = mustDecode(testTriples)

func TestGraphWhere(t *testing.T) {
	tests := []struct {
		patterns []rdf.TriplePattern
		want     string
	}{
		{ // 1
			mustParsePatterns(""),
			"",
		},
		{ // 2
			mustParsePatterns("?s ?p ?o ."),
			mustEncode(testGraph),
		},
		{ // 3
			mustParsePatterns("?pub <hasMainTitle> ?title ."),
			`
			<w1> <hasMainTitle> "Le Cosmicomiche" .
			<p1> <hasMainTitle> "The Complete Cosmicomics" .
			<w2> <hasMainTitle> "Il barone rampante" .
			<p2> <hasMainTitle> "Klatrebaronen" .`,
		},
		{ // 4
			mustParsePatterns(`?p2 <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .`),
			`<p2> <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .`,
		},
		{ // 5
			mustParsePatterns("<p2> ?p ?o ."),
			`
			<p2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Publication> .
			<p2> <isPublicationOf> <w2> .
			<p2> <hasMainTitle> "Klatrebaronen" .
			<p2> <hasPublishYear> "1961"^^<http://www.w3.org/2001/XMLSchema#gYear> .
			<p2> <isPublishedBy> <c2> .
			<p2> <hasContributor> _:1 .`,
		},
		{ // 6
			mustParsePatterns(`?s ?p "Penguin" .`),
			`<c1> <hasName> "Penguin" .`,
		},
		{ // 7
			mustParsePatterns(`?s ?p "Penguin"@en .`),
			"",
		},
		{ // 8
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
		{ // 9
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
		{ // 10
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
		{ // 11
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
		{ // 12
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
	impls := newGraphImplementations()

	for _, impl := range impls {
		t.Run(fmt.Sprintf("%s implementation", impl.name), func(t *testing.T) {
			defer impl.closeFn()
			impl.graph.Insert(mustTriples(testTriples)...)
			for i, test := range tests {
				got, err := impl.graph.Where(test.patterns...)
				if err != nil {
					t.Fatalf("%d:got:\n%v\nwant:\n%v", i+1, err, test.want)
				} else if !got.(*memory.Graph).Eq(mustDecode(test.want)) {
					t.Fatalf("%d:got:\n%v\nwant:\n%v", i+1, mustEncode(got.(*memory.Graph)), test.want)
				}
			}
		})
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
	impls := newGraphImplementations()
	for _, impl := range impls {
		t.Run(fmt.Sprintf("%s implementation", impl.name), func(t *testing.T) {
			defer impl.closeFn()
			impl.graph.Insert(mustTriples(testTriples)...)

			for _, test := range tests {
				want := mustParseSolutions(test.want)
				got, _ := impl.graph.Select(test.vars, test.patterns...)
				if !solutionsEq(got, want) {
					t.Fatalf("got:\n%v\nwant:\n%v", got, want)
				}
			}
		})
	}
}

func TestGraphDescribe(t *testing.T) {
	tests := []struct {
		node rdf.NamedNode
		want string
	}{
		{
			rdf.NewNamedNode("a1"),
			`<a1> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Person> .
			 <a1> <hasName> "Italo Calvino" .
			 <a1> <hasBirthYear> "1923"^^<http://www.w3.org/2001/XMLSchema#gYear> .
			 <a1> <hasDeathYear> "1985"^^<http://www.w3.org/2001/XMLSchema#gYear> .`,
		},
		{
			rdf.NewNamedNode("w2"),
			`<w2> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Work> .
			 <w2> <hasMainTitle> "Il barone rampante" .
			 <w2> <hasPublicationYear> "1957"^^<http://www.w3.org/2001/XMLSchema#gYear> .
			 <w2> <hasContributor> _:c5 .

			 _:c5 <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <Contribution> .
			 _:c5 <hasRole> <author> .
			 _:c5 <hasAgent> <a1> .`,
		},
	}
	impls := newGraphImplementations()
	for _, impl := range impls {
		t.Run(fmt.Sprintf("%s implementation", impl.name), func(t *testing.T) {
			defer impl.closeFn()
			impl.graph.Insert(mustTriples(testTriples)...)

			for _, test := range tests {
				want := mustDecode(test.want)
				got, _ := impl.graph.Describe(test.node)
				if !want.Eq(got.(*memory.Graph)) {
					t.Fatalf("got:\n%v\nwant:\n%v", mustEncode(got.(*memory.Graph)), test.want)
				}
			}
		})
	}
}
