package sudoku

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExclue(t *testing.T) {
	testcases := []struct {
		arr, nums, expected []byte
	}{
		{
			arr:      []byte{1},
			nums:     []byte{},
			expected: []byte{1},
		},
		{
			arr:      []byte{1, 2},
			nums:     []byte{2},
			expected: []byte{1},
		},
		{
			arr:      []byte{1, 2, 3},
			nums:     []byte{1},
			expected: []byte{2, 3},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			actual := exclude(tc.arr, tc.nums...)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestSolve(t *testing.T) {
	testcases := []struct {
		board, expected [][]byte
	}{
		{
			board: [][]byte{
				{5, 3, 0 /**/, 0, 7, 0 /**/, 0, 0, 0},
				{6, 0, 0 /**/, 1, 9, 5 /**/, 0, 0, 0},
				{0, 9, 8 /**/, 0, 0, 0 /**/, 0, 6, 0},
				/************************************/
				{8, 0, 0 /**/, 0, 6, 0 /**/, 0, 0, 3},
				{4, 0, 0 /**/, 8, 0, 3 /**/, 0, 0, 1},
				{7, 0, 0 /**/, 0, 2, 0 /**/, 0, 0, 6},
				/************************************/
				{0, 6, 0 /**/, 0, 0, 0 /**/, 2, 8, 0},
				{0, 0, 0 /**/, 4, 1, 9 /**/, 0, 0, 5},
				{0, 0, 0 /**/, 0, 8, 0 /**/, 0, 7, 9},
			},
			expected: [][]byte{
				{5, 3, 4 /**/, 6, 7, 8 /**/, 9, 1, 2},
				{6, 7, 2 /**/, 1, 9, 5 /**/, 3, 4, 8},
				{1, 9, 8 /**/, 3, 4, 2 /**/, 5, 6, 7},
				/************************************/
				{8, 5, 9 /**/, 7, 6, 1 /**/, 4, 2, 3},
				{4, 2, 6 /**/, 8, 5, 3 /**/, 7, 9, 1},
				{7, 1, 3 /**/, 9, 2, 4 /**/, 8, 5, 6},
				/************************************/
				{9, 6, 1 /**/, 5, 3, 7 /**/, 2, 8, 4},
				{2, 8, 7 /**/, 4, 1, 9 /**/, 6, 3, 5},
				{3, 4, 5 /**/, 2, 8, 6 /**/, 1, 7, 9},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			Solve(tc.board)
			assert.Equal(t, tc.expected, tc.board)
		})
	}
}
