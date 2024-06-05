package pred_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/pred"
	"github.com/jtcaraball/goexpdt/query/pred/internal/testtable"
)

func runSubsumptionVarVar(t *testing.T, id int, tc testtable.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = pred.SubsumptionVarVar{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.WithVar{
			I: y,
			Q: logop.And{
				Q1: logop.And{
					Q1: pred.SubsumptionVarConst{x, c1},
					Q2: pred.SubsumptionConstVar{c1, x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: pred.SubsumptionVarConst{y, c2},
						Q2: pred.SubsumptionConstVar{c2, y},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestSubsumptionVarVar_Encoding(t *testing.T) {
	for i, tc := range testtable.SubsumptionPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc, false)
		})
	}
}

func TestNotSubsumptionVarVar_Encoding(t *testing.T) {
	for i, tc := range testtable.SubsumptionNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc, true)
		})
	}
}

func TestSubsumptionVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QVar("y")

	f := pred.SubsumptionVarVar{x, y}
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
