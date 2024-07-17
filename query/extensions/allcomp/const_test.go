package allcomp_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/allcomp"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func runAllCompConst(t *testing.T, id int, tc test.OTRecord, val, neg bool) {
	tree := AllCompTree()
	ctx := query.BasicQContext(tree)

	c := query.QConst{Val: tc.Val}

	var f test.Encodable = allcomp.Const{I: c, LeafValue: val}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func runGuardedAllCompConst(
	t *testing.T,
	id int,
	tc test.OTRecord,
	val, neg bool,
) {
	tree := AllCompTree()
	ctx := query.BasicQContext(tree)

	c := query.QConst{ID: "x"}
	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.Val)

	var f test.Encodable = allcomp.Const{I: c, LeafValue: val}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.ExpCode)
}

func TestConst_Encoding_AllPos(t *testing.T) {
	for i, tc := range AllPosPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompConst(t, i, tc, true, false)
		})
	}
}

func TestConst_Encoding_AllPos_Guarded(t *testing.T) {
	for i, tc := range AllPosPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc, true, false)
		})
	}
}

func TestNotConst_Encoding_AllPos(t *testing.T) {
	for i, tc := range AllPosNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompConst(t, i, tc, true, true)
		})
	}
}

func TestNotConst_Encoding_AllPos_Guarded(t *testing.T) {
	for i, tc := range AllPosNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc, true, true)
		})
	}
}

func TestConst_Encoding_AllNeg(t *testing.T) {
	for i, tc := range AllNegPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompConst(t, i, tc, false, false)
		})
	}
}

func TestConst_Encoding_AllNeg_Guraded(t *testing.T) {
	for i, tc := range AllNegPTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc, false, false)
		})
	}
}

func TestNotConst_Encoding_AllNeg(t *testing.T) {
	for i, tc := range AllNegNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runAllCompConst(
				t,
				i,
				tc,
				false,
				true,
			)
		})
	}
}

func TestNotConst_Encoding_AllNeg_Guraded(t *testing.T) {
	for i, tc := range AllNegNTT {
		t.Run(tc.Name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc, false, true)
		})
	}
}

func TestConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := allcomp.Const{I: x}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}

	f := allcomp.Const{I: x}
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
