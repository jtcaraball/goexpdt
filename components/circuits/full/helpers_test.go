package full

import "stratifoiled/components"

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
