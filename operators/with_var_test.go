package operators

import (
	"slices"
	"goexpdt/base"
	"testing"
)

func TestWithVar_Encoding(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	trivial := base.NewTrivial(false)
	context := base.NewContext(1, nil)
	component := WithVar(x, WithVar(y, trivial))
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
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestWithVar_Encoding_AddVarToScope(t *testing.T) {
	x := base.Var("x")
	context := base.NewContext(1, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{
			Target: "T",
			Value: base.Const{base.BOT},
			Rep: "1",
		},
	)
	component := WithVar(x, base.NewTrivial(true))
	expScope := []string{"x"}
	_, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	resultingScopes := context.Guards[0].InScope
	if !slices.Equal[[]string](resultingScopes, expScope){
		t.Errorf(
			"Var not included in guard scope. Expected %s but got %s",
			expScope,
			resultingScopes,
		)
	}
}

func TestWithVar_Encoding_ScopedVariable(t *testing.T) {
	x := base.Var("x")
	context := base.NewContext(1, nil)
	context.Guards = append(
		context.Guards,
		base.Guard{
			Target: "T",
			Value: base.Const{base.BOT},
			Rep: "1",
		},
	)
	component := WithVar(x, base.NewTrivial(true))
	_, err := component.Encoding(context)
	if err != nil {
		t.Errorf("CNF encoding error. %s", err.Error())
		return
	}
	for key := range context.GetFeatVars() {
		if key.Name != "x1" {
			t.Errorf("Wrong scoped var name. Expected x1 but got %s", key.Name)
			return
		}
	}
}

func TestWithVar_Simplified(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	trivial := base.NewTrivial(false)
	context := base.NewContext(1, nil)
	component := WithVar(x, WithVar(y, trivial))
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
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestWithVar_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	trivial := base.NewTrivial(false)
	childComp := WithVar(y, trivial)
	component := WithVar(x, childComp)
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
	x := base.NewVar("x")
	trivial := base.NewTrivial(false)
	component := WithVar(x, trivial)
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
