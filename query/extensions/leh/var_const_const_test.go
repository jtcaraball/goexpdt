package leh_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/leh"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runLEHVarConstConst(t *testing.T, id int, tc test.TTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	var f test.Encodable = leh.VarConstConst{
		I1:                    x,
		I2:                    c2,
		I3:                    c3,
		HammingDistanceVarGen: test.VarGenHammingDistance,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: x, I2: c1},
				Q2: subsumption.ConstVar{I1: c1, I2: x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLEHVarConstConst(
	t *testing.T,
	id int,
	tc test.TTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{ID: "y"}
	z := query.QConst{ID: "z"}

	c1 := query.QConst{Val: tc.Val1}

	ctx.AddScope("y")
	_ = ctx.SetScope(1, tc.Val2)
	ctx.AddScope("z")
	_ = ctx.SetScope(2, tc.Val3)

	var f test.Encodable = leh.VarConstConst{
		I1:                    x,
		I2:                    y,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: x, I2: c1},
				Q2: subsumption.ConstVar{I1: c1, I2: x},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestVarConstConst_Encoding(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarConstConst(t, i, tc, false)
		})
	}
}

func TestNotVarConstConst_Encoding(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarConstConst(t, i, tc, true)
		})
	}
}

func TestVarConstConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHVarConstConst(t, i, tc, false)
		})
	}
}

func TestNotVarConstConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHVarConstConst(t, i, tc, true)
		})
	}
}

func TestVarConstConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	z := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := leh.VarConstConst{I1: x, I2: y, I3: z}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstConst_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	z := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := leh.VarConstConst{I1: x, I2: y, I3: z}
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
