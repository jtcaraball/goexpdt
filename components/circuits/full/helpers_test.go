package full

import (
	"testing"
	"stratifoiled/sfdtest"
	"stratifoiled/components"
)

const DIM = 3

var tests = []struct {
	name string
	val components.Const
	expCode int
}{
	{
		name: "(_,1,0)",
		val: components.Const{components.BOT, components.ONE, components.ZERO},
		expCode: 20,
	},
	{
		name: "(1,_,1)",
		val: components.Const{components.ONE, components.BOT, components.ONE},
		expCode: 20,
	},
	{
		name: "(0,1,_)",
		val: components.Const{components.ZERO, components.ONE, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,1,1)",
		val: components.Const{components.ZERO, components.ONE, components.ONE},
		expCode: 10,
	},
}

func encodeAndRun(
	t *testing.T,
	formula components.Component,
	context *components.Context,
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
	sfdtest.RunFormulaTest(t, id, expCode, filePath)
}
