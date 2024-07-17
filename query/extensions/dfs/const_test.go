package dfs_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/dfs"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func runDFSConst(t *testing.T, id int, tc test.OTRecord, neg bool) {
	tree := DFSTree()
	ctx := query.BasicQContext(tree)

	c := query.QConst{Val: tc.Val}

	var f test.Encodable = dfs.Const{I: c}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedDFSConst(t *testing.T, id int, tc test.OTRecord, neg bool) {
	tree := DFSTree()
	ctx := query.BasicQContext(tree)

	c := query.QConst{ID: "x"}
	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val)

	var f test.Encodable = dfs.Const{I: c}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConst_Encoding(t *testing.T) {
	for i, tc := range DFSPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runDFSConst(t, i, tc, false)
		})
	}
}

func TestNotConst_Encoding(t *testing.T) {
	for i, tc := range DFSNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runDFSConst(t, i, tc, true)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range DFSPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedDFSConst(t, i, tc, false)
		})
	}
}

func TestNotConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range DFSNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedDFSConst(t, i, tc, true)
		})
	}
}

func TestConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := dfs.Const{I: x}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}

	f := dfs.Const{I: x}
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
