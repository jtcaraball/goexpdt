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
