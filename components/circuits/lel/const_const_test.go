package lel

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"stratifoiled/sfdtest"
	"testing"
)

const constConstSUFIX = "lel.constconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 instances.Const,
	simplify bool,
) {
	var err error
	var formula components.Component
	context := components.NewContext(DIM, nil)
	formula = ConstConst(c1, c2)
	filePath := sfdtest.CNFName(constConstSUFIX, id, simplify)
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

func TestConstConst_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, constConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := ConstConst(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, constConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := ConstConst(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_GetChildren(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := ConstConst(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstConst_IsTrivial(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	y := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := ConstConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
