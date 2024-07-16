package dfs_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/dfs"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func runDFSVar(t *testing.T, id int, tc test.OTRecord, neg bool) {
	tree := DFSTree()
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	c := query.QConst{Val: tc.Val}

	var f test.Encodable = dfs.Var{I: x}
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
	for i, tc := range DFSPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runDFSVar(t, i, tc, false)
		})
	}
}

func TestNotVar_Encoding(t *testing.T) {
	for i, tc := range DFSNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runDFSVar(t, i, tc, true)
		})
	}
}

func TestVarVar_Encoding_NilCtx(t *testing.T) {
	x := query.QVar("x")

	f := dfs.Var{I: x}
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
