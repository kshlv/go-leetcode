// Package palindromeint ...
//
// Original problem:
// https://leetcode.com/problems/palindrome-number/
package palindromeint

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

// IsPalindrome has an actual signature of
// isPalindrome(x int) bool
func IsPalindrome(x int) bool {
	if x >= 0 {
		card := 1
		for x/pow(10, card) > 0 {
			card++
		}
		if card == 1 {
			return true
		}

		for i := 0; i < card/2; i++ {
			left := (x % pow(10, card-i)) / pow(10, card-i-1)
			right := (x % pow(10, i+1)) / pow(10, i)
			if left != right {
				return false
			}
		}

		return true
	}

	return false
}
