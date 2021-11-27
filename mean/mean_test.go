package mean

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMean(t *testing.T) {
	testcases := []struct {
		name     string
		input    [2][]int
		expected float64
	}{
		// m + n == 0
		{
			name:     "[][] (m = 0, n = 0)",
			input:    [...][]int{{}, {}},
			expected: 0.0,
		},
		// m + n == 1
		{
			name:     "[0][] (m = 1, n = 0)",
			input:    [...][]int{{0}, {}},
			expected: 0.0,
		},
		{
			name:     "[][0] (m = 0, n = 1)",
			input:    [...][]int{{}, {0}},
			expected: 0.0,
		},
		{
			name:     "[2][] (m = 1, n = 0)",
			input:    [...][]int{{2}, {}},
			expected: 2.0,
		},
		{
			name:     "[][1] (m = 0, n = 1)",
			input:    [...][]int{{}, {1}},
			expected: 1.0,
		},
		// m + n == 2
		{
			name:     "[0][0] (m = 1, n = 1)",
			input:    [...][]int{{0}, {0}},
			expected: 0.0,
		},
		{
			name:     "[1][0] (m = 1, n = 1)",
			input:    [...][]int{{1}, {0}},
			expected: 0.5,
		},
		{
			name:     "[1,1][] (m = 2, n = 0)",
			input:    [...][]int{{1, 1}, {}},
			expected: 1.0,
		},
		{
			name:     "[0,1][] (m = 2, n = 0)",
			input:    [...][]int{{0, 1}, {}},
			expected: 0.5,
		},
		// m + n == 3
		{
			name:     "[0,0][1] (m = 2, n = 1)",
			input:    [...][]int{{0, 0}, {1}},
			expected: 0.0,
		},
		// m + n == 4
		{
			name:     "[0,0][0,0] (m = 2, n = 2)",
			input:    [...][]int{{0}, {}},
			expected: 0.0,
		},
		{
			name:     "[0,0][0,1] (m = 2, n = 2)",
			input:    [...][]int{{0, 0}, {0, 1}},
			expected: 0.0,
		},
		{
			name:     "[0,1][1,1] (m = 2, n = 2)",
			input:    [...][]int{{0, 1}, {1, 1}},
			expected: 1.0,
		},
		{
			name:     "[1,1][1,1] (m = 2, n = 2)",
			input:    [...][]int{{1, 1}, {1, 1}},
			expected: 1.0,
		},
		{
			name:     "[1,1][0,0] (m = 2, n = 2)",
			input:    [...][]int{{1, 1}, {0, 0}},
			expected: 0.5,
		},
		// m + n == 5
		{
			name:     "[1,1][1,1,1] (m = 2, n = 3)",
			input:    [...][]int{{1, 1}, {1, 1, 1}},
			expected: 1.0,
		},
		{
			name:     "[1,1][1,1,1] (m = 2, n = 3)",
			input:    [...][]int{{1, 1}, {1, 1, 1}},
			expected: 1.0,
		},
		// m + n == 6
		{
			name:     "[1,1,1][1,1,1] (m = 3, n = 3)",
			input:    [...][]int{{1, 1, 1}, {1, 1, 1}},
			expected: 1.0,
		},
		// problem examples 1, 2
		{
			name:     "[1,3][2]",
			input:    [...][]int{{1, 3}, {2}},
			expected: 2.0,
		},
		{
			name:     "[1,2][3,4]",
			input:    [...][]int{{1, 2}, {3, 4}},
			expected: 2.5,
		},
		// and some tests from failed submissions
		{
			name:     "[3][-2,-1]",
			input:    [...][]int{{3}, {-2, -1}},
			expected: -1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			m := Mean(tc.input[0], tc.input[1])
			assert.Equal(t, tc.expected, m)
		})
	}
}
