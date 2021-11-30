// Package zigzagconv ...
//
// Original problem:
// https://leetcode.com/problems/zigzag-conversion/
package zigzagconv

// Convert is ...
// convert(s string, numRows int) string
func Convert(s string, numRows int) string {
	if numRows == 1 || len(s) < numRows {
		return s
	}

	grid := make([][]rune, numRows, numRows)
	for i := range grid {
		grid[i] = []rune{}
	}

	// z is for zigzag, but it actually is a direction
	z := 1
	row := 0
	for _, c := range s {
		grid[row] = append(grid[row], c)
		row += z
		if row == 0 || row == numRows-1 {
			z *= -1
		}
	}

	r := []rune{}
	for row := range grid {
		for col := range grid[row] {
			r = append(r, grid[row][col])
		}
	}

	return string(r)
}
