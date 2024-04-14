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
	varVarConstSUFIX        = "leh.varvarconst"
	guardedVarVarConstSUFIX = "leh.Gvarvarconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHVarVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = VarVarConst(x, y, c3)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(
					subsumption.VarConst(x, c1),
					subsumption.ConstVar(c1, x),
				),
				operators.And(
					operators.And(
						subsumption.VarConst(y, c2),
						subsumption.ConstVar(c2, y),
					),
					circuit,
				),
			),
		),
	)
	filePath := solver.CNFName(varVarConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x", "y")
}

func runGuardedLEHVarVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	neg, simplify bool,
) {
	x := base.Var("x")
	y := base.Var("y")
	z := base.GuardedConst("z")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "z", Value: c3, Idx: 1},
	)
	var circuit base.Component = VarVarConst(x, y, z)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(
					subsumption.VarConst(x, c1),
					subsumption.ConstVar(c1, x),
				),
				operators.And(
					operators.And(
						subsumption.VarConst(y, c2),
						subsumption.ConstVar(c2, y),
					),
					circuit,
				),
			),
		),
	)
	filePath := solver.CNFName(guardedVarVarConstSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x#z#1", "y#z#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVarConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, varVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarVarConst(
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

func TestNotVarVarConst_Encoding(t *testing.T) {
	solver.AddCleanup(t, varVarConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarVarConst(
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

func TestVarVarConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarVarConst(
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

func TestNotVarVarConst_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarVarConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarVarConst(
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

func TestVarVarConst_Encoding_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarVarConst(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarVarConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, varVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarVarConst(
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

func TestNotVarVarConst_Simplified(t *testing.T) {
	solver.AddCleanup(t, varVarConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarVarConst(
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

func TestVarVarConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarVarConst(
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

func TestNotVarVarConst_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedVarVarConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarVarConst(
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

func TestVarVarConst_Simplified_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarVarConst(x, y, z)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarVarConst_GetChildren(t *testing.T) {
	x := base.Var("x")
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarVarConst(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarVarConst_IsTrivial(t *testing.T) {
	x := base.Var("x")
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarVarConst(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
