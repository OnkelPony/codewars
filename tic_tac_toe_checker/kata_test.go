package kata

import (
	"testing"
)

func TestIsSolved(t *testing.T) {
	cases := []struct {
		board [3][3]int
		want  int
	}{
		{[3][3]int{
			{0, 0, 1},
			{0, 1, 2},
			{2, 1, 0},
		}, -1},
		{[3][3]int{
			{1, 1, 1},
			{0, 2, 2},
			{0, 0, 0},
		}, 1},
		{[3][3]int{
			{2, 1, 2},
			{2, 1, 1},
			{1, 1, 2},
		}, 1},
		{[3][3]int{
			{2, 1, 2},
			{2, 1, 1},
			{1, 2, 1},
		}, 0},
	}

	for _, c := range cases {
		got := IsSolved(c.board)
		if got != c.want {
			t.Errorf("IsSolved(%v) => %d, want %d", c.board, got, c.want)
		}
	}
}
