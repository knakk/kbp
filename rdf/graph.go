package rdf

import "errors"

// Exported errors that Graph methods can return.
var (
	ErrNotFound = errors.New("not found")
)

// Graph represents a RDF Graph - a colllection of RDF triples which
// can be queried for subgraphs and mutated by adding and removing triples.
type Graph interface {
	// Describe returns a Graph containing the data about the given nodes.
	// Specifically, it should contain all the triples where a given node is subject,
	// as well all triples of blank nodes which are outgoing relation of the node.
	Describe(...NamedNode) (Graph, error)

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
	// Update returns the number of triples deleted, inserted, and an error should it occur.
	Update(del []TriplePattern, ins []TriplePattern, where []TriplePattern) (int, int, error)

	// Insert inserts the given triples into the Graph, reutrning the number of triples
	// inserted, and an error should it occur.
	Insert(ins ...Triple) (int, error)

	// Delete deletes the given triples into the Graph, reutrning the number of triples
	// deleted, and an error should it occur.
	Delete(ins ...Triple) (int, error)

	Where(...TriplePattern) (Graph, error)
	//Construct([]TriplePattern, ...TriplePattern) (Graph, error)

	Select([]Variable, ...TriplePattern) ([][]Node, error)

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
