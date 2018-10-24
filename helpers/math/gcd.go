package math

import "math"

func Gcd(a, b int32) int32 {
	var c int32 = 1

	for b != 0 {
		c = a % b
		a = b
		b = c
	}

	return int32(math.Abs(float64(a)))
}
