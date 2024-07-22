package isnode_test

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
)

var IsNodePTT = []test.OTRecord{
	{
		Dim:     3,
		Name:    "(_,_,_)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,_,_)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,_)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,0,_)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,1)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,0,0)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,0,1)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,1,_)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,_,0)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,0,1)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,1)",
		Val:     []query.FeatV{query.ONE, query.ONE, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,1,0)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,1,1)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,0,0)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,0)",
		Val:     []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,0)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,1)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,0,0)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,1,0)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,0,1)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,1,1)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,_)",
		Val:     []query.FeatV{query.ONE, query.ONE, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,0,_)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,0,_)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,1,_)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,_,0)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,_,1)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.ONE},
		ExpCode: 20,
	},
}

var IsNodeNTT = []test.OTRecord{
	{
		Dim:     3,
		Name:    "(_,_,_)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,_,_)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,_)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,0,_)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,1)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,0,0)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,0,1)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,1,_)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,_,0)",
		Val:     []query.FeatV{query.ONE, query.BOT, query.ZERO},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,0,1)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(1,1,1)",
		Val:     []query.FeatV{query.ONE, query.ONE, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,1,0)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,1,1)",
		Val:     []query.FeatV{query.ZERO, query.ONE, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,0,0)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,0)",
		Val:     []query.FeatV{query.ONE, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,0)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,1)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,0,0)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,1,0)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,0,1)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,1,1)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,1,_)",
		Val:     []query.FeatV{query.ONE, query.ONE, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(1,0,_)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,0,_)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,1,_)",
		Val:     []query.FeatV{query.BOT, query.ONE, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,_,0)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.ZERO},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,_,1)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.ONE},
		ExpCode: 10,
	},
}

func IsNodeTree() query.Model {
	// root: _,_,_
	//   node_1: 0,_,_
	//       node_3: 0,0,_
	//           leaf_1: 0,0,0
	//           leaf_2: 0,0,1
	//       leaf_3: 0,1,_
	//   node_2: 1,_,_
	//       leaf_4: 1,_,0
	//       node_4: 1,_,1
	//           leaf_5: 1,0,1
	//           leaf_6: 1,1,1
	nodes := []query.Node{
		{Feat: 0, ZChild: 1, OChild: 6},                              // root
		{Feat: 1, ZChild: 2, OChild: 5},                              // node1
		{Feat: 2, ZChild: 3, OChild: 4},                              // node3
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf1
		{Value: false, ZChild: query.NoChild, OChild: query.NoChild}, // leaf2
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf3
		{Feat: 2, ZChild: 7, OChild: 8},                              // node2
		{Value: false, ZChild: query.NoChild, OChild: query.NoChild}, // leaf4
		{Feat: 1, ZChild: 9, OChild: 10},                             // node4
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf5
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf6
	}

	t, err := test.NewMockTree(3, nodes)
	if err != nil {
		panic(err)
	}

	return t
}
