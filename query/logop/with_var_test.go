package logop_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func TestWithVar_Encoding(t *testing.T) {
	tree, err := test.NewMockTree(1, nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	cmp := logop.WithVar{x, logop.WithVar{y, test.Trivial(false)}}

	ncnf, err := cmp.Encoding(ctx)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
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

func TestWithVar_Encoding_Scoping(t *testing.T) {
	tree, err := test.NewMockTree(1, nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := query.BasicQContext(tree)
	ctx.AddScope("T")
	_ = ctx.SetScope(1, []query.FeatV{query.BOT})

	x := query.QVar("x")
	cmp := logop.WithVar{x, test.Trivial(true)}

	_, err = cmp.Encoding(ctx)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}

	scp := ctx.ScopeVar(x)
	if scp == x {
		t.Error(
			"QVar not included in guard scope. Guarded QVar is the same",
			" as original",
		)
	}
}

func TestWithVar_Encoding_Nil(t *testing.T) {
	tree, _ := test.NewMockTree(1, nil)
	ctx := query.BasicQContext(tree)

	vcmp := logop.WithVar{query.QVar("x"), test.Trivial(true)}
	icmp := logop.WithVar{query.QVar("x"), nil}

	ce := "WithVar: Invalid encoding with nil ctx"
	che := "WithVar: Invalid encoding of nil child"

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
