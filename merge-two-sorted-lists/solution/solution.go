package solution

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var first, counter *ListNode

	if l1.Val < l2.Val {
		counter = l1
		l1 = l1.Next
	} else {
		counter = l2
		l2 = l2.Next
	}
	first = counter
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			counter.Next = l1
			l1 = l1.Next
		} else {
			counter.Next = l2
			l2 = l2.Next
		}
		counter = counter.Next
	}
	if l1 != nil && l2 != nil {
		panic("For loop exited before one of the lists was drained")
	}
	if l1 != nil {
		counter.Next = l1
	}
	if l2 != nil {
		counter.Next = l2
	}

	return first
}
