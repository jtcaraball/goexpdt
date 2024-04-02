package subsumption

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/operators"
	"testing"
)

const constConstSUFIX = "subsumtpion.constconst"
const guardedConstConstSUFIX = "subsumtpion.Gconstconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	context := base.NewContext(DIM, nil)
	// Define formula
	var formula base.Component = ConstConst(c1, c2)
	if neg {
		formula = operators.Not(formula)
	}
	// Run it
	filePath := test.CNFName(constConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedSubsumptionConstConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	x := base.GuardedConst("x")
	y := base.GuardedConst("y")
	context := base.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{Target: "x", Value: c1, Idx: 1},
		base.Guard{Target: "y", Value: c2, Idx: 2},
	)
	// Define formula
	var formula base.Component = ConstConst(x, y)
	if neg {
		formula = operators.Not(formula)
	}
	// Run it
	filePath := test.CNFName(guardedConstConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConstConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(
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

func TestConstConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstConst(
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

func TestNotConstConst_Encoding(t *testing.T) {
	test.AddCleanup(t, constConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(
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

func TestNotConstConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstConst(
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

func TestConstConst_Encoding_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	context := base.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_Simplified(t *testing.T) {
	test.AddCleanup(t, constConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionConstConst(
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

func TestConstConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedConstConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionConstConst(
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

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	context := base.NewContext(4, nil)
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConstConst_GetChildren(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestConstConst_IsTrivial(t *testing.T) {
	x := base.Const{base.BOT, base.BOT, base.BOT}
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := ConstConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
