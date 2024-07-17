package leh_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/leh"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runLEHVarVarConst(t *testing.T, id int, tc test.TTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	var f test.Encodable = leh.VarVarConst{
		I1:                    x,
		I2:                    y,
		I3:                    c3,
		HammingDistanceVarGen: test.VarGenHammingDistance,
		EqualFeatureVarGen:    test.VarGenEqualFeat,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.WithVar{
			I: y,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{I1: x, I2: c1},
					Q2: subsumption.ConstVar{I1: c1, I2: x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: subsumption.VarConst{I1: y, I2: c2},
						Q2: subsumption.ConstVar{I1: c2, I2: y},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedLEHVarVarConst(
	t *testing.T,
	id int,
	tc test.TTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	z := query.QConst{ID: "z"}

	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	ctx.AddScope("z")
	_ = ctx.SetScope(1, tc.Val3)

	var f test.Encodable = leh.VarVarConst{
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
			I: y,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{I1: x, I2: c1},
					Q2: subsumption.ConstVar{I1: c1, I2: x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: subsumption.VarConst{I1: y, I2: c2},
						Q2: subsumption.ConstVar{I1: c2, I2: y},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestVarVarConst_Encoding(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarVarConst(t, i, tc, false)
		})
	}
}

func TestNotVarVarConst_Encoding(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarVarConst(t, i, tc, true)
		})
	}
}

func TestVarVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHVarVarConst(t, i, tc, false)
		})
	}
}

func TestNotVarVarConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHVarVarConst(t, i, tc, true)
		})
	}
}

func TestVarVarConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	z := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := leh.VarVarConst{I1: x, I2: y, I3: z}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarVarConst_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QVar("y")
	z := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := leh.VarVarConst{I1: x, I2: y, I3: z}
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
