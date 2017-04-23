package internal

import (
	"reflect"
	"testing"
)

func u32p(s, p, o int) EncTriplePattern {
	return EncTriplePattern{uint32(s), uint32(p), uint32(o)}
}

func TestGroupPatterns(t *testing.T) {
	tests := []struct {
		patterns []EncTriplePattern
		want     [][]EncTriplePattern
	}{
		{
			nil,
			nil,
		},
		{
			[]EncTriplePattern{u32p(100, 101, 102)},
			[][]EncTriplePattern{
				{u32p(100, 101, 102)}},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102)}},
		},
		{
			[]EncTriplePattern{u32p(100, 101, 102), u32p(200, 201, 202)},
			[][]EncTriplePattern{
				{u32p(100, 101, 102)},
				{u32p(200, 201, 202)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102), u32p(2, 101, 1)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102), u32p(2, 101, 1)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102), u32p(2, 101, 3)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102)},
				{u32p(2, 101, 3)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102), u32p(3, 101, 4), u32p(2, 101, 1)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102), u32p(2, 101, 1)},
				{u32p(3, 101, 4)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 2), u32p(2, 101, 3), u32p(4, 101, 5), u32p(5, 101, 6), u32p(5, 101, 6)},
			[][]EncTriplePattern{
				{u32p(1, 101, 2), u32p(2, 101, 3)},
				{u32p(4, 101, 5), u32p(5, 101, 6), u32p(5, 101, 6)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 2), u32p(2, 101, 3), u32p(4, 101, 5), u32p(5, 101, 6), u32p(5, 0, 6)},
			[][]EncTriplePattern{
				{u32p(1, 101, 2), u32p(2, 101, 3)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102), u32p(2, 101, 102), u32p(3, 101, 102)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102)},
				{u32p(2, 101, 102)},
				{u32p(3, 101, 102)},
			},
		},
	}

	for _, test := range tests {
		if got := GroupPatterns(test.patterns); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("GroupPatterns(%v) => %v; want %v", test.patterns, got, test.want)
		}
	}
}
