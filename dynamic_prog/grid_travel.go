package dyn_prog

import (
    "fmt"
)

func gridTravel(m, n int, ways map[string]int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	key := fmt.Sprintf("%d,%d", m, n)
	if _, ok := ways[key]; ok {
        return ways[key]
    }
	ways[key] = gridTravel(m - 1, n, ways) + gridTravel(m, n - 1, ways)
	return ways[key]
}