package cons

import (
	"goexpdt/base"
	"goexpdt/internal/test/solver"
	"goexpdt/internal/test/context"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"testing"
)

const varVarSUFIX = "cons.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runConsVarVar(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	neg, simplify bool,
) {
	x := base.NewVar("x")
	y := base.NewVar("y")
	ctx := base.NewContext(DIM, nil)
	var circuit base.Component = VarVar(x, y)
	if neg {
		circuit = operators.Not(circuit)
	}
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
					circuit,
				),
			),
		),
	)
	filePath := solver.CNFName(varVarSUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x", "y")
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsVarVar(t, i, tc.expCode, tc.val1, tc.val2, false, false)
		})
	}
}

func TestNotVarVar_Encoding(t *testing.T) {
	solver.AddCleanup(t, varVarSUFIX, false)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runConsVarVar(t, i, tc.expCode, tc.val1, tc.val2, true, false)
		})
	}
}

func TestVarVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runConsVarVar(t, i, tc.expCode, tc.val1, tc.val2, false, true)
		})
	}
}

func TestNotVarVar_Simplified(t *testing.T) {
	solver.AddCleanup(t, varVarSUFIX, true)
	for i, tc := range notTests {
		t.Run(tc.name, func(t *testing.T) {
			runConsVarVar(t, i, tc.expCode, tc.val1, tc.val2, true, true)
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
