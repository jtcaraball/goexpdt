package lel

import (
	"stratifoiled/components"
	"stratifoiled/components/circuits/subsumption"
	"stratifoiled/components/instances"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const constVarSUFIX = "lel.constvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELConstVar(
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
				subsumption.VarConst(x, c2),
				subsumption.ConstVar(c2, x),
			),
			ConstVar(c1, x),
		),
	)
	filePath := sfdtest.CNFName(constVarSUFIX, id, simplify)
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

func TestConstVar_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstVar(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestConstVar_Encoding_WrongDim(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.NewVar("y")
	formula := ConstVar(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, constVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstVar(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestConstVar_Simplified_WrongDim(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.NewVar("y")
	formula := ConstVar(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_GetChildren(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.NewVar("y")
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
	y := instances.NewVar("y")
	formula := ConstVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
