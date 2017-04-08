package rdf

import "errors"

// Exported errors that Graph methods can return.
var (
	ErrNotFound = errors.New("not found")
)

// Graph represents a RDF Graph - a colllection of RDF triples which
// can be queried for subgraphs and mutated by adding and removing triples.
type Graph interface {

	// Insert adds one or more triples in to the Graph, returning the
	// number of triples added that where not already present, or an error.
	Insert(...Triple) (int, error)

	// Delete removes one or more triples from the Graph, returning the
	// number of triples succesfully removed, or an error.
	Delete(...Triple) (int, error)

	Where(...TriplePattern) (Graph, error)
	//Construct([]TriplePattern, ...TriplePattern) (Graph, error)

	Select([]Variable, ...TriplePattern) ([][]Node, error)

	// Eq tests if two graphs are equal (isomorphic).
	Eq(Graph) (bool, error)

	// Stats returns statistics about the Graph.
	Stats() (Stats, error)
}

// Stats represents statistics about a graph.
type Stats struct {
	// Counts
	NumNodes      int
	NumNamedNodes int
	NumLiterals   int
	NumBlankNodes int
	NumSubjects   int
	NumPredicates int
	NumObjects    int
	NumTriples    int
	// Distributions
	Predicates map[string]int
}
