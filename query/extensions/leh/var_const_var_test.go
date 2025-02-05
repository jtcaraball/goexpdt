package leh_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/leh"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runLEHVarConstVar(t *testing.T, id int, tc test.TTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	z := query.QVar("z")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	var f test.Encodable = leh.VarConstVar{
		I1:                    x,
		I2:                    c2,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
		EqualFeatureVarGen:    test.VarGenEqualFeat,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.WithVar{
			I: z,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{I1: x, I2: c1},
					Q2: subsumption.ConstVar{I1: c1, I2: x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: subsumption.VarConst{I1: z, I2: c3},
						Q2: subsumption.ConstVar{I1: c3, I2: z},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLEHVarConstVar(
	t *testing.T,
	id int,
	tc test.TTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{ID: "y"}
	z := query.QVar("z")

	c1 := query.QConst{Val: tc.Val1}
	c3 := query.QConst{Val: tc.Val3}

	ctx.AddScope("y")
	_ = ctx.SetScope(1, tc.Val2)

	var f test.Encodable = leh.VarConstVar{
		I1:                    x,
		I2:                    y,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
		EqualFeatureVarGen:    test.VarGenEqualFeat,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.WithVar{
			I: z,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{I1: x, I2: c1},
					Q2: subsumption.ConstVar{I1: c1, I2: x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: subsumption.VarConst{I1: z, I2: c3},
						Q2: subsumption.ConstVar{I1: c3, I2: z},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConstVar_Encoding(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarConstVar(t, i, tc, false)
		})
	}
}

func TestVarConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHVarConstVar(t, i, tc, false)
		})
	}
}

func TestNotVarConstVar_Encoding(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarConstVar(t, i, tc, true)
		})
	}
}

func TestNotVarConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHVarConstVar(t, i, tc, true)
		})
	}
}

func TestVarConstVar_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	z := query.QVar("z")

	f := leh.VarConstVar{I1: x, I2: y, I3: z}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	z := query.QVar("z")

	f := leh.VarConstVar{I1: x, I2: y, I3: z}
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
