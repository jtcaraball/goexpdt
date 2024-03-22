package lel

import (
	"stratifoiled/base"
	"stratifoiled/circuits/subsumption"
	"stratifoiled/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varVarSUFIX = "lel.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	simplify bool,
) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	context := base.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(
					subsumption.VarConst(x, c1),
					subsumption.ConstVar(c1, x),
				),
				operators.And(
					operators.And(
						subsumption.VarConst(y, c2),
						subsumption.ConstVar(c2, y),
					),
					VarVar(x, y),
				),
			),
		),
	)
	filePath := sfdtest.CNFName(varVarSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVar_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarVar(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarVar_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarVar(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarVar_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	formula := VarVar(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarVar_IsTrivial(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	formula := VarVar(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
