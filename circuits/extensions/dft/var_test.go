package dft

import (
	"goexpdt/base"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"goexpdt/circuits/internal/test"
	"testing"
)

const varSUFIX = "dft.var"

// =========================== //
//           HELPERS           //
// =========================== //

func runDFTVar(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	context := base.NewContext(DIM, genTree())
	var circuit base.Component = Var(x)
	if neg {
		circuit = operators.Not(circuit)
	}
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.ConstVar(c, x),
				subsumption.VarConst(x, c),
			),
			circuit,
		),
	)
	filePath := test.CNFName(varSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encoding(t *testing.T) {
	test.AddCleanup(t, varSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runDFTVar(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestNotVar_Encoding(t *testing.T) {
	test.AddCleanup(t, varSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runDFTVar(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestVar_Simplified(t *testing.T) {
	test.AddCleanup(t, varSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runDFTVar(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestNotVar_Simplified(t *testing.T) {
	test.AddCleanup(t, varSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runDFTVar(t, i, tc.expCode, tc.val, true, true)
		})
	}
}

func TestVar_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	formula := Var(x)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVar_IsTrivial(t *testing.T) {
	x := base.NewVar("x")
	formula := Var(x)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
