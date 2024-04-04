package lel

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const (
	constVarSUFIX        = "lel.constvar"
	guardedConstVarSUFIX = "lel.Gconstvar"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runLELConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	x := base.NewVar("x")
	context := base.NewContext(DIM, nil)
	// Define circuit
	var circuit base.Component = ConstVar(c1, x)
	if neg {
		circuit = operators.Not(circuit)
	}
	// Define formula
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
	// Run it
	filePath := test.CNFName(constVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x")
}

func runGuardedLELConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	x := base.NewVar("x")
	y := base.GuardedConst("y")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "y", Value: c1, Idx: 1},
	)
	// Define circuit
	var circuit base.Component = ConstVar(y, x)
	if neg {
		circuit = operators.Not(circuit)
	}
	// Define formula
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
	// Run it
	filePath := test.CNFName(guardedConstVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x#y#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstVar_Encoding(t *testing.T) {
	test.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstVar(t, i, tc.expCode, tc.val1, tc.val2, false, false)
		})
	}
}

func TestConstVar_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELConstVar(
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

func TestNotConstVar_Encoding(t *testing.T) {
	test.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstVar(t, i, tc.expCode, tc.val1, tc.val2, true, false)
		})
	}
}

func TestNotConstVar_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELConstVar(
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
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVar_Simplified(t *testing.T) {
	test.AddCleanup(t, constVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstVar(t, i, tc.expCode, tc.val1, tc.val2, false, true)
		})
	}
}

func TestConstVar_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELConstVar(
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

func TestNotConstVar_Simplified(t *testing.T) {
	test.AddCleanup(t, constVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runLELConstVar(t, i, tc.expCode, tc.val1, tc.val2, true, true)
		})
	}
}

func TestNotConstVar_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELConstVar(
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
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
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
