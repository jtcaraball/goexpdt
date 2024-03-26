package isnode

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"testing"
)

const constSUFIX = "isnode.const"
const guardedConstSUFIX = "isnode.Gconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runIsNodeConst(
	t *testing.T,
	id, expCode int,
	c base.Const,
	simplify bool,
) {
	context := base.NewContext(DIM, genTree())
	formula := Const(c)
	filePath := test.CNFName(constSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedIsNodeConst(
	t *testing.T,
	id, expCode int,
	c base.Const,
	simplify bool,
) {
	x := base.GuardedConst("x")
	context := base.NewContext(DIM, genTree())
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c, Rep: "1"},
	)
	formula := Const(x)
	filePath := test.CNFName(guardedConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeConst(t, i, tc.expCode, tc.val, false)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc.expCode, tc.val, false)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	formula := Const(x)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Simplified(t *testing.T) {
	test.AddCleanup(t, varSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeConst(t, i, tc.expCode, tc.val, true)
		})
	}
}

func TestConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc.expCode, tc.val, true)
		})
	}
}

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	formula := Const(x)
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
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
	x := base.Const{base.BOT, base.BOT, base.BOT}
	formula := Const(x)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
