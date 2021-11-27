// Package reverseint implements digits reversion in 32-bit integer
//
// Original problem:
// https://leetcode.com/problems/reverse-integer/
package reverseint

func pow(x, y int) int {
	if x == 0 {
		return 0
	}
	if y == 0 {
		return 1
	}

	result := 1
	for i := 0; i < y; i++ {
		result *= x
	}

	return result
}

// Reverse is actually pretty self-explanatory; its original signature is
// reverse(x int) int
func Reverse(x int) (result int) {
	if x == 0 {
		return result
	}

	for count := 0; x/pow(10, count) != 0; count++ {
		inc := (x % pow(10, count+1)) / pow(10, count)
		result += inc

		result *= 10

	}

	result /= 10

	if result < -1<<31 || result >= 1<<31 {
		return 0
	}

	return
}
