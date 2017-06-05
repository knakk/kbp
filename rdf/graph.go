package rdf

// DescribeMode represents an algorithm which Describe method can use.
type DescribeMode int

// Available DescribeModes:
const (
	// DescForward describes a node by including its literal properties and outgoing relations.
	// A blank node relation will be expanded, so that all triples where the blank node is subject
	// will be included.
	DescForward DescribeMode = iota

	// DescForwardRecursive describes a node by including all its literal properties and outgoing relations,
	// and recursivly calls Describe with the same algorithm on the relations.
	DescForwardRecursive

	// DescSymmetricRecursive describes a node first like DescForwardRecursive, then, recursivly for each
	// incomming relation, calls DescForwardRecursive on the subject of the reltion.
	DescSymmetricRecursive

	// TODO:
	// DescBackward
	// DescBackwardRecursive
	// DescSymmetric
)

// TODO consider splitting interface into Graph (read-only) and UpdatableGraph

// Graph represents a RDF Graph - a colllection of RDF triples which
// can be queried for subgraphs and mutated by adding and removing triples.
type Graph interface {

	// Describe returns a Graph which describes the given nodes, according
	// to the DescribeMode algorithm.
	Describe(DescribeMode, ...NamedNode) (Graph, error)

	// DescribeW works like Describe, except that it will encode the Graph
	// using the given Encoder, instead of returning it.
	DescribeW(*Encoder, DescribeMode, ...NamedNode) error

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

	Select([]Variable, ...TriplePattern) (Bindings, error)

	// Stats returns statistics about the Graph.
	Stats() (Stats, error)
}

// Bindings represents the results from a Select query.
type Bindings struct {
	Vars []Variable
	Rows [][]Node
}

// FirstBound returns the first bound value (Node) of the given variable,
// along with a boolean stating if there is one or not.
func (b Bindings) FirstBound(v Variable) (Node, bool) {
	for i, bv := range b.Vars {
		if bv == v {
			for _, row := range b.Rows {
				if row[i] != nil {
					return row[i], true
				}
			}
			return nil, false
		}
	}
	return nil, false
}

// AllBound returns all the bound values (Nodes) for the given variable.
func (b Bindings) AllBound(v Variable) (res []Node) {
	for i, bv := range b.Vars {
		if bv == v {
			for _, row := range b.Rows {
				if row[i] != nil {
					res = append(res, row[i])
				}
			}
			return res
		}
	}
	return res
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
