package isnode_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/internal/testtable"
	"github.com/jtcaraball/goexpdt/query/predicates/isnode"
)

func runIsNodeConst(t *testing.T, id int, tc testtable.OTRecord, neg bool) {
	tree := testtable.IsNodeTree()
	ctx := query.BasicQContext(tree)

	c := query.QConst{Val: tc.Val}

	var f test.Encodable = isnode.Const{I: c}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedIsNodeConst(
	t *testing.T,
	id int,
	tc testtable.OTRecord,
	neg bool,
) {
	tree := testtable.IsNodeTree()
	ctx := query.BasicQContext(tree)

	c := query.QConst{ID: "x"}
	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val)

	var f test.Encodable = isnode.Const{I: c}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestConst_Encoding(t *testing.T) {
	for i, tc := range testtable.IsNodePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runIsNodeConst(t, i, tc, false)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range testtable.IsNodePTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc, false)
		})
	}
}

func TestNotConst_Encoding(t *testing.T) {
	for i, tc := range testtable.IsNodeNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runIsNodeConst(t, i, tc, true)
		})
	}
}

func TestNotConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range testtable.IsNodeNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc, true)
		})
	}
}

func TestConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := isnode.Const{I: x}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}

	f := isnode.Const{I: x}
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
