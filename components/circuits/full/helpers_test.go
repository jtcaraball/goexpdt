package full

import "stratifoiled/components/instances"

const DIM = 3

var tests = []struct {
	name string
	val instances.Const
	expCode int
}{
	{
		name: "(_,1,0)",
		val: instances.Const{instances.BOT, instances.ONE, instances.ZERO},
		expCode: 20,
	},
	{
		name: "(1,_,1)",
		val: instances.Const{instances.ONE, instances.BOT, instances.ONE},
		expCode: 20,
	},
	{
		name: "(0,1,_)",
		val: instances.Const{instances.ZERO, instances.ONE, instances.BOT},
		expCode: 20,
	},
	{
		name: "(0,1,1)",
		val: instances.Const{instances.ZERO, instances.ONE, instances.ONE},
		expCode: 10,
	},
}
