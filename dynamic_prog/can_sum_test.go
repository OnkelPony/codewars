package dyn_prog

import (
	"testing"
)

func TestCanSum(t *testing.T) {
	cases := []struct {
		targetSum int
		numbers   []int
		want      bool
	}{
		{8, []int{2, 3, 5}, true},
		{8, []int{3, 6}, false},
		{18, []int{2, 3, 5, 7}, true},
		{108, []int{3, 5, 7}, true},
		{3700017, []int{7, 17, 23, 29, 37}, true},
	}
	for _, c := range cases {
		memo := make(map[int](bool))
		got := canSum(c.targetSum, c.numbers, memo)
		if got != c.want {
			t.Errorf("canSum(%d, %v) => %t, want %t", c.targetSum, c.numbers, got, c.want)
		}
	}
}
