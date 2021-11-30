package kreverse

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	testcases := []struct {
		vals, expected []int
	}{
		{
			vals:     []int{1},
			expected: []int{1},
		},
		{
			vals:     []int{1, 2},
			expected: []int{2, 1},
		},
		{
			vals:     []int{1, 2, 3},
			expected: []int{3, 2, 1},
		},
		{
			vals:     []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			vals:     []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			assert.Equal(t, tc.expected, reverse(tc.vals))
		})
	}
}

func TestKReverse(t *testing.T) {
	testcases := []struct {
		head, expected *ListNode
		k              int
	}{
		{
			head:     &ListNode{},
			k:        1,
			expected: &ListNode{},
		},
		{
			head:     &ListNode{1, &ListNode{2, nil}},
			k:        2,
			expected: &ListNode{2, &ListNode{1, nil}},
		},
		// Examples
		{
			head:     &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:        3,
			expected: &ListNode{3, &ListNode{2, &ListNode{1, &ListNode{5, &ListNode{5, nil}}}}},
		},
		{
			head:     &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:        2,
			expected: &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}},
		},
		{
			head:     &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			k:        1,
			expected: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
		},
		{
			head:     &ListNode{1, nil},
			k:        1,
			expected: &ListNode{1, nil},
		},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			rev := KReverse(tc.head, tc.k)
			ie, ir := tc.expected, rev
			assert.Equal(t, ie.Val, ir.Val)
			for ie.Next != nil && ir.Next != nil {
				ie = ie.Next
				ir = ir.Next
			}
		})
	}

}
