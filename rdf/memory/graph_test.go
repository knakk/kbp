package memory

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"sort"
	"testing"

	"github.com/go-test/deep"
	"github.com/knakk/kbp/rdf"
)

// TODO move to knakk/kbp/rdf/internal
func TestGroupPatternsByVariable(t *testing.T) {
	tests := []struct {
		patterns []encPattern
		want     [][]encPattern
	}{
		{
			[]encPattern{},
			[][]encPattern{},
		},
		{
			[]encPattern{{-1, -2, -3, 99}},
			[][]encPattern{
				{{-1, -2, -3, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-2, 2, 3, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}},
				{{-2, 2, 3, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-1, 2, 3, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-1, 2, 3, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-1, 2, -2, 99}, {-2, 2, 3, 99}, {-3, 4, 1, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}},
				{{-3, 4, 1, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-1, 2, -2, 99}, {-2, 2, 3, 99}, {-3, 4, -1, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-3, 4, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}, {-3, 4, -1, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-3, 4, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}},
			},
		},
		{
			[]encPattern{{-1, 1, -2, 99}, {-1, 2, -3, 99}, {-3, 3, 4, 99}, {-2, 5, 6, 99}, {-2, 7, -4, 99}, {-4, 8, -5, 99}},
			[][]encPattern{
				{{-4, 8, -5, 99}, {-3, 3, 4, 99}, {-2, 7, -4, 99}, {-2, 5, 6, 99}, {-1, 2, -3, 99}, {-1, 1, -2, 99}},
			},
		},
	}

	for _, test := range tests {
		groups := groupPatternsByVariable(test.patterns)
		for n := range groups {
			sort.Slice(groups[n], func(i, j int) bool {
				return fmt.Sprintf("%d%d%d", groups[n][i][0], groups[n][i][1], groups[n][i][2]) > fmt.Sprintf("%d%d%d", groups[n][j][0], groups[n][j][1], groups[n][j][2])
			})
		}
		if !reflect.DeepEqual(groups, test.want) {
			t.Fatalf("groupPatternsByVariable(%v) => %v; want %v", test.patterns, groups, test.want)
		}
	}
}

func TestDecode(t *testing.T) {
	type contribution struct {
		Role  string `rdf:"->hasRole;->hasLabel"`
		Agent string `rdf:"->hasAgent;->hasName"`
	}
	type binding struct {
		Name string `rdf:"->hasLabel"`
	}
	type publication struct {
		ID            string         `rdf:"id"`
		WorkID        string         `rdf:"->isPublicationOf;id"`
		MainTitle     string         `rdf:"->hasMainTitle"`
		Subtitle      string         `rdf:"->hasSubtitle"`
		Year          int            `rdf:"->wasPublishedYear"`
		Binding       *binding       `rdf:"->hasBinding"`
		Subjects      []string       `rdf:">>hasTag"`
		Genres        []string       `rdf:"->isPublicationOf;>>hasGenre;->hasLabel"`
		WorkSubjects  []string       `rdf:"->isPublicationOf;>>hasTag"`
		Contributions []contribution `rdf:">>hasContribution"`
		Fans          []string       `rdf:"<<isFanOf;->hasName"`
		Similar       *publication   `rdf:"<-isSimilarTo"`
	}

	g := mustDecode(`
		<book1> <hasMainTitle> "Das Kapital" .
		<book1> <hasSubtitle> "Kritik der politischen Ökonomie" .
		<book1> <wasPublishedYear> "1867"^^<http://www.w3.org/2001/XMLSchema#gYear> .
		<book1> <hasContribution> _:c1 .
		<book1> <hasTag> "Politics" .
		<book1> <hasTag> "Economy" .
		<book1> <hasTag> "Work" .
		<book1> <isPublicationOf> <work1> .
		<book1> <hasBinding> <binding1> .
		<binding1> <hasLabel> "Hardcover" .
		<work1> <hasTag> "Politikk" .
		<work1> <hasMainTitle> "Das Kapital" .
		<work1> <hasGenre> <genre1> .
		<genre1> <hasLabel> "Komedie" .
		_:c1 <hasAgent> <person1> .
		_:c1 <hasRole> <author> .
		<book1> <hasContribution> _:c2 .
		_:c2 <hasAgent> <person2> .
		_:c2 <hasRole> <contributor> .
		<person1> <hasName> "Karl Marx" .
		<person2> <hasName> "Friedrich Engels" .
		<person3> <hasName> "Ole" .
		<person3> <isFanOf> <book1> .
		<person1> <isFanOf> <book1> .
		<author> <hasLabel> "forfatter" .
		<contributor> <hasLabel> "bidragsyter" .`)

	var p publication
	if err := g.Decode(&p, rdf.NewNamedNode("book1"), rdf.NewNamedNode("")); err != nil {
		t.Fatal(err)
	}

	want := publication{
		ID:           "book1",
		WorkID:       "work1",
		MainTitle:    "Das Kapital",
		Subtitle:     "Kritik der politischen Ökonomie",
		Year:         1867,
		Subjects:     []string{"Economy", "Politics", "Work"},
		WorkSubjects: []string{"Politikk"},
		Contributions: []contribution{
			{
				Role:  "bidragsyter",
				Agent: "Friedrich Engels",
			},
			{
				Role:  "forfatter",
				Agent: "Karl Marx",
			},
		},
		Fans:    []string{"Karl Marx", "Ole"},
		Genres:  []string{"Komedie"},
		Binding: &binding{Name: "Hardcover"},
		Similar: nil,
	}
	sort.Slice(p.Contributions, func(i, j int) bool {
		return p.Contributions[i].Agent < p.Contributions[j].Agent
	})
	sort.Strings(p.Subjects)
	sort.Strings(p.Fans)
	if diff := deep.Equal(p, want); diff != nil {
		t.Error(diff)
	}
}

func decodeGraph(d *rdf.Decoder) (*Graph, error) {
	g := NewGraph()
	bnodeTriples := make(map[rdf.BlankNode][]rdf.Triple)

	for tr, err := d.Decode(); err != io.EOF; tr, err = d.Decode() {
		if err != nil {
			return g, err
		}
		switch t := tr.Subject.(type) {
		case rdf.BlankNode:
			bnodeTriples[t] = append(bnodeTriples[t], tr)
			continue
		}
		switch t := tr.Object.(type) {
		case rdf.BlankNode:
			bnodeTriples[t] = append(bnodeTriples[t], tr)
			continue
		}
		if _, err := g.Insert(tr); err != nil {
			return nil, err
		}
	}

	// Insert triples with bnodes in batches by ID, so that they get assigned
	// the same (new) blank node ID in the Graph
	for _, trs := range bnodeTriples {
		if _, err := g.Insert(trs...); err != nil {
			return nil, err
		}
	}
	return g, nil
}

func mustDecode(s string) *Graph {
	dec := rdf.NewDecoder(bytes.NewBufferString(s))
	g, err := decodeGraph(dec)
	if err != nil {
		panic("mustDecode: " + err.Error())
	}
	return g
}
