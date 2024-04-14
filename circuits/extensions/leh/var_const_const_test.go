package leh

import (
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/internal/test/context"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const (
	varConstConstSUFIX        = "leh.varconstconst"
	guardedVarConstConstSUFIX = "leh.Gvarconstconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHVarConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = VarConstConst(x, c2, c3)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			circuit,
		),
	)
	filePath := solver.CNFName(varConstConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x")
}

func runGuardedLEHVarConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.Var("x")
	y := base.GuardedConst("y")
	z := base.GuardedConst("z")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "y", Value: c2, Idx: 1},
		base.Guard{Target: "z", Value: c3, Idx: 1},
	)
	var circuit base.Component = VarConstConst(x, y, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			circuit,
		),
	)
	filePath := solver.CNFName(guardedVarConstConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x#y#1#z#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConstConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, varConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstConst(
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

func TestNotVarConstConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, varConstConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstConst(
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

func TestVarConstConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstConst(
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

func TestNotVarConstConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstConst(
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

func TestVarConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConstConst(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, varConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstConst(
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

func TestNotVarConstConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, varConstConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstConst(
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

func TestVarConstConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstConst(
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

func TestNotVarConstConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstConst(
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

func TestVarConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConstConst(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstConst_GetChildren(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConstConst(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarConstConst_IsTrivial(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConstConst(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
