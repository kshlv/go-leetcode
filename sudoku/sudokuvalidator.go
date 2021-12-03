// Package sudoku implements a simple validation of
// a sudoku matrix 9x9
//
// Original problem:
// https://leetcode.com/problems/valid-sudoku/
package sudoku

// Validate has an actual signature of
// isValidSudoku(board [][]byte) bool
func Validate(board [][]byte) bool {
	for _, row := range board {
		m := map[byte]bool{}
		for _, b := range row {
			if (0 < b && b <= 9) || ('1' <= b && b <= '9') {
				if _, ok := m[b]; ok {
					return false
				}
				m[b] = true
			}
		}
	}

	for j := 0; j < len(board); j++ {
		m := map[byte]bool{}
		for i := 0; i < len(board[j]); i++ {
			b := board[i][j]
			if (0 < b && b <= 9) || ('1' <= b && b <= '9') {
				if _, ok := m[b]; ok {
					return false
				}
				m[b] = true
			}
		}
	}

	for i := 0; i < len(board)/3; i++ {
		for j := 0; j < len(board)/3; j++ {
			square := map[byte]bool{}
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					b := board[3*i+m][3*j+n]
					if (0 < b && b <= 9) || ('1' <= b && b <= '9') {
						if _, ok := square[b]; ok {
							return false
						}
						square[b] = true
					}
				}
			}
		}
	}

	return true
}
