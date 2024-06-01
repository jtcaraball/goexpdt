package logop_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func buildFAGTree() query.Model {
	nodes := []query.Node{
		{Feat: 0, ZQ: 1, OQ: 2},
		{Value: false, ZQ: query.NoQ, OQ: query.NoQ},
		{Value: false, ZQ: query.NoQ, OQ: query.NoQ},
	}

	// If this returns an error I kill myself.
	t, err := test.NewMockTree(1, nodes)
	if err != nil {
		panic(err)
	}

	return t
}

func TestForAllGuarded_Encoding(t *testing.T) {
	ctx := query.BasicQContext(buildFAGTree())

	x := query.Const{ID: "x"}
	y := query.Var("y")
	cmp := logop.ForAllGuarded{x, logop.WithVar{y, test.Trivial(true)}}

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
		{4, 5, 6},
		{-4, -5},
		{-4, -6},
		{-5, -6},
		{7, 8, 9},
		{-7, -8},
		{-7, -9},
		{-8, -9},
	}

	test.ValidClauses(t, sc, cc, esc, ecc)
}

func TestForAllGuarded_Encoding_Nil(t *testing.T) {
	ctx := query.BasicQContext(buildFAGTree())

	vcmp := logop.ForAllGuarded{query.Const{ID: "x"}, test.Trivial(true)}
	icmp := logop.ForAllGuarded{query.Const{ID: "x"}, nil}

	ce := "ForAllGuarded: Invalid encoding with nil ctx"
	che := "ForAllGuarded: Invalid encoding of nil child"

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
