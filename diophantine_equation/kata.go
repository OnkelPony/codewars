package kata

import "math"

func Solequa(n int) [][]int {
	result := [][]int{}
	var maxx int
	if n%2 == 0 {
		maxx = n / 2
	} else {
		maxx = n/2 + 1
	}
	for x := maxx; x >= 0; x -= 2 {
		y := int(math.Sqrt(float64(x*x-n) / 4))
		if x*x-4*y*y == n {
			result = append(result, []int{x, y})
		}
	}
	return result
}
