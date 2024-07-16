package allcomp_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/allcomp"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runAllCompVar(t *testing.T, id int, tc test.OTRecord, val, neg bool) {
	tree := AllCompTree()
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c := query.QConst{Val: tc.Val}

	var f test.Encodable = allcomp.Var{
		I:               x,
		LeafValue:       val,
		ReachNodeVarGen: test.VarGenNodeReach,
	}

	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.ConstVar{I1: c, I2: x},
				Q2: subsumption.VarConst{I1: x, I2: c},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestVar_Encoding_AllPos(t *testing.T) {
	for i, tc := range AllPosPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompVar(t, i, tc, true, false)
		})
	}
}

func TestNotVar_Encoding_AllPos(t *testing.T) {
	for i, tc := range AllPosNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompVar(t, i, tc, true, true)
		})
	}
}

func TestVar_Encoding_AllNeg(t *testing.T) {
	for i, tc := range AllNegPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompVar(t, i, tc, false, false)
		})
	}
}

func TestNotVar_Encoding_AllNeg(t *testing.T) {
	for i, tc := range AllNegNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompVar(t, i, tc, false, true)
		})
	}
}

func TestVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")

	f := allcomp.Var{I: x}
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
