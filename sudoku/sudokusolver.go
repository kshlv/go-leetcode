// Package sudoku solves
// a classic sudoku matrix 9x9
//
// Original problem:
// https://leetcode.com/problems/sudoku-solver/
package sudoku

import "fmt"

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

func isDigit(b byte) bool {
	if (1 <= b && b <= 9) || ('1' <= b && b <= '9') {
		return true
	}
	return false
}

var boardSize = 9

// Solve has an actual signature of
// solveSudoku(board [][]byte)
//
// I see three phases to this algorythm:
// 1. Fill [][][]byte-matrix (call it `options``) of possible values for each cell.
// 2. Iterate through options:
//  * every cell (actually, []byte) with only one value (len() == 1) are considered to be solved;
//  * add the value to the corresponding cell of the `board`;
//  * remove the value from `options`' 3x3 box, vertical and horizontal.
// 3. Repeat step 2 until all cells are solved
func Solve(board [][]byte) {
	// if board's height or length are inappropriate, return
	if len(board) != boardSize {
		return
	}
	for i := range board {
		if len(board[i]) != boardSize {
			return
		}
	}
	// opts for possible digits in the cells
	opts := make([][][]byte, boardSize, boardSize)
	// Phase 1
	// collect row candidates
	for i, row := range board {
		// pres <-> [numbers] presented [in the row]
		pres := []byte{}
		opts[i] = make([][]byte, boardSize, boardSize)
		for j, b := range row {
			if isDigit(b) {
				pres = append(pres, b)
			} else {
				if opts[i][j] == nil {
					if b == 0 {
						opts[i][j] = []byte{1, 2, 3, 4, 5, 6, 7, 8, 9}
					}
					if b == '.' {
						opts[i][j] = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}
					}

				}
			}
		}
		for j := range row {
			if b := board[i][j]; !isDigit(b) {
				opts[i][j] = exclude(opts[i][j], pres...)
			}
		}
	}
	// collect column candidates
	for i := 0; i < boardSize; i++ {
		pres := []byte{}
		for j := 0; j < boardSize; j++ {
			if b := board[j][i]; isDigit(b) {
				pres = append(pres, b)
			}
		}
		for j := 0; j < boardSize; j++ {
			if b := board[j][i]; !isDigit(b) {
				opts[j][i] = exclude(opts[j][i], pres...)
			}
		}
	}
	// collect box candidates
	for i := 0; i < boardSize/3; i++ {
		for j := 0; j < boardSize/3; j++ {
			pres := []byte{}
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if b := board[3*i+m][3*j+n]; isDigit(b) {
						pres = append(pres, b)
					}
				}
			}
			for m := 0; m < 3; m++ {
				for n := 0; n < 3; n++ {
					if b := board[3*i+m][3*j+n]; !isDigit(b) {
						opts[3*i+m][3*j+n] = exclude(opts[3*i+m][3*j+n], pres...)
					}
				}
			}
		}
	}
	// Phase 2
	// iteratively fill the board with lonely candidates
	boardFilled := false
	for !boardFilled {
		for i := 0; i < len(opts); i++ {
			for j := 0; j < len(opts[i]); j++ {
				if len(opts[i][j]) == 1 {
					sol := opts[i][j][0]
					board[i][j] = sol
					// remove the candidate from options' 3x3
					for m := 0; m < 3; m++ {
						for n := 0; n < 3; n++ {
							b := board[3*(i/3)+m][3*(j/3)+n]
							if b == 0 {
								opts[3*(i/3)+m][3*(j/3)+n] = exclude(opts[3*(i/3)+m][3*(j/3)+n], sol)
							}
						}
					}
					// remove the candidate from options' horizontal
					for col := 0; col < 9; col++ {
						if opts[i][col] != nil && len(opts[i][col]) > 0 {
							opts[i][col] = exclude(opts[i][col], sol)
						}
					}
					// remove the candidate from options' vertical
					for row := 0; row < 9; row++ {
						if opts[row][j] != nil && len(opts[row][j]) > 0 {
							opts[row][j] = exclude(opts[row][j], sol)
						}
					}
				}
			}
		}
		// Phase 3
		// decide whether shall we stop or continue
		count := 0
		for _, row := range board {
			for _, cell := range row {
				fmt.Printf("%v", cell)
				if isDigit(cell) {
					count++
				}
			}
		}
		if count == boardSize*boardSize {
			boardFilled = true
		}
	}
}
