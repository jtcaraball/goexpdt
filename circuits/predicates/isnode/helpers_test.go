package isnode

import (
	"goexpdt/base"
	"goexpdt/trees"
)

const DIM = 3

type testS struct {
	name    string
	val     base.Const
	expCode int
}

var tests = []testS{
	{
		name:    "(_,_,_)",
		val:     base.Const{base.BOT, base.BOT, base.BOT},
		expCode: 10,
	},
	{
		name:    "(0,_,_)",
		val:     base.Const{base.ZERO, base.BOT, base.BOT},
		expCode: 10,
	},
	{
		name:    "(1,_,_)",
		val:     base.Const{base.ONE, base.BOT, base.BOT},
		expCode: 10,
	},
	{
		name:    "(0,0,_)",
		val:     base.Const{base.ZERO, base.ZERO, base.BOT},
		expCode: 10,
	},
	{
		name:    "(1,_,1)",
		val:     base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 10,
	},
	{
		name:    "(0,0,0)",
		val:     base.Const{base.ZERO, base.ZERO, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(0,0,1)",
		val:     base.Const{base.ZERO, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(0,1,_)",
		val:     base.Const{base.ZERO, base.ONE, base.BOT},
		expCode: 10,
	},
	{
		name:    "(1,_,0)",
		val:     base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,0,1)",
		val:     base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,1,1)",
		val:     base.Const{base.ONE, base.ONE, base.ONE},
		expCode: 10,
	},
	{
		name:    "(0,1,0)",
		val:     base.Const{base.ZERO, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(0,1,1)",
		val:     base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,0,0)",
		val:     base.Const{base.ONE, base.ZERO, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,1,0)",
		val:     base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(_,_,0)",
		val:     base.Const{base.BOT, base.BOT, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(_,_,1)",
		val:     base.Const{base.BOT, base.BOT, base.ONE},
		expCode: 20,
	},
	{
		name:    "(_,0,0)",
		val:     base.Const{base.BOT, base.ZERO, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(_,1,0)",
		val:     base.Const{base.BOT, base.ONE, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(_,0,1)",
		val:     base.Const{base.BOT, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(_,1,1)",
		val:     base.Const{base.BOT, base.ONE, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,1,_)",
		val:     base.Const{base.ONE, base.ONE, base.BOT},
		expCode: 20,
	},
	{
		name:    "(1,0,_)",
		val:     base.Const{base.ONE, base.ZERO, base.BOT},
		expCode: 20,
	},
	{
		name:    "(_,0,_)",
		val:     base.Const{base.BOT, base.ZERO, base.BOT},
		expCode: 20,
	},
	{
		name:    "(_,1,_)",
		val:     base.Const{base.BOT, base.ONE, base.BOT},
		expCode: 20,
	},
	{
		name:    "(0,_,0)",
		val:     base.Const{base.ZERO, base.BOT, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(0,_,1)",
		val:     base.Const{base.ZERO, base.BOT, base.ONE},
		expCode: 20,
	},
}

var notTests = []testS{
	{
		name:    "(_,_,_)",
		val:     base.Const{base.BOT, base.BOT, base.BOT},
		expCode: 20,
	},
	{
		name:    "(0,_,_)",
		val:     base.Const{base.ZERO, base.BOT, base.BOT},
		expCode: 20,
	},
	{
		name:    "(1,_,_)",
		val:     base.Const{base.ONE, base.BOT, base.BOT},
		expCode: 20,
	},
	{
		name:    "(0,0,_)",
		val:     base.Const{base.ZERO, base.ZERO, base.BOT},
		expCode: 20,
	},
	{
		name:    "(1,_,1)",
		val:     base.Const{base.ONE, base.BOT, base.ONE},
		expCode: 20,
	},
	{
		name:    "(0,0,0)",
		val:     base.Const{base.ZERO, base.ZERO, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(0,0,1)",
		val:     base.Const{base.ZERO, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(0,1,_)",
		val:     base.Const{base.ZERO, base.ONE, base.BOT},
		expCode: 20,
	},
	{
		name:    "(1,_,0)",
		val:     base.Const{base.ONE, base.BOT, base.ZERO},
		expCode: 20,
	},
	{
		name:    "(1,0,1)",
		val:     base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 20,
	},
	{
		name:    "(1,1,1)",
		val:     base.Const{base.ONE, base.ONE, base.ONE},
		expCode: 20,
	},
	{
		name:    "(0,1,0)",
		val:     base.Const{base.ZERO, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(0,1,1)",
		val:     base.Const{base.ZERO, base.ONE, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,0,0)",
		val:     base.Const{base.ONE, base.ZERO, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(1,1,0)",
		val:     base.Const{base.ONE, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(_,_,0)",
		val:     base.Const{base.BOT, base.BOT, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(_,_,1)",
		val:     base.Const{base.BOT, base.BOT, base.ONE},
		expCode: 10,
	},
	{
		name:    "(_,0,0)",
		val:     base.Const{base.BOT, base.ZERO, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(_,1,0)",
		val:     base.Const{base.BOT, base.ONE, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(_,0,1)",
		val:     base.Const{base.BOT, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name:    "(_,1,1)",
		val:     base.Const{base.BOT, base.ONE, base.ONE},
		expCode: 10,
	},
	{
		name:    "(1,1,_)",
		val:     base.Const{base.ONE, base.ONE, base.BOT},
		expCode: 10,
	},
	{
		name:    "(1,0,_)",
		val:     base.Const{base.ONE, base.ZERO, base.BOT},
		expCode: 10,
	},
	{
		name:    "(_,0,_)",
		val:     base.Const{base.BOT, base.ZERO, base.BOT},
		expCode: 10,
	},
	{
		name:    "(_,1,_)",
		val:     base.Const{base.BOT, base.ONE, base.BOT},
		expCode: 10,
	},
	{
		name:    "(0,_,0)",
		val:     base.Const{base.ZERO, base.BOT, base.ZERO},
		expCode: 10,
	},
	{
		name:    "(0,_,1)",
		val:     base.Const{base.ZERO, base.BOT, base.ONE},
		expCode: 10,
	},
}

func genTree() *trees.Tree {
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
	leaf1 := &trees.Node{ID: 5, Value: true}
	leaf2 := &trees.Node{ID: 6, Value: false}
	leaf3 := &trees.Node{ID: 7, Value: true}
	leaf4 := &trees.Node{ID: 8, Value: false}
	leaf5 := &trees.Node{ID: 9, Value: true}
	leaf6 := &trees.Node{ID: 10, Value: true}
	node4 := &trees.Node{ID: 4, Feat: 1, LChild: leaf5, RChild: leaf6}
	node3 := &trees.Node{ID: 3, Feat: 2, LChild: leaf1, RChild: leaf2}
	node2 := &trees.Node{ID: 2, Feat: 2, LChild: leaf4, RChild: node4}
	node1 := &trees.Node{ID: 1, Feat: 1, LChild: node3, RChild: leaf3}
	root := &trees.Node{ID: 0, Feat: 0, LChild: node1, RChild: node2}
	return &trees.Tree{
		Root:      root,
		NodeCount: 11,
		FeatCount: 3,
		PosLeafs:  []*trees.Node{leaf1, leaf3, leaf5, leaf6},
		NegLeafs:  []*trees.Node{leaf2, leaf4},
	}
}
