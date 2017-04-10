package rdf

import "errors"

// Exported errors that Graph methods can return.
var (
	ErrNotFound = errors.New("not found")
)

// Graph represents a RDF Graph - a colllection of RDF triples which
// can be queried for subgraphs and mutated by adding and removing triples.
type Graph interface {

	// Update performs an update query which inserts and/or deletes triples
	// from the Graph.
	//
	// The bindings for all solutions matching the where patterns will be substituted
	// with the variables in the delete patterns to remove triples, and then in the
	// insert patterns to insert triples, in that order.
	//
	// If the insert patterns are nil, only the delete operation is performed, or if
	// the delete patterns are nil, only the insert opreation is performed.
	//
	// If all inserte and delete patterns are concrete (ie. contains no variables), the
	// where patterns will be ignored.
	//
	// Update returns the number of triples deleted, inserted, or an error should it occur.
	Update(del []TriplePattern, ins []TriplePattern, where []TriplePattern) (int, int, error)

	Where(...TriplePattern) (Graph, error)
	//Construct([]TriplePattern, ...TriplePattern) (Graph, error)

	Select([]Variable, ...TriplePattern) ([][]Node, error)

	// Eq tests if two graphs are equal (isomorphic).
	Eq(Graph) (bool, error)

	// Stats returns statistics about the Graph.
	Stats() (Stats, error)
}

// Insert is a convenience function which performs an update query which only inserts
// triples.
// TODO consider add to Graph interface
func Insert(g Graph, trs ...Triple) (int, error) {
	patterns := make([]TriplePattern, len(trs))
	for i, t := range trs {
		patterns[i] = t.ToTriplePattern()
	}
	_, n, err := g.Update(nil, patterns, nil)
	return n, err
}

// Delete is a convenience function which performs an update query which only deletes
// triples.
// TODO consider add to Graph interface
func Delete(g Graph, trs ...Triple) (int, error) {
	patterns := make([]TriplePattern, len(trs))
	for i, t := range trs {
		patterns[i] = t.ToTriplePattern()
	}
	n, _, err := g.Update(patterns, nil, nil)
	return n, err
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
