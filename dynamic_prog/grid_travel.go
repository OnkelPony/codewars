package dyn_prog

func gridTravel(m, n int) int {
	if m == 1 && n == 1 {
		return 1
	}
	if m == 0 || n == 0 {
		return 0
	}
	return gridTravel(m - 1, n) + gridTravel(m, n - 1)
}