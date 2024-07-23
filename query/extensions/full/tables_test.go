package full_test

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
)

var FullPTT = []test.OTRecord{
	{
		Dim: 3,
		Name:    "(_,1,0)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim: 3,
		Name:    "(1,_,1)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim: 3,
		Name:    "(0,1,_)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.BOT},
		ExpCode: 20,
	},
	{
		Dim: 3,
		Name:    "(0,1,1)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.ONE},
		ExpCode: 10,
	},
}

var FullNTT = []test.OTRecord{
	{
		Dim: 3,
		Name:    "(_,1,0)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim: 3,
		Name:    "(1,_,1)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim: 3,
		Name:    "(0,1,_)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.BOT},
		ExpCode: 10,
	},
	{
		Dim: 3,
		Name:    "(0,1,1)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.ONE},
		ExpCode: 20,
	},
}
