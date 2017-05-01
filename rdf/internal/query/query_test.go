package query

import (
	"reflect"
	"testing"

	"github.com/RoaringBitmap/roaring"
)

func u32p(s, p, o, c int) EncTriplePattern {
	return EncTriplePattern{uint32(s), uint32(p), uint32(o), uint32(c)}
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
			[]EncTriplePattern{u32p(100, 101, 102, 99)},
			[][]EncTriplePattern{
				{u32p(100, 101, 102, 99)}},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102, 99)}},
		},
		{
			[]EncTriplePattern{u32p(1, 2, 0, 0)},
			nil,
		},
		{
			[]EncTriplePattern{u32p(100, 101, 102, 99), u32p(200, 201, 202, 99)},
			[][]EncTriplePattern{
				{u32p(100, 101, 102, 99)},
				{u32p(200, 201, 202, 99)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102, 99), u32p(2, 101, 1, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102, 99), u32p(2, 101, 1, 99)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102, 99), u32p(2, 101, 3, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102, 99)},
				{u32p(2, 101, 3, 99)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102, 99), u32p(3, 101, 4, 99), u32p(2, 101, 1, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102, 99), u32p(2, 101, 1, 99)},
				{u32p(3, 101, 4, 99)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 2, 99), u32p(2, 101, 3, 99), u32p(4, 101, 5, 99), u32p(5, 101, 6, 99), u32p(5, 101, 6, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 2, 99), u32p(2, 101, 3, 99)},
				{u32p(4, 101, 5, 99), u32p(5, 101, 6, 99), u32p(5, 101, 6, 99)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 2, 99), u32p(2, 101, 3, 99), u32p(4, 101, 5, 99), u32p(5, 101, 6, 99), u32p(5, 0, 6, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 2, 99), u32p(2, 101, 3, 99)},
			},
		},
		{
			[]EncTriplePattern{u32p(1, 101, 102, 99), u32p(2, 101, 102, 99), u32p(3, 101, 102, 99)},
			[][]EncTriplePattern{
				{u32p(1, 101, 102, 99)},
				{u32p(2, 101, 102, 99)},
				{u32p(3, 101, 102, 99)},
			},
		},
	}

	for _, test := range tests {
		if got := GroupPatterns(test.patterns); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("GroupPatterns(%v) => %v; want %v", test.patterns, got, test.want)
		}
	}
}

func TestJoinSolutions(t *testing.T) {
	tests := []struct {
		a, b, want EncSolutions
	}{
		{ // 1:
			a:    EncSolutions{},
			b:    EncSolutions{},
			want: EncSolutions{},
		},
		{ // 2:
			a: EncSolutions{},
			b: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
			want: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
		},
		{ // 3:
			a: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
			b: EncSolutions{},
			want: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
		},
		{ // 4:
			a: EncSolutions{
				Vars: []uint32{1, 4},
				Rows: [][]uint32{{102, 210}},
			},
			b: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
			want: EncSolutions{
				Vars: []uint32{1, 4, 2, 3},
				Rows: [][]uint32{{102, 210, 103, 104}},
			},
		},
	}

	for i, test := range tests {
		bound := make(map[uint32]*roaring.Bitmap)
		// TODO fill bound with encsolutions a & b
		// TODO test bound
		if got := test.a.Join(test.b, bound); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("%d: %v join  %v => %v; want %v", i+1, test.a, test.b, got, test.want)
		}
	}
}

func TestSolutionsProject(t *testing.T) {
	tests := []struct {
		input EncSolutions
		vars  []uint32
		want  EncSolutions
	}{
		/*{ // 1:
			input: EncSolutions{},
			vars:  nil,
			want:  EncSolutions{},
		},*/
		{ // 2:
			input: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
			vars: []uint32{2},
			want: EncSolutions{
				Vars: []uint32{2},
				Rows: [][]uint32{{101}, {103}},
			},
		},
		{ // 3:
			input: EncSolutions{
				Vars: []uint32{1, 2, 3},
				Rows: [][]uint32{{100, 101, 102}, {102, 103, 104}},
			},
			vars: []uint32{1, 3},
			want: EncSolutions{
				Vars: []uint32{1, 3},
				Rows: [][]uint32{{100, 102}, {102, 104}},
			},
		},
	}

	for i, test := range tests {

		if got := test.input.Project(test.vars); !reflect.DeepEqual(got, test.want) {
			t.Fatalf("%d: %v.Project(%v) => %v; want %v", i+1, test.input, test.vars, got, test.want)
		}
	}
}
