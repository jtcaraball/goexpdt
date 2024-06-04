package subsumption_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/circuit/predicates/subsumption"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func runSubsumptionConstVar(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	y := query.QVar("y")
	c1 := query.QConst{Val: tc.val1}
	c2 := query.QConst{Val: tc.val2}

	var f test.Encodable = subsumption.ConstVar{c1, y}
	if neg {
		f = logop.Not{Q: f}
	}
	// Define formula
	f = logop.WithVar{
		I: y,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{y, c2},
				Q2: subsumption.ConstVar{c2, y},
			},
			Q2: f,
		},
	}
	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func runGuardedSubsumptionConstVar(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{ID: "x"}
	y := query.QVar("y")
	c2 := query.QConst{Val: tc.val2}

	ctx.AddScope("x")
	_ = ctx.SetScope(1, tc.val1)

	var f test.Encodable = subsumption.ConstVar{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: y,
		Q: logop.And{
			Q1: logop.And{
				Q1: subsumption.VarConst{y, c2},
				Q2: subsumption.ConstVar{c2, y},
			},
			Q2: f,
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func TestConstVar_Encoding(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstVar(t, i, tc, false)
		})
	}
}

func TestConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstVar(t, i, tc, false)
		})
	}
}

func TestNotConstVar_Encoding(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstVar(t, i, tc, true)
		})
	}
}

func TestNotConstVar_Encoding_Guarded(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstVar(t, i, tc, true)
		})
	}
}

func TestConstVar_Encoding_WrongDim(t *testing.T) {
	tree, _ := test.NewMockTree(4, nil)
	ctx := query.BasicQContext(tree)

	x := query.QConst{Val: []query.FeatV{query.BOT, query.BOT, query.BOT}}
	y := query.QVar("y")

	f := subsumption.ConstVar{x, y}

	_, err := f.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_Encoding_NilCtx(t *testing.T) {
	x := query.QConst{Val: []query.FeatV{query.BOT}}
	y := query.QVar("y")

	f := subsumption.ConstVar{x, y}
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
