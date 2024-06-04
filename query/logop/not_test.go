package logop_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func TestNot_Encoding(t *testing.T) {
	tree, err := test.NewMockTree(1, nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	cmp := logop.Not{logop.WithVar{x, test.Trivial(false)}}

	ncnf, err := cmp.Encoding(ctx)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sc, cc := ncnf.Clauses()
	esc := []cnf.Clause{}
	ecc := []cnf.Clause{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
	}

	test.ValidClauses(t, sc, cc, esc, ecc)
}

func TestNot_Encoding_Nil(t *testing.T) {
	tree, _ := test.NewMockTree(1, nil)
	ctx := query.BasicQContext(tree)

	vcmp := logop.Not{test.Trivial(true)}
	icmp := logop.Not{nil}

	ce := "Not: Invalid encoding with nil ctx"
	che := "Not: Invalid encoding of nil child"

	_, err := vcmp.Encoding(nil)
	if err == nil {
		t.Error("Nil context encoding error not caught.")
	} else if err.Error() != ce {
		t.Errorf(
			"Incorrect error for nil context encoding. Expected %s but got %s",
			ce,
			err.Error(),
		)
	}

	_, err = icmp.Encoding(ctx)
	if err == nil {
		t.Error("Nil child encoding error not caught.")
	} else if err.Error() != che {
		t.Errorf(
			"Incorrect error for nil context encoding. Expected %s but got %s",
			che,
			err.Error(),
		)
	}
}
