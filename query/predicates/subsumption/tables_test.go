package subsumption_test

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
)

// SubsumptionPTT is the subsumption positive test cases table.
var SubsumptionPTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_):(1,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,_):(0,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,0,0):(1,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,_):(1,0,1)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,0):(1,_,1)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,1):(1,1,1)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,0,0):(1,0,1)",
		Val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,0,0):(_,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 20,
	},
}

// SubsumptionNTT is the subsumption negated test cases table.
var SubsumptionNTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_):(1,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,_):(0,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,0,0):(1,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,_):(1,0,1)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,0):(1,_,1)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,1):(1,1,1)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,0,0):(1,0,1)",
		Val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,0,0):(_,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 10,
	},
}
