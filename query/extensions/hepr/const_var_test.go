package hepr_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/hepr"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runHEPRConstVar(t *testing.T, id int, tc HEPRRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	y := query.QVar("y")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = hepr.ConstVar{
		I1:                          c1,
		I2:                          y,
		Preference:                  tc.Pref,
		FeaturePreferenceRankVarGen: test.VarGenFeaturePreference,
	}
	if neg {
		f = logop.Not{Q: f}
	}
	f = logop.WithVar{
		I: y,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: y, I2: c2},
				Q2: subsumption.ConstVar{I1: c2, I2: y},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedHEPRConstVar(t *testing.T, id int, tc HEPRRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QVar("y")
	c2 := query.QConst{Val: tc.Val2}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val1)

	var f test.Encodable = hepr.ConstVar{
		I1:                          x,
		I2:                          y,
		Preference:                  tc.Pref,
		FeaturePreferenceRankVarGen: test.VarGenFeaturePreference,
	}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: y,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{I1: y, I2: c2},
				Q2: subsumption.ConstVar{I1: c2, I2: y},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestConstVar_Encoding(t *testing.T) {
	for i, tc := range HEPRPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runHEPRConstVar(t, i, tc, false)
		})
	}
}

func TestConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range HEPRPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedHEPRConstVar(t, i, tc, false)
		})
	}
}

func TestNotConstVar_Encoding(t *testing.T) {
	for i, tc := range HEPRNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runHEPRConstVar(t, i, tc, true)
		})
	}
}

func TestNotConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range HEPRNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedHEPRConstVar(t, i, tc, true)
		})
	}
}

func TestConstVar_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QVar("y")

	f := hepr.ConstVar{I1: x, I2: y, Preference: nil}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_Encoding_WrongPref(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QVar("y")

	f := hepr.ConstVar{I1: x, I2: y, Preference: []int{1, 2, 1}}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}
	y := query.QVar("y")

	f := hepr.ConstVar{I1: x, I2: y, Preference: nil}
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
