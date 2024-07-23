package logop_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func TestForAllGuarded_Encoding(t *testing.T) {
	tree, _ := test.NewMockTree(
		1,
		[]query.Node{
			{Feat: 0, ZChild: 1, OChild: 2},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
		},
	)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QVar("y")
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

type fagVerifyConst struct {
	I query.QConst
}

func (f fagVerifyConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	sc, _ := ctx.ScopeConst(f.I)
	ncnf := cnf.CNF{}

	// Lets assign cnf variables to our witnesses. This will be used to check
	// if the iterator is correctly walking the model.
	ctx.CNFVar(query.QVar(""), 0, int(query.BOT))  // 1
	ctx.CNFVar(query.QVar(""), 0, int(query.ONE))  // 2
	ctx.CNFVar(query.QVar(""), 0, int(query.ZERO)) // 3

	clause := cnf.Clause{}
	for _, ft := range sc.Val {
		clause = append(clause, ctx.CNFVar(query.QVar(""), 0, int(ft)))
	}

	return ncnf.AppendConsistency(clause), nil
}

func TestForAllGuarded_Encoding_Iterator(t *testing.T) {
	// This is not ideal but the test assumes the order in which ForAllGuarded
	// traverses the model. In particular it will fail if it does not use DFS.
	tree, _ := test.NewMockTree(
		2,
		[]query.Node{
			// (_, _)
			{Feat: 0, ZChild: 2, OChild: 1},
			// (1, _)
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			// (0, _)
			{Feat: 1, ZChild: 3, OChild: 4},
			// (0, 0)
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			// (0, 1)
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
		},
	)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	cmp := logop.ForAllGuarded{x, fagVerifyConst{I: x}}

	ncnf, err := cmp.Encoding(ctx)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}

	sc, cc := ncnf.Clauses()
	esc := []cnf.Clause{}
	ecc := []cnf.Clause{
		{1, 1}, // (_, _)
		{3, 1}, // (1, _)
		{3, 3}, // (0, _)
		{3, 2}, // (0, 0)
		{2, 1}, // (0, 1)
	}

	test.ValidClauses(t, sc, cc, esc, ecc)
}

func TestForAllGuarded_Encoding_Nil(t *testing.T) {
	ctx := query.BasicQContext(nil)

	vcmp := logop.ForAllGuarded{query.QConst{ID: "x"}, test.Trivial(true)}
	icmp := logop.ForAllGuarded{query.QConst{ID: "x"}, nil}

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
