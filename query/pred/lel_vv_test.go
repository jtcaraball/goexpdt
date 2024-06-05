package pred_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/pred"
	"github.com/jtcaraball/goexpdt/query/pred/internal/testtable"
)

func runLELVarVar(t *testing.T, id int, tc testtable.BTRecord, neg bool) {
	tree, _ := test.NewMockTree(tc.Dim, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	c1 := query.QConst{Val: tc.Val1}
	c2 := query.QConst{Val: tc.Val2}

	var f test.Encodable = pred.LELVarVar{x, y, test.VarGenBotCount}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.WithVar{
			I: y,
			Q: logop.And{
				Q1: logop.And{
					Q1: pred.SubsumptionVarConst{I1: x, I2: c1},
					Q2: pred.SubsumptionConstVar{I1: c1, I2: x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: pred.SubsumptionVarConst{I1: y, I2: c2},
						Q2: pred.SubsumptionConstVar{I1: c2, I2: y},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestVarVar_Encoding(t *testing.T) {
	for i, tc := range testtable.LELPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLELVarVar(t, i, tc, false)
		})
	}
}

func TestNotVarVar_Encoding(t *testing.T) {
	for i, tc := range testtable.LELNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runLELVarVar(t, i, tc, true)
		})
	}
}

func TestVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")
	y := query.QVar("y")

	f := pred.LELVarVar{x, y, test.VarGenBotCount}
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
