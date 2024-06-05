package pred_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/pred"
	"github.com/jtcaraball/goexpdt/query/pred/internal/testtable"
)

func runLELVarConst(t *testing.T, id int, tc testtable.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = pred.LELVarConst{x, c2, test.VarGenBotCount}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: pred.SubsumptionVarConst{I1: x, I2: c1},
				Q2: pred.SubsumptionConstVar{I1: c1, I2: x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLELVarConst(
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

	var f test.Encodable = pred.LELVarConst{x, y, test.VarGenBotCount}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: pred.SubsumptionVarConst{I1: x, I2: c1},
				Q2: pred.SubsumptionConstVar{I1: c1, I2: x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestLELVarConst_Encoding(t *testing.T) {
	for i, tc := range testtable.LELPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLELVarConst(t, i, tc, false)
		})
	}
}

func TestLELVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range testtable.LELPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLELVarConst(t, i, tc, false)
		})
	}
}

func TestNotLELVarConst_Encoding(t *testing.T) {
	for i, tc := range testtable.LELNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLELVarConst(t, i, tc, true)
		})
	}
}

func TestNotLELVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range testtable.LELNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLELVarConst(t, i, tc, true)
		})
	}
}

func TestLELVarConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := pred.LELVarConst{x, y, test.VarGenBotCount}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestLELVarConst_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT}}

	f := pred.LELVarConst{x, y, test.VarGenBotCount}
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
