package rdf

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/knakk/kbp/rdf"
	"github.com/knakk/kbp/rdf/memory"
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

func dataset(set string) *memory.Graph {
	if g, ok := datasets[set]; ok {
		return g
	}

	containsBnodes := true
	switch set {
	case "small", "medium":
	case "large":
		containsBnodes = false
	default:
		panic("dataset: set must be one of small, medium, large")
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
	var triples []rdf.Triple
	g := memory.NewGraph()
	dec := rdf.NewDecoder(r)
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

var datasets = make(map[string]*memory.Graph)

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

func BenchmarkDecode(b *testing.B) {
	small := datasetAsByte(b, "small")
	var triples []rdf.Triple
	for n := 0; n < b.N; n++ {
		dec := rdf.NewDecoder(bytes.NewBuffer(small))
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
		g := memory.NewGraph()
		g.Insert(dataset("small").Triples()...)
	}
}

func BenchmarkGraphWhereSmall(b *testing.B) {
	benchmarks := []struct {
		name     string
		patterns []rdf.TriplePattern
	}{
		{"match all", mustParsePatterns("?s ?p ?o .")},
		{"match S", mustParsePatterns("<http://www.gutenberg.org/ebooks/39110> ?p ?o .")},
		{"match O", mustParsePatterns("?s ?p <http://purl.org/dc/terms/IMT> .")},
		{"match P", mustParsePatterns("?s <http://purl.org/dc/terms/modified> ?o .")},
		{"match SPO (existing)", mustParsePatterns(`<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/extent> "492484"^^<http://www.w3.org/2001/XMLSchema#integer> .`)},
		{"match SPO (missing)", mustParsePatterns(`<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/extent> "-492484"^^<http://www.w3.org/2001/XMLSchema#integer> .`)},
		{"match SP", mustParsePatterns("<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> ?o .")},
		{"match PO", mustParsePatterns(`?s <http://purl.org/dc/terms/publisher> "Project Gutenberg" .`)},
		{"match SO", mustParsePatterns(`<http://www.gutenberg.org/ebooks/39110> ?p "Project Gutenberg" .`)},
		{"complex", mustParsePatterns(`
				?work <hasMainTitle> "Le Cosmicomiche" .
				?book <isPublicationOf> ?work.
				?book <hasContributor> ?contrib .
				?contrib <hasRole> <translator> .
				?contrib <hasAgent> ?person .
				?person <hasName> ?translator .`,
		)},
	}
	var results rdf.Graph
	g := dataset("small")
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				results, _ = g.Where(bm.patterns...)
			}
		})
	}
}

func BenchmarkGraphWhereMedium(b *testing.B) {
	benchmarks := []struct {
		name     string
		patterns []rdf.TriplePattern
	}{
		{"match all", mustParsePatterns("?s ?p ?o .")},
		{"match S", mustParsePatterns("<http://data.linkedopendata.it/musei/resource/Museo_della_Permanente-Milano> ?p ?o .")},
		{"match O", mustParsePatterns("?s ?p <http://www.w3.org/2006/vcard/ns#Tel> .")},
		{"match P", mustParsePatterns("?s <http://www.w3.org/2006/vcard/ns#geo> ?o .")},
		{"match SPO (existing)", mustParsePatterns(`<http://data.linkedopendata.it/musei/resource/Museo_della_Regina-Cattolica> <http://www.w3.org/2006/vcard/ns#url> <http://www.cattolica.net> .`)},
		{"match SPO (missing)", mustParsePatterns(`<http://data.linkedopendata.it/musei/resource/Museo_della_Regina-Cattolica> <http://www.w3.org/2006/vcard/ns#url> <http://xyz.cattolica.net>  .`)},
		{"match SP", mustParsePatterns("<http://data.linkedopendata.it/musei/resource/Aboca_Museum-Sansepolcro> <http://www.w3.org/2000/01/rdf-schema#label> ?o .")},
		{"match PO", mustParsePatterns(`?s <http://www.w3.org/2006/vcard/ns#latitude> "43.9659588" .`)},
		{"match SO", mustParsePatterns(`<http://data.linkedopendata.it/musei/resource/Cuglieri> ?p <http://sws.geonames.org/3177719/about.rdf>  .`)},
		{"complex", mustParsePatterns(`
			?loc <http://www.w3.org/2006/vcard/ns#latitude> "45.6030908" .
			?mus <http://www.w3.org/2006/vcard/ns#geo> ?loc .
			?mus <http://www.w3.org/2000/01/rdf-schema#label> ?label .
			?mus <http://www.w3.org/2004/02/skos/core#subject> <http://dbpedia.org/resource/Category:Art_museums_and_galleries> .`,
		)},
	}
	var results rdf.Graph
	g := dataset("medium")
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				results, _ = g.Where(bm.patterns...)
			}
		})
	}
}

func BenchmarkGraphWhereLarge(b *testing.B) {
	benchmarks := []struct {
		name     string
		patterns []rdf.TriplePattern
	}{
		//{"match all", mustParsePatterns("?s ?p ?o .")},
		{"match S", mustParsePatterns("<http://lexvo.org/id/iso639-3/eng> ?p ?o .")},
		{"match O", mustParsePatterns("?s ?p <http://lexvo.org/id/iso639-3/eng> .")},
		{"match P", mustParsePatterns("?s <http://lexvo.org/ontology#nearlySameAs> ?o .")},
		{"match SPO (existing)", mustParsePatterns(`<http://lexvo.org/id/iso639-3/eng> <http://www.w3.org/2000/01/rdf-schema#label> "Engliš"@wae .`)},
		{"match SPO (missing)", mustParsePatterns(`<http://lexvo.org/id/iso639-3/eng> <http://www.w3.org/2000/01/rdf-schema#label> "Engliš"@en .`)},
		{"match SP", mustParsePatterns("<http://lexvo.org/id/iso639-3/eng> <http://www.w3.org/2000/01/rdf-schema#label> ?o .")},
		{"match PO", mustParsePatterns(`?s <http://www.w3.org/2002/07/owl#sameAs> <http://lexvo.org/id/iso639-3/nor> .`)},
		{"match SO", mustParsePatterns(`<http://lexvo.org/id/iso639-3/nor> ?p "ኖርዌጂያን"@tig .`)},
	}
	var results rdf.Graph
	g := dataset("large")
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				results, _ = g.Where(bm.patterns...)
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
