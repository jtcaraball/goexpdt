package pred_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/pred"
	"github.com/jtcaraball/goexpdt/query/pred/internal/testtable"
)

func runSubsumptionVarConst(
	t *testing.T,
	id int,
	tc testtable.BTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = pred.SubsumptionVarConst{x, c2}

	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: pred.SubsumptionVarConst{x, c1},
				Q2: pred.SubsumptionConstVar{c1, x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedSubsumptionVarConst(
	t *testing.T,
	id int,
	tc testtable.BTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{ID: "y"}
	c1 := query.QConst{Val: tc.Val1}

	ctx.AddScope("y")
	_ = ctx.SetScope(1, tc.Val2)

	var f test.Encodable = pred.SubsumptionVarConst{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: pred.SubsumptionVarConst{x, c1},
				Q2: pred.SubsumptionConstVar{c1, x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestSubsumptionVarConst_Encoding(t *testing.T) {
	for i, tc := range testtable.SubsumptionPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runSubsumptionVarConst(t, i, tc, false)
		})
	}
}

func TestSubsumptionVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range testtable.SubsumptionPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(t, i, tc, false)
		})
	}
}

func TestNotSubsumptionVarConst_Encoding(t *testing.T) {
	for i, tc := range testtable.SubsumptionNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runSubsumptionVarConst(t, i, tc, true)
		})
	}
}

func TestNotSubsumptionVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range testtable.SubsumptionNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(t, i, tc, true)
		})
	}
}

func TestSubsumptionVarConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := pred.SubsumptionVarConst{x, y}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestSubsumptionVarConst_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT}}

	f := pred.SubsumptionVarConst{x, y}
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
