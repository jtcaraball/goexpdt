package full

import "stratifoiled/base"

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
