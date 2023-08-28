package kata

import "math"

func IsPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	} else if n%2 == 0 || n <= 1 {
		return false
	}
	upper := math.Ceil(math.Sqrt(float64(n)))
	for i := 3; i <= int(upper); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
