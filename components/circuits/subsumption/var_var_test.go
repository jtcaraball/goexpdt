package subsumption

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varVarSUFIX = "subsumtpion.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2 instances.Const,
	simplify bool,
) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	context := components.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(VarConst(x, c1), ConstVar(c1, x)),
				operators.And(
					operators.And(VarConst(y, c2), ConstVar(c2, y)),
					VarVar(x, y),
				),
			),
		),
	)
	filePath := sfdtest.CNFName(varVarSUFIX, id, simplify)
	if simplify {
		simpleFormula, err := formula.Simplified()
		if err != nil {
			t.Errorf("Formula simplification error: %s", err.Error())
		}
		sfdtest.RunFormulaTest(t, id, expCode, simpleFormula, context, filePath)
		return
	}
	sfdtest.RunFormulaTest(t, id, expCode, formula, context, filePath)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVar_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarVar_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}
