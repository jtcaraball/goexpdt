package full

import (
	"stratifoiled/components"
	"stratifoiled/sfdtest"
	"testing"
)

const constSUFIX = "full.const"
const guardedConstSUFIX = "full.Gconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runFullConst(
	t *testing.T,
	id, expCode int,
	c components.Const,
	simplify bool,
) {
	context := components.NewContext(DIM, nil)
	formula := Const(c)
	filePath := sfdtest.CNFName(varSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedFullConst(
	t *testing.T,
	id, expCode int,
	c components.Const,
	simplify bool,
) {
	x := components.GuardedConst("x")
	context := components.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		components.Guard{Target: "x", Value: c, Rep: "1"},
	)
	formula := Const(x)
	filePath := sfdtest.CNFName(guardedConstSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConst_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, constSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, false)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	sfdtest.AddCleanup(t, guardedConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedFullConst(t, i, tc.expCode, tc.val, false)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := components.Const{components.BOT, components.BOT, components.BOT}
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
	x := components.Const{components.BOT, components.BOT, components.BOT}
	formula := Const(x)
	context := components.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_GetChildren(t *testing.T) {
	x := components.Const{components.BOT, components.BOT, components.BOT}
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
	x := components.Const{components.BOT, components.BOT, components.BOT}
	formula := Const(x)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
