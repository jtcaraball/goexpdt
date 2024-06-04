package subsumption_test

import (
	"testing"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/circuit/predicates/subsumption"
	"github.com/jtcaraball/goexpdt/query/internal/test"
	"github.com/jtcaraball/goexpdt/query/logop"
)

func runSubsumptionVarVar(t *testing.T, id int, tc testS, neg bool) {
	tree, _ := test.NewMockTree(DIM, nil)
	ctx := query.BasicQContext(tree)

	x := query.QVar("x")
	y := query.QVar("y")
	c1 := query.QConst{Val: tc.val1}
	c2 := query.QConst{Val: tc.val2}

	var f test.Encodable = subsumption.VarVar{x, y}
	if neg {
		f = logop.Not{Q: f}
	}

	f = logop.WithVar{
		I: x,
		Q: logop.WithVar{
			I: y,
			Q: logop.And{
				Q1: logop.And{
					Q1: subsumption.VarConst{x, c1},
					Q2: subsumption.ConstVar{c1, x},
				},
				Q2: logop.And{
					Q1: logop.And{
						Q1: subsumption.VarConst{y, c2},
						Q2: subsumption.ConstVar{c2, y},
					},
					Q2: f,
				},
			},
		},
	}

	test.EncodeAndRun(t, f, ctx, id, tc.expCode)
}

func TestVarVar_Encoding(t *testing.T) {
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc, false)
		})
	}
}

func TestNotVarVar_Encoding(t *testing.T) {
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc, true)
		})
	}
}
