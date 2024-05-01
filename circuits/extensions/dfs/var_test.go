package dfs

import (
	"goexpdt/base"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/internal/test/context"
	"goexpdt/internal/test/solver"
	"goexpdt/operators"
	"testing"
)

const varSUFIX = "dfs.var"

// =========================== //
//           HELPERS           //
// =========================== //

func runDFSVar(
	t *testing.T,
	id, expCode int,
	c base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, genTree())
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
			runDFSVar(t, i, tc.expCode, tc.val, false, false)
		})
	}
}

func TestNotVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runDFSVar(t, i, tc.expCode, tc.val, true, false)
		})
	}
}

func TestVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runDFSVar(t, i, tc.expCode, tc.val, false, true)
		})
	}
}

func TestNotVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runDFSVar(t, i, tc.expCode, tc.val, true, true)
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