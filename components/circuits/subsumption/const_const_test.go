package subsumption

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"stratifoiled/sfdtest"
	"testing"
)

const constConstSUFIX = "subsumtpion.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 instances.Const,
	simplify bool,
) {
	context := components.NewContext(DIM, nil)
	formula := ConstConst(c1, c2)
	filePath := sfdtest.CNFName(constConstSUFIX, id, simplify)
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

func TestConstConst_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, constConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestConstConst_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, constConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}
