package lel

import (
	"stratifoiled/components"
	"stratifoiled/sfdtest"
	"testing"
)

const constConstSUFIX = "lel.constconst"
const guardedConstConstSUFIX = "lel.Gconstconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 components.Const,
	simplify bool,
) {
	context := components.NewContext(DIM, nil)
	formula := ConstConst(c1, c2)
	filePath := sfdtest.CNFName(constConstSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedLELConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 components.Const,
	simplify bool,
) {
	x := components.GuardedConst("x")
	y := components.GuardedConst("y")
	context := components.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		components.Guard{Target: "x", Value: c1, Rep: "1"},
		components.Guard{Target: "y", Value: c2, Rep: "2"},
	)
	formula := ConstConst(x, y)
	filePath := sfdtest.CNFName(guardedConstConstSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
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

func TestConstConst_Encoding_Guarded(t *testing.T) {
	sfdtest.AddCleanup(t, guardedConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELConstConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := components.Const{components.BOT, components.BOT, components.BOT}
	y := components.Const{components.BOT, components.BOT, components.BOT}
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
	x := components.Const{components.BOT, components.BOT, components.BOT}
	y := components.Const{components.BOT, components.BOT, components.BOT}
	formula := ConstConst(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_GetChildren(t *testing.T) {
	x := components.Const{components.BOT, components.BOT, components.BOT}
	y := components.Const{components.BOT, components.BOT, components.BOT}
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
	x := components.Const{components.BOT, components.BOT, components.BOT}
	y := components.Const{components.BOT, components.BOT, components.BOT}
	formula := ConstConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
