package clauses

import (
	"slices"
	"testing"
)

func SlicesEq(c1, c2 [][]int) bool {
	return slices.EqualFunc(
		c1,
		c2,
		func(l1, l2 []int) bool {
			return slices.Equal(l1, l2)
		},
	)
}

func ValidateClauses(
	t *testing.T,
	sClauses, cClauses, expSClauses, expCClauses [][]int,
) {
	t.Helper()
	if !SlicesEq(sClauses, expSClauses) {
		t.Errorf(
			"Semantic clauses not equal. Expected %d but got %d",
			sClauses,
			expSClauses,
		)
	}
	if !SlicesEq(cClauses, expCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expCClauses,
		)
	}
}
