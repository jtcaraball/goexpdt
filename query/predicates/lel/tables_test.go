package lel_test

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
)

var LELPTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_):(1,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,_):(1,1,0)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,_):(_,_,1)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.BOT, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,1):(0,1,1)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ZERO, query.ONE, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,0):(_,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		Val2:    []query.FeatV{query.BOT, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,0):(1,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,1):(_,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		ExpCode: 20,
	},
}

var LELNTT = []test.BTRecord{
	{
		Dim:     3,
		Name:    "(1,_,_):(1,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,_):(1,1,0)",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,_):(_,_,1)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.BOT, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,1):(0,1,1)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ZERO, query.ONE, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,0):(_,_,0)",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ZERO},
		Val2:    []query.FeatV{query.BOT, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,0):(1,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,1):(_,_,_)",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.BOT, query.BOT, query.BOT},
		ExpCode: 10,
	},
}
