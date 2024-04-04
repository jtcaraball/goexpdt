package leh

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/operators"
	"testing"
)

const (
	constConstConstSUFIX        = "leh.constconstconst"
	guardedConstConstConstSUFIX = "leh.Gconstconstconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHConstConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	context := base.NewContext(DIM, nil)
	var formula base.Component = ConstConstConst(c1, c2, c3)
	if neg {
		formula = operators.Not(formula)
	}
	filePath := test.CNFName(constConstConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedLEHConstConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.GuardedConst("y")
	z := base.GuardedConst("z")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
		base.Guard{Target: "y", Value: c2, Idx: 2},
		base.Guard{Target: "z", Value: c3, Idx: 3},
	)
	var formula base.Component = ConstConstConst(x, y, z)
	if neg {
		formula = operators.Not(formula)
	}
	filePath := test.CNFName(guardedConstConstConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstConstConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				false,
				false,
			)
		})
	}
}

func TestNotConstConstConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constConstConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				true,
				false,
			)
		})
	}
}

func TestConstConstConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				false,
				false,
			)
		})
	}
}

func TestNotConstConstConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				true,
				false,
			)
		})
	}
}

func TestConstConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConstConst(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConstConst_Simplified(t *testing.T) {
	test.AddCleanup(t, constConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				false,
				true,
			)
		})
	}
}

func TestNotConstConstConst_Simplified(t *testing.T) {
	test.AddCleanup(t, constConstConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				true,
				true,
			)
		})
	}
}

func TestConstConstConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				false,
				true,
			)
		})
	}
}

func TestNotConstConstConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				true,
				true,
			)
		})
	}
}

func TestConstConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConstConst(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConstConst_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConstConst(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstConstConst_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConstConst(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
