package leh

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const constVarVarSUFIX = "leh.constvarvar"
const guardedConstVarVarSUFIX = "leh.Gconstvarvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHConstVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	simplify bool,
) {
	y := base.NewVar("y")
	z := base.NewVar("z")
	context := base.NewContext(DIM, nil)
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
					ConstVarVar(c1, y, z),
				),
			),
		),
	)
	filePath := test.CNFName(constVarVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "y", "z")
}

func runGuardedLEHConstVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.Var("y")
	z := base.Var("z")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
	)
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
					ConstVarVar(x, y, z),
				),
			),
		),
	)
	filePath := test.CNFName(guardedConstVarVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "y#x#1", "z#x#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstVarVar_Encoding(t *testing.T) {
	test.AddCleanup(t, constVarVarSUFIX, false)
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
			)
		})
	}
}

func TestConstVarVar_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarVarSUFIX, false)
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
			)
		})
	}
}

func TestConstVarVar_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Var("z")
	formula := ConstVarVar(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVarVar_Simplified(t *testing.T) {
	test.AddCleanup(t, constVarVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarVar(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
				true,
			)
		})
	}
}

func TestConstVarVar_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarVar(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.val3,
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
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
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


