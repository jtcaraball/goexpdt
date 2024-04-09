package subsumption

import "goexpdt/base"

const DIM = 3

type testS struct {
	name    string
	val1    base.Const
	val2    base.Const
	expCode int
}

var tests = []testS{
	{
		name:    "(1,_,_):(1,_,0)",
		val1:    base.Const{base.ONE, base.BOT, base.BOT},
		val2:    base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,_,_):(0,_,0)",
		val1:    base.Const{base.ONE, base.BOT, base.BOT},
		val2:    base.Const{base.ZERO, base.BOT, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,0,0):(1,_,_)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.BOT, base.BOT},
		expCode: 20,
	},
	{
		name:    "(_,_,_):(1,0,1)",
		val1:    base.Const{base.BOT, base.BOT, base.BOT},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(_,_,0):(1,_,1)",
		val1:    base.Const{base.BOT, base.BOT, base.ZERO},
		val2:    base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,1,1):(1,1,1)",
		val1:    base.Const{base.ONE, base.ONE, base.ONE},
		val2:    base.Const{base.ONE, base.ONE, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,0,0):(1,0,1)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,0,0):(_,_,_)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 20,
	},
}

var notTests = []testS{
	{
		name:    "(1,_,_):(1,_,0)",
		val1:    base.Const{base.ONE, base.BOT, base.BOT},
		val2:    base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,_,_):(0,_,0)",
		val1:    base.Const{base.ONE, base.BOT, base.BOT},
		val2:    base.Const{base.ZERO, base.BOT, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,0,0):(1,_,_)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.BOT, base.BOT},
		expCode: 10,
	},
	{
		name:    "(_,_,_):(1,0,1)",
		val1:    base.Const{base.BOT, base.BOT, base.BOT},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(_,_,0):(1,_,1)",
		val1:    base.Const{base.BOT, base.BOT, base.ZERO},
		val2:    base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,1,1):(1,1,1)",
		val1:    base.Const{base.ONE, base.ONE, base.ONE},
		val2:    base.Const{base.ONE, base.ONE, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,0,0):(1,0,1)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,0,0):(_,_,_)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 10,
	},
}
