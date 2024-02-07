package subsumption

import "stratifoiled/components/instances"

const DIM = 3

var tests = []struct {
	name string
	val1 instances.Const
	val2 instances.Const
	expCode int
}{
	{
		name: "(1,_,_):(1,_,0)",
		val1: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ONE, instances.BOT, instances.ZERO},
		expCode: 10,
	},
	{
		name: "(1,_,_):(0,_,0)",
		val1: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ZERO, instances.BOT, instances.ZERO},
		expCode: 20,
	},
	{
		name: "(1,0,0):(1,_,_)",
		val1: instances.Const{instances.ONE, instances.ZERO, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		expCode: 20,
	},
	{
		name: "(_,_,_):(1,0,1)",
		val1: instances.Const{instances.BOT, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ONE, instances.ZERO, instances.ONE},
		expCode: 10,
	},
	{
		name: "(_,_,0):(1,_,1)",
		val1: instances.Const{instances.BOT, instances.BOT, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.BOT, instances.ONE},
		expCode: 20,
	},
	{
		name: "(1,1,1):(1,1,1)",
		val1: instances.Const{instances.ONE, instances.ONE, instances.ONE},
		val2: instances.Const{instances.ONE, instances.ONE, instances.ONE},
		expCode: 10,
	},
	{
		name: "(1,0,0):(1,0,1)",
		val1: instances.Const{instances.ONE, instances.ZERO, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.ZERO, instances.ONE},
		expCode: 20,
	},
}
