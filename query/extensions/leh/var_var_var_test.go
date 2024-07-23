package leh_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/leh"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runLEHVarVarVar(t *testing.T, id int, tc test.TTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	z := query.QVar("z")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}
	c3 := query.QConst{Val: tc.Val3}

	var f test.Encodable = leh.VarVarVar{
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
			Q: logop.WithVar{
				I: z,
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
						Q2: logop.And{
							Q1: logop.And{
								Q1: subsumption.VarConst{I1: z, I2: c3},
								Q2: subsumption.ConstVar{I1: c3, I2: z},
							},
							Q2: f,
						},
					},
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestVarVarVar_Encoding(t *testing.T) {
	for i, tc := range LEHPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarVarVar(t, i, tc, false)
		})
	}
}

func TestNotVarVarVar_Encoding(t *testing.T) {
	for i, tc := range LEHNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLEHVarVarVar(t, i, tc, true)
		})
	}
}

func TestVarVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QVar("y")
	z := query.QVar("z")

	f := leh.VarVarVar{I1: x, I2: y, I3: z}
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
