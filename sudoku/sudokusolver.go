// Package sudoku solves
// a classic sudoku matrix 9x9
//
// Original problem:
// https://leetcode.com/problems/sudoku-solver/
package sudoku

func exclude(arr []byte, nums ...byte) []byte {
	result := []byte{}

	var inNum = func(b byte) bool {
		for _, num := range nums {
			if b == num {
				return true
			}
		}
		return false
	}

	for _, b := range arr {
		if inNum(b) {
			continue
		}
		result = append(result, b)
	}

	return result
}

// Solve has an actual signature of
// solveSudoku(board [][]byte)
//
// I see three phases to this algorythm
// 1. Fill [][][]byte-matrix of possible values for each cell
// 2. Every cell with only one value considered to be solved
//  * don't forget to remove it from box, vertical and horizontal
// 3. Repeat 2 until all cells are solved
func Solve(board [][]byte) [][][]byte {
	// options for possible digits in the cells
	options := make([][][]byte, 9, 9)

	for i, row := range board {
		// pres = [numbers] presented
		pres := []byte{}
		options[i] = make([][]byte, 9, 9)
		for j, b := range row {
			if (0 < b && b <= 9) || ('1' <= b && b <= '9') {
				pres = append(pres, b)
			} else {
				if options[i][j] == nil {
					options[i][j] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
				}
			}
		}
		for j := range row {
			if board[i][j] == 0 {
				options[i][j] = exclude(options[i][j], pres...)
			}
		}

	}

	for i := 0; i < len(board); i++ {
		pres := []byte{}
		for j := 0; j < len(board[i]); j++ {
			b := board[j][i]
			if (0 < b && b <= 9) || ('1' <= b && b <= '9') {
				pres = append(pres, b)
			}
		}
		for j := 0; j < len(board[i]); j++ {
			if board[j][i] == 0 {
				options[j][i] = exclude(options[j][i], pres...)
			}
		}
	}

	for i := 0; i < len(board)/3; i++ {
		for j := 0; j < len(board)/3; j++ {
			pres := []byte{}
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					b := board[3*i+m][3*j+n]
					if (0 < b && b <= 9) || ('1' <= b && b <= '9') {
						pres = append(pres, b)
					}
				}
			}

			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					b := board[3*i+m][3*j+n]
					if b == 0 {
						options[3*i+m][3*j+n] = exclude(options[3*i+m][3*j+n], pres...)
					}
				}
			}
		}
	}

	return options
}
