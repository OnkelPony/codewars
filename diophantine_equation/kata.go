package kata

import "math"

func Solequa(n int) [][]int {
	result := [][]int{}
	maxx := n/2 + 1
	for x := maxx; x >= 0; x-- {
		y := int(math.Sqrt(float64(x*x-n) / 4))
		if x*x-4*y*y == n {
			result = append(result, []int{x, y})
		}
	}
	return result
}
