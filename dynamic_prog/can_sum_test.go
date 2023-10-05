package dyn_prog

import (
	"testing"
)

func TestCanSum(t *testing.T) {
	cases := []struct {
		targetSum    int
		numbers []int
		want bool
	}{
		{8, []int{2, 3, 5}, true},
		{8, []int{3, 6}, false},
	}
	for _, c := range cases {
		got := canSum(c.targetSum, c.numbers)
		if got != c.want {
			t.Errorf("canSum(%d, %v) => %t, want %t",c.targetSum, c.numbers, got, c.want)
		}
	}
}
