package leh_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/leh"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runLEHConstVarVar(t *testing.T, id int, tc test.TTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	y := query.QVar("y")
	z := query.QVar("z")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	var f test.Encodable = leh.ConstVarVar{
		I1:                    c1,
		I2:                    y,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: y,
		Q: logop.WithVar{
			I: z,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{I1: y, I2: c2},
					Q2: subsumption.ConstVar{I1: c2, I2: y},
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

func runGuardedLEHConstVarVar(
	t *testing.T,
	id int,
	tc test.TTRecord,
	neg bool,
) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QVar("y")
	z := query.QVar("z")

	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val1)

	var f test.Encodable = leh.ConstVarVar{
		I1:                    x,
		I2:                    y,
		I3:                    z,
		HammingDistanceVarGen: test.VarGenHammingDistance,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: y,
		Q: logop.WithVar{
			I: z,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{I1: y, I2: c2},
					Q2: subsumption.ConstVar{I1: c2, I2: y},
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

func TestConstVarVar_Encoding(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHConstVarVar(t, i, tc, false)
		})
	}
}

func TestConstVarVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHConstVarVar(t, i, tc, false)
		})
	}
}

func TestNotConstVarVar_Encoding(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHConstVarVar(t, i, tc, true)
		})
	}
}

func TestNotConstVarVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedLEHConstVarVar(t, i, tc, true)
		})
	}
}

func TestConstVarVar_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QVar("y")
	z := query.QVar("z")

	f := leh.ConstVarVar{I1: x, I2: y, I3: z}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QVar("y")
	z := query.QVar("z")

	f := leh.ConstVarVar{I1: x, I2: y, I3: z}
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
