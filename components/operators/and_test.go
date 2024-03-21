package operators

import (
	"stratifoiled/components"
	"stratifoiled/sfdtest"
	"testing"
)

func TestAnd_Encoding_DTrue(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(true))
	childY := WithVar(y, components.NewTrivial(true))
	context := components.NewContext(1, nil)
	component := And(childX, childY)
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
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
	sfdtest.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Encoding_DFalse(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(false))
	childY := WithVar(y, components.NewTrivial(false))
	context := components.NewContext(1, nil)
	component := And(childX, childY)
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}, {}}
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
	sfdtest.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Encoding_Mixed(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(true))
	childY := WithVar(y, components.NewTrivial(false))
	context := components.NewContext(1, nil)
	component := And(childX, childY)
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
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
	sfdtest.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Simplified_DTrue(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(true))
	childY := WithVar(y, components.NewTrivial(true))
	context := components.NewContext(1, nil)
	component := And(childX, childY)
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
	expSClauses := [][]int{}
	expCClauses := [][]int{}
	sfdtest.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Simplified_DFalse(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(false))
	childY := WithVar(y, components.NewTrivial(false))
	context := components.NewContext(1, nil)
	component := And(childX, childY)
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

func TestAnd_Simplified_Mixed(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(true))
	childY := WithVar(y, components.NewTrivial(false))
	context := components.NewContext(1, nil)
	component := And(childX, childY)
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

func TestAnd_Children(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(true))
	childY := WithVar(y, components.NewTrivial(false))
	component := And(childX, childY)
	compChildren := component.GetChildren()
	expCompChildren := []*withVar{childX, childY}
	if len(compChildren) != len(expCompChildren) {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			len(expCompChildren),
			len(compChildren),
		)
		return
	}
	for i, elem := range compChildren {
		if elem != expCompChildren[i] {
			t.Errorf(
				"Wrong children. Expected pointer %p but got %p",
				expCompChildren,
				compChildren,
			)
		}
	}
}

func TestAnd_IsTrivial(t *testing.T) {
	x := components.NewVar("x")
	y := components.NewVar("y")
	childX := WithVar(x, components.NewTrivial(true))
	childY := WithVar(y, components.NewTrivial(false))
	component := And(childX, childY)
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
