package components

import (
	"slices"
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

func TestCNF_Negate_Empty(t *testing.T) {
	expectedMClauses := [][]int{{}}
	expectedCClauses := [][]int{{1, 2, 3}}
	cnf := &CNF{}
	cnf.AppendConsistency(expectedCClauses[0])
	cnf.Negate()
	mClauses, cClauses := cnf.Clauses()
	if !clausesEq(mClauses, expectedMClauses) {
		t.Errorf(
			"Meaning clauses not equal. Expected %d but got %d",
			mClauses,
			expectedMClauses,
		)
	}
	if !clausesEq(cClauses, expectedCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expectedCClauses,
		)
	}
}

func TestCNF_Negate_SingleClauseEmpty(t *testing.T) {
	expectedClauses := [][]int{}
	cnf := &CNF{}
	cnf.AppendMeaning([]int{})
	cnf.Negate()
	mClauses, _ := cnf.Clauses()
	if !clausesEq(mClauses, expectedClauses) {
		t.Errorf(
			"Meaning clauses not equal. Expected %d but got %d",
			mClauses,
			expectedClauses,
		)
	}
}

func TestCNF_Negate_SingleLiteral(t *testing.T) {
	expectedMClauses := [][]int{{-1}}
	expectedCClauses := [][]int{{1, 2, 3}}
	cnf := &CNF{}
	cnf.AppendMeaning([]int{1})
	cnf.AppendConsistency(expectedCClauses[0])
	cnf.Negate()
	mClauses, cClauses := cnf.Clauses()
	if !clausesEq(mClauses, expectedMClauses) {
		t.Errorf(
			"Meaning clauses not equal. Expected %d but got %d",
			mClauses,
			expectedMClauses,
		)
	}
	if !clausesEq(cClauses, expectedCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expectedCClauses,
		)
	}
}

func TestCNF_Negate_SingleClauses(t *testing.T) {
	mClausesToAdd := [][]int{{1, 2}}
	cClausesToAdd := [][]int{{1, 2, 3}}
	expectedMClauses := [][]int{{4}}
	expectedCClauses := [][]int{
		{1, 2, 3},
		{-1, -4},
		{-2, -4},
		{1, 2, 4},
	}
	cnf := &CNF{}
	cnf.ExtendMeaning(mClausesToAdd)
	cnf.ExtendConsistency(cClausesToAdd)
	cnf.Negate()
	mClauses, cClauses := cnf.Clauses()
	if !clausesEq(mClauses, expectedMClauses) {
		t.Errorf(
			"Meaning clauses not equal. Expected %d but got %d",
			mClauses,
			expectedMClauses,
		)
	}
	if !clausesEq(cClauses, expectedCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expectedCClauses,
		)
	}
}

func TestCNF_Negate_MultipleClauses(t *testing.T) {
	mClausesToAdd := [][]int{{1, 2}, {-2, 3}}
	cClausesToAdd := [][]int{{1, 2, 3}}
	expectedMClauses := [][]int{{6}}
	expectedCClauses := [][]int{
		{1, 2, 3},
		{-1, -4},
		{-2, -4},
		{1, 2, 4},
		{2, -5},
		{-3, -5},
		{-2, 3, 5},
		{-4, 6},
		{-5, 6},
		{4, 5, -6},
	}
	cnf := &CNF{}
	cnf.ExtendMeaning(mClausesToAdd)
	cnf.ExtendConsistency(cClausesToAdd)
	cnf.Negate()
	mClauses, cClauses := cnf.Clauses()
	if !clausesEq(mClauses, expectedMClauses) {
		t.Errorf(
			"Meaning clauses not equal. Expected %d but got %d",
			mClauses,
			expectedMClauses,
		)
	}
	if !clausesEq(cClauses, expectedCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expectedCClauses,
		)
	}
}
