package powfloatint

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Constraints:
// -100.0 < x < 100.0
// -231 <= n <= 231-1
// -104 <= x^n <= 104
func TestPow(t *testing.T) {
	testcases := []struct {
		x, expected float64
		n           int
	}{
		// Examples
		{
			x:        2.0,
			n:        -2,
			expected: 0.25,
		},
		{
			x:        2.0,
			n:        10,
			expected: 1024.0,
		},
		{
			x:        2.1,
			n:        3,
			expected: math.Pow(2.1, 3.0),
		},
		// Timeout
		{
			x:        1.00000,
			n:        -2147483648,
			expected: 1.0,
		},
		{
			x:        0.999999999,
			n:        -2147483648,
			expected: math.Pow(0.999999999, -2147483648),
		},
		// Wrong answer
		{
			x:        -1.00000,
			n:        -2147483648,
			expected: 1.0,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Pow(%f, %d) -> %f", tc.x, tc.n, tc.expected), func(t *testing.T) {
			assert.Equal(t, tc.expected, Pow(tc.x, tc.n))
		})
	}
}
