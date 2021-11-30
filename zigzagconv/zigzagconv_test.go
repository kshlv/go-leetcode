package zigzagconv

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	testcases := []struct {
		s, expected string
		numRows     int
	}{
		{
			s:        "",
			numRows:  0,
			expected: "",
		},
		{
			s:        "A",
			numRows:  2,
			expected: "A",
		},
		// Examples
		{
			s:       "PAYPALISHIRING",
			numRows: 3,
			// P   A   H   N
			//  A P L S I I G
			//   Y   I   R
			expected: "PAHNAPLSIIGYIR",
		},
		{
			s:       "PAYPALISHIRING",
			numRows: 4,
			// P     I     N
			//  A   L S   I G
			//   Y A   H R
			//    P     I
			expected: "PINALSIGYAHRPI",
		},
		{
			s:        "A",
			numRows:  1,
			expected: "A",
		},
		// Wrong
		{
			s:        "AB",
			numRows:  1,
			expected: "AB",
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			conv := Convert(tc.s, tc.numRows)
			assert.Equal(t, tc.expected, conv)
		})
	}
}
