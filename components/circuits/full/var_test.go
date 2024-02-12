package full

import (
	"stratifoiled/components"
	"stratifoiled/components/circuits/subsumption"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varSUFIX = "full.var"

// =========================== //
//           HELPERS           //
// =========================== //

func runFullVar(
	t *testing.T,
	id, expCode int,
	c components.Const,
	simplify bool,
) {
	x := components.NewVar("x")
	context := components.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.ConstVar(c, x),
				subsumption.VarConst(x, c),
			),
			Var(x),
		),
	)
	filePath := sfdtest.CNFName(varSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVar_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullVar(t, i, tc.expCode, tc.val, false)
		})
	}
}

func TestVar_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runFullVar(t, i, tc.expCode, tc.val, true)
		})
	}
}

func TestVar_GetChildren(t *testing.T) {
	x := components.NewVar("x")
	formula := Var(x)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVar_IsTrivial(t *testing.T) {
	x := components.NewVar("x")
	formula := Var(x)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
