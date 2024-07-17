package leh_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/leh"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runLEHConstConstVar(t *testing.T, id int, tc test.TTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	z := query.QVar("z")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	var f test.Encodable = leh.ConstConstVar{
		I1:                    c1,
		I2:                    c2,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: z,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: z, I2: c3},
				Q2: subsumption.ConstVar{I1: c3, I2: z},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLEHConstConstVar(
	t *testing.T,
	id int,
	tc test.TTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QConst{ID: "y"}
	z := query.QVar("z")

	c3 := query.QConst{Val: tc.Val3}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val1)
	ctx.AddScope("y")
	_ = ctx.SetScope(2, tc.Val2)

	var f test.Encodable = leh.ConstConstVar{
		I1:                    x,
		I2:                    y,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: z,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: z, I2: c3},
				Q2: subsumption.ConstVar{I1: c3, I2: z},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestConstConstVar_Encoding(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHConstConstVar(t, i, tc, false)
		})
	}
}

func TestConstConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHConstConstVar(t, i, tc, false)
		})
	}
}

func TestNotConstConstVar_Encoding(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHConstConstVar(t, i, tc, true)
		})
	}
}

func TestNotConstConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHConstConstVar(t, i, tc, true)
		})
	}
}

func TestConstConstVar_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	z := query.QVar("z")

	f := leh.ConstConstVar{I1: x, I2: y, I3: z}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConstVar_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	z := query.QVar("z")

	f := leh.ConstConstVar{I1: x, I2: y, I3: z}
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
