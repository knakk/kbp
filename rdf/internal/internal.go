package internal

const (
	maxVariables = 99
)

// EncTriplePattern is a TriplePattern where the nodes are encoded as uint32.
//
// The nodes are encoded as this:
// 0    Node is not in the graph, so the pattern will never match any triples,
//      and neither will any of the patterns sharing variables with it.
// 1-99 The numbers up to 99 (up to maxVariables constant) is reserved for variables.
// 100- Node ID from the Graphs node dictionary.
type EncTriplePattern [3]uint32

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
		res = append(res, patterns)
		return res
	}

	// index of variable to group number
	v2g := make(map[uint32]int)

	for len(patterns) > 0 {
		// Is the pattern concrete?
		if patterns[0][0] > maxVariables &&
			patterns[0][1] > maxVariables &&
			patterns[0][2] > maxVariables {
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
				if id > 0 && id < maxVariables {
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
			for _, id := range pattern {
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
		if id > 0 && id < maxVariables {
			v2g[id] = group
		}
	}
}

type EncSolutions struct {
	Vars []uint32
	Rows [][]uint32
}

//func (s EncSolutions) OrderBy(v uint32) {}
//func (s EncSolutions) Project(vs uint32...) {}
//func (s EncSolutions) Distinct() {}

type Operator int

const (
	OpBGP Operator = iota
	OpJoin
	OpLeftJoin
	OpFilter
	OpUnion
)
