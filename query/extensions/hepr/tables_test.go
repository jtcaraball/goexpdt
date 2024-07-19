package hepr_test

import "github.com/jtcaraball/goexpdt/query"

type HEPRRecord struct {
	Dim     int
	Name    string
	Val1    []query.FeatV
	Val2    []query.FeatV
	Pref    []int
	ExpCode int
}

var HEPRPTT = []HEPRRecord{
	{
		Dim:     4,
		Name:    "(_,_,_,_):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(_,_,_,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(_,1,_,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.BOT, query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(1,1,_,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(1,1,1,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(_,_,1,1):(_,0,0,_):[0,2,3]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.BOT, query.ZERO, query.ZERO, query.BOT},
		Pref:    []int{0, 2, 3},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(_,_,0,1):(_,0,0,_):[0,1,3]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.BOT, query.ZERO, query.ZERO, query.BOT},
		Pref:    []int{0, 1, 3},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(1,1,0,1):(1,0,0,1):[0,1,3,2]",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ZERO, query.ONE},
		Pref:    []int{0, 1, 3, 2},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(1,_,0,1):(1,0,0,_):[0,3,2,1]",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ZERO, query.BOT},
		Pref:    []int{0, 3, 2, 1},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(_,_,0,1):(1,0,_,_):[0,3,2,1]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.BOT, query.BOT},
		Pref:    []int{0, 3, 2, 1},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(0,1,_,1):(1,_,0,1):[3,1,0]",
		Val1:    []query.FeatV{query.ZERO, query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO, query.ONE},
		Pref:    []int{3, 1, 0},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(0,_,1,1):(1,0,_,1):[3,0,1]",
		Val1:    []query.FeatV{query.ZERO, query.BOT, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.BOT, query.ONE},
		Pref:    []int{3, 0, 1},
		ExpCode: 20,
	},
}

var HEPRNTT = []HEPRRecord{
	{
		Dim:     4,
		Name:    "(_,_,_,_):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT, query.BOT},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(_,_,_,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(_,1,_,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.BOT, query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(1,1,_,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(1,1,1,1):(0,0,0,0):[0]",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.ZERO, query.ZERO, query.ZERO, query.ZERO},
		Pref:    []int{0},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(_,_,1,1):(_,0,0,_):[0,2,3]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.BOT, query.ZERO, query.ZERO, query.BOT},
		Pref:    []int{0, 2, 3},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(_,_,0,1):(_,0,0,_):[0,1,3]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.BOT, query.ZERO, query.ZERO, query.BOT},
		Pref:    []int{0, 1, 3},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(1,1,0,1):(1,0,0,1):[0,1,3,2]",
		Val1:    []query.FeatV{query.ONE, query.ONE, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ZERO, query.ONE},
		Pref:    []int{0, 1, 3, 2},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(1,_,0,1):(1,0,0,_):[0,3,2,1]",
		Val1:    []query.FeatV{query.ONE, query.BOT, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.ZERO, query.BOT},
		Pref:    []int{0, 3, 2, 1},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(_,_,0,1):(1,0,_,_):[0,3,2,1]",
		Val1:    []query.FeatV{query.BOT, query.BOT, query.ZERO, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.BOT, query.BOT},
		Pref:    []int{0, 3, 2, 1},
		ExpCode: 10,
	},
	{
		Dim:     4,
		Name:    "(0,1,_,1):(1,_,0,1):[3,1,0]",
		Val1:    []query.FeatV{query.ZERO, query.ONE, query.BOT, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.BOT, query.ZERO, query.ONE},
		Pref:    []int{3, 1, 0},
		ExpCode: 20,
	},
	{
		Dim:     4,
		Name:    "(0,_,1,1):(1,0,_,1):[3,0,1]",
		Val1:    []query.FeatV{query.ZERO, query.BOT, query.ONE, query.ONE},
		Val2:    []query.FeatV{query.ONE, query.ZERO, query.BOT, query.ONE},
		Pref:    []int{3, 0, 1},
		ExpCode: 10,
	},
}
