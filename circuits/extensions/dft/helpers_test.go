package dft

import (
	"goexpdt/base"
	"goexpdt/trees"
)

const DIM = 3

var tests = []struct {
	name string
	val base.Const
	expCode int
}{
	{
		name: "(0,0,0)",
		val: base.Const{base.ZERO, base.ZERO, base.ZERO},
		expCode: 10,
	},
	{
		name: "(1,0,0)",
		val: base.Const{base.ONE, base.ZERO, base.ZERO},
		expCode: 10,
	},
	{
		name: "(1,0,1)",
		val: base.Const{base.ONE, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name: "(_,1,_)",
		val: base.Const{base.BOT, base.ONE, base.BOT},
		expCode: 10,
	},
	{
		name: "(_,0,_)",
		val: base.Const{base.BOT, base.ZERO, base.BOT},
		expCode: 10,
	},
	{
		name: "(_,0,1)",
		val: base.Const{base.BOT, base.ZERO, base.ONE},
		expCode: 10,
	},
	{
		name: "(1,0,_)",
		val: base.Const{base.ONE, base.ZERO, base.BOT},
		expCode: 10,
	},
	{
		name: "(0,_,_)",
		val: base.Const{base.ZERO, base.BOT, base.BOT},
		expCode: 20,
	},
	{
		name: "(0,_,1)",
		val: base.Const{base.ZERO, base.BOT, base.ONE},
		expCode: 20,
	},
	{
		name: "(_,_,1)",
		val: base.Const{base.BOT, base.BOT, base.ONE},
		expCode: 20,
	},
}

func genTree() *trees.Tree {
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
    // FULL, _,*,_, _,*,* and *,*,_ are DFT

    leaf1 := &trees.Node{ID: 7, Value: false}
    leaf2 := &trees.Node{ID: 8, Value: false}
    leaf3 := &trees.Node{ID: 9, Value: true}
    leaf4 := &trees.Node{ID: 10, Value: true}
    leaf5 := &trees.Node{ID: 11, Value: false}
    leaf6 := &trees.Node{ID: 12, Value: true}
    leaf7 := &trees.Node{ID: 13, Value: false}
    leaf8 := &trees.Node{ID: 14, Value: true}
    node6 := &trees.Node{ID: 6, Feat: 1, LChild: leaf7, RChild: leaf8}
    node5 := &trees.Node{ID: 5, Feat: 1, LChild: leaf5, RChild: leaf6}
    node4 := &trees.Node{ID: 4, Feat: 2, LChild: leaf3, RChild: leaf4}
    node3 := &trees.Node{ID: 3, Feat: 2, LChild: leaf1, RChild: leaf2}
    node2 := &trees.Node{ID: 2, Feat: 2, LChild: node5, RChild: node6}
    node1 := &trees.Node{ID: 1, Feat: 1, LChild: node3, RChild: node4}
    root := &trees.Node{ID: 0, Feat: 0, LChild: node1, RChild: node2}

    node1.Parent = root
    node2.Parent = root
    node3.Parent = node1
    node4.Parent = node1
    node5.Parent = node2
    node6.Parent = node2
    leaf1.Parent = node3
    leaf2.Parent = node3
    leaf3.Parent = node4
    leaf4.Parent = node4
    leaf5.Parent = node5
    leaf6.Parent = node5
    leaf7.Parent = node6
    leaf8.Parent = node6

    return &trees.Tree{
		Root: root,
		NodeCount: 15,
		FeatCount: 3,
		PosLeafs: []*trees.Node{leaf3, leaf4, leaf6, leaf8},
		NegLeafs: []*trees.Node{leaf1, leaf2, leaf5, leaf7},
	}
}
