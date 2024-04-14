package full

import (
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/operators"
	"testing"
)

const (
	constSUFIX        = "full.const"
	guardedConstSUFIX = "full.Gconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runFullConst(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	ctx := base.NewContext(DIM, nil)
	var formula base.Component = Const(c)
	if neg {
		formula = operators.Not(formula)
	}
	filePath := solver.CNFName(constSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
}

func runGuardedFullConst(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	x := base.GuardedConst("x")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "x", Value: c, Idx: 1},
	)
	var formula base.Component = Const(x)
	if neg {
		formula = operators.Not(formula)
	}
	filePath := solver.CNFName(guardedConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, constSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestNotConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, constSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedFullConst(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestNotConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedFullConst(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	formula := Const(x)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, constSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestNotConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, constSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runFullConst(t, i, tc.expCode, tc.val, true, true)
		})
	}
}

func TestConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedFullConst(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestNotConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedFullConst(t, i, tc.expCode, tc.val, true, true)
		})
	}
}

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	formula := Const(x)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
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
