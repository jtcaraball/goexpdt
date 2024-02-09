package lel

import "stratifoiled/components"

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
		name: "(_,_,_):(1,1,0)",
		val1: components.Const{components.BOT, components.BOT, components.BOT},
		val2: components.Const{components.ONE, components.ONE, components.ZERO},
		expCode: 10,
	},
	{
		name: "(1,_,_):(_,_,1)",
		val1: components.Const{components.ONE, components.BOT, components.BOT},
		val2: components.Const{components.BOT, components.BOT, components.ONE},
		expCode: 10,
	},
	{
		name: "(1,1,1):(0,1,1)",
		val1: components.Const{components.ONE, components.ONE, components.ZERO},
		val2: components.Const{components.ZERO, components.ONE, components.ONE},
		expCode: 10,
	},
	{
		name: "(1,_,0):(_,_,0)",
		val1: components.Const{components.ONE, components.BOT, components.ZERO},
		val2: components.Const{components.BOT, components.BOT, components.ZERO},
		expCode: 20,
	},
	{
		name: "(1,1,0):(1,_,_)",
		val1: components.Const{components.ONE, components.ONE, components.ZERO},
		val2: components.Const{components.ONE, components.BOT, components.BOT},
		expCode: 20,
	},
}
