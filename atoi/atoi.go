// Package atoi implements a solution to the self-explanatory problem
//
// Original problem:
// https://leetcode.com/problems/string-to-integer-atoi
package atoi

import "unicode"

// Atoi is ...
func Atoi(s string) int {
	var mult int64 = 0
	var num int64 = 0
	numStarted := false
	for _, r := range []rune(s) {
		switch r {
		case ' ', '\t', '\n', '\r':
			if numStarted {
				return int(mult * num)
			}
			continue
		case '+':
			if numStarted {
				return int(mult * num)
			}
			mult = 1
			numStarted = true
		case '-':
			if numStarted {
				return int(mult * num)
			}
			numStarted = true
			mult = -1

		// TODO: pls optimize this
		case '1':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 1
		case '2':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 2
		case '3':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 3
		case '4':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 4
		case '5':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 5
		case '6':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 6
		case '7':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 7
		case '8':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 8
		case '9':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num*10 + 9
		case '0':
			if mult == 0 {
				mult = 1
			}
			numStarted = true
			num = num * 10
		}

		if unicode.IsLetter(r) || r == '.' {
			return int(mult * num)
		}

		if mult*num < -1<<31 {
			return int(mult * 1 << 31)
		}
		if mult*num > (1<<31)-1 {
			return int(mult*(1<<31) - 1)
		}
	}

	return int(mult * num)
}
