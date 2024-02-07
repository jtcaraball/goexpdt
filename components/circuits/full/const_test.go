package full

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"stratifoiled/sfdtest"
	"testing"
)

const constSUFIX = "full.const"

// =========================== //
//           HELPERS           //
// =========================== //

func runFullConst(
	t *testing.T,
	id, expCode int,
	c instances.Const,
	simplify bool,
) {
	var err error
	var formula components.Component
	context := components.NewContext(DIM, nil)
	formula = Const(c)
	filePath := sfdtest.CNFName(varSUFIX, id, simplify)
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

func TestConst_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, false)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := Const(x)
	context := components.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, true)
		})
	}
}

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := Const(x)
	context := components.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_GetChildren(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := Const(x)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConst_IsTrivial(t *testing.T) {
	x := instances.Const{instances.BOT, instances.BOT, instances.BOT}
	formula := Const(x)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
