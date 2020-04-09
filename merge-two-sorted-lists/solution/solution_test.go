package solution

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createListNodeSequence(n []int) *ListNode {
	var next, l *ListNode
	for i := len(n); i > 0; i-- {
		l = &ListNode{n[i-1], next}
		next = l
	}

	return l
}

func createSlice(l *ListNode) []int {
	s := make([]int, 0, 32)
	ln := l
	for ln != nil {
		s = append(s, ln.Val)
		ln = ln.Next
	}
	return s
}

func validateAnswer(l1 *ListNode, l2 *ListNode) error {
	ln1 := l1
	ln2 := l2
	pos := 0
	for ln1 != nil {
		if ln1.Val != ln2.Val {
			return fmt.Errorf("l1 (%v) is not equal to l2 (%v)", createSlice(l1), createSlice(l2))
		}
		ln1 = ln1.Next
		ln2 = ln2.Next
		pos++
	}
	return nil
}

func TestMergeTwoLists(t *testing.T) {
	type test struct {
		l1 []int
		l2 []int
		a  []int
	}
	tests := []test{
		{
			l1: []int{1, 2, 4},
			l2: []int{1, 3, 4},
			a:  []int{1, 1, 2, 3, 4, 4},
		},
		{
			l1: []int{},
			l2: []int{1, 2, 3, 4},
			a:  []int{1, 2, 3, 4},
		},
		{
			l1: []int{1, 2, 3, 4},
			l2: []int{},
			a:  []int{1, 2, 3, 4},
		},
	}
	for _, tN := range tests {
		if err := validateAnswer(mergeTwoLists(createListNodeSequence(tN.l1), createListNodeSequence(tN.l2)), createListNodeSequence(tN.a)); err != nil {
			t.Error(err)
		}
	}
}
