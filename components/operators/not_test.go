package operators

import (
	"testing"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

func TestNot_Encoding(t *testing.T) {
	x := instances.NewVar("x")
	trivial := components.NewTrivial(false)
	context := components.NewContext(1, nil)
	component := &Not{ child: &WithVar{ instance: x, child: trivial } }
	encCNF := component.Encoding(context)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestNot_Simplified(t *testing.T) {
	x := instances.NewVar("x")
	trivial := components.NewTrivial(true)
	context := components.NewContext(1, nil)
	component := &Not{ child: &WithVar{ instance: x, child: trivial } }
	encCNF := component.Simplified().Encoding(context)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestNot_GetChildren(t *testing.T) {
	x := instances.NewVar("x")
	trivial := components.NewTrivial(true)
	childComp := &WithVar{ instance: x, child: trivial }
	component := &Not{ child: childComp }
	compChildren := component.GetChildren()
	if len(compChildren) != 1 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			1,
			len(compChildren),
		)
		return
	}
	if compChildren[0] != childComp {
		t.Errorf(
			"Wrong children. Expected pointer %p but got %p",
			childComp,
			compChildren[0],
		)
	}
}

func TestNot_IsTrivial(t *testing.T) {
	x := instances.NewVar("x")
	trivial := components.NewTrivial(false)
	component := &Not{ child: &WithVar{ instance: x, child: trivial } }
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong is trivial value. Expected %t but got %t", false, true)
	}
}
