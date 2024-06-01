package test

import (
	"slices"
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
)

func ValidClauses(t *testing.T, sc, cc, esc, ecc []cnf.Clause) {
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
