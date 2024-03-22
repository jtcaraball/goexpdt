package subsumption

import (
	"stratifoiled/base"
	"stratifoiled/operators"
	"stratifoiled/internal/test"
	"testing"
)

const constVarSUFIX = "subsumtpion.constvar"
const guardedConstVarSUFIX = "subsumtpion.Gconstvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	simplify bool,
) {
	y := base.NewVar("y")
	context := base.NewContext(DIM, nil)
	formula := operators.WithVar(
		y,
		operators.And(
			operators.And(VarConst(y, c2), ConstVar(c2, y)),
			ConstVar(c1, y),
		),
	)
	filePath := test.CNFName(constVarSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedSubsumptionConstVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	simplify bool,
) {
	x := base.GuardedConst("x")
	y := base.Var("y")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c1, Rep: "1"},
	)
	formula := operators.WithVar(
		y,
		operators.And(
			operators.And(VarConst(y, c2), ConstVar(c2, y)),
			ConstVar(x, y),
		),
	)
	filePath := test.CNFName(guardedConstVarSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstVar_Encoding(t *testing.T) {
	test.AddCleanup(t, constVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstVar(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestConstVar_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstVar(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
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
			runSubsumptionConstVar(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestConstVar_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstVar(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
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
	y := base.NewVar("x")
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
	y := base.NewVar("x")
	formula := ConstVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
