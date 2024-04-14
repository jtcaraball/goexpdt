package allcomp

import (
	"fmt"
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/internal/test/context"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"goexpdt/trees"
	"testing"
)

const varSUFIX = "allComp.var"

// =========================== //
//           HELPERS           //
// =========================== //

func runAllCompVar(
	t *testing.T,
	id, expCode int,
	c base.Const,
	tree *trees.Tree,
	leafValue, neg, simplify bool,
) {
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, tree)
	var circuit base.Component = Var(x, leafValue)
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
	filePath := solver.CNFName(compVarSufix(leafValue), id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x")
}

func compVarSufix(val bool) string {
	return varSUFIX + fmt.Sprintf("%t", val)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encoding_AllPos(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(true), false)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, true, false, false)
		})
	}
}

func TestNotVar_Encoding_AllPos(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(true), false)
	tree := genTree()
	for i, tc := range allPosNotTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, true, true, false)
		})
	}
}

func TestVar_Encoding_AllNeg(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(false), false)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, false, false, false)
		})
	}
}

func TestNotVar_Encoding_AllNeg(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(false), false)
	tree := genTree()
	for i, tc := range allNegNotTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, false, true, false)
		})
	}
}

func TestVar_Simplified_AllPos(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(true), true)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, true, false, true)
		})
	}
}

func TestNotVar_Simplified_AllPos(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(true), true)
	tree := genTree()
	for i, tc := range allPosNotTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, true, true, true)
		})
	}
}

func TestVar_Simplified_AllNeg(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(false), true)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, false, false, true)
		})
	}
}

func TestNotVar_Simplified_AllNeg(t *testing.T) {
	solver.AddCleanup(t, compVarSufix(false), true)
	tree := genTree()
	for i, tc := range allNegNotTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, false, true, true)
		})
	}
}

func TestVar_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	formula := Var(x, true)
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
	formula := Var(x, true)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
