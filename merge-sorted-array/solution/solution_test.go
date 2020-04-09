package solution

import (
	"fmt"
	"testing"
)

func validateAnswer(l1 []int, l2 []int) error {
	if len(l1) != len(l2) {
		return fmt.Errorf("List are not equal length: l1 == %d and l2 == %d", len(l1), len(l2))
	}

	for i := 0; i < len(l1); i++ {
		if l1[i] != l2[i] {
			return fmt.Errorf("List are not equal at pos %d: l1 == %d and l2 == %d", i, l1[i], l2[i])
		}
	}

	return nil
}

func TestMergeTwoLists(t *testing.T) {
	type test struct {
		l1 []int
		m  int
		l2 []int
		n  int
		a  []int
	}
	tests := []test{
		{
			l1: []int{1, 2, 3, 0, 0, 0},
			m:  3,
			l2: []int{2, 5, 6},
			n:  3,
			a:  []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tN := range tests {
		merge(tN.l1, tN.m, tN.l2, tN.n)
		if err := validateAnswer(tN.l1, tN.a); err != nil {
			t.Error(fmt.Sprintf("Failed with error '%v' with values {l1: '%v' l2: '%v'}", err, tN.l1, tN.a))
		}
	}
}
