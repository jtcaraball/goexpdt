package lel

import (
	"stratifoiled/components"
	"stratifoiled/components/circuits/subsumption"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varConstSUFIX = "lel.varconst"
const guardedVarConstSUFIX = "lel.Gvarconst"

// =========================== //
//           HELPERS           //
// =========================== //

func runLELVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2 components.Const,
	simplify bool,
) {
	x := components.NewVar("x")
	context := components.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			VarConst(x, c2),
		),
	)
	filePath := sfdtest.CNFName(varConstSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

func runGuardedLELVarConst(
	t *testing.T,
	id, expCode int,
	c1, c2 components.Const,
	simplify bool,
) {
	x := components.NewVar("x")
	y := components.GuardedConst("y")
	context := components.NewContext(DIM, nil)
	context.Guards = append(
		context.Guards,
		components.Guard{Target: "y", Value: c2, Rep: "1"},
	)
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			VarConst(x, y),
		),
	)
	filePath := sfdtest.CNFName(guardedVarConstSUFIX, id, simplify)
	encodeAndRun(t, formula, context, filePath, id, expCode, simplify)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarConst_Encoding(t *testing.T) {
	sfdtest.AddCleanup(t, varConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarConst_Encoding_Guarded(t *testing.T) {
	sfdtest.AddCleanup(t, guardedVarConstSUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELVarConst(t, i, tc.expCode, tc.val1, tc.val2, false)
		})
	}
}

func TestVarConst_Encoding_WrongDim(t *testing.T) {
	x := components.NewVar("x")
	y := components.Const{components.BOT, components.BOT, components.BOT}
	formula := VarConst(x, y)
	context := components.NewContext(4, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func TestVarConst_Simplified(t *testing.T) {
	sfdtest.AddCleanup(t, varConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runLELVarConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarConst_Simplified_Guarded(t *testing.T) {
	sfdtest.AddCleanup(t, guardedVarConstSUFIX, true)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runGuardedLELVarConst(t, i, tc.expCode, tc.val1, tc.val2, true)
		})
	}
}

func TestVarConst_GetChildren(t *testing.T) {
	x := components.NewVar("x")
	y := components.Const{components.BOT, components.BOT, components.BOT}
	formula := VarConst(x, y)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func TestVarConst_IsTrivial(t *testing.T) {
	x := components.NewVar("x")
	y := components.Const{components.BOT, components.BOT, components.BOT}
	formula := VarConst(x, y)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
