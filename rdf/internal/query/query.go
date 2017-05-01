package query

import "github.com/RoaringBitmap/roaring"

const (
	MaxVariables = 99
)

// EncTriplePattern is a TriplePattern where the nodes are encoded as uint32 IDs.
//
// The node IDs for subject, predicate and object nodes occupy the first 3 slots
// in the array, and the last represents the estimated cardinality for the pattern.
//
// Some node ID numbers have special meanings:
// 0    Node is not in the graph, so the pattern will never match any triples,
//      and neither will any of the patterns sharing variables with it.
// 1-99 The numbers up to 99 (up to MaxVariables) is reserved for variables.
// 100- Node ID from the Graphs node dictionary.
type EncTriplePattern [4]uint32

// IsConcrete returns true if the pattern has no variables.
func (ep EncTriplePattern) IsConcrete() bool {
	return ep[0] > MaxVariables && ep[1] > MaxVariables && ep[2] > MaxVariables
}

// GroupPatterns does two things: 1) separates a slice of EncTriplePatterns into
// groups where patterns share variables, either direct or transitive, and 2)
// prunes any groups with non-existing nodes (identified by node ID 0).
func GroupPatterns(patterns []EncTriplePattern) (res [][]EncTriplePattern) {

	// Fast path for no patterns
	if len(patterns) == 0 {
		return res
	}

	// Fast path for single pattern
	if len(patterns) == 1 {
		if patterns[0][3] == 0 {
			// 0 in estimated cardinality means it has non-existing node
			return res
		}
		res = append(res, patterns)
		return res
	}

	// index of variable to group number
	v2g := make(map[uint32]int)

	for len(patterns) > 0 {
		// Is the pattern concrete?
		if patterns[0][0] > MaxVariables &&
			patterns[0][1] > MaxVariables &&
			patterns[0][2] > MaxVariables {
			// Pattern has no variables, so it belong to its own group.
			res = append(res, []EncTriplePattern{patterns[0]})
			patterns = patterns[1:]
			continue
		}

		// Index the first pattern. Either, it's the first pattern with
		// variables, or it's a pattern which does not share variables
		// with the patterns indexed so far.
		res = append(res, []EncTriplePattern{patterns[0]})
		indexVariables(patterns[0], len(res), v2g)
		patterns = patterns[1:]

	again:
		indexedOne := false
		for i := 0; i < len(patterns); i++ {
			for _, id := range patterns[i] {
				if id > 0 && id < MaxVariables {
					if group, ok := v2g[id]; ok {
						indexVariables(patterns[i], group, v2g)
						res[group-1] = append(res[group-1], patterns[i])
						patterns = append(patterns[:i], patterns[i+1:]...)
						i--
						indexedOne = true
						break
					}
				}
			}
		}
		if indexedOne {
			// Loop again until there are no patterns with indexed variables
			goto again
		}

	}

	// Remove any groups containing non-exsiting nodes.
pruneGroups:
	for i := 0; i < len(res); i++ {
		for _, pattern := range res[i] {
			for _, id := range pattern[:3] {
				if id == 0 {
					res = append(res[:i], res[i+1:]...)
					i--
					continue pruneGroups
				}
			}
		}
	}

	return res
}

func indexVariables(p EncTriplePattern, group int, v2g map[uint32]int) {
	for _, id := range p {
		if id > 0 && id < MaxVariables {
			v2g[id] = group
		}
	}
}

type EncSolutions struct {
	Vars []uint32
	Rows [][]uint32
}

// compatible tests if two solutions rows are compatible, meaning that
// shared variables are bound to the same nodes. Rows without shared
// variables are always compatible.
func compatible(pairs [][2]int, rowa, rowb []uint32) bool {
	for _, p := range pairs {
		if rowa[p[0]] != rowb[p[1]] {
			return false
		}
	}
	return true
}

func includeNode(nodes []uint32, find uint32) bool {
	for _, node := range nodes {
		if node == find {
			return true
		}
	}
	return false
}

func (s EncSolutions) Join(other EncSolutions, bound map[uint32]*roaring.Bitmap) EncSolutions {
	// bound map must not be nil

	if len(s.Vars) == 0 {
		// We're joining with the empty solution set
		return other
	}

	if len(other.Vars) == 0 {
		return s
	}

	var pairs [][2]int // the indices of shared variables in a & b
	for i, va := range s.Vars {
		for j, vb := range other.Vars {
			if va == vb {
				pairs = append(pairs, [2]int{i, j})
			}
		}
	}
	var addVars []int
	for i, v := range other.Vars {
		if !includeNode(s.Vars, v) {
			s.Vars = append(s.Vars, v)
			addVars = append(addVars, i)
		}
	}

	var out [][]uint32
	for _, a := range s.Rows {
		for _, b := range other.Rows {
			if compatible(pairs, a, b) {
				row := make([]uint32, 0, len(a)+len(addVars))
				row = append(row, a...)
				for _, v := range addVars {
					row = append(row, b[v])
				}
				out = append(out, row)
			} else {
				// TODO remove from bound
			}
		}
	}

	s.Rows = out
	return s
}

func (s EncSolutions) Project(vars []uint32) EncSolutions {
	rows := make([][]uint32, 0, len(s.Rows))
	var svars []int
	for _, v := range vars {
		for i, vs := range s.Vars {
			if v == vs {
				svars = append(svars, i)
			}
		}
	}
	for _, row := range s.Rows {
		newrow := make([]uint32, 0, len(svars))
		for _, i := range svars {
			newrow = append(newrow, row[i])
		}
		rows = append(rows, newrow)
	}
	return EncSolutions{
		Vars: vars,
		Rows: rows,
	}
}

//func (s EncSolutions) OrderBy(v uint32) {}
//func (s EncSolutions) Project(vs uint32...) {}
//func (s EncSolutions) Distinct() {}

type Operator int

const (
	OpBGP      Operator = iota
	OpJoin              // SPARQL: .
	OpLeftJoin          // SPARQL: OPTIONAL
	OpFilter
	OpUnion
	OpDiff // SPARQL: MINUS
)
