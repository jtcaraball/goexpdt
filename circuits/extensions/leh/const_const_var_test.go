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
	constConstVarSUFIX        = "leh.constconstvar"
	guardedConstConstVarSUFIX = "leh.Gconstconstvar"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHConstConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = ConstConstVar(c1, c2, x)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c3),
				subsumption.ConstVar(c3, x),
			),
			circuit,
		),
	)
	filePath := solver.CNFName(constConstVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x")
}

func runGuardedLEHConstConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.GuardedConst("y")
	z := base.Var("z")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
		base.Guard{Target: "y", Value: c2, Idx: 1},
	)
	var circuit base.Component = ConstConstVar(x, y, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		z,
		operators.And(
			operators.And(
				subsumption.VarConst(z, c3),
				subsumption.ConstVar(c3, z),
			),
			circuit,
		),
	)
	filePath := solver.CNFName(guardedConstConstVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "z#x#1#y#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstConstVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, constConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstVar(
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

func TestNotConstConstVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, constConstVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstVar(
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

func TestConstConstVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstVar(
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

func TestNotConstConstVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstVar(
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

func TestConstConstVar_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := ConstConstVar(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConstVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, constConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstVar(
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

func TestNotConstConstVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, constConstVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstConstVar(
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

func TestConstConstVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstVar(
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

func TestNotConstConstVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstConstVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstConstVar(
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

func TestConstConstVar_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := ConstConstVar(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConstVar_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := ConstConstVar(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstConstVar_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := ConstConstVar(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
