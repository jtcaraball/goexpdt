package leh

import "goexpdt/base"

const DIM = 3

type testS struct {
	name    string
	val1    base.Const
	val2    base.Const
	val3    base.Const
	expCode int
}

var tests = []testS{
	{
		name:    "(0,0,0),(0,0,1),(1,1,0)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(0,0,0),(0,0,1),(0,1,1)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 10,
	},
	{
		name:    "(0,0,0),(0,0,1),(1,0,0)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ONE, base.ZERO, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(0,0,0),(0,0,1),(0,1,0)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(0,0,0),(0,0,1),(0,0,1)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(0,0,0),(1,0,1),(0,0,1)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,_,_),(1,1,0),(1,1,0)",
		val1:    base.Const{base.ONE, base.BOT, base.BOT},
		val2:    base.Const{base.ONE, base.ONE, base.ZERO},
		val3:    base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,0,0),(1,1,_),(1,1,0)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ONE, base.BOT},
		val3:    base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,1,1),(1,1,0),(1,_,0)",
		val1:    base.Const{base.ONE, base.ONE, base.ONE},
		val2:    base.Const{base.ONE, base.ONE, base.ZERO},
		val3:    base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 20,
	},
}

var notTests = []testS{
	{
		name:    "(0,0,0),(0,0,1),(1,1,0)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(0,0,0),(0,0,1),(0,1,1)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 20,
	},
	{
		name:    "(0,0,0),(0,0,1),(1,0,0)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ONE, base.ZERO, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(0,0,0),(0,0,1),(0,1,0)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(0,0,0),(0,0,1),(0,0,1)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ZERO, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(0,0,0),(1,0,1),(0,0,1)",
		val1:    base.Const{base.ZERO, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ZERO, base.ONE},
		val3:    base.Const{base.ZERO, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,_,_),(1,1,0),(1,1,0)",
		val1:    base.Const{base.ONE, base.BOT, base.BOT},
		val2:    base.Const{base.ONE, base.ONE, base.ZERO},
		val3:    base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,0,0),(1,1,_),(1,1,0)",
		val1:    base.Const{base.ONE, base.ZERO, base.ZERO},
		val2:    base.Const{base.ONE, base.ONE, base.BOT},
		val3:    base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,1,1),(1,1,0),(1,_,0)",
		val1:    base.Const{base.ONE, base.ONE, base.ONE},
		val2:    base.Const{base.ONE, base.ONE, base.ZERO},
		val3:    base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 10,
	},
}
