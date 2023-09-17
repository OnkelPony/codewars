package kata

import (
	"math/big"
)

func LastDigit(n1, n2 string) int {
	var result uint
	base := uint(n1[len(n1)-1] - '0')
	
	if base == 0 || base == 1 || base == 5 || base == 6 {
		return int(base)
	}
	hugeExponent, _ := new(big.Int).SetString(n2, 10)
	smallExponent := big.NewInt(0)
	smallExponent.Mod(hugeExponent, big.NewInt(4))
	exponent := (*smallExponent).Int64()
	if exponent == 0 && len(n2) > 1 {
		exponent = 4
	}
	result = UintPow(uint(base), uint(exponent))
	return int(result % 10)
}

func UintPow(base, exponent uint) uint {
	result := uint(1)
	for i := uint(0); i < exponent; i++ {
		result *= base
	}
	return result
}
