package sparql

import (
	"encoding/json"
	"io"
	"time"

	"github.com/knakk/kbp/rdf"
)

// DateFormat is the expected layout of the xsd:DateTime values. You can override
// it if your triple store uses a different layout.
var DateFormat = time.RFC3339

// Results holds the parsed results of a application/sparql-results+json response.
type Results struct {
	Head    header
	Results results
}

type header struct {
	Link []string
	Vars []string
}

type results struct {
	Distinct bool
	Ordered  bool
	Bindings []map[string]binding
}

type binding struct {
	Type     string // "uri", "literal", "typed-literal" or "bnode"
	Value    string
	Lang     string `json:"xml:lang"`
	DataType string
}

// ParseJSON takes an application/sparql-results+json response and parses it
// into a Results struct.
func ParseJSON(r io.Reader) (*Results, error) {
	var res Results
	err := json.NewDecoder(r).Decode(&res)

	return &res, err
}

// Bindings returns a map of the bound variables in the SPARQL response, where
// each variable points to one or more RDF terms.
func (r *Results) Bindings() map[string][]rdf.Node {
	rb := make(map[string][]rdf.Node)
	for _, v := range r.Head.Vars {
		for _, b := range r.Results.Bindings {
			if t := termFromJSON(b[v]); t != nil {
				rb[v] = append(rb[v], t)
			}
		}
	}

	return rb
}

// Solutions returns a slice of the query solutions, each containing a map
// of all bindings to RDF terms.
func (r *Results) Solutions() []map[string]rdf.Node {
	var rs []map[string]rdf.Node

	for _, s := range r.Results.Bindings {
		solution := make(map[string]rdf.Node)
		for k, v := range s {
			if t := termFromJSON(v); t != nil {
				solution[k] = t
			}
		}
		rs = append(rs, solution)
	}

	return rs
}

// termFromJSON converts a SPARQL json result binding into a rdf.Node. Any
// parsing errors on typed-literal will result in a xsd:string-typed RDF term.
func termFromJSON(b binding) rdf.Node {
	switch b.Type {
	case "bnode":
		return rdf.NewBlankNode(b.Value)
	case "uri":
		return rdf.NewNamedNode(b.Value)
	case "literal":
		// Untyped literals are typed as xsd:string
		if b.Lang != "" {
			return rdf.NewLangLiteral(b.Value, b.Lang)
		}
		return rdf.NewStrLiteral(b.Value)
	case "typed-literal":
		return rdf.NewTypedLiteral(b.Value, rdf.NewNamedNode(b.DataType))
	default:
		//panic("unknown term type")
		return nil
	}
}
