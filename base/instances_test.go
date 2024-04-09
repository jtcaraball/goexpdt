package base

import (
	"slices"
	"testing"
)

func slicesEq(c1, c2 [][]int) bool {
	return slices.EqualFunc(
		c1,
		c2,
		func(l1, l2 []int) bool {
			return slices.Equal(l1, l2)
		},
	)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encode(t *testing.T) {
	var Variable Var = "testVar"
	ctx := NewContext(3, nil)
	encCNF := Variable.Encoding(ctx)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{-1, -2},
		{-1, -3},
		{-2, -3},
		{-4, -5},
		{-4, -6},
		{-5, -6},
		{-7, -8},
		{-7, -9},
		{-8, -9},
	}
	if !slicesEq(sClauses, expSClauses) {
		t.Errorf(
			"Clauses not equal. Expected %d but got %d",
			expSClauses,
			sClauses,
		)
	}
	if !slicesEq(cClauses, expCClauses) {
		t.Errorf(
			"Clauses not equal. Expected %d but got %d",
			expCClauses,
			cClauses,
		)
	}
}
