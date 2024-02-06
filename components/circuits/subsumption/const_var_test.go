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
	var err error
	var formula components.Component
	y := instances.NewVar("y")
	context := components.NewContext(DIM, nil)
	formula = operators.WithVar(
		y,
		operators.And(
			operators.And(VarConst(y, c2), ConstVar(c2, y)),
			ConstVar(c1, y),
		),
	)
	filePath := sfdtest.CNFName(constVarSUFIX, id, simplify)
	if simplify {
		formula, err = formula.Simplified()
		if err != nil {
			t.Errorf("Formula simplification error. %s", err.Error())
			return
		}
	}
	cnf, err := formula.Encoding(context)
	if err = cnf.ToFile(filePath); err != nil {
		t.Errorf("CNF writing error. %s", err.Error())
		return
	}
	sfdtest.RunFormulaTest(t, id, expCode, filePath)
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

func TestConstVar_GetChildren(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.NewVar("x")
	formula := ConstVar(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstVar_IsTrivial(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.NewVar("x")
	formula := ConstVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
