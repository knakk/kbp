package memory

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

// TODO move to knakk/kbp/rdf/internal
func TestGroupPatternsByVariable(t *testing.T) {
	tests := []struct {
		patterns []encPattern
		want     [][]encPattern
	}{
		{
			[]encPattern{},
			[][]encPattern{},
		},
		{
			[]encPattern{{-1, -2, -3, 99}},
			[][]encPattern{
				{{-1, -2, -3, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-2, 2, 3, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}},
				{{-2, 2, 3, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-1, 2, 3, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-1, 2, 3, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-1, 2, -2, 99}, {-2, 2, 3, 99}, {-3, 4, 1, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}},
				{{-3, 4, 1, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-1, 2, -2, 99}, {-2, 2, 3, 99}, {-3, 4, -1, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-3, 4, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}},
			},
		},
		{
			[]encPattern{{1, 2, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}, {-3, 4, -1, 99}},
			[][]encPattern{
				{{1, 2, -1, 99}, {-3, 4, -1, 99}, {-2, 2, 3, 99}, {-1, 2, -2, 99}},
			},
		},
		{
			[]encPattern{{-1, 1, -2, 99}, {-1, 2, -3, 99}, {-3, 3, 4, 99}, {-2, 5, 6, 99}, {-2, 7, -4, 99}, {-4, 8, -5, 99}},
			[][]encPattern{
				{{-4, 8, -5, 99}, {-3, 3, 4, 99}, {-2, 7, -4, 99}, {-2, 5, 6, 99}, {-1, 2, -3, 99}, {-1, 1, -2, 99}},
			},
		},
	}

	for _, test := range tests {
		groups := groupPatternsByVariable(test.patterns)
		for n := range groups {
			sort.Slice(groups[n], func(i, j int) bool {
				return fmt.Sprintf("%d%d%d", groups[n][i][0], groups[n][i][1], groups[n][i][2]) > fmt.Sprintf("%d%d%d", groups[n][j][0], groups[n][j][1], groups[n][j][2])
			})
		}
		if !reflect.DeepEqual(groups, test.want) {
			t.Fatalf("groupPatternsByVariable(%v) => %v; want %v", test.patterns, groups, test.want)
		}
	}
}
