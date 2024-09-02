package test

import (
	"strconv"
	"strings"
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
)

func equalClauseSlice(cs1, cs2 []cnf.Clause) bool {
	if len(cs1) != len(cs2) {
		return false
	}

	var builder strings.Builder
	ct1 := make(map[string]struct{}, len(cs1))
	ct2 := make(map[string]struct{}, len(cs2))

	for _, c := range cs1 {
		ct1[clauseId(c, builder)] = struct{}{}
	}
	for _, c := range cs2 {
		ct2[clauseId(c, builder)] = struct{}{}
	}

	for k := range ct1 {
		if _, ok := ct2[k]; !ok {
			return false
		}
	}

	return true
}

func clauseId(c cnf.Clause, b strings.Builder) string {
	for _, v := range c {
		b.WriteRune(rune(v))
	}

	id := strconv.Itoa(len(c)) + b.String()
	b.Reset()

	return id
}

func ValidClauses(t *testing.T, sc, cc, esc, ecc []cnf.Clause) {
	t.Helper()
	if !equalClauseSlice(sc, esc) {
		t.Errorf("Semantic clauses not equal. Expected %d but got %d", esc, sc)
	}
	if !equalClauseSlice(cc, ecc) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			ecc,
			cc,
		)
	}
}
