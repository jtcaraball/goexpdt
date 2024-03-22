package base

import (
	"testing"
	"stratifoiled/internal/test"
)

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encode(t *testing.T) {
	var Variable Var = "testVar"
	ctx := NewContext(1, nil)
	encCNF := Variable.Encoding(ctx)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
	}
	if !test.ClausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Clauses not equal. Expected %d but got %d",
			expSClauses,
			sClauses,
		)
	}
	if !test.ClausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Clauses not equal. Expected %d but got %d",
			expCClauses,
			cClauses,
		)
	}
}

