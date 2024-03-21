package operators

import (
	"stratifoiled/components"
	"stratifoiled/sfdtest"
	"stratifoiled/trees"
	"testing"
)

// =========================== //
//           HELPERS           //
// =========================== //

func buildFAGTree() *trees.Tree {
	// Tree
	// root: _
	//   leaf_1: 0
	//	 leaf_2: 1
	leaf1 := &trees.Node{ID: 1}
	leaf2 := &trees.Node{ID: 2}
	root := &trees.Node{ID: 0, Feat: 0, LChild: leaf1, RChild: leaf2}
	return &trees.Tree{
		Root: root,
		NodeCount: 3,
		FeatCount: 3,
		NegLeafs: []*trees.Node{leaf1, leaf2},
	}
}

// =========================== //
//            TESTS            //
// =========================== //

func TestForAllGuarded_Encoding(t *testing.T) {
	x := components.GuardedConst("x")
	y := components.Var("y")
	component := ForAllGuarded(
		x,
		WithVar(y, components.NewTrivial(true)),
	)
	context := components.NewContext(1, buildFAGTree())
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{
		{1, 2, 3}, {-1, -2}, {-1, -3}, {-2, -3},
		{4, 5, 6}, {-4, -5}, {-4, -6}, {-5, -6},
		{7, 8, 9}, {-7, -8}, {-7, -9}, {-8, -9},
	}
	for key := range context.GetVars() {
		t.Log(key.Name)
	}
	sfdtest.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestForAllGuarded_Simplified(t *testing.T) {
	x := components.GuardedConst("x")
	y := components.NewVar("y")
	trivial := components.NewTrivial(false)
	context := components.NewContext(1, nil)
	component := ForAllGuarded(x, WithVar(y, trivial))
	simpleComponent, err := component.Simplified(context)
	if err != nil {
		t.Errorf("Simplification error. %s", err.Error())
		return
	}
	encCNF, err := simpleComponent.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	sfdtest.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestForAllGuarded_GetChildren(t *testing.T) {
	x := components.GuardedConst("x")
	trivial := components.NewTrivial(false)
	component := ForAllGuarded(x, trivial)
	compChildren := component.GetChildren()
	if len(compChildren) != 1 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			1,
			len(compChildren),
		)
		return
	}
	if compChildren[0] != trivial {
		t.Errorf(
			"Wrong children. Expected pointer %p but got %p",
			trivial,
			compChildren[0],
		)
	}
}

func TestForAllGuarded_IsTrivial(t *testing.T) {
	x := components.GuardedConst("x")
	trivial := components.NewTrivial(false)
	component := ForAllGuarded(x, trivial)
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
