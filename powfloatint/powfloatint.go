// Package powfloatint implements a basic power operation
// x^n with types float64 and int respectively
//
// Original problem:
// https://leetcode.com/problems/powx-n/
package powfloatint

// Pow has an actual signature of
// myPow(x float64, n int) float64
func Pow(x float64, n int) float64 {
	if n == 0 || x == 1.0 {
		return 1.0
	}

	if x == -1.0 {
		if n%2 == 0 {
			return 1.0
		}
		return -1.0
	}

	n64 := int64(n)
	var neg bool
	if n64 < 0 {
		neg = true
		n64 *= -1
	}

	result := 1.0
	for i := int64(1); i <= n64; i++ {
		result *= x
	}

	if neg {
		result = 1 / result
	}

	return result
}
