package isnode_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/pred/internal/testtable"
	"github.com/jtcaraball/goexpdt/query/pred/isnode"
	"github.com/jtcaraball/goexpdt/query/pred/subsumption"
)

func runIsNodeVar(t *testing.T, id int, tc testtable.OTRecord, neg bool) {
	tree := testtable.IsNodeTree()
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c := query.QConst{Val: tc.Val}

	// Define circuit
	var f test.Encodable = isnode.Var{
		I:               x,
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

func TestVar_Encoding(t *testing.T) {
	for i, tc := range testtable.IsNodePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runIsNodeVar(t, i, tc, false)
		})
	}
}

func TestNotVar_Encoding(t *testing.T) {
	for i, tc := range testtable.IsNodeNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runIsNodeVar(t, i, tc, true)
		})
	}
}

func TestVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")

	f := isnode.Var{I: x}
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
