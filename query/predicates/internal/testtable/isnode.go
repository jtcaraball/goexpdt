package testtable

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
)

var IsNodePTT = []OTRecord{
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

var IsNodeNTT = []OTRecord{
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
	// leaf1 := &trees.Node{ID: 5, Value: true}
	// leaf2 := &trees.Node{ID: 6, Value: false}
	// leaf3 := &trees.Node{ID: 7, Value: true}
	// leaf4 := &trees.Node{ID: 8, Value: false}
	// leaf5 := &trees.Node{ID: 9, Value: true}
	// leaf6 := &trees.Node{ID: 10, Value: true}
	// node4 := &trees.Node{ID: 4, Feat: 1, LChild: leaf5, RChild: leaf6}
	// node3 := &trees.Node{ID: 3, Feat: 2, LChild: leaf1, RChild: leaf2}
	// node2 := &trees.Node{ID: 2, Feat: 2, LChild: leaf4, RChild: node4}
	// node1 := &trees.Node{ID: 1, Feat: 1, LChild: node3, RChild: leaf3}
	// root := &trees.Node{ID: 0, Feat: 0, LChild: node1, RChild: node2}
	// return &trees.Tree{
	// 	Root:      root,
	// 	NodeCount: 11,
	// 	FeatCount: 3,
	// 	PosLeafs:  []*trees.Node{leaf1, leaf3, leaf5, leaf6},
	// 	NegLeafs:  []*trees.Node{leaf2, leaf4},
	// }
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
