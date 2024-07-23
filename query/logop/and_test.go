package logop_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func TestAnd_Encoding_DTrue(t *testing.T) {
	tree, _ := test.NewMockTree(1, nil)
	ctx := query.BasicQContext(tree)

	cx := logop.WithVar{query.QVar("x"), test.Trivial(true)}
	cy := logop.WithVar{query.QVar("y"), test.Trivial(true)}
	cmp := logop.And{cx, cy}

	ncnf, err := cmp.Encoding(ctx)
	if err != nil {
		t.Fatalf("CNF encoding error. %s", err.Error())
	}

	sc, cc := ncnf.Clauses()
	esc := []cnf.Clause{}
	ecc := []cnf.Clause{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
		{4, 5, 6},
		{-4, -5},
		{-4, -6},
		{-5, -6},
	}
	test.ValidClauses(t, sc, cc, esc, ecc)
}

func TestAnd_Encoding_DFalse(t *testing.T) {
	tree, _ := test.NewMockTree(1, nil)
	ctx := query.BasicQContext(tree)

	cx := logop.WithVar{query.QVar("x"), test.Trivial(false)}
	cy := logop.WithVar{query.QVar("y"), test.Trivial(false)}
	cmp := logop.And{cx, cy}

	ncnf, err := cmp.Encoding(ctx)
	if err != nil {
		t.Fatalf("CNF encoding error. %s", err.Error())
	}

	sc, cc := ncnf.Clauses()
	esc := []cnf.Clause{{}, {}}
	ecc := []cnf.Clause{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
		{4, 5, 6},
		{-4, -5},
		{-4, -6},
		{-5, -6},
	}
	test.ValidClauses(t, sc, cc, esc, ecc)
}

func TestAnd_Encoding_Mixed(t *testing.T) {
	tree, _ := test.NewMockTree(1, nil)
	ctx := query.BasicQContext(tree)

	cx := logop.WithVar{query.QVar("x"), test.Trivial(true)}
	cy := logop.WithVar{query.QVar("y"), test.Trivial(false)}
	cmp := logop.And{cx, cy}

	ncnf, err := cmp.Encoding(ctx)
	if err != nil {
		t.Fatalf("CNF encoding error. %s", err.Error())
	}

	sc, cc := ncnf.Clauses()
	esc := cnf.NegClauses
	ecc := []cnf.Clause{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
		{4, 5, 6},
		{-4, -5},
		{-4, -6},
		{-5, -6},
	}
	test.ValidClauses(t, sc, cc, esc, ecc)
}

func TestAnd_Encoding_Nil(t *testing.T) {
	tree, _ := test.NewMockTree(1, nil)
	ctx := query.BasicQContext(tree)

	vcmp := logop.And{test.Trivial(true), test.Trivial(true)}
	icmp1 := logop.And{nil, test.Trivial(true)}
	icmp2 := logop.And{test.Trivial(true), nil}

	ce := "And: Invalid encoding with nil ctx"
	che1 := "And: Invalid encoding of nil child (1)"
	che2 := "And: Invalid encoding of nil child (2)"

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

	_, err = icmp1.Encoding(ctx)
	if err == nil {
		t.Error("Nil child encoding error not caught.")
	} else if err.Error() != che1 {
		t.Errorf(
			"Incorrect error for nil context encoding. Expected %s but got %s",
			che1,
			err.Error(),
		)
	}

	_, err = icmp2.Encoding(ctx)
	if err == nil {
		t.Error("Nil child encoding error not caught.")
	} else if err.Error() != che2 {
		t.Errorf(
			"Incorrect error for nil context encoding. Expected %s but got %s",
			che2,
			err.Error(),
		)
	}
}
