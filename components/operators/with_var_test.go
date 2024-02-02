package operators

import (
	"slices"
	"testing"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           HELPERS           //
// =========================== //

func clausesEq(c1, c2 [][]int) bool {
	return slices.EqualFunc[[][]int](
		c1,
		c2,
		func (l1, l2 []int) bool {
			return slices.Equal[[]int](l1, l2)
		},
	)
}

func errorInClauses(
	t *testing.T,
	sClauses, cClauses, expSClauses, expCClauses [][]int,
) {
	if !clausesEq(sClauses, expSClauses) {
		t.Errorf(
			"Semantic clauses not equal. Expected %d but got %d",
			sClauses,
			expSClauses,
		)
	}
	if !clausesEq(cClauses, expCClauses) {
		t.Errorf(
			"Consistency clauses not equal. Expected %d but got %d",
			cClauses,
			expCClauses,
		)
	}
}

// =========================== //
//            TESTS            //
// =========================== //

func TestWithVar_Encoding(t *testing.T) {
	var x instances.Var = "x"
	var y instances.Var = "y"
	var trivial components.Trivial = false
	context := components.NewContext(1, nil)
	component := &WithVar{
		instance: x,
		child: &WithVar{instance: y, child: &trivial},
	}
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

func TestWithVar_Simplified(t *testing.T) {
	var x instances.Var = "x"
	var y instances.Var = "y"
	var trivial components.Trivial = false
	context := components.NewContext(1, nil)
	component := &WithVar{
		instance: x,
		child: &WithVar{instance: y, child: &trivial},
	}
	encCNF := component.Simplified().Encoding(context)
	sClauses, cClauses := encCNF.Clauses()
	expSClauses := [][]int{{}}
	expCClauses := [][]int{}
	errorInClauses(t, sClauses, cClauses, expSClauses, expCClauses)
}

func TestWithVar_GetChildren(t *testing.T) {
	var x instances.Var = "x"
	var y instances.Var = "y"
	var trivial components.Trivial = false
	childComp := &WithVar{instance: y, child: &trivial}
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
	var x instances.Var = "x"
	var trivial components.Trivial = false
	component := &WithVar{
		instance: x,
		child: &trivial,
	}
	isTrivial, _ := component.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong is trivial value. Expected %t but got %t", false, true)
	}
}
