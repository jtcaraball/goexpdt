package sfdtest

import (
	"slices"
	"testing"
)

func ClausesEq(c1, c2 [][]int) bool {
	return slices.EqualFunc[[][]int](
		c1,
		c2,
		func (l1, l2 []int) bool {
			return slices.Equal[[]int](l1, l2)
		},
	)
}

func ErrorInClauses(
	t *testing.T,
	sClauses, cClauses, expSClauses, expCClauses [][]int,
) {
	t.Helper()
	if !ClausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Semantic clauses not equal. Expected %d but got %d",
			expSClauses,
			sClauses,
		)
	}
	if !ClausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			expCClauses,
			cClauses,
		)
	}
}
