package lel

import "stratifoiled/base"

const DIM = 3

var tests = []struct {
	name string
	val1 base.Const
	val2 base.Const
	expCode int
}{
	{
		name: "(1,_,_):(1,_,0)",
		val1: base.Const{base.ONE, base.BOT, base.BOT},
		val2: base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 10,
	},
	{
		name: "(_,_,_):(1,1,0)",
		val1: base.Const{base.BOT, base.BOT, base.BOT},
		val2: base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name: "(1,_,_):(_,_,1)",
		val1: base.Const{base.ONE, base.BOT, base.BOT},
		val2: base.Const{base.BOT, base.BOT, base.ONE},
		expCode: 10,
	},
	{
		name: "(1,1,1):(0,1,1)",
		val1: base.Const{base.ONE, base.ONE, base.ZERO},
		val2: base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 10,
	},
	{
		name: "(1,_,0):(_,_,0)",
		val1: base.Const{base.ONE, base.BOT, base.ZERO},
		val2: base.Const{base.BOT, base.BOT, base.ZERO},
		expCode: 20,
	},
	{
		name: "(1,1,0):(1,_,_)",
		val1: base.Const{base.ONE, base.ONE, base.ZERO},
		val2: base.Const{base.ONE, base.BOT, base.BOT},
		expCode: 20,
	},
}
