package subsumption_test

import "github.com/jtcaraball/goexpdt/query"

const DIM = 3

type testS struct {
	name    string
	val1    []query.FeatV
	val2    []query.FeatV
	expCode int
}

var tests = []testS{
	{
		name:    "(1,_,_):(1,_,0)",
		val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,_,_):(0,_,0)",
		val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,0,0):(1,_,_)",
		val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		expCode: 20,
	},
	{
		name:    "(_,_,_):(1,0,1)",
		val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		expCode: 10,
	},
	{
		name:    "(_,_,0):(1,_,1)",
		val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		expCode: 20,
	},
	{
		name:    "(1,1,1):(1,1,1)",
		val1:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		val2:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		expCode: 10,
	},
	{
		name:    "(1,0,0):(1,0,1)",
		val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		expCode: 20,
	},
	{
		name:    "(1,0,0):(_,_,_)",
		val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		expCode: 20,
	},
}

var notTests = []testS{
	{
		name:    "(1,_,_):(1,_,0)",
		val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,_,_):(0,_,0)",
		val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,0,0):(1,_,_)",
		val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		expCode: 10,
	},
	{
		name:    "(_,_,_):(1,0,1)",
		val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		expCode: 20,
	},
	{
		name:    "(_,_,0):(1,_,1)",
		val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		expCode: 10,
	},
	{
		name:    "(1,1,1):(1,1,1)",
		val1:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		val2:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		expCode: 20,
	},
	{
		name:    "(1,0,0):(1,0,1)",
		val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		expCode: 10,
	},
	{
		name:    "(1,0,0):(_,_,_)",
		val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		expCode: 10,
	},
}
