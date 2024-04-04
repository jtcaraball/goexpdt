package subsumption

import (
	"goexpdt/base"
	"goexpdt/operators"
	"goexpdt/circuits/internal/test"
	"testing"
)

const varConstSUFIX = "subsumtpion.varconst"
const guardedVarConstSUFIX = "subsumtpion.Gvarconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	x := base.NewVar("x")
	context := base.NewContext(DIM, nil)
	// Define circuit
	var circuit base.Component = VarConst(x, c2)
	if neg {
		circuit = operators.Not(circuit)
	}
	// Define formula
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(VarConst(x, c1), ConstVar(c1, x)),
			circuit,
		),
	)
	// Run it
	filePath := test.CNFName(varConstSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x")
}

func runGuardedSubsumptionVarConst(
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
		base.Guard{Target: "y", Value: c2, Idx: 1},
	)
	// Define circuit
	var circuit base.Component = VarConst(x, y)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(VarConst(x, c1), ConstVar(c1, x)),
			circuit,
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
			runSubsumptionVarConst(
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

func TestVarConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(
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

func TestNotVarConst_Encoding(t *testing.T) {
	test.AddCleanup(t, varConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarConst(
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

func TestNotVarConst_Encoding_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(
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

func TestVarConst_Encoding_WrongDim(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
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
			runSubsumptionVarConst(
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

func TestVarConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(
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

func TestNotVarConst_Simplified(t *testing.T) {
	test.AddCleanup(t, varConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarConst(
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

func TestNotVarConst_Simplified_Guarded(t *testing.T) {
	test.AddCleanup(t, guardedVarConstSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedSubsumptionVarConst(
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

func TestVarConst_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
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
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := VarConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
