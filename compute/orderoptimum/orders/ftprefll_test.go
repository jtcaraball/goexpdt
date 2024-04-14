package orders

import (
	"goexpdt/base"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/internal/test/context"
	"goexpdt/internal/test/solver"
	"goexpdt/operators"
	"testing"
)

const (
	DIM           = 4
	TESTFILESUFIX = "ftprefll"
)

// =========================== //
//           HELPERS           //
// =========================== //

var tests = []struct {
	name    string
	val1    base.Const
	val2    base.Const
	pref    []int
	expCode int
}{
	{
		name:    "(_,_,_,_):(0,0,0,0)",
		val1:    base.Const{base.BOT, base.BOT, base.BOT, base.BOT},
		val2:    base.Const{base.ZERO, base.ZERO, base.ZERO, base.ZERO},
		pref:    []int{0},
		expCode: 10,
	},
	{
		name:    "(_,_,_,1):(0,0,0,0)",
		val1:    base.Const{base.BOT, base.BOT, base.BOT, base.ONE},
		val2:    base.Const{base.ZERO, base.ZERO, base.ZERO, base.ZERO},
		pref:    []int{0},
		expCode: 10,
	},
	{
		name:    "(_,1,_,1):(0,0,0,0)",
		val1:    base.Const{base.BOT, base.ONE, base.BOT, base.ONE},
		val2:    base.Const{base.ZERO, base.ZERO, base.ZERO, base.ZERO},
		pref:    []int{0},
		expCode: 10,
	},
	{
		name:    "(1,1,_,1):(0,0,0,0)",
		val1:    base.Const{base.ONE, base.ONE, base.BOT, base.ONE},
		val2:    base.Const{base.ZERO, base.ZERO, base.ZERO, base.ZERO},
		pref:    []int{0},
		expCode: 10,
	},
	{
		name:    "(1,1,1,1):(0,0,0,0)",
		val1:    base.Const{base.ONE, base.ONE, base.ONE, base.ONE},
		val2:    base.Const{base.ZERO, base.ZERO, base.ZERO, base.ZERO},
		pref:    []int{0},
		expCode: 20,
	},
	{
		name:    "(_,_,1,1):(_,0,0,_)",
		val1:    base.Const{base.BOT, base.BOT, base.ONE, base.ONE},
		val2:    base.Const{base.BOT, base.ZERO, base.ZERO, base.BOT},
		pref:    []int{0, 2, 3},
		expCode: 10,
	},
	{
		name:    "(_,_,0,1):(_,0,0,_)",
		val1:    base.Const{base.BOT, base.BOT, base.ZERO, base.ONE},
		val2:    base.Const{base.BOT, base.ZERO, base.ZERO, base.BOT},
		pref:    []int{0, 1, 3},
		expCode: 20,
	},
	{
		name:    "(1,1,0,1):(1,0,0,1)",
		val1:    base.Const{base.ONE, base.ONE, base.ZERO, base.ONE},
		val2:    base.Const{base.ONE, base.ZERO, base.ZERO, base.ONE},
		pref:    []int{0, 1, 3, 2},
		expCode: 20,
	},
	{
		name:    "(1,_,0,1):(1,0,0,_)",
		val1:    base.Const{base.ONE, base.BOT, base.ZERO, base.ONE},
		val2:    base.Const{base.ONE, base.ZERO, base.ZERO, base.BOT},
		pref:    []int{0, 3, 2, 1},
		expCode: 10,
	},
	{
		name:    "(_,_,0,1):(1,0,_,_)",
		val1:    base.Const{base.BOT, base.BOT, base.ZERO, base.ONE},
		val2:    base.Const{base.ONE, base.ZERO, base.BOT, base.BOT},
		pref:    []int{0, 3, 2, 1},
		expCode: 20,
	},
	{
		name:    "(0,1,_,1):(1,_,0,1)",
		val1:    base.Const{base.ZERO, base.ONE, base.BOT, base.ONE},
		val2:    base.Const{base.ONE, base.BOT, base.ZERO, base.ONE},
		pref:    []int{3, 1, 0},
		expCode: 10,
	},
	{
		name:    "(0,_,1,1):(1,0,_,1)",
		val1:    base.Const{base.ZERO, base.BOT, base.ONE, base.ONE},
		val2:    base.Const{base.ONE, base.ZERO, base.BOT, base.ONE},
		pref:    []int{3, 0, 1},
		expCode: 20,
	},
}

func runTest(
	t *testing.T,
	id, expCode int,
	c1, c2 base.Const,
	pref []int,
	simplify bool,
) {
	// Define variable and context
	x := base.NewVar("x")
	ctx := base.NewContext(DIM, nil)
	// Define formula
	formula := operators.WithVar(
		x,
		operators.And(
			operators.And(
				subsumption.VarConst(x, c1),
				subsumption.ConstVar(c1, x),
			),
			FTPrefLL(x, c2, pref),
		),
	)
	// Run it
	filePath := solver.CNFName(TESTFILESUFIX, id, simplify)
	solver.EncodeAndRun(t, formula, ctx, filePath, id, expCode, simplify)
	context.OnlyFeatVariables(t, ctx, "x")
}

// =========================== //
//            TESTS            //
// =========================== //

func Test_Encoding(t *testing.T) {
	solver.AddCleanup(t, TESTFILESUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runTest(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.pref,
				false,
			)
		})
	}
}

func Test_Simplified(t *testing.T) {
	solver.AddCleanup(t, TESTFILESUFIX, false)
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runTest(
				t,
				i,
				tc.expCode,
				tc.val1,
				tc.val2,
				tc.pref,
				true,
			)
		})
	}
}

func Test_Encoding_WrongDim(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT, base.BOT}
	formula := FTPrefLL(x, y, []int{})
	context := base.NewContext(3, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected constant wrong dimension error")
	}
}

func Test_Encoding_InvalidFT(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT, base.BOT}
	formula := FTPrefLL(x, y, []int{4})
	context := base.NewContext(3, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error("Error not cached. Expected invalid feature preference error")
	}
}

func Test_Encoding_DupFT(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT, base.BOT}
	formula := FTPrefLL(x, y, []int{1, 1})
	context := base.NewContext(3, nil)
	_, err := formula.Encoding(context)
	if err == nil {
		t.Error(
			"Error not cached. Expected duplicated feature preference error",
		)
	}
}

func Test_GetChildren(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT, base.BOT}
	formula := FTPrefLL(x, y, nil)
	children := formula.GetChildren()
	if len(children) != 0 {
		t.Errorf(
			"Wrong number of children. Expected %d but got %d",
			0,
			len(children),
		)
	}
}

func Test_IsTrivial(t *testing.T) {
	x := base.NewVar("x")
	y := base.Const{base.BOT, base.BOT, base.BOT}
	formula := FTPrefLL(x, y, nil)
	isTrivial, _ := formula.IsTrivial()
	if isTrivial {
		t.Errorf("Wrong IsTrivial value. Expected %t but got %t", false, true)
	}
}
