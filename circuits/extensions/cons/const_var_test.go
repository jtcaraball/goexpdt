package cons

import (
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/internal/test/context"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const (
	constVarSUFIX        = "cons.constvar"
	guardedConstVarSUFIX = "cons.Gconstvar"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runConsConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = ConstVar(c1, x)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c2),
				subsumption.ConstVar(c2, x),
			),
			circuit,
		),
	)
	filePath := solver.CNFName(constVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x")
}

func runGuardedConsConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.NewVar("y")
	ctx := base.NewContext(DIM, nil)
	ctx.Guards = append(
		ctx.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
	)
	var circuit base.Component = ConstVar(x, y)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		y,
		operators.And(
			operators.And(
				subsumption.VarConst(y, c2),
				subsumption.ConstVar(c2, y),
			),
			circuit,
		),
	)
	filePath := solver.CNFName(guardedConstVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "y#x#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstVar(t, i, tc.expCode, tc.val1, tc.val2, false, false)
		})
	}
}

func TestNotConstVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstVar(t, i, tc.expCode, tc.val1, tc.val2, true, false)
		})
	}
}

func TestConstVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstVar(
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

func TestNotConstVar_Encoding_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstVar(
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

func TestConstVar_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := ConstVar(x, y)
	ctx := base.NewContext(4, nil)
	_, err := formula.Encoding(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, constVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstVar(t, i, tc.expCode, tc.val1, tc.val2, false, true)
		})
	}
}

func TestNotConstVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, constVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runConsConstVar(t, i, tc.expCode, tc.val1, tc.val2, true, true)
		})
	}
}

func TestConstVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstVar(
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

func TestNotConstVar_Simplified_Guarded(t *testing.T) {
	solver.AddCleanup(t, guardedConstVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsConstVar(
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

func TestConstVar_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := ConstVar(x, y)
	ctx := base.NewContext(4, nil)
	_, err := formula.Simplified(ctx)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := ConstVar(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstVar_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := ConstVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
