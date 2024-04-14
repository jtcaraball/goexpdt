package operators

import (
	"goexpdt/base"
	"goexpdt/internal/test/clauses"
	"testing"
)

func TestOr_Encoding_DTrue(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(true))
	childY := WithVar(y, base.NewTrivial(true))
	context := base.NewContext(1, nil)
	component := Or(childX, childY)
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{}
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestOr_Encoding_DFalse(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(false))
	childY := WithVar(y, base.NewTrivial(false))
	context := base.NewContext(1, nil)
	component := Or(childX, childY)
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestOr_Encoding_Mixed(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(true))
	childY := WithVar(y, base.NewTrivial(false))
	context := base.NewContext(1, nil)
	component := Or(childX, childY)
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{}
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestOr_Simplified_DTrue(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(true))
	childY := WithVar(y, base.NewTrivial(true))
	context := base.NewContext(1, nil)
	component := Or(childX, childY)
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
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestOr_Simplified_DFalse(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(false))
	childY := WithVar(y, base.NewTrivial(false))
	context := base.NewContext(1, nil)
	component := Or(childX, childY)
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
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestOr_Simplified_Mixed(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(true))
	childY := WithVar(y, base.NewTrivial(false))
	context := base.NewContext(1, nil)
	component := Or(childX, childY)
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
	clauses.ValidateClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestOr_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(true))
	childY := WithVar(y, base.NewTrivial(false))
	component := Or(childX, childY)
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

func TestOr_IsTrivial(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	childX := WithVar(x, base.NewTrivial(true))
	childY := WithVar(y, base.NewTrivial(false))
	component := Or(childX, childY)
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
