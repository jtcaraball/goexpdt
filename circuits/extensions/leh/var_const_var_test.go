package leh

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
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
	context := base.NewContext(DIM, nil)
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
	filePath := test.CNFName(varConstVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x", "z")
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
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
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
	filePath := test.CNFName(guardedVarConstVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x#y#1", "z#y#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConstVar_Encoding(t *testing.T) {
	test.AddCleanup(t, varConstVarSUFIX, false)
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
	test.AddCleanup(t, varConstVarSUFIX, false)
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
	test.AddCleanup(t, guardedVarConstVarSUFIX, false)
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
	test.AddCleanup(t, guardedVarConstVarSUFIX, false)
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
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstVar_Simplified(t *testing.T) {
	test.AddCleanup(t, varConstVarSUFIX, true)
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
	test.AddCleanup(t, varConstVarSUFIX, true)
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
	test.AddCleanup(t, guardedVarConstVarSUFIX, true)
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
	test.AddCleanup(t, guardedVarConstVarSUFIX, true)
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
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
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
