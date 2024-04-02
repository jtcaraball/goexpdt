package leh

import (
	"goexpdt/base"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"goexpdt/circuits/internal/test"
	"testing"
)

const varVarVarSUFIX = "leh.varvarvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runLEHVarVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2, c3 base.Const,
	simplify bool,
) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	z := base.NewVar("z")
	context := base.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.WithVar(
				z,
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
						operators.And(
							operators.And(
								subsumption.VarConst(z, c3),
								subsumption.ConstVar(c3, z),
							),
							VarVarVar(x, y, z),
						),
					),
				),
			),
		),
	)
	filePath := test.CNFName(varVarVarSUFIX, id, simplify)
	test.EncodeAndRun(t, formula, context, filePath, id, expCode, simplify)
	test.OnlyFeatVariables(t, context, "x", "y", "z")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVarVar_Encoding(t *testing.T) {
	test.AddCleanup(t, varVarVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarVarVar(t, i, tc.expCode, tc.val1, tc.val2, tc.val3, false)
		})
	}
}

func TestVarVarVar_Simplified(t *testing.T) {
	test.AddCleanup(t, varVarVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLEHVarVarVar(t, i, tc.expCode, tc.val1, tc.val2, tc.val3, true)
		})
	}
}

func TestVarVarVar_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	z := base.NewVar("z")
	formula := VarVarVar(x, y, z)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarVarVar_IsTrivial(t *testing.T) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	z := base.NewVar("z")
	formula := VarVarVar(x, y, z)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
