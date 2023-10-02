package dyn_prog

import (
	"testing"
)

func TestFibo(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{30, 832040},
		{40, 102334155},
		{50, 12586269025},
		{60, 1548008755920},
		{70, 190392490709135},
		{80, 23416728348467685},
        {90, 2880067194370816120},
		{100, 3736710778780434371},
	}
	var fibos = make(map[int]int)
	for _, c := range cases {
		got := Fibo(c.n, fibos)
		if got != c.want {
			t.Errorf("Fibo(%d) => %d, want %d", c.n, got, c.want)
		}
	}
}
