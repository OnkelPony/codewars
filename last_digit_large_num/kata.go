package kata

import (
	"math/big"
)


func LastDigit(n1, n2 string) int {
  var result int
  base := int(n1[len(n1)-1] - '0')
  if base == 0 && len(n1) == 1 {
    return 0
  }
  if base == 5 {
    return 5
  }
  hugeExponent, _  := new(big.Int).SetString(n2, 10)
  exponent := big.NewInt(0)
  exponent.Mod(hugeExponent, big.NewInt(12))
  result = IntPow((base), int((*exponent).Int64()))
  return result % 10
}

func IntPow(base, exponent int) int {
  result := 1
  for i := 0; i < exponent; i++ {
    result *= base
  }
  return result
}