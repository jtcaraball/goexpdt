package full

import (
	"testing"
	"stratifoiled/internal/test"
	"stratifoiled/base"
)

const DIM = 3

var tests = []struct {
	name string
	val base.Const
	expCode int
}{
	{
		name: "(_,1,0)",
		val: base.Const{base.BOT, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name: "(1,_,1)",
		val: base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 20,
	},
	{
		name: "(0,1,_)",
		val: base.Const{base.ZERO, base.ONE, base.BOT},
		expCode: 20,
	},
	{
		name: "(0,1,1)",
		val: base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 10,
	},
}

func encodeAndRun(
	t *testing.T,
	formula base.Component,
	context *base.Context,
	filePath string,
	id, expCode int,
	simplify bool,
) {
	var err error
	if simplify {
		formula, err = formula.Simplified(context)
		if err != nil {
			t.Errorf("Formula simplification error. %s", err.Error())
			return
		}
	}
	cnf, err := formula.Encoding(context)
	if err != nil {
		t.Errorf("Formula encoding error. %s", err.Error())
		return
	}
	if err = cnf.ToFile(filePath); err != nil {
		t.Errorf("CNF writing error. %s", err.Error())
		return
	}
	test.RunFormulaTest(t, id, expCode, filePath)
}
