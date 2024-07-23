package cnf

import (
	"os"
	"slices"
	"testing"
)

func validTopV(t *testing.T, got, expected int) {
	if got != expected {
		t.Errorf("NV not equal. Expected %d but got %d", expected, got)
	}
}

func validClauses(t *testing.T, sc, cc, esc, ecc []Clause) {
	t.Helper()
	if !slices.EqualFunc(sc, esc, slices.Equal) {
		t.Errorf("Semantic clauses not equal. Expected %d but got %d", esc, sc)
	}
	if !slices.EqualFunc(cc, ecc, slices.Equal) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			ecc,
			cc,
		)
	}
}

func TestCNF_Negate_Empty(t *testing.T) {
	expectedSClauses := NegClauses
	expectedCClauses := []Clause{}
	tcnf := CNF{tv: 0, cClauses: expectedCClauses}
	tcnf = tcnf.Negate()
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 0)
}

func TestCNF_Negate_SingleClauseEmpty(t *testing.T) {
	sInitClauses := NegClauses
	cInitClauses := []Clause{{1, 2, 3}}
	expectedSClauses := []Clause{}
	expectedCClauses := cInitClauses
	tcnf := CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	tcnf = tcnf.Negate()
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 3)
}

func TestCNF_Negate_SingleLiteral(t *testing.T) {
	sInitClauses := []Clause{{1}}
	cInitClauses := []Clause{{1, 2, 3}}
	expectedSClauses := []Clause{{-1}}
	expectedCClauses := cInitClauses
	tcnf := CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	tcnf = tcnf.Negate()
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 3)
}

func TestCNF_Negate_SingleClause(t *testing.T) {
	sInitClauses := []Clause{{1, 2}}
	cInitClauses := []Clause{{1, 2, 3}}
	expectedSClauses := []Clause{{4}}
	expectedCClauses := []Clause{
		{1, 2, 3},
		{-1, -4},
		{-2, -4},
		{1, 2, 4},
	}
	tcnf := CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	tcnf = tcnf.Negate()
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 4)
}

func TestCNF_Negate_MultipleClauses(t *testing.T) {
	sInitClauses := []Clause{{1, 2}, {-2, 3}}
	cInitClauses := []Clause{{1, 2, 3}}
	expectedSClauses := []Clause{{6}}
	expectedCClauses := []Clause{
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
	tcnf := CNF{tv: 3, sClauses: sInitClauses, cClauses: cInitClauses}
	tcnf = tcnf.Negate()
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 6)
}

func TestCNF_Conjunctions(t *testing.T) {
	expectedSClauses := []Clause{{1, 2}, {2, 3}, {-3, 4}}
	expectedCClauses := []Clause{{1, 2}}
	tcnf1SInitClauses := []Clause{{1, 2}, {2, 3}}
	tcnf2SInitClauses := []Clause{{-3, 4}}
	tcnf1CInitClauses := []Clause{{1, 2}}
	tcnf1 := CNF{tv: 3, sClauses: tcnf1SInitClauses, cClauses: tcnf1CInitClauses}
	tcnf2 := CNF{tv: 4, sClauses: tcnf2SInitClauses}
	tcnf1 = tcnf1.Conjunction(tcnf2)
	sClauses, cClauses := tcnf1.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf1.tv, 4)
}

func TestCNF_ExtendSemantic(t *testing.T) {
	expectedSClauses := []Clause{{1, 2}, {2, 3}, {-3, 4}}
	expectedCClauses := []Clause{}
	tcnfSInitClauses := []Clause{{1, 2}}
	tcnfCInitClauses := []Clause{}
	tcnf := CNF{tv: 2, sClauses: tcnfSInitClauses, cClauses: tcnfCInitClauses}
	tcnf = tcnf.AppendSemantics(Clause{2, 3}, Clause{-3, 4})
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 4)
}

func TestCNF_ExtendConsistency(t *testing.T) {
	expectedSClauses := []Clause{{1, 2}}
	expectedCClauses := []Clause{{2, 3}, {-3, 4}}
	tcnfSInitClauses := []Clause{{1, 2}}
	tcnfCInitClauses := []Clause{}
	tcnf := CNF{tv: 2, sClauses: tcnfSInitClauses, cClauses: tcnfCInitClauses}
	tcnf = tcnf.AppendConsistency(Clause{2, 3}, Clause{-3, 4})
	sClauses, cClauses := tcnf.Clauses()
	validClauses(t, sClauses, cClauses, expectedSClauses, expectedCClauses)
	validTopV(t, tcnf.tv, 4)
}

func TestCNF_ToBytes(t *testing.T) {
	tcnfSClauses := []Clause{{1, 2}, {3, 4}, {-1, 2}}
	tcnfCClauses := []Clause{{1, 2, 3, 4}, {-1, -2}, {4}}
	tcnf := FromClauses(tcnfSClauses)
	tcnf = tcnf.AppendConsistency(tcnfCClauses...)
	expBytes := []byte("p cnf 4 6\n1 2 0\n3 4 0\n-1 2 0\n1 2 3 4 0\n-1 -2 0\n4 0\n")
	tcnfBytes := tcnf.ToBytes()
	if !slices.Equal(expBytes, tcnfBytes) {
		t.Errorf(
			"CNF as bytes not equal.\nExpected\n%sbut got\n%s",
			expBytes,
			tcnfBytes,
		)
	}
}

func TestCNF_ToFile(t *testing.T) {
	tcnfFileName := "tcnfToFile"
	t.Cleanup(
		func() {
			os.Remove(tcnfFileName)
		},
	)
	tcnfSClauses := []Clause{{1, 2}, {3, 4}, {-1, 2}}
	tcnfCClauses := []Clause{{1, 2, 3, 4}, {-1, -2}, {4}}
	tcnf := FromClauses(tcnfSClauses)
	tcnf = tcnf.AppendConsistency(tcnfCClauses...)
	if err := tcnf.ToFile(tcnfFileName); err != nil {
		t.Errorf("File writing error. %s", err.Error())
	}
	expBytes := []byte("p cnf 4 6\n1 2 0\n3 4 0\n-1 2 0\n1 2 3 4 0\n-1 -2 0\n4 0\n")
	tcnfBytes, err := os.ReadFile(tcnfFileName)
	if err != nil {
		t.Errorf("File reading error. %s", err.Error())
	}
	if !slices.Equal(expBytes, tcnfBytes) {
		t.Errorf(
			"CNF as bytes not equal.\nExpected\n%sbut got\n%s",
			expBytes,
			tcnfBytes,
		)
	}
}
