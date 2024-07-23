package hepr_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/hepr"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runHEPRVarVar(t *testing.T, id int, tc HEPRRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = hepr.VarVar{
		I1:                          x,
		I2:                          y,
		Preference:                  tc.Pref,
		FeaturePreferenceRankVarGen: test.VarGenFeaturePreference,
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

func TestVarVar_Encoding(t *testing.T) {
	for i, tc := range HEPRPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runHEPRVarVar(t, i, tc, false)
		})
	}
}

func TestNotVarVar_Encoding(t *testing.T) {
	for i, tc := range HEPRNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runHEPRVarVar(t, i, tc, true)
		})
	}
}

func TestVarVar_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")

	f := hepr.VarVar{I1: x, I2: y, Preference: nil}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarVar_Encoding_WrongPref(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")

	f := hepr.VarVar{I1: x, I2: y, Preference: []int{1, 2, 1}}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QVar("y")

	f := hepr.VarVar{I1: x, I2: y, Preference: nil}
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
