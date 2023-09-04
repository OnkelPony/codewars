package main

import (
	"fmt"
	"math"
)

func isDisarium(n uint64) bool {
  
  return sum == n
}

func nextDisarium(n uint64) uint64 {
  var sum uint64
  for i := n + 1; i <= math.MaxUint64; i++ {
    for j, digit := range fmt.Sprint(n) {
      sum += uint64(math.Pow(float64(digit - '0') , float64(j + 1)))
    }
  }
  return 0
}

func main() {
  fmt.Println(nextDisarium(3000))
}
