package full

import "stratifoiled/components/instances"

const DIM = 3

var tests = []struct {
	name string
	val instances.Const
	expCode int
}{
	{
		name: "BOT1",
		val: instances.Const{instances.BOT, instances.ONE, instances.ZERO},
		expCode: 20,
	},
	{
		name: "BOT2",
		val: instances.Const{instances.ONE, instances.BOT, instances.ONE},
		expCode: 20,
	},
	{
		name: "BOT3",
		val: instances.Const{instances.ZERO, instances.ONE, instances.BOT},
		expCode: 20,
	},
	{
		name: "FULL",
		val: instances.Const{instances.ZERO, instances.ONE, instances.ONE},
		expCode: 10,
	},
}
