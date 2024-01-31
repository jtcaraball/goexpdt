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

func errorInClauses(
	t *testing.T,
	sClauses, cClauses, expSClauses, expCClauses [][]int,
) {
	if !clausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Semantic clauses not equal. Expected %d but got %d",
			sClauses,
			expSClauses,
		)
	}
	if !clausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expCClauses,
		)
	}
}

// =========================== //
//            TESTS            //
// =========================== //

func TestCNF_Negate_Empty(t *testing.T) {
	expectedSClauses := [][]int{{}}
	expectedCClauses := [][]int{}
	cnf := &CNF{tv: 3, cClauses: expectedCClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 3 {
		t.Errorf("NV not equal. Expected %d but got %d", 3, cnf.tv)
	}
}

func TestCNF_Negate_SingleClauseEmpty(t *testing.T) {
	sInitClauses := [][]int{{}}
	cInitClauses := [][]int{{1, 2, 3}}
	expectedSClauses := [][]int{}
	expectedCClauses := [][]int{}
	cnf := &CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 0 {
		t.Errorf("NV not equal. Expected %d but got %d", 0, cnf.tv)
	}
}

func TestCNF_Negate_SingleLiteral(t *testing.T) {
	sInitClauses := [][]int{{1}}
	cInitClauses := [][]int{{1, 2, 3}}
	expectedSClauses := [][]int{{-1}}
	expectedCClauses := cInitClauses
	cnf := &CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 3 {
		t.Errorf("NV not equal. Expected %d but got %d", 3, cnf.tv)
	}
}

func TestCNF_Negate_SingleClause(t *testing.T) {
	sInitClauses := [][]int{{1, 2}}
	cInitClauses := [][]int{{1, 2, 3}}
	expectedSClauses := [][]int{{4}}
	expectedCClauses := [][]int{
		{1, 2, 3},
		{-1, -4},
		{-2, -4},
		{1, 2, 4},
	}
	cnf := &CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 4 {
		t.Errorf("NV not equal. Expected %d but got %d", 4, cnf.tv)
	}
}

func TestCNF_Negate_MultipleClauses(t *testing.T) {
	sInitClauses := [][]int{{1, 2}, {-2, 3}}
	cInitClauses := [][]int{{1, 2, 3}}
	expectedSClauses := [][]int{{6}}
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
	cnf := &CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 6 {
		t.Errorf("NV not equal. Expected %d but got %d", 6, cnf.tv)
	}
}

func TestCNF_Conjunctions(t *testing.T) {
	expectedSClauses := [][]int{{1, 2}, {2, 3}, {-3, 4}}
	expectedCClauses := [][]int{{1, 2}}
	cnf1SInitClauses := [][]int{{1, 2}, {2, 3}}
	cnf2SInitClauses := [][]int{{-3, 4}}
	cnf1CInitClauses := [][]int{{1, 2}}
	cnf1 := &CNF{tv: 3, sClauses: cnf1SInitClauses, cClauses: cnf1CInitClauses}
	cnf2 := &CNF{tv: 4, sClauses: cnf2SInitClauses}
	cnf1.Conjunction(cnf2)
	sClauses, cClauses := cnf1.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf1.tv != 4 {
		t.Errorf("NV not equal. Expected %d but got %d", 4, cnf1.tv)
	}
}

func TestCNF_ExtendSemantic(t *testing.T) {
	expectedSClauses := [][]int{{1, 2}, {2, 3}, {-3, 4}}
	expectedCClauses := [][]int{}
	cnfSInitClauses := [][]int{{1, 2}}
	cnfCInitClauses := [][]int{}
	cnf := &CNF{tv: 2, sClauses: cnfSInitClauses, cClauses: cnfCInitClauses}
	cnf.ExtendSemantics([][]int{{2, 3}, {-3, 4}})
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 4 {
		t.Errorf("NV not equal. Expected %d but got %d", 4, cnf.tv)
	}
}

func TestCNF_ExtendConsistency(t *testing.T) {
	expectedSClauses := [][]int{{1, 2}}
	expectedCClauses := [][]int{{2, 3}, {-3, 4}}
	cnfSInitClauses := [][]int{{1, 2}}
	cnfCInitClauses := [][]int{}
	cnf := &CNF{tv: 2, sClauses: cnfSInitClauses, cClauses: cnfCInitClauses}
	cnf.ExtendConsistency([][]int{{2, 3}, {-3, 4}})
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	if cnf.tv != 4 {
		t.Errorf("NV not equal. Expected %d but got %d", 4, cnf.tv)
	}
}
