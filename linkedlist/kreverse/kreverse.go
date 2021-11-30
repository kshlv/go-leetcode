// Package kreverse implements a reversal
// on groups of k on a linked list
//
// Original problem:
// https://leetcode.com/problems/reverse-nodes-in-k-group/
package kreverse

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// List ...
type List struct {
	Head   *ListNode
	Length int
}

func length(list *ListNode) int {
	if list == nil {
		return 0
	}

	count := 1
	cur := list
	for cur.Next != nil {
		cur = cur.Next
		count++
	}

	return count
}

// New ...
func New(h *ListNode) *List {
	return &List{
		Length: length(h),
		Head:   h,
	}
}

func reverse(vals []int) []int {
	result := []int{}
	for _, val := range vals {
		result = append(result, val)
	}

	l := len(vals)
	for i := 0; i < l/2; i++ {
		result[i], result[l-i-1] = result[l-i-1], result[i]
	}

	return result
}

// Reverse ...
func (l *List) Reverse(k int) {
	rab, tur := l.Head, l.Head
	for p := 0; p < l.Length/k; p++ {
		arr := []int{}
		for i := 0; i < k; i++ {
			arr = append(arr, rab.Val)
			rab = rab.Next
		}
		rarr := reverse(arr)
		for i := 0; i < k; i++ {
			tur.Val = rarr[i]
			tur = tur.Next
		}
	}
}

// KReverse is ...
// reverseKGroup(head *ListNode, k int) *ListNode
func KReverse(head *ListNode, k int) *ListNode {
	l := New(head)
	l.Reverse(k)
	return l.Head
}
