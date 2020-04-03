package solution

import (
	"math/big"
	"strconv"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	n := big.NewInt(0)
	c := numCount(l1)
	if s := numCount(l2); s > c {
		c = s
	}

	n = n.Add(extractNumber(l1, c), extractNumber(l2, c))

	return createListNode(*n)
}

func createListNode(n big.Int) *ListNode {
	if n.Cmp(big.NewInt(0)) == 0 {
		return &ListNode{0, nil}
	}
	list := make([]int64, 0, len(n.String()))
	number := n.String()

	y := 0
	for y < len(number) {
		letter, err := strconv.Atoi(string(number[y]))
		if err != nil {
			panic(err)
		}
		list = append(list, int64(letter))
		y++
	}

	var next, l *ListNode
	for i := 0; i < len(list); i++ {
		l = &ListNode{int(list[i]), next}
		next = l
	}

	return l
}

func extractNumber(l *ListNode, s int64) *big.Int {
	a := big.NewInt(0)
	n := numCount(l)
	m := big.NewInt(1)

	ln := l
	for ln != nil {
		z := big.NewInt(0)
		a = a.Add(a, z.Mul(big.NewInt(int64(ln.Val)), m))
		m.Mul(m, big.NewInt(10))
		ln = ln.Next
	}

	for (s - n) > 0 {
		m.Mul(m, big.NewInt(10))
		s--
	}

	return a
}

func numCount(l *ListNode) int64 {
	n := int64(0)

	ln := l
	for ln != nil {
		n++
		ln = ln.Next
	}

	return n
}
