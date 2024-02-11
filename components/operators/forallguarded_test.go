package operators

import (
	"testing"
	"stratifoiled/sfdtest"
	"stratifoiled/components"
)

func TestForAllGuarded_Encoding(t *testing.T) {
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
