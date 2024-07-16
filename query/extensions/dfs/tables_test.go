package dfs_test

import (
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/internal/test"
)

var DFSPTT = []test.OTRecord{
	{
		Dim:     3,
		Name:    "(0,0,0)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.ZERO},
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
		Name:    "(1,0,1)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.ONE},
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
		Name:    "(_,0,_)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.BOT},
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
		Name:    "(1,0,_)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,_,_)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,_,1)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.ONE},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(_,_,1)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.ONE},
		ExpCode: 20,
	},
}

var DFSNTT = []test.OTRecord{
	{
		Dim:     3,
		Name:    "(0,0,0)",
		Val:     []query.FeatV{query.ZERO, query.ZERO, query.ZERO},
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
		Name:    "(1,0,1)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.ONE},
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
		Name:    "(_,0,_)",
		Val:     []query.FeatV{query.BOT, query.ZERO, query.BOT},
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
		Name:    "(1,0,_)",
		Val:     []query.FeatV{query.ONE, query.ZERO, query.BOT},
		ExpCode: 20,
	},
	{
		Dim:     3,
		Name:    "(0,_,_)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.BOT},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(0,_,1)",
		Val:     []query.FeatV{query.ZERO, query.BOT, query.ONE},
		ExpCode: 10,
	},
	{
		Dim:     3,
		Name:    "(_,_,1)",
		Val:     []query.FeatV{query.BOT, query.BOT, query.ONE},
		ExpCode: 10,
	},
}

func DFSTree() query.Model {
	// In tree:
	// root: _,_,_
	//   node_1: 0,_,_
	//       node_3: 0,0,_
	//           leaf_1: 0,0,0 (False)
	//           leaf_2: 0,0,1 (False)
	//       node_4: 0,1,_
	//           leaf_3: 0,1,0 (True)
	//           leaf_4: 0,1,1 (True)
	//   node_2: 1,_,_
	//       node_5: 1,_,0
	//           leaf_5: 1,0,0 (False)
	//           leaf_6: 1,1,0 (True)
	//       node_6: 1,_,1
	//           leaf_7: 1,0,1 (False)
	//           leaf_8: 1,1,1 (True)
	// FULL, _,*,_, _,*,* and *,*,_ are DFS

	// leaf1 := &trees.Node{ID: 7, Value: false}
	// leaf2 := &trees.Node{ID: 8, Value: false}
	// leaf3 := &trees.Node{ID: 9, Value: true}
	// leaf4 := &trees.Node{ID: 10, Value: true}
	// leaf5 := &trees.Node{ID: 11, Value: false}
	// leaf6 := &trees.Node{ID: 12, Value: true}
	// leaf7 := &trees.Node{ID: 13, Value: false}
	// leaf8 := &trees.Node{ID: 14, Value: true}
	// node6 := &trees.Node{ID: 6, Feat: 1, LChild: leaf7, RChild: leaf8}
	// node5 := &trees.Node{ID: 5, Feat: 1, LChild: leaf5, RChild: leaf6}
	// node4 := &trees.Node{ID: 4, Feat: 2, LChild: leaf3, RChild: leaf4}
	// node3 := &trees.Node{ID: 3, Feat: 2, LChild: leaf1, RChild: leaf2}
	// node2 := &trees.Node{ID: 2, Feat: 2, LChild: node5, RChild: node6}
	// node1 := &trees.Node{ID: 1, Feat: 1, LChild: node3, RChild: node4}
	// root := &trees.Node{ID: 0, Feat: 0, LChild: node1, RChild: node2}

	// node1.Parent = root
	// node2.Parent = root
	// node3.Parent = node1
	// node4.Parent = node1
	// node5.Parent = node2
	// node6.Parent = node2
	// leaf1.Parent = node3
	// leaf2.Parent = node3
	// leaf3.Parent = node4
	// leaf4.Parent = node4
	// leaf5.Parent = node5
	// leaf6.Parent = node5
	// leaf7.Parent = node6
	// leaf8.Parent = node6

	// return &trees.Tree{
	// 	Root:      root,
	// 	NodeCount: 15,
	// 	FeatCount: 3,
	// 	PosLeafs:  []*trees.Node{leaf3, leaf4, leaf6, leaf8},
	// 	NegLeafs:  []*trees.Node{leaf1, leaf2, leaf5, leaf7},
	// }

	nodes := []query.Node{
		{Feat: 0, ZChild: 1, OChild: 8},                              // root
		{Feat: 1, ZChild: 2, OChild: 5},                              // node1
		{Feat: 2, ZChild: 3, OChild: 4},                              // node3
		{Value: false, ZChild: query.NoChild, OChild: query.NoChild}, // leaf1
		{Value: false, ZChild: query.NoChild, OChild: query.NoChild}, // leaf2
		{Feat: 2, ZChild: 6, OChild: 7},                              // node4
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf3
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf4
		{Feat: 2, ZChild: 9, OChild: 12},                             // node2
		{Feat: 1, ZChild: 10, OChild: 11},                            // node5
		{Value: false, ZChild: query.NoChild, OChild: query.NoChild}, // leaf5
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf6
		{Feat: 1, ZChild: 13, OChild: 14},                            // node6
		{Value: false, ZChild: query.NoChild, OChild: query.NoChild}, // leaf7
		{Value: true, ZChild: query.NoChild, OChild: query.NoChild},  // leaf8
	}

	t, err := test.NewMockTree(3, nodes)
	if err != nil {
		panic(err)
	}

	return t
}
