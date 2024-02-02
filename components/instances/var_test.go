package instances

import (
	"slices"
	"stratifoiled/components"
	"testing"
)

// =========================== //
//           HELPERS           //
// =========================== //

func clausesEq(c1, c2 [][]int) bool {
	return slices.EqualFunc[[][]int](
		c1,
		c2,
		func (l1, l2 []int) bool {
			return slices.Equal[[]int](l1, l2)
		},
	)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encode(t *testing.T) {
	var Variable Var = "testVar"
	ctx := components.NewContext(1, nil)
	encCNF := Variable.Encoding(ctx)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
	}
	if !clausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Clauses not equal. Expected %d but got %d",
			expSClauses,
			sClauses,
		)
	}
	if !clausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Clauses not equal. Expected %d but got %d",
			expCClauses,
			cClauses,
		)
	}
}