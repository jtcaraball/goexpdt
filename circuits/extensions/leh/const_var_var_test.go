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
	constVarVarSUFIX        = "leh.constvarvar"
	guardedConstVarVarSUFIX = "leh.Gconstvarvar"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHConstVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	y := base.NewVar("y")
	z := base.NewVar("z")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = ConstVarVar(c1, y, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		y,
		operators.WithVar(
			z,
			operators.And(
				operators.And(
					subsumption.VarConst(y, c2),
					subsumption.ConstVar(c2, y),
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
	filePath := solver.CNFName(constVarVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "y", "z")
}

func runGuardedLEHConstVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.Var("y")
	z := base.Var("z")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
	)
	var circuit base.Component = ConstVarVar(x, y, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		y,
		operators.WithVar(
			z,
			operators.And(
				operators.And(
					subsumption.VarConst(y, c2),
					subsumption.ConstVar(c2, y),
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
	filePath := solver.CNFName(guardedConstVarVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "y#x#1", "z#x#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstVarVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, constVarVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarVar(
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

func TestNotConstVarVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, constVarVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarVar(
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

func TestConstVarVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarVar(
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

func TestNotConstVarVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarVar(
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

func TestConstVarVar_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Var("z")
	formula := ConstVarVar(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVarVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, constVarVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarVar(
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

func TestNotConstVarVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, constVarVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarVar(
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

func TestConstVarVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarVar(
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

func TestNotConstVarVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarVar(
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

func TestConstVarVar_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Var("z")
	formula := ConstVarVar(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVarVar_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Var("z")
	formula := ConstVarVar(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstVarVar_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Var("z")
	formula := ConstVarVar(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
