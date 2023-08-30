package kata

import (
	"math"
	"strconv"
)


func SumDigPow(a, b uint64) []uint64 {
	var result []uint64
	for i := a; i <= b; i++ {
		j := i
		sum := uint64(0)
		exp := len(strconv.FormatUint(i, 10))
		for j > 0 {
            sum += uint64(math.Pow(float64(j % 10), float64(exp)))
            j /= 10
            exp--
        }
		if sum == i {
			result = append(result, i)
		}
	}
	return result
}
