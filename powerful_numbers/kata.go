package kata

import (
	"fmt"
	"math"
	"strconv"
)
type Power struct {
	base uint64
	exp uint64
}
func SumDigPow(a, b uint64) []uint64 {
	
	allPowers := makePowers(a, b)
	var result []uint64 
	for i := a; i <= b; i++ {
		strI := fmt.Sprint(i)
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


func makePowers(a, b uint64) map[Power]uint64 {
	
}