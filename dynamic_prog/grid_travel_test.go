package dyn_prog

import (
	"testing"
)

func TestGridTravel(t *testing.T) {
	cases := []struct {
		m,n    int
		want int
	}{
		{1, 1, 1},
		{2, 3, 3},
		{3, 3, 6},
		{4, 5, 35},
		{5, 5, 70},
        {6, 6, 252},
        {7, 7, 924},
        {8, 8, 3432},
        {9, 9, 12870},
		{12, 15, 4457400},
		{18, 18, 2333606220},
	}
	ways := make(map[string]int)
	for _, c := range cases {
		got := gridTravel(c.m, c.n, ways)
		if got != c.want {
			t.Errorf("gridTravel(%d, %d) => %d, want %d",c.m, c.n, got, c.want)
		}
	}
}
