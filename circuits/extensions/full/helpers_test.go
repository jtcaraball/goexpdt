package full

import "goexpdt/base"

const DIM = 3

type testS struct {
	name    string
	val     base.Const
	expCode int
}

var tests = []testS{
	{
		name:    "(_,1,0)",
		val:     base.Const{base.BOT, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,_,1)",
		val:     base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 20,
	},
	{
		name:    "(0,1,_)",
		val:     base.Const{base.ZERO, base.ONE, base.BOT},
		expCode: 20,
	},
	{
		name:    "(0,1,1)",
		val:     base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 10,
	},
}

var notTests = []testS{
	{
		name:    "(_,1,0)",
		val:     base.Const{base.BOT, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,_,1)",
		val:     base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 10,
	},
	{
		name:    "(0,1,_)",
		val:     base.Const{base.ZERO, base.ONE, base.BOT},
		expCode: 10,
	},
	{
		name:    "(0,1,1)",
		val:     base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 20,
	},
}
