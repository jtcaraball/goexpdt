package isnode

import (
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/internal/test/context"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const varSUFIX = "isnode.var"

// =========================== //
//           HELPERS           //
// =========================== //

func runIsNodeVar(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	// Define variable and ctx
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, genTree())
	// Define circuit
	var circuit base.Component = Var(x)
	if neg {
		circuit = operators.Not(circuit)
	}
	// Define formula
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
	// Run it
	filePath := solver.CNFName(varSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeVar(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeVar(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestNotVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeVar(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestNotVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runIsNodeVar(t, i, tc.expCode, tc.val, true, true)
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
