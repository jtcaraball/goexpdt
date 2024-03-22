package cnf

import (
	"os"
	"slices"
	"stratifoiled/internal/test"
	"testing"
)

// =========================== //
//           HELPERS           //
// =========================== //

func errorInClauses(
	t *testing.T,
	sClauses, cClauses, expSClauses, expCClauses [][]int,
	topv, expTopV int,
) {
	t.Helper()
	if !test.ClausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Semantic clauses not equal. Expected %d but got %d",
			sClauses,
			expSClauses,
		)
	}
	if !test.ClausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expCClauses,
		)
	}
	if topv != expTopV {
		t.Errorf("NV not equal. Expected %d but got %d", expTopV, topv)
	}
}

// =========================== //
//            TESTS            //
// =========================== //

func TestCNF_Negate_Empty(t *testing.T) {
	expectedSClauses := [][]int{{}}
	expectedCClauses := [][]int{}
	cnf := &CNF{tv: 0, cClauses: expectedCClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		0,
	)
}

func TestCNF_Negate_SingleClauseEmpty(t *testing.T) {
	sInitClauses := [][]int{{}}
	cInitClauses := [][]int{{1, 2, 3}}
	expectedSClauses := [][]int{}
	expectedCClauses := [][]int{}
	cnf := &CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		3,
	)
}

func TestCNF_Negate_SingleLiteral(t *testing.T) {
	sInitClauses := [][]int{{1}}
	cInitClauses := [][]int{{1, 2, 3}}
	expectedSClauses := [][]int{{-1}}
	expectedCClauses := cInitClauses
	cnf := &CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	cnf.Negate()
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		3,
	)
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
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		4,
	)
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
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		6,
	)
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
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf1.tv,
		4,
	)
}

func TestCNF_ExtendSemantic(t *testing.T) {
	expectedSClauses := [][]int{{1, 2}, {2, 3}, {-3, 4}}
	expectedCClauses := [][]int{}
	cnfSInitClauses := [][]int{{1, 2}}
	cnfCInitClauses := [][]int{}
	cnf := &CNF{tv: 2, sClauses: cnfSInitClauses, cClauses: cnfCInitClauses}
	cnf.ExtendSemantics([][]int{{2, 3}, {-3, 4}})
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		4,
	)
}

func TestCNF_ExtendConsistency(t *testing.T) {
	expectedSClauses := [][]int{{1, 2}}
	expectedCClauses := [][]int{{2, 3}, {-3, 4}}
	cnfSInitClauses := [][]int{{1, 2}}
	cnfCInitClauses := [][]int{}
	cnf := &CNF{tv: 2, sClauses: cnfSInitClauses, cClauses: cnfCInitClauses}
	cnf.ExtendConsistency([][]int{{2, 3}, {-3, 4}})
	sClauses, cClauses := cnf.Clauses()
	errorInClauses(
		t,
		sClauses,
		cClauses,
		expectedSClauses,
		expectedCClauses,
		cnf.tv,
		4,
	)
}

func TestCNF_ToBytes(t *testing.T) {
	cnfSClauses := [][]int{{1, 2}, {3, 4}, {-1, 2}}
	cnfCClauses := [][]int{{1, 2, 3, 4}, {-1, -2}, {4}}
	cnf := CNFFromClauses(cnfSClauses)
	cnf.ExtendConsistency(cnfCClauses)
	expBytes := []byte("p cnf 4 6\n1 2 0\n3 4 0\n-1 2 0\n1 2 3 4 0\n-1 -2 0\n4 0\n")
	cnfBytes := cnf.ToBytes()
	if !slices.Equal[[]byte](expBytes, cnfBytes) {
		t.Errorf(
			"CNF as bytes not equal. Expected %s but got %s",
			expBytes,
			cnfBytes,
		)
	}
}

func TestCNF_ToFile(t *testing.T) {
	cnfFileName := "cnfToFile"
	t.Cleanup(
		func() {
			os.Remove(cnfFileName)
		},
	)
	cnfSClauses := [][]int{{1, 2}, {3, 4}, {-1, 2}}
	cnfCClauses := [][]int{{1, 2, 3, 4}, {-1, -2}, {4}}
	cnf := CNFFromClauses(cnfSClauses)
	cnf.ExtendConsistency(cnfCClauses)
	if err := cnf.ToFile(cnfFileName); err != nil {
		t.Errorf("File writing error. %s", err.Error())
	}
	expBytes := []byte("p cnf 4 6\n1 2 0\n3 4 0\n-1 2 0\n1 2 3 4 0\n-1 -2 0\n4 0\n")
	cnfBytes, err := os.ReadFile(cnfFileName)
	if err != nil {
		t.Errorf("File reading error. %s", err.Error())
	}
	if !slices.Equal[[]byte](expBytes, cnfBytes) {
		t.Errorf(
			"CNF as bytes not equal. Expected %s but got %s",
			expBytes,
			cnfBytes,
		)
	}
}
