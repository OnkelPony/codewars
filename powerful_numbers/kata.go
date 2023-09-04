package kata

import (
	"fmt"
)

func SumDigPow(a, b uint64) []uint64 {
	var result []uint64
	for i := a; i <= b; i++ {
		sum := uint64(0)
		for exp, digit := range fmt.Sprint(i) {
			sum += uintPow(uint64(digit-'0'), uint64(exp+1))
		}
		if sum == i {
			result = append(result, i)
		}
	}
	return result
}

func uintPow(base, exp uint64) uint64 {
	result := uint64(1)
	for i := uint64(0); i < exp; i++ {
		result *= base
	}
	return result
}
