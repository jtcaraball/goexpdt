package subsumption_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/circuit/predicates/subsumption"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func runSubsumptionVarConst(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.val1}
	c2 := query.QConst{Val: tc.val2}

	var f test.Encodable = subsumption.VarConst{x, c2}

	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{x, c1},
				Q2: subsumption.ConstVar{c1, x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func runGuardedSubsumptionVarConst(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{ID: "y"}
	c1 := query.QConst{Val: tc.val1}

	ctx.AddScope("y")
	_ = ctx.SetScope(1, tc.val2)

	var f test.Encodable = subsumption.VarConst{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{x, c1},
				Q2: subsumption.ConstVar{c1, x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func TestVarConst_Encoding(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarConst(t, i, tc, false)
		})
	}
}

func TestVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(t, i, tc, false)
		})
	}
}

func TestNotVarConst_Encoding(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarConst(t, i, tc, true)
		})
	}
}

func TestNotVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(t, i, tc, true)
		})
	}
}

func TestVarConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := subsumption.VarConst{x, y}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT}}

	f := subsumption.VarConst{x, y}
	e := "Invalid encoding with nil ctx"

	_, err := f.Encoding(nil)
	if err == nil {
		t.Error("Nil context encoding error not caught.")
	} else if err.Error() != e {
		t.Errorf(
			"Incorrect error for nil context encoding. Expected %s but got %s",
			e,
			err.Error(),
		)
	}
}
