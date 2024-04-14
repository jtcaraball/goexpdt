package base

import (
	"goexpdt/internal/test/clauses"
	"testing"
)

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
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}
