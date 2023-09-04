package main

import (
	"fmt"
)

func isDisarium(n uint64) bool {
	var sum uint64
	for i, digit := range fmt.Sprint(n) {
		sum += uintPow(uint64(digit-'0'), uint64(i+1))
		//fmt.Println("sum = ", sum)
	}
	return sum == n
}

func nextDisarium(n uint64) uint64 {
	for i := n + 1; i < 12157692622039625638; i++ {
		if isDisarium(i) {
			return i
		}
	}
	return 0
}

func uintPow(base, exp uint64) uint64 {
  result := uint64(1)
	for i := uint64(0); i < exp; i++ {
		result *= base
	}
	return result
}

func main() {
	fmt.Println(nextDisarium(2646798))
}
