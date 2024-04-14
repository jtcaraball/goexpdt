package cons

import (
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/operators"
	"testing"
)

const (
	constConstSUFIX        = "cons.constconst"
	guardedConstConstSUFIX = "cons.Gconstconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runConsConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	ctx := base.NewContext(DIM, nil)
	var formula base.Component = ConstConst(c1, c2)
	if neg {
		formula = operators.Not(formula)
	}
	filePath := solver.CNFName(constConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
}

func runGuardedConsConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.GuardedConst("y")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
		base.Guard{Target: "y", Value: c2, Idx: 2},
	)
	var formula base.Component = ConstConst(x, y)
	if neg {
		formula = operators.Not(formula)
	}
	filePath := solver.CNFName(guardedConstConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, constConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstConst(t, i, tc.expCode, tc.val1, tc.val2, false, false)
		})
	}
}

func TestNotConstConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, constConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstConst(t, i, tc.expCode, tc.val1, tc.val2, true, false)
		})
	}
}

func TestConstConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				false,
				false,
			)
		})
	}
}

func TestNotConstConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				true,
				false,
			)
		})
	}
}

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, constConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstConst(t, i, tc.expCode, tc.val1, tc.val2, false, true)
		})
	}
}

func TestNotConstConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, constConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstConst(t, i, tc.expCode, tc.val1, tc.val2, true, true)
		})
	}
}

func TestConstConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				false,
				true,
			)
		})
	}
}

func TestNotConstConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstConst(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				true,
				true,
			)
		})
	}
}

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
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
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
