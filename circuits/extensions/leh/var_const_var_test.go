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
	varConstVarSUFIX        = "leh.varconstvar"
	guardedVarConstVarSUFIX = "leh.Gvarconstvar"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHVarConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	z := base.NewVar("z")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = VarConstVar(x, c2, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.WithVar(
			z,
			operators.And(
				operators.And(
					subsumption.VarConst(x, c1),
					subsumption.ConstVar(c1, x),
				),
				operators.And(
					operators.And(
						subsumption.VarConst(z, c3),
						subsumption.ConstVar(c3, z),
					),
					circuit,
				),
			),
		),
	)
	filePath := solver.CNFName(varConstVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x", "z")
}

func runGuardedLEHVarConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.Var("x")
	y := base.GuardedConst("y")
	z := base.Var("z")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "y", Value: c2, Idx: 1},
	)
	var circuit base.Component = VarConstVar(x, y, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.WithVar(
			z,
			operators.And(
				operators.And(
					subsumption.VarConst(x, c1),
					subsumption.ConstVar(c1, x),
				),
				operators.And(
					operators.And(
						subsumption.VarConst(z, c3),
						subsumption.ConstVar(c3, z),
					),
					circuit,
				),
			),
		),
	)
	filePath := solver.CNFName(guardedVarConstVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x#y#1", "z#y#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConstVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstVar(
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

func TestNotVarConstVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varConstVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstVar(
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

func TestVarConstVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstVar(
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

func TestNotVarConstVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstVar(
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

func TestVarConstVar_Encoding_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := VarConstVar(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstVar(
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

func TestNotVarConstVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varConstVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstVar(
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

func TestVarConstVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstVar(
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

func TestNotVarConstVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarConstVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstVar(
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

func TestVarConstVar_Simplified_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := VarConstVar(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstVar_GetChildren(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := VarConstVar(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarConstVar_IsTrivial(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Var("z")
	formula := VarConstVar(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
