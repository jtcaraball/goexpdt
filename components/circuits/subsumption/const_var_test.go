package subsumption

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const constVarSUFIX = "subsumtpion.constvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 instances.Const,
	simplify bool,
) {
	y := instances.NewVar("y")
	context := components.NewContext(DIM, nil)
	formula := operators.WithVar(
		y,
		operators.And(
			operators.And(VarConst(y, c2), ConstVar(c2, y)),
			ConstVar(c1, y),
		),
	)
	filePath := sfdtest.CNFName(constVarSUFIX, id, simplify)
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

func TestConstVar_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstVar(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestConstVar_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, constVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstVar(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}
