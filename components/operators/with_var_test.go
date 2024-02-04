package operators

import (
	"testing"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

func TestWithVar_Encoding(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	trivial := components.NewTrivial(false)
	context := components.NewContext(1, nil)
	component := &WithVar{
		instance: x,
		child: &WithVar{instance: y, child: trivial},
	}
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error: %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{
		{1, 2, 3},
		{-1, -2},
		{-1, -3},
		{-2, -3},
		{4, 5, 6},
		{-4, -5},
		{-4, -6},
		{-5, -6},
	}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestWithVar_Simplified(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	trivial := components.NewTrivial(false)
	context := components.NewContext(1, nil)
	component := &WithVar{
		instance: x,
		child: &WithVar{instance: y, child: trivial},
	}
	simpleComponent, err := component.Simplified()
	if err != nil {
		t.Errorf("Simplification error: %s", err.Error())
		return
	}
	encCNF, err := simpleComponent.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error: %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestWithVar_GetChildren(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	trivial := components.NewTrivial(false)
	childComp := &WithVar{instance: y, child: trivial}
	component := &WithVar{
		instance: x,
		child: childComp,
	}
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

func TestWithVar_IsTrivial(t *testing.T) {
	x := instances.NewVar("x")
	trivial := components.NewTrivial(false)
	component := &WithVar{
		instance: x,
		child: trivial,
	}
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong is trivial value. Expected %t but got %t", false, true)
	}
}
