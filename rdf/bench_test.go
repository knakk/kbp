package rdf

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

// dataset returns a Graph from an n-triples dataset.
// The available datasets in testdata folder are:
//
// - small (119 triples)
//   A book catalog entry from the gutenberg.org collection.
//   http://www.gutenberg.org/ebooks/39110.rdf
//
// - medium (49905 triples)
//   Museums in italy
//   http://www.linkedopendata.it/datasets/musei
//
// - large (726674 triples)
//   Lexvo.org data dump.
//   http://www.lexvo.org/resources/lexvo_latest.rdf.gz
//
// - lubm (99567 triples)
//   Lehigh University Benchmark, generated with -index 0 -seed 0
//   http://swat.cse.lehigh.edu/projects/lubm/
func dataset(set string) *Graph {
	if g, ok := datasets[set]; ok {
		return g
	}

	containsBnodes := true
	switch set {
	case "small", "medium":
	case "large", "lubm":
		containsBnodes = false
	default:
		panic("dataset: set must be one of small, medium, large, lubm")
	}

	f, err := os.Open("testdata/" + set + ".nt.gz")
	if err != nil {
		panic("dataset: " + err.Error())
	}
	defer f.Close()
	r, err := gzip.NewReader(f)
	if err != nil {
		panic("dataset: " + err.Error())
	}
	defer r.Close()
	var triples []Triple
	g := NewGraph()
	dec := NewDecoder(r)
	for tr, err := dec.Decode(); err != io.EOF; tr, err = dec.Decode() {
		if err != nil {
			panic("dataset: " + err.Error())
		}
		if containsBnodes {
			triples = append(triples, tr)
		} else {
			g.Insert(tr)
		}
	}
	if containsBnodes {
		g.Insert(triples...)
	}
	datasets[set] = g
	return datasets[set]
}

var datasets = make(map[string]*Graph)

func datasetAsByte(b *testing.B, dataset string) []byte {
	f, err := os.Open("testdata/" + dataset + ".nt.gz")
	if err != nil {
		b.Fatal(err)
	}
	defer f.Close()
	r, err := gzip.NewReader(f)
	if err != nil {
		b.Fatal(err)
	}
	defer r.Close()
	bset, err := ioutil.ReadAll(r)
	if err != nil {
		b.Fatal(err)
	}
	return bset
}

func BenchmarkScanner(b *testing.B) {
	var tok token
	small := datasetAsByte(b, "small")
	for n := 0; n < b.N; n++ {
		s := newScanner(bytes.NewBuffer(small))
		for {
			tok = s.Scan()
			if tok.Type == tokenEOF {
				break
			}
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	small := datasetAsByte(b, "small")
	var triples []Triple
	for n := 0; n < b.N; n++ {
		dec := NewDecoder(bytes.NewBuffer(small))
		for tr, err := dec.Decode(); err != io.EOF; tr, err = dec.Decode() {
			if err != nil {
				b.Fatal(err)
			}
			triples = append(triples, tr)
		}
	}
}

func BenchmarkGraphInsert(b *testing.B) {
	dataset("small") // ensure it is parsed and cached in datasets
	for n := 0; n < b.N; n++ {
		g := NewGraph()
		g.Insert(dataset("small").Triples()...)
	}
}

func BenchmarkGraphWhereSmall(b *testing.B) {
	benchmarks := []struct {
		name     string
		patterns []TriplePattern
	}{
		{"match all", mustParsePatterns("?s ?p ?o .")},
		{"match subject", mustParsePatterns("<http://www.gutenberg.org/ebooks/39110> ?p ?o .")},
		{"match object", mustParsePatterns("?s ?p <http://purl.org/dc/terms/IMT> .")},
		{"match predicate", mustParsePatterns("?s <http://purl.org/dc/terms/modified> ?o .")},
		{"match stored triple", mustParsePatterns(`<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/extent> "492484"^^<http://www.w3.org/2001/XMLSchema#integer> .`)},
		{"match missing triple", mustParsePatterns(`<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/extent> "-492484"^^<http://www.w3.org/2001/XMLSchema#integer> .`)},
	}
	var results *Graph
	g := dataset("small")
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				results = g.Where(bm.patterns...)
			}
		})
	}
}

func BenchmarkGraphEqSmall(b *testing.B) {
	g := dataset("small")
	for n := 0; n < b.N; n++ {
		g.Eq(g)
	}
}

func BenchmarkGraphEqMedium(b *testing.B) {
	g := dataset("medium")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g.Eq(g)
	}
}

func BenchmarkGraphEqLarge(b *testing.B) {
	g := dataset("large")
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g.Eq(g)
	}
}
