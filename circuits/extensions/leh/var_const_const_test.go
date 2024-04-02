package leh

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const varConstConstSUFIX = "leh.varconstconst"
const guardedVarConstConstSUFIX = "leh.Gvarconstconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHVarConstConst(
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
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			VarConstConst(x, c2, c3),
		),
	)
	filePath := test.CNFName(varConstConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x")
}

func runGuardedLEHVarConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	simplify bool,
) {
	x := base.Var("x")
	y := base.GuardedConst("y")
	z := base.GuardedConst("z")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "y", Value: c2, Idx: 1},
		base.Guard{Target: "z", Value: c3, Idx: 1},
	)
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			VarConstConst(x, y, z),
		),
	)
	filePath := test.CNFName(guardedVarConstConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x#y#1#z#1")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConstConst_Encoding(t *testing.T) {
	test.AddCleanup(t, varConstConstSUFIX, false)
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
			)
		})
	}
}

func TestVarConstConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstConstSUFIX, false)
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
			)
		})
	}
}

func TestVarConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConstConst(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConstConst_Simplified(t *testing.T) {
	test.AddCleanup(t, varConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarConstConst(
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

func TestVarConstConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLEHVarConstConst(
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

func TestVarConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Var("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	z := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConstConst(x, y, z)
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
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
