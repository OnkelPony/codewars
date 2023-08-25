package kata

func Number(busStops [][2]int) int {
	var result int
	for _, inOut := range busStops {
		result += inOut[0]
		result -= inOut[1]
		if result < 0 {
			result = 0
		}
	}
	return result
}
