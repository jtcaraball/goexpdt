package subsumption_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/circuit/predicate/subsumption"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func runSubsumptionConstConst(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	var f test.Encodable
	f = subsumption.ConstConst{
		query.QConst{Val: tc.val1},
		query.QConst{Val: tc.val2},
	}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func runGuardedSubsumptionConstConst(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QConst{ID: "y"}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.val1)
	ctx.AddScope("y")
	_ = ctx.SetScope(2, tc.val2)

	var f test.Encodable
	f = subsumption.ConstConst{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func TestConstConst_Encoding(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(t, i, tc, false)
		})
	}
}

func TestConstConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstConst(t, i, tc, false)
		})
	}
}

func TestNotConstConst_Encoding(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(t, i, tc, true)
		})
	}
}

func TestNotConstConst_Encoding_Guarded(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstConst(t, i, tc, true)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}

	f := subsumption.ConstConst{x, y}
	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}
	y := query.QConst{Val: []query.FeatV{query.BOT}}

	f := subsumption.ConstConst{x, y}
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
