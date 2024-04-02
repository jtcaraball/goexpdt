package leh

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const constVarConstSUFIX = "leh.constvarconst"
const guardedConstVarConstSUFIX = "leh.Gconstvarconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHConstVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	simplify bool,
) {
	x := base.NewVar("x")
	context := base.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c2),
				subsumption.ConstVar(c2, x),
			),
			ConstVarConst(c1, x, c3),
		),
	)
	filePath := test.CNFName(constVarConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x")
}

func runGuardedLEHConstVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.Var("y")
	z := base.GuardedConst("z")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
		base.Guard{Target: "z", Value: c3, Idx: 1},
	)
	formula := operators.WithVar(
		y,
		operators.And(
			operators.And(
				subsumption.VarConst(y, c2),
				subsumption.ConstVar(c2, y),
			),
			ConstVarConst(x, y, z),
		),
	)
	filePath := test.CNFName(guardedConstVarConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "y#x#1#z#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstVarConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarConst(
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

func TestConstVarConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarConst(
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

func TestConstVarConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstVarConst(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVarConst_Simplified(t *testing.T) {
	test.AddCleanup(t, constVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHConstVarConst(
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

func TestConstVarConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHConstVarConst(
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

func TestConstVarConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstVarConst(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstVarConst_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstVarConst(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstVarConst_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Var("y")
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstVarConst(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
