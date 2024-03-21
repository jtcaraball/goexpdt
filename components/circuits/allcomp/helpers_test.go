package allcomp

import (
	"testing"
	"stratifoiled/trees"
	"stratifoiled/sfdtest"
	"stratifoiled/components"
)

const DIM = 3

var allPosTests = []struct {
	name string
	val components.Const
	expCode int
}{
	{
		name: "(_,_,_)",
		val: components.Const{components.BOT, components.BOT, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,_,_)",
		val: components.Const{components.ZERO, components.BOT, components.BOT},
		expCode: 10,
	},
	{
		name: "(1,_,_)",
		val: components.Const{components.ONE, components.BOT, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,0,_)",
		val: components.Const{components.ZERO, components.ZERO, components.BOT},
		expCode: 10,
	},
	{
		name: "(0,1,_)",
		val: components.Const{components.ZERO, components.ONE, components.BOT},
		expCode: 10,
	},
	{
		name: "(1,0,_)",
		val: components.Const{components.ONE, components.ZERO, components.BOT},
		expCode: 20,
	},
	{
		name: "(1,1,_)",
		val: components.Const{components.ONE, components.ONE, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,0,0)",
		val: components.Const{components.ZERO, components.ZERO, components.ZERO},
		expCode: 10,
	},
	{
		name: "(0,0,1)",
		val: components.Const{components.ZERO, components.ZERO, components.ONE},
		expCode: 10,
	},
	{
		name: "(0,1,0)",
		val: components.Const{components.ZERO, components.ONE, components.ZERO},
		expCode: 10,
	},
	{
		name: "(0,1,1)",
		val: components.Const{components.ZERO, components.ONE, components.ONE},
		expCode: 10,
	},
	{
		name: "(1,0,0)",
		val: components.Const{components.ONE, components.ZERO, components.ZERO},
		expCode: 20,
	},
	{
		name: "(1,0,1)",
		val: components.Const{components.ONE, components.ZERO, components.ONE},
		expCode: 10,
	},
	{
		name: "(1,1,0)",
		val: components.Const{components.ONE, components.ONE, components.ZERO},
		expCode: 20,
	},
	{
		name: "(1,1,1)",
		val: components.Const{components.ONE, components.ONE, components.ONE},
		expCode: 20,
	},
	{
		name: "(_,0,_)",
		val: components.Const{components.BOT, components.ZERO, components.BOT},
		expCode: 20,
	},
	{
		name: "(_,1,_)",
		val: components.Const{components.BOT, components.ONE, components.BOT},
		expCode: 20,
	},
	{
		name: "(_,0,0)",
		val: components.Const{components.BOT, components.ZERO, components.ZERO},
		expCode: 20,
	},
	{
		name: "(_,0,1)",
		val: components.Const{components.BOT, components.ZERO, components.ONE},
		expCode: 10,
	},
	{
		name: "(_,1,0)",
		val: components.Const{components.BOT, components.ONE, components.ZERO},
		expCode: 20,
	},
	{
		name: "(_,1,1)",
		val: components.Const{components.BOT, components.ONE, components.ONE},
		expCode: 20,
	},
	{
		name: "(_,_,0)",
		val: components.Const{components.BOT, components.BOT, components.ZERO},
		expCode: 20,
	},
	{
		name: "(_,_,1)",
		val: components.Const{components.BOT, components.BOT, components.ONE},
		expCode: 20,
	},
	{
		name: "(0,_,0)",
		val: components.Const{components.ZERO, components.BOT, components.ZERO},
		expCode: 10,
	},
	{
		name: "(0,_,1)",
		val: components.Const{components.ZERO, components.BOT, components.ONE},
		expCode: 10,
	},
	{
		name: "(1,_,0)",
		val: components.Const{components.ONE, components.BOT, components.ZERO},
		expCode: 20,
	},
	{
		name: "(1,_,1)",
		val: components.Const{components.ONE, components.BOT, components.ONE},
		expCode: 20,
	},
}

var allNegTests = []struct{
	name string
	val components.Const
	expCode int
}{
	{
		name: "(_,_,_)",
		val: components.Const{components.BOT, components.BOT, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,_,_)",
		val: components.Const{components.ZERO, components.BOT, components.BOT},
		expCode: 20,
	},
	{
		name: "(1,_,_)",
		val: components.Const{components.ONE, components.BOT, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,0,_)",
		val: components.Const{components.ZERO, components.ZERO, components.BOT},
		expCode: 20,
	},
	{
		name: "(0,1,_)",
		val: components.Const{components.ZERO, components.ONE, components.BOT},
		expCode: 20,
	},
	{
		name: "(1,0,_)",
		val: components.Const{components.ONE, components.ZERO, components.BOT},
		expCode: 20,
	},
	{
		name: "(1,1,_)",
		val: components.Const{components.ONE, components.ONE, components.BOT},
		expCode: 10,
	},
	{
		name: "(0,0,0)",
		val: components.Const{components.ZERO, components.ZERO, components.ZERO},
		expCode: 20,
	},
	{
		name: "(0,0,1)",
		val: components.Const{components.ZERO, components.ZERO, components.ONE},
		expCode: 20,
	},
	{
		name: "(0,1,0)",
		val: components.Const{components.ZERO, components.ONE, components.ZERO},
		expCode: 20,
	},
	{
		name: "(0,1,1)",
		val: components.Const{components.ZERO, components.ONE, components.ONE},
		expCode: 20,
	},
	{
		name: "(1,0,0)",
		val: components.Const{components.ONE, components.ZERO, components.ZERO},
		expCode: 10,
	},
	{
		name: "(1,0,1)",
		val: components.Const{components.ONE, components.ZERO, components.ONE},
		expCode: 20,
	},
	{
		name: "(1,1,0)",
		val: components.Const{components.ONE, components.ONE, components.ZERO},
		expCode: 10,
	},
	{
		name: "(1,1,1)",
		val: components.Const{components.ONE, components.ONE, components.ONE},
		expCode: 10,
	},
	{
		name: "(_,0,_)",
		val: components.Const{components.BOT, components.ZERO, components.BOT},
		expCode: 20,
	},
	{
		name: "(_,1,_)",
		val: components.Const{components.BOT, components.ONE, components.BOT},
		expCode: 20,
	},
	{
		name: "(_,0,0)",
		val: components.Const{components.BOT, components.ZERO, components.ZERO},
		expCode: 20,
	},
	{
		name: "(_,0,1)",
		val: components.Const{components.BOT, components.ZERO, components.ONE},
		expCode: 20,
	},
	{
		name: "(_,1,0)",
		val: components.Const{components.BOT, components.ONE, components.ZERO},
		expCode: 20,
	},
	{
		name: "(_,1,1)",
		val: components.Const{components.BOT, components.ONE, components.ONE},
		expCode: 20,
	},
	{
		name: "(_,_,0)",
		val: components.Const{components.BOT, components.BOT, components.ZERO},
		expCode: 20,
	},
	{
		name: "(_,_,1)",
		val: components.Const{components.BOT, components.BOT, components.ONE},
		expCode: 20,
	},
	{
		name: "(0,_,0)",
		val: components.Const{components.ZERO, components.BOT, components.ZERO},
		expCode: 20,
	},
	{
		name: "(0,_,1)",
		val: components.Const{components.ZERO, components.BOT, components.ONE},
		expCode: 20,
	},
	{
		name: "(1,_,0)",
		val: components.Const{components.ONE, components.BOT, components.ZERO},
		expCode: 10,
	},
	{
		name: "(1,_,1)",
		val: components.Const{components.ONE, components.BOT, components.ONE},
		expCode: 20,
	},
}

func genTree() *trees.Tree {
    // root: _,_,_
    //   node_1: 0,_,_
    //       node_3: 0,0,_
    //           leaf_1: 0,0,0 (True)
    //           leaf_2: 0,0,1 (True)
    //       leaf_3: 0,1,_ (True)
    //   node_2: 1,_,_
    //       leaf_4: 1,_,0 (False)
    //       node_4: 1,_,1
    //           leaf_5: 1,0,1 (True)
    //           leaf_6: 1,1,1 (False)
	leaf1 := &trees.Node{ID: 5, Value: true}
	leaf2 := &trees.Node{ID: 6, Value: true}
	leaf3 := &trees.Node{ID: 7, Value: true}
	leaf4 := &trees.Node{ID: 8, Value: false}
	leaf5 := &trees.Node{ID: 9, Value: true}
	leaf6 := &trees.Node{ID: 10, Value: false}
	node4 := &trees.Node{ID: 4, Feat: 1, LChild: leaf5, RChild: leaf6}
	node3 := &trees.Node{ID: 3, Feat: 2, LChild: leaf1, RChild: leaf2}
	node2 := &trees.Node{ID: 2, Feat: 2, LChild: leaf4, RChild: node4}
	node1 := &trees.Node{ID: 1, Feat: 1, LChild: node3, RChild: leaf3}
	root := &trees.Node{ID: 0, Feat: 0, LChild: node1, RChild: node2}
	return &trees.Tree{
		Root: root,
		NodeCount: 11,
		FeatCount: 3,
		PosLeafs: []*trees.Node{leaf1, leaf2, leaf3, leaf5},
		NegLeafs: []*trees.Node{leaf4, leaf6},
	}
}

func encodeAndRun(
	t *testing.T,
	formula components.Component,
	context *components.Context,
	filePath string,
	id, expCode int,
	simplify bool,
) {
	var err error
	if simplify {
		formula, err = formula.Simplified(context)
		if err != nil {
			t.Errorf("Formula simplification error. %s", err.Error())
			return
		}
	}
	cnf, err := formula.Encoding(context)
	if err != nil {
		t.Errorf("Formula encoding error. %s", err.Error())
		return
	}
	if err = cnf.ToFile(filePath); err != nil {
		t.Errorf("CNF writing error. %s", err.Error())
		return
	}
	sfdtest.RunFormulaTest(t, id, expCode, filePath)
}
