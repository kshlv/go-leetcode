package reverseint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	testcases := []struct {
		x, y, expected int
	}{
		{
			x:        0,
			y:        0,
			expected: 0,
		},
		{
			x:        1,
			y:        0,
			expected: 1,
		},
		{
			x:        -1,
			y:        1,
			expected: -1,
		},
	}

	for _, tc := range testcases {
		t.Run("whatever", func(t *testing.T) {
			pow := pow(tc.x, tc.y)
			assert.Equal(t, tc.expected, pow)
		})
	}
}

// Constraints:
//
//          -1<<31 <= x <= 1<<31 - 1
//          -2**31 <= x <= 2**31 - 1
//  -2_147_483_648 <= x <= 2_147_483_647
func TestReverse(t *testing.T) {
	testcases := []struct {
		name        string
		x, expected int
	}{
		{
			name:     "100 → 001 → 1",
			x:        100,
			expected: 1,
		},
		{
			name:     "102 → 201",
			expected: 102,
			x:        201,
		},
		// The problem's examples
		{
			name:     "123 → 321",
			x:        123,
			expected: 321,
		},
		{
			name:     "-123 → -321",
			x:        -123,
			expected: -321,
		},
		{
			name:     "120 → 021 → 21",
			x:        120,
			expected: 21,
		},
		{
			name:     "0 → 0",
			x:        0,
			expected: 0,
		},
		// After-submit
		{
			name:     "1_534_236_469 !→ 9_646_324_351: Output overflows int32",
			x:        1_534_236_469,
			expected: 0,
		},
		{
			name:     "-2_147_483_412 → -2_143_847_412",
			x:        -2_147_483_412,
			expected: -2_143_847_412, // -21_438_474_120
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			rev := Reverse(tc.x)
			assert.Equal(t, tc.expected, rev)
		})
	}
}
