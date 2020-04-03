package solution

import (
	"fmt"
	"math/big"
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
	s := make([]int, 0, numCount(l))
	ln := l
	for ln != nil {
		s = append(s, ln.Val)
		ln = ln.Next
	}
	return s
}

func validateAnswer(l1 *ListNode, l2 *ListNode) error {
	if numCount(l1) != numCount(l2) {
		return fmt.Errorf("length l1 (%v) is not equal to l2 (%v)", createSlice(l1), createSlice(l2))
	}
	ln1 := l1
	ln2 := l2
	pos := 0
	for int64(pos) < numCount(ln1) {
		if ln1.Val != ln2.Val {
			return fmt.Errorf("l1 (%v) is not equal to l2 (%v)", createSlice(l1), createSlice(l2))
		}
		ln1 = ln1.Next
		ln2 = ln2.Next
		pos++
	}
	return nil
}

func TestAddTwoNumbers(t *testing.T) {
	type test struct {
		l1 []int
		l2 []int
		a  []int
	}
	tests := []test{
		{
			l1: []int{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			l2: []int{5, 6, 4},
			a:  []int{6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
		},
		{
			l1: []int{0},
			l2: []int{0},
			a:  []int{0},
		},
		{
			l1: []int{2, 4, 3},
			l2: []int{5, 6, 4},
			a:  []int{7, 0, 8},
		},
	}
	for _, tN := range tests {
		if err := validateAnswer(addTwoNumbers(createListNodeSequence(tN.l1), createListNodeSequence(tN.l2)), createListNodeSequence(tN.a)); err != nil {
			t.Error(err)
		}
	}
}

func TestNumCount(t *testing.T) {
	type test struct {
		l []int
		a int64
	}
	tests := []test{
		{[]int{1, 2, 3}, 3},
		{[]int{1, 2, 3, 4}, 4},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 20},
	}
	for _, tN := range tests {
		if numCount(createListNodeSequence(tN.l)) != tN.a {
			t.Fail()
		}
	}
}

func TestExtractNumber(t *testing.T) {
	type test struct {
		n []int
		a *big.Int
	}
	tests := []test{
		{
			n: []int{6, 4, 5},
			a: big.NewInt(546),
		},
		{
			n: []int{6, 4, 5, 0},
			a: big.NewInt(546),
		},
		{
			n: []int{0, 6, 4, 5},
			a: big.NewInt(5460),
		},
	}
	for _, tN := range tests {
		if n := extractNumber(createListNodeSequence(tN.n), int64(len(tN.n))); n.Cmp(tN.a) != 0 {
			t.Errorf("n == %s, but expected is %s", n.String(), tN.a.String())
		}
	}
}

func TestCreateListNode(t *testing.T) {
	type test struct {
		n *big.Int
		a []int
	}
	tests := []test{
		{
			n: big.NewInt(546),
			a: []int{6, 5, 4},
		},
	}
	for _, tN := range tests {
		if n := createListNode(*tN.n); n.Val != tN.a[0] {
			t.Errorf("n == %d, but expected is %d", n.Val, tN.a)
		}
	}
}
