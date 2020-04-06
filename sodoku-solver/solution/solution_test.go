package solution

import (
	"bytes"
	"testing"
)

var (
	puzzle1 = [][]byte{
		[]byte("53..7...."),
		[]byte("6..195..."),
		[]byte(".98....6."),
		[]byte("8...6...3"),
		[]byte("4..8.3..1"),
		[]byte("7...2...6"),
		[]byte(".6....28."),
		[]byte("...419..5"),
		[]byte("....8..79"),
	}
	solution1 = [][]byte{
		[]byte("534678912"),
		[]byte("672195348"),
		[]byte("198342567"),
		[]byte("859761423"),
		[]byte("426853791"),
		[]byte("713924856"),
		[]byte("961537284"),
		[]byte("287419635"),
		[]byte("345286179"),
	}
)

func cmpBoards(b1, b2 [][]byte) bool {
	for i, v := range b1 {
		if bytes.Compare(v, b2[i]) != 0 {
			return false
		}
	}
	return true
}

func TestSolveSudoku(t *testing.T) {
	type test struct {
		input  boardType
		expect boardType
	}
	tests := []test{
		{puzzle1, solution1},
	}
	for _, tN := range tests {
		if solveSudoku(tN.input); cmpBoards(tN.input, tN.expect) == false {
			t.Errorf("Got:\n%s\nWant:\n%s", tN.input.String(), tN.expect.String())
		}
	}
}

func TestValidSet(t *testing.T) {
	type test struct {
		set    []uint8
		expect bool
	}
	tests := []test{
		{[]uint8{1, 2, 3, 3, 4, 5, 6, 7, 8}, false},
		{[]uint8{1, 2, 3, 9, 4, 5, 6, 7, 8}, true},
		{[]uint8{1, 2, 0, 9, 4, 5, 0, 0, 8}, true},
	}
	for _, tN := range tests {
		if got := validSet(tN.set); got != tN.expect {
			t.Errorf("%v should be %t but got %t", tN.set, tN.expect, got)
		}
	}
}

func TestBlock(t *testing.T) {
	type test struct {
		p      pos
		expect []pos
	}
	tests := []test{
		{pos{4, 5}, []pos{{3, 3}, {4, 3}, {5, 3}, {3, 4}, {4, 4}, {5, 4}, {3, 5}, {4, 5}, {5, 5}}},
	}
	equal := func(a []pos, b []pos) bool {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	for _, tN := range tests {
		if got := block(tN.p); !equal(got, tN.expect) {
			t.Errorf("%v should be %v but got %v", tN.p, tN.expect, got)
		}
	}
}

func BenchmarkSolveSodoku(b *testing.B) {
	b.Run("example code", func(b *testing.B) {
		b.ReportAllocs()
		for n := 0; n < b.N; n++ {
			solveSudoku(puzzle1)
		}
	})
}
