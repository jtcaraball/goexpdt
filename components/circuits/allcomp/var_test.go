package allcomp

import (
	"fmt"
	"stratifoiled/components"
	"stratifoiled/components/circuits/subsumption"
	"stratifoiled/components/instances"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"stratifoiled/trees"
	"testing"
)

const varSUFIX = "allComp.var"

// =========================== //
//           HELPERS           //
// =========================== //

func runAllCompVar(
	t *testing.T,
	id, expCode int,
	c instances.Const,
	tree *trees.Tree,
	leafValue bool,
	simplify bool,
) {
	var err error
	var formula components.Component
	x := instances.NewVar("x")
	context := components.NewContext(DIM, tree)
	formula = operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.ConstVar(c, x),
				subsumption.VarConst(x, c),
			),
			Var(x, leafValue),
		),
	)
	filePath := sfdtest.CNFName(compVarSufix(leafValue), id, simplify)
	if simplify {
		formula, err = formula.Simplified(context)
		if err != nil {
			t.Errorf("Formula simplification error. %s", err.Error())
			return
		}
	}
	cnf, err := formula.Encoding(context)
	if err != nil {
		t.Errorf("Formula encoding error. %s", err.Error())
		return
	}
	if err = cnf.ToFile(filePath); err != nil {
		t.Errorf("CNF writing error. %s", err.Error())
		return
	}
	sfdtest.RunFormulaTest(t, id, expCode, filePath)
}

func compVarSufix(val bool) string {
	return varSUFIX + fmt.Sprintf("%t", val)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encoding_AllPos(t *testing.T) {
	sfdtest.AddCleanup(t, compVarSufix(true), false)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, true, false)
		})
	}
}

func TestVar_Encoding_AllNeg(t *testing.T) {
	sfdtest.AddCleanup(t, compVarSufix(false), false)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, false, false)
		})
	}
}

func TestVar_Simplified_AllPos(t *testing.T) {
	sfdtest.AddCleanup(t, compVarSufix(true), true)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, true, true)
		})
	}
}

func TestVar_Simplified_AllNeg(t *testing.T) {
	sfdtest.AddCleanup(t, compVarSufix(false), true)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompVar(t, i, tc.expCode, tc.val, tree, false, true)
		})
	}
}

func TestVar_GetChildren(t *testing.T) {
	x := instances.NewVar("x")
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
	x := instances.NewVar("x")
	formula := Var(x, true)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
