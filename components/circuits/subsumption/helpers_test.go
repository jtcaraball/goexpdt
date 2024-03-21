package subsumption

import (
	"testing"
	"stratifoiled/sfdtest"
	"stratifoiled/components"
)

const DIM = 3

var tests = []struct {
	name string
	val1 components.Const
	val2 components.Const
	expCode int
}{
	{
		name: "(1,_,_):(1,_,0)",
		val1: components.Const{components.ONE, components.BOT, components.BOT},
		val2: components.Const{components.ONE, components.BOT, components.ZERO},
		expCode: 10,
	},
	{
		name: "(1,_,_):(0,_,0)",
		val1: components.Const{components.ONE, components.BOT, components.BOT},
		val2: components.Const{components.ZERO, components.BOT, components.ZERO},
		expCode: 20,
	},
	{
		name: "(1,0,0):(1,_,_)",
		val1: components.Const{components.ONE, components.ZERO, components.ZERO},
		val2: components.Const{components.ONE, components.BOT, components.BOT},
		expCode: 20,
	},
	{
		name: "(_,_,_):(1,0,1)",
		val1: components.Const{components.BOT, components.BOT, components.BOT},
		val2: components.Const{components.ONE, components.ZERO, components.ONE},
		expCode: 10,
	},
	{
		name: "(_,_,0):(1,_,1)",
		val1: components.Const{components.BOT, components.BOT, components.ZERO},
		val2: components.Const{components.ONE, components.BOT, components.ONE},
		expCode: 20,
	},
	{
		name: "(1,1,1):(1,1,1)",
		val1: components.Const{components.ONE, components.ONE, components.ONE},
		val2: components.Const{components.ONE, components.ONE, components.ONE},
		expCode: 10,
	},
	{
		name: "(1,0,0):(1,0,1)",
		val1: components.Const{components.ONE, components.ZERO, components.ZERO},
		val2: components.Const{components.ONE, components.ZERO, components.ONE},
		expCode: 20,
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
