package subsumption

import "stratifoiled/components/instances"

const (
	DIM = 3
	CNFPATH = "tmpCNF"
	SOLVER = "/kissat"
)

var tests = []struct {
	name string
	val1 instances.Const
	val2 instances.Const
	expCode int
}{
	{
		// x = Constant((Symbol.ONE, Symbol.BOT, Symbol.BOT))
		// y = Constant((Symbol.ONE, Symbol.BOT, Symbol.ZERO))
		// 10
		name: "BOT BOT True",
		val1: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ONE, instances.BOT, instances.ZERO},
		expCode: 10,
	},
	{
		// x = Constant((Symbol.ONE, Symbol.BOT, Symbol.BOT))
		// y = Constant((Symbol.ZERO, Symbol.BOT, Symbol.ZERO))
		// 20
		name: "BOT BOT False",
		val1: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ZERO, instances.BOT, instances.ZERO},
		expCode: 20,
	},
	{
		// x = Constant((Symbol.ONE, Symbol.ZERO, Symbol.ZERO))
		// y = Constant((Symbol.ONE, Symbol.BOT, Symbol.BOT))
		// 20
		name: "nBOT BOT False",
		val1: instances.Const{instances.ONE, instances.ZERO, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.BOT, instances.BOT},
		expCode: 20,
	},
	{
		// x = Constant((Symbol.BOT, Symbol.BOT, Symbol.BOT))
		// y = Constant((Symbol.ONE, Symbol.ZERO, Symbol.ONE))
		// 10
		name: "BOT nBOT True",
		val1: instances.Const{instances.BOT, instances.BOT, instances.BOT},
		val2: instances.Const{instances.ONE, instances.ZERO, instances.ONE},
		expCode: 10,
	},
	{
		// x = Constant((Symbol.BOT, Symbol.BOT, Symbol.ZERO))
		// y = Constant((Symbol.ONE, Symbol.BOT, Symbol.ONE))
		// 20
		name: "BOT nBOT False",
		val1: instances.Const{instances.BOT, instances.BOT, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.BOT, instances.ONE},
		expCode: 20,
	},
		// y = Constant((Symbol.ONE, Symbol.ONE, Symbol.ONE))
		// x = Constant((Symbol.ONE, Symbol.ONE, Symbol.ONE))
		// 10
	{
		name: "nBOT nBOT True",
		val1: instances.Const{instances.ONE, instances.ONE, instances.ONE},
		val2: instances.Const{instances.ONE, instances.ONE, instances.ONE},
		expCode: 10,
	},
	{
		// x = Constant((Symbol.ONE, Symbol.ZERO, Symbol.ZERO))
		// y = Constant((Symbol.ONE, Symbol.ZERO, Symbol.ONE))
		// 20
		name: "nBOT nBOT False",
		val1: instances.Const{instances.ONE, instances.ZERO, instances.ZERO},
		val2: instances.Const{instances.ONE, instances.ZERO, instances.ONE},
		expCode: 20,
	},
}
