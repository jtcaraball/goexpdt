package lel

import (
	"stratifoiled/components"
	"stratifoiled/components/circuits/subsumption"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varVarSUFIX = "lel.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2 components.Const,
	simplify bool,
) {
	var err error
	var formula components.Component
	x := components.NewVar("x")
	y := components.NewVar("y")
	context := components.NewContext(DIM, nil)
	formula = operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(
					subsumption.VarConst(x, c1),
					subsumption.ConstVar(c1, x),
				),
				operators.And(
					operators.And(
						subsumption.VarConst(y, c2),
						subsumption.ConstVar(c2, y),
					),
					VarVar(x, y),
				),
			),
		),
	)
	filePath := sfdtest.CNFName(varVarSUFIX, id, simplify)
	if simplify {
		formula, err = formula.Simplified(context)
		if err != nil {
			t.Errorf("Formula simplification error. %s", err.Error())
			return
		}
	}
	cnf, err := formula.Encoding(context)
	if err != nil {
		t.Errorf("Formula encoding error. %s", err.Error())
		return
	}
	if err = cnf.ToFile(filePath); err != nil {
		t.Errorf("CNF writing error. %s", err.Error())
		return
	}
	sfdtest.RunFormulaTest(t, id, expCode, filePath)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVar_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarVar(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarVar_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarVar(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarVar_GetChildren(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	formula := VarVar(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarVar_IsTrivial(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	formula := VarVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
