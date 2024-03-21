package allcomp

import (
	"fmt"
	"stratifoiled/components"
	"stratifoiled/sfdtest"
	"stratifoiled/trees"
	"testing"
)

const (
	constSUFIX = "allComp.const"
	guardedConstSUFIX = "allComp.Gconst"
)

// =========================== //
//           HELPERS           //
// =========================== //

func runAllCompConst(
	t *testing.T,
	id, expCode int,
	c components.Const,
	tree *trees.Tree,
	leafValue bool,
	simplify bool,
) {
	context := components.NewContext(DIM, tree)
	formula := Const(c, leafValue)
	filePath := sfdtest.CNFName(compConstSufix(leafValue), id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedAllCompConst(
	t *testing.T,
	id, expCode int,
	c components.Const,
	tree *trees.Tree,
	leafValue bool,
	simplify bool,
) {
	context := components.NewContext(DIM, tree)
	formula := Const(c, leafValue)
	filePath := sfdtest.CNFName(compGuardedConstSufix(leafValue), id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func compConstSufix(val bool) string {
	return constSUFIX + fmt.Sprintf("%t", val)
}

func compGuardedConstSufix(val bool) string {
	return guardedConstSUFIX + fmt.Sprintf("%t", val)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestConst_Encoding_AllPos(t *testing.T) {
	sfdtest.AddCleanup(t, compConstSufix(true), false)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompConst(t, i, tc.expCode, tc.val, tree, true, false)
		})
	}
}

func TestConst_Encoding_AllPos_Guarded(t *testing.T) {
	sfdtest.AddCleanup(t, compGuardedConstSufix(true), false)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc.expCode, tc.val, tree, true, false)
		})
	}
}

func TestConst_Encoding_AllNeg(t *testing.T) {
	sfdtest.AddCleanup(t, compConstSufix(false), false)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompConst(t, i, tc.expCode, tc.val, tree, false, false)
		})
	}
}

func TestConst_Encoding_AllNeg_Guraded(t *testing.T) {
	sfdtest.AddCleanup(t, compGuardedConstSufix(false), false)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc.expCode, tc.val, tree, false, false)
		})
	}
}

func TestConst_Encoding_WrongDim(t *testing.T) {
	x := components.Const{components.BOT, components.BOT, components.BOT}
	formula := Const(x, true)
	context := components.NewContext(4, &trees.Tree{Root: &trees.Node{}})
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_Simplified_AllPos(t *testing.T) {
	sfdtest.AddCleanup(t, compConstSufix(true), true)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompConst(t, i, tc.expCode, tc.val, tree, true, true)
		})
	}
}

func TestConst_Simplified_AllPos_Guarded(t *testing.T) {
	sfdtest.AddCleanup(t, compGuardedConstSufix(true), true)
	tree := genTree()
	for i, tc := range allPosTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc.expCode, tc.val, tree, true, true)
		})
	}
}

func TestConst_Simplified_AllNeg(t *testing.T) {
	sfdtest.AddCleanup(t, compConstSufix(false), true)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runAllCompConst(t, i, tc.expCode, tc.val, tree, false, true)
		})
	}
}

func TestConst_Simplified_AllNeg_Guraded(t *testing.T) {
	sfdtest.AddCleanup(t, compGuardedConstSufix(false), true)
	tree := genTree()
	for i, tc := range allNegTests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedAllCompConst(t, i, tc.expCode, tc.val, tree, false, true)
		})
	}
}

func TestConstConst_Simplified_WrongDim(t *testing.T) {
	x := components.Const{components.BOT, components.BOT, components.BOT}
	formula := Const(x, true)
	context := components.NewContext(4, &trees.Tree{Root: &trees.Node{}})
	_, err := formula.Simplified(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestConst_GetChildren(t *testing.T) {
	x := components.NewVar("x")
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

func TestConst_IsTrivial(t *testing.T) {
	x := components.NewVar("x")
	formula := Var(x, true)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
