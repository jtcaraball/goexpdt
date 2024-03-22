package operators

import (
	"testing"
	"stratifoiled/base"
	"stratifoiled/internal/test"
)

func TestNot_Encoding(t *testing.T) {
	x := base.NewVar("x")
	trivial := base.NewTrivial(false)
	context := base.NewContext(1, nil)
	component := Not(WithVar(x, trivial))
	encCNF, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{}
	expCClauses := [][]int{}
	test.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestNot_Simplified(t *testing.T) {
	x := base.NewVar("x")
	trivial := base.NewTrivial(true)
	context := base.NewContext(1, nil)
	component := Not(WithVar(x, trivial))
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
	test.ErrorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestNot_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	trivial := base.NewTrivial(true)
	childComp := WithVar(x, trivial)
	component := Not(childComp)
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
	x := base.NewVar("x")
	trivial := base.NewTrivial(false)
	component := Not(WithVar(x, trivial))
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
