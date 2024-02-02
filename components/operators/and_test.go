package operators

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"testing"
)

func TestAnd_Encoding_DTrue(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(true)}
	childY := &WithVar{instance: y, child: components.NewTrivial(true)}
	context := components.NewContext(1, nil)
	component := &And{ child1: childX, child2: childY }
	encCNF := component.Encoding(context)
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
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Encoding_DFalse(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(false)}
	childY := &WithVar{instance: y, child: components.NewTrivial(false)}
	context := components.NewContext(1, nil)
	component := &And{ child1: childX, child2: childY }
	encCNF := component.Encoding(context)
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
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Encoding_Mixed(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(true)}
	childY := &WithVar{instance: y, child: components.NewTrivial(false)}
	context := components.NewContext(1, nil)
	component := &And{ child1: childX, child2: childY }
	encCNF := component.Encoding(context)
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

func TestAnd_Simplified_DTrue(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(true)}
	childY := &WithVar{instance: y, child: components.NewTrivial(true)}
	context := components.NewContext(1, nil)
	component := &And{ child1: childX, child2: childY }
	encCNF := component.Simplified().Encoding(context)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Simplified_DFalse(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(false)}
	childY := &WithVar{instance: y, child: components.NewTrivial(false)}
	context := components.NewContext(1, nil)
	component := &And{ child1: childX, child2: childY }
	encCNF := component.Simplified().Encoding(context)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Simplified_Mixed(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(true)}
	childY := &WithVar{instance: y, child: components.NewTrivial(false)}
	context := components.NewContext(1, nil)
	component := &And{ child1: childX, child2: childY }
	encCNF := component.Simplified().Encoding(context)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestAnd_Children(t *testing.T) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(true)}
	childY := &WithVar{instance: y, child: components.NewTrivial(false)}
	component := &And{ child1: childX, child2: childY }
	compChildren := component.GetChildren()
	expCompChildren := []*WithVar{childX, childY}
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
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	childX := &WithVar{instance: x, child: components.NewTrivial(true)}
	childY := &WithVar{instance: y, child: components.NewTrivial(false)}
	component := &And{ child1: childX, child2: childY }
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong is trivial value. Expected %t but got %t", false, true)
	}
}
