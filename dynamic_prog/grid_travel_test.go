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
		// {18, 18, 2333606220},
	}
	for _, c := range cases {
		got := gridTravel(c.m, c.n)
		if got != c.want {
			t.Errorf("gridTravel(%d, %d) => %d, want %d",c.m, c.n, got, c.want)
		}
	}
}
