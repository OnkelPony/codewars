package kata

import (
	"math"
	"strconv"
)


func SumDigPow(a, b uint64) []uint64 {
	var result []uint64
	for i := a; i <= b; i++ {
		strI := strconv.FormatUint(i, 10)
		sum := uint64(0)
		exp := 1
		for _, strDigit := range strI {
			digit, _ := strconv.Atoi(string(strDigit))
            sum += uint64(math.Pow(float64(digit), float64(exp)))
            exp++
        }
		if sum == i {
			result = append(result, i)
		}
	}
	return result
}
