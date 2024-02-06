package lel

import "stratifoiled/components/instances"

const DIM = 3

var tests = []struct {
	name string
	val1 instances.Const
	val2 instances.Const
	expCode int
}{
	{
		name: "less BOT BOT",
		val1: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ONE, instances.BOT, instances.ZERO},
		expCode: 10,
	},
	{
		name: "less BOT nBOT",
		val1: instances.Const{instances.BOT, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ONE, instances.ONE, instances.ZERO},
		expCode: 10,
	},
	{
		name: "equal BOT BOT",
		val1: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		val2: instances.Const{instances.BOT, instances.BOT, instances.ONE},
		expCode: 10,
	},
	{
		name: "equal nBOT nBOT",
		val1: instances.Const{instances.ONE, instances.ONE, instances.ZERO},
		val2: instances.Const{instances.ZERO, instances.ONE, instances.ONE},
		expCode: 10,
	},
	{
		name: "greater BOT BOT",
		val1: instances.Const{instances.ONE, instances.BOT, instances.ZERO},
		val2: instances.Const{instances.BOT, instances.BOT, instances.ZERO},
		expCode: 20,
	},
	{
		name: "greater BOT nBOT",
		val1: instances.Const{instances.ONE, instances.ONE, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		expCode: 20,
	},
}
