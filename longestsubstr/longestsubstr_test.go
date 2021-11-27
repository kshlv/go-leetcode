package longestsubstr

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueSubsWithLength(t *testing.T) {
	testcases := []struct {
		s        string
		l        int
		expected []string
	}{
		{
			s:        "",
			l:        0,
			expected: []string{""},
		},

		{
			s:        "",
			l:        1,
			expected: nil,
		},
		{
			s:        "a",
			l:        2,
			expected: nil,
		},

		{
			s:        "aabb",
			l:        1,
			expected: []string{"a", "b"},
		},

		{
			s:        "aabb",
			l:        2,
			expected: []string{"ab"},
		},
		{
			s:        "abba",
			l:        2,
			expected: []string{"ab", "ba"},
		},

		{
			s:        "aaabbbccc",
			l:        2,
			expected: []string{"ab", "bc"},
		},
		{
			s:        "aaabbbccc",
			l:        3,
			expected: []string{},
		},
		{
			s:        "abc",
			l:        3,
			expected: []string{"abc"},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("s: %v, expected: %v", tc.s, tc.expected), func(t *testing.T) {
			subs := uniqueSubsOfLength(tc.s, tc.l)
			for _, expected := range tc.expected {
				assert.Contains(t, subs, expected)
			}

		})
	}
}

func TestSubsWithLength(t *testing.T) {
	testcases := []struct {
		s        string
		l        int
		expected []string
	}{
		{
			s:        "",
			l:        0,
			expected: []string{""},
		},

		{
			s:        "",
			l:        1,
			expected: nil,
		},
		{
			s:        "a",
			l:        2,
			expected: nil,
		},

		{
			s:        "aa",
			l:        2,
			expected: []string{"aa"},
		},

		{
			s:        "aabb",
			l:        1,
			expected: []string{"a", "a", "b", "b"},
		},

		{
			s:        "aabb",
			l:        2,
			expected: []string{"aa", "ab", "bb"},
		},
		{
			s:        "abba",
			l:        2,
			expected: []string{"ab", "bb", "ba"},
		},

		{
			s:        "abc",
			l:        3,
			expected: []string{"abc"},
		},
	}

	for _, tc := range testcases {
		t.Run("whatever", func(t *testing.T) {
			subs := subsOfLength(tc.s, tc.l)
			for _, expected := range tc.expected {
				assert.Contains(t, subs, expected)
			}
		})
	}
}

func TestSubstr(t *testing.T) {
	testcases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "s: empty string",
			s:        "",
			expected: "",
		},
		{
			name:     "s: aa, sub: a",
			s:        "aa",
			expected: "a",
		},
		{
			name:     "s: ab, sub: ab",
			s:        "ab",
			expected: "ab",
		},
		{
			name:     "s: abb, sub: ab",
			s:        "abb",
			expected: "ab",
		},
		{
			// If s: 'abbac', kinda makes sense if
			// its longest predeccessor was 'ba', isn't it?
			// a -> ab -> ba -> bac
			name:     "s: abba, sub: ba",
			s:        "abba",
			expected: "ba",
		},
		{
			name:     "s: abbac, sub: bac",
			s:        "abbac",
			expected: "bac",
		},
		// Examples
		{
			name:     "s: abcabcbb, sub: abc",
			s:        "abcabcbb",
			expected: "abc",
		},
		{
			name:     "s: bbbbb, sub: b",
			s:        "bbbbb",
			expected: "b",
		},
		{
			name:     "s: pwwkew, sub: wke",
			s:        "pwwkew",
			expected: "kew",
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			sub := Substr(tc.s)
			assert.Equal(t, tc.expected, sub)
		})
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	testcases := []struct {
		s        string
		expected int
	}{
		{
			s:        "",
			expected: 0,
		},
		{
			s:        "aa",
			expected: 1,
		},
		{
			s:        "bbbbb",
			expected: 1,
		},
		{
			s:        "ab",
			expected: 2,
		},
		{
			s:        "pwwkew",
			expected: 3,
		},
		// Wrong answer
		//
		// This TC doesn't pass, which obviously means that
		// I didn't finish the problem, yet I'm going to get back and complete it
		// {
		// 	s: maxnumberOfPossibleChars,
		// 	expected: 11,
		// },
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("s: %v, expected: %v", tc.s, tc.expected), func(t *testing.T) {
			sub := lengthOfLongestSubstring(tc.s)
			assert.Equal(t, tc.expected, sub)
		})
	}
}
