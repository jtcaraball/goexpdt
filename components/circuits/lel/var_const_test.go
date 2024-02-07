package lel

import (
	"stratifoiled/components"
	"stratifoiled/components/circuits/subsumption"
	"stratifoiled/components/instances"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varConstSUFIX = "lel.varconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2 instances.Const,
	simplify bool,
) {
	var err error
	var formula components.Component
	x := instances.NewVar("x")
	context := components.NewContext(DIM, nil)
	formula = operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			VarConst(x, c2),
		),
	)
	filePath := sfdtest.CNFName(varConstSUFIX, id, simplify)
	if simplify {
		formula, err = formula.Simplified(context)
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

func TestVarConst_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarConst_Encoding_WrongDim(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := VarConst(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarConst_Simplified_WrongDim(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := VarConst(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_GetChildren(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := VarConst(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarConst_IsTrivial(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := VarConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}