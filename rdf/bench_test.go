package rdf

import (
	"bytes"
	"io"
	"testing"
)

const ebook = `<http://www.gutenberg.org/ebooks/39110> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.gutenberg.org/2009/pgterms/ebook> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/rights> "Public domain in the USA." .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/subject> _:N43534d4a9c6d488ba6aed0828a451c3f .
_:N43534d4a9c6d488ba6aed0828a451c3f <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/LCSH> .
_:N43534d4a9c6d488ba6aed0828a451c3f <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "Caricatures and cartoons -- Great Britain" .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/ebooks/39110.kindle.noimages> .
<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/modified> "2015-11-20T05:00:16.163417"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/format> _:Nb62fbc49752e4d9698fc5a8f48f3712f .
_:Nb62fbc49752e4d9698fc5a8f48f3712f <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/x-mobipocket-ebook"^^<http://purl.org/dc/terms/IMT> .
_:Nb62fbc49752e4d9698fc5a8f48f3712f <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110.kindle.noimages> <http://purl.org/dc/terms/extent> "492484"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/ebooks/39110.epub.images> .
<http://www.gutenberg.org/ebooks/39110.epub.images> <http://purl.org/dc/terms/extent> "3503547"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110.epub.images> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/ebooks/39110.epub.images> <http://purl.org/dc/terms/modified> "2015-11-20T04:59:43.613830"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110.epub.images> <http://purl.org/dc/terms/format> _:N5796c62c589a4921b783b1f2eb27993b .
_:N5796c62c589a4921b783b1f2eb27993b <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/epub+zip"^^<http://purl.org/dc/terms/IMT> .
_:N5796c62c589a4921b783b1f2eb27993b <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/type> _:Nc6618f74de634115af3c445f0c099866 .
_:Nc6618f74de634115af3c445f0c099866 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/DCMIType> .
_:Nc6618f74de634115af3c445f0c099866 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "Text" .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/files/39110/39110.txt> .
<http://www.gutenberg.org/files/39110/39110.txt> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/files/39110/39110.txt> <http://purl.org/dc/terms/extent> "27883"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/files/39110/39110.txt> <http://purl.org/dc/terms/modified> "2012-03-11T18:56:44"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/files/39110/39110.txt> <http://purl.org/dc/terms/format> _:N1208b828e87f4221bb7e0f1b0d948d4d .
_:N1208b828e87f4221bb7e0f1b0d948d4d <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "text/plain; charset=us-ascii"^^<http://purl.org/dc/terms/IMT> .
_:N1208b828e87f4221bb7e0f1b0d948d4d <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/language> _:Nb4eefe6de8bd4da3ba0f6d183d20b69d .
_:Nb4eefe6de8bd4da3ba0f6d183d20b69d <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "en"^^<http://purl.org/dc/terms/RFC4646> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/title> "Phil May's Gutter-Snipes: 50 Original Sketches in Pen & Ink" .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/issued> "2012-03-12"^^<http://www.w3.org/2001/XMLSchema#date> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/publisher> "Project Gutenberg" .
<http://www.gutenberg.org/ebooks/39110> <http://www.gutenberg.org/2009/pgterms/marc901> "file:///public/vhost/g/gutenberg/html/files/39110/39110-h/images/th_cover.png" .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/subject> _:N1e27993323d943e5acd204754c7b3b6b .
_:N1e27993323d943e5acd204754c7b3b6b <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/LCC> .
_:N1e27993323d943e5acd204754c7b3b6b <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "NC" .
<http://www.gutenberg.org/ebooks/39110> <http://id.loc.gov/vocabulary/relators/ill> <2009/agents/37403> .
<2009/agents/37403> <http://www.gutenberg.org/2009/pgterms/alias> "Summers, Charlie" .
<2009/agents/37403> <http://www.gutenberg.org/2009/pgterms/name> "May, Phil" .
<2009/agents/37403> <http://www.gutenberg.org/2009/pgterms/webpage> <http://en.wikipedia.org/wiki/Phil_May_%28caricaturist%29> .
<2009/agents/37403> <http://www.gutenberg.org/2009/pgterms/birthdate> "1864"^^<http://www.w3.org/2001/XMLSchema#integer> .
<2009/agents/37403> <http://www.gutenberg.org/2009/pgterms/deathdate> "1903"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/ebooks/39110.rdf> .
<http://www.gutenberg.org/ebooks/39110.rdf> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/ebooks/39110.rdf> <http://purl.org/dc/terms/format> _:Ne578ddf736b4492bb33af6323d39c1cd .
_:Ne578ddf736b4492bb33af6323d39c1cd <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:Ne578ddf736b4492bb33af6323d39c1cd <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/rdf+xml"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110.rdf> <http://purl.org/dc/terms/extent> "13053"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110.rdf> <http://purl.org/dc/terms/modified> "2015-12-03T07:23:35.357006"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/ebooks/39110.txt.utf-8> .
<http://www.gutenberg.org/ebooks/39110.txt.utf-8> <http://purl.org/dc/terms/format> _:N542eeb0cee7948e0b63b719894948d12 .
_:N542eeb0cee7948e0b63b719894948d12 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:N542eeb0cee7948e0b63b719894948d12 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "text/plain"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110.txt.utf-8> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/ebooks/39110.txt.utf-8> <http://purl.org/dc/terms/extent> "27855"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110.txt.utf-8> <http://purl.org/dc/terms/modified> "2015-11-20T04:59:37.097112"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/files/39110/39110-h/39110-h.htm> .
<http://www.gutenberg.org/files/39110/39110-h/39110-h.htm> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/files/39110/39110-h/39110-h.htm> <http://purl.org/dc/terms/format> _:N5be9985c2b314f00962a30301c053a55 .
_:N5be9985c2b314f00962a30301c053a55 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:N5be9985c2b314f00962a30301c053a55 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "text/html; charset=iso-8859-1"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/files/39110/39110-h/39110-h.htm> <http://purl.org/dc/terms/extent> "48819"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/files/39110/39110-h/39110-h.htm> <http://purl.org/dc/terms/modified> "2012-03-11T18:56:44"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/ebooks/39110.epub.noimages> .
<http://www.gutenberg.org/ebooks/39110.epub.noimages> <http://purl.org/dc/terms/modified> "2015-11-20T04:59:44.517805"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110.epub.noimages> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/ebooks/39110.epub.noimages> <http://purl.org/dc/terms/format> _:N7aec5ee19ebc42649d89987c05cf71a2 .
_:N7aec5ee19ebc42649d89987c05cf71a2 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:N7aec5ee19ebc42649d89987c05cf71a2 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/epub+zip"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110.epub.noimages> <http://purl.org/dc/terms/extent> "174391"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/license> <http://www.gutenberg.org/license> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/ebooks/39110.kindle.images> .
<http://www.gutenberg.org/ebooks/39110.kindle.images> <http://purl.org/dc/terms/format> _:Ne9140760cca2480980624ccbd599d79d .
_:Ne9140760cca2480980624ccbd599d79d <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:Ne9140760cca2480980624ccbd599d79d <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/x-mobipocket-ebook"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/ebooks/39110.kindle.images> <http://purl.org/dc/terms/modified> "2015-11-20T05:00:14.723504"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110.kindle.images> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/ebooks/39110.kindle.images> <http://purl.org/dc/terms/extent> "7514120"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110> <http://www.gutenberg.org/2009/pgterms/downloads> "15"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/files/39110/39110.zip> .
<http://www.gutenberg.org/files/39110/39110.zip> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/files/39110/39110.zip> <http://purl.org/dc/terms/extent> "9350"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/files/39110/39110.zip> <http://purl.org/dc/terms/format> _:N72a9817d16df40778dbdcb7720924f11 .
_:N72a9817d16df40778dbdcb7720924f11 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "text/plain; charset=us-ascii"^^<http://purl.org/dc/terms/IMT> .
_:N72a9817d16df40778dbdcb7720924f11 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/files/39110/39110.zip> <http://purl.org/dc/terms/format> _:Nb7e798e365d145188849dc359cda661f .
_:Nb7e798e365d145188849dc359cda661f <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/zip"^^<http://purl.org/dc/terms/IMT> .
_:Nb7e798e365d145188849dc359cda661f <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/files/39110/39110.zip> <http://purl.org/dc/terms/modified> "2012-03-11T18:57:02"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/files/39110/39110-h.zip> .
<http://www.gutenberg.org/files/39110/39110-h.zip> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/files/39110/39110-h.zip> <http://purl.org/dc/terms/format> _:N31f1318054c44b0a893ee51f6e13ff40 .
_:N31f1318054c44b0a893ee51f6e13ff40 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:N31f1318054c44b0a893ee51f6e13ff40 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "text/html; charset=iso-8859-1"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/files/39110/39110-h.zip> <http://purl.org/dc/terms/format> _:Nb54fa037df8a42f89e985bd326845601 .
_:Nb54fa037df8a42f89e985bd326845601 <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:Nb54fa037df8a42f89e985bd326845601 <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "application/zip"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/files/39110/39110-h.zip> <http://purl.org/dc/terms/modified> "2012-03-11T18:57:04"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/files/39110/39110-h.zip> <http://purl.org/dc/terms/extent> "13813773"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/cache/epub/39110/pg39110.cover.small.jpg> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.small.jpg> <http://purl.org/dc/terms/extent> "2824"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.small.jpg> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.small.jpg> <http://purl.org/dc/terms/format> _:N7dd5520cdd88406588254789487daa0d .
_:N7dd5520cdd88406588254789487daa0d <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
_:N7dd5520cdd88406588254789487daa0d <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "image/jpeg"^^<http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.small.jpg> <http://purl.org/dc/terms/modified> "2015-11-20T05:00:16.270411"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/ebooks/39110> <http://purl.org/dc/terms/hasFormat> <http://www.gutenberg.org/cache/epub/39110/pg39110.cover.medium.jpg> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.medium.jpg> <http://purl.org/dc/terms/extent> "17289"^^<http://www.w3.org/2001/XMLSchema#integer> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.medium.jpg> <http://purl.org/dc/terms/isFormatOf> <http://www.gutenberg.org/ebooks/39110> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.medium.jpg> <http://purl.org/dc/terms/format> _:N5e6c495623354b508380a673a9a0fb6c .
_:N5e6c495623354b508380a673a9a0fb6c <http://www.w3.org/1999/02/22-rdf-syntax-ns#value> "image/jpeg"^^<http://purl.org/dc/terms/IMT> .
_:N5e6c495623354b508380a673a9a0fb6c <http://purl.org/dc/dcam/memberOf> <http://purl.org/dc/terms/IMT> .
<http://www.gutenberg.org/cache/epub/39110/pg39110.cover.medium.jpg> <http://purl.org/dc/terms/modified> "2015-11-20T05:00:16.365417"^^<http://www.w3.org/2001/XMLSchema#dateTime> .
<http://www.gutenberg.org/> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://web.resource.org/cc/Work> .
<http://www.gutenberg.org/> <http://web.resource.org/cc/license> <http://www.gnu.org/licenses/gpl.html> .
<http://www.gutenberg.org/> <http://www.w3.org/2000/01/rdf-schema#comment> "Archives containing the RDF files for *all* our books can be downloaded at\n            http://www.gutenberg.org/wiki/Gutenberg:Feeds#The_Complete_Project_Gutenberg_Catalog" .
<http://en.wikipedia.org/wiki/Phil_May_%28caricaturist%29> <http://purl.org/dc/terms/description> "en.wikipedia" .`

func BenchmarkScanner(b *testing.B) {
	var tok token
	for n := 0; n < b.N; n++ {
		s := newScanner(bytes.NewBufferString(ebook))
		for {
			tok = s.Scan()
			if tok.Type == tokenEOF {
				break
			}
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	var triples []Triple
	for n := 0; n < b.N; n++ {
		dec := NewDecoder(bytes.NewBufferString(ebook))
		for tr, err := dec.Decode(); err != io.EOF; tr, err = dec.Decode() {
			if err != nil {
				b.Fatal(err)
			}
			triples = append(triples, tr)
		}
	}
}

func BenchmarkGraphInsert(b *testing.B) {
	triples := mustTriples(ebook)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		g := NewGraph()
		g.Insert(triples...)
	}
}

func BenchmarkGraphWhere(b *testing.B) {
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
	g := mustDecode(ebook)
	var results *Graph
	b.ResetTimer()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				results = g.Where(bm.patterns...)
			}
		})
	}
}

func BenchmarkGraphEq(b *testing.B) {
	graph1 := mustDecode(ebook)
	graph2 := mustDecode(ebook)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		graph1.Eq(graph2)
	}
}
