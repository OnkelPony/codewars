package kata

func Solequa(n int) [][]int {
	result := [][]int{}
	for x := n / 2 + 1; x >= 0; x-- {
		for y := x / 2 + 1; y >= 0; y-- {
			if (x-2*y)*(x+2*y) == n {
				result = append(result, []int{x, y})
			}
			if (x-2*y)*(x+2*y) > n {
				break
            }
		}
	}
	return result
}
