package palindromeint

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	testcases := []struct {
		x        int
		expected bool
	}{
		{
			x:        0,
			expected: true,
		},
		{
			x:        1,
			expected: true,
		},
		{
			x:        2122,
			expected: false,
		},
		{
			x:        2112,
			expected: true,
		},
		{
			x:        10201,
			expected: true,
		},
		{
			x:        1_234_554_321,
			expected: true,
		},
		// Examples
		{
			x:        121,
			expected: true,
		},
		{
			x:        -121,
			expected: false,
		},
		{
			x:        10,
			expected: false,
		},
		{
			x:        -101,
			expected: false,
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("x: %d", tc.x), func(t *testing.T) {
			assert.Equal(t, tc.expected, IsPalindrome(tc.x))
		})
	}
}
