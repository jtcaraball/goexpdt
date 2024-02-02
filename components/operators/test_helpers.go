package operators

import (
	"slices"
	"testing"
)

func clausesEq(c1, c2 [][]int) bool {
	return slices.EqualFunc[[][]int](
		c1,
		c2,
		func (l1, l2 []int) bool {
			return slices.Equal[[]int](l1, l2)
		},
	)
}

func componentsEq(c1, c2 []*any) bool {
	if len(c1) != len(c2) {
		return false
	}
	for i, elem1 := range c1 {
		if elem1 != c2[i] {
			return false
		}
	}
	return true
}

func errorInClauses(
	t *testing.T,
	sClauses, cClauses, expSClauses, expCClauses [][]int,
) {
	t.Helper()
	if !clausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Semantic clauses not equal. Expected %d but got %d",
			expSClauses,
			sClauses,
		)
	}
	if !clausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			expCClauses,
			cClauses,
		)
	}
}
