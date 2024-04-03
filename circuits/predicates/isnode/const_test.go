package isnode

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/operators"
	"testing"
)

const (
	constSUFIX        = "isnode.const"
	guardedConstSUFIX = "isnode.Gconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runIsNodeConst(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	context := base.NewContext(DIM, genTree())
	// Define formula
	var formula base.Component = Const(c)
	if neg {
		formula = operators.Not(formula)
	}
	// Run it
	filePath := test.CNFName(constSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedIsNodeConst(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	x := base.GuardedConst("x")
	context := base.NewContext(DIM, genTree())
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c, Idx: 1},
	)
	// Define formula
	var formula base.Component = Const(x)
	if neg {
		formula = operators.Not(formula)
	}
	// Run it
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
			runIsNodeConst(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestNotConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeConst(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestNotConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestConst_Encoding_WrongDim(t *testing.T) {
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
			runIsNodeConst(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestNotConst_Simplified(t *testing.T) {
	test.AddCleanup(t, varSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeConst(t, i, tc.expCode, tc.val, true, true)
		})
	}
}

func TestNotConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedIsNodeConst(t, i, tc.expCode, tc.val, true, true)
		})
	}
}

func TestConst_Simplified_WrongDim(t *testing.T) {
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
