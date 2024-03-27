package cons

import (
	"goexpdt/base"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"goexpdt/circuits/internal/test"
	"testing"
)

const varConstSUFIX = "cons.varconst"
const guardedVarConstSUFIX = "cons.Gvarconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runConsVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
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
			VarConst(c1, x),
		),
	)
	filePath := test.CNFName(varConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x")
}

func runGuardedConsVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	simplify bool,
) {
	x := base.NewVar("x")
	y := base.GuardedConst("y")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "y", Value: c1, Idx: 1},
	)
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c2),
				subsumption.ConstVar(c2, x),
			),
			VarConst(y, x),
		),
	)
	filePath := test.CNFName(guardedVarConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x#y#1", "y")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConst_Encoding(t *testing.T) {
	test.AddCleanup(t, varConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsVarConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsVarConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := VarConst(x, y)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_Simplified(t *testing.T) {
	test.AddCleanup(t, varConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsVarConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedConsVarConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := VarConst(x, y)
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := VarConst(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarConst_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.NewVar("y")
	formula := VarConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
