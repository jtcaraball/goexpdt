package subsumption

import (
	"goexpdt/base"
	"goexpdt/circuits/internal/test"
	"goexpdt/operators"
	"testing"
)

const varVarSUFIX = "subsumtpion.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	// Define variable and context
	x := base.NewVar("x")
	y := base.NewVar("y")
	context := base.NewContext(DIM, nil)
	// Define circuit
	var circuit base.Component = VarVar(x, y)
	if neg {
		circuit = operators.Not(circuit)
	}
	// Define formula
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(VarConst(x, c1), ConstVar(c1, x)),
				operators.And(
					operators.And(VarConst(y, c2), ConstVar(c2, y)),
					circuit,
				),
			),
		),
	)
	// Run it
	filePath := test.CNFName(varVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x", "y")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVar_Encoding(t *testing.T) {
	test.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(
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

func TestVarVar_Simplified(t *testing.T) {
	test.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(
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

func TestNotVarVar_Encoding(t *testing.T) {
	test.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(
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

func TestNotVarVar_Simplified(t *testing.T) {
	test.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(
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

func TestVarVar_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	formula := VarVar(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarVar_IsTrivial(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	formula := VarVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
