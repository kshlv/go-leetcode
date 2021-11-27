package atoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var whitespaces []byte = []byte{' ', '\t', '\n', '\r'}

func TestAtoi(t *testing.T) {
	testcases := []struct {
		name, a  string
		expected int
	}{
		{
			a:        "+-+-+",
			expected: 0,
		},
		{
			a:        "-+-+-",
			expected: 0,
		},
		{
			name:     "\"0\"",
			a:        "0",
			expected: 0,
		},
		{
			name:     "\"0032\" → int 32",
			a:        "0032",
			expected: 32,
		},
		{
			name:     "\"213456789\" → int 213456789",
			a:        "213456789",
			expected: 213_456_789,
		},
		// Examples
		{
			name:     "\"42\" → int 42",
			a:        "42",
			expected: 42,
		},
		{
			name:     "\"   -42\" → int -42",
			a:        "   -42",
			expected: -42,
		},
		{
			name:     "\"4193 with words\" → 4193",
			a:        "4193 with words",
			expected: 4193,
		},
		{
			name:     "\"words and 987\" → 0",
			a:        "words and 987",
			expected: 0,
		},
		{
			name:     "\"-91283472332\" → -91_283_472_332 → (clamp to int lower bound) → -2_147_483_648",
			a:        "-91283472332",
			expected: -2_147_483_648,
		},
		// Wrong answer
		{
			name:     "π",
			a:        "3.14159",
			expected: 3,
		},
		{
			name:     "\"+-12\" → int -12",
			a:        "+-12",
			expected: 0,
		},
		{
			name:     "\"21474836460\" → 21_474_836_460 > 1<<31-1 → 2_147_483_647",
			a:        "21474836460",
			expected: 2_147_483_647,
		},
		{
			name:     "\"   +0 123\" → 0",
			a:        "   +0 123",
			expected: 0,
		},
		{
			name:     "\"9223372036854775808\" → 9_223_372_036_854_775_808",
			a:        "9223372036854775808",
			expected: 2_147_483_647,
		},
		{
			a:        "-5-",
			expected: -5,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			i := Atoi(tc.a)
			assert.Equal(t, tc.expected, i)
		})
	}
}
