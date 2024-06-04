package tree

import (
	"os"
	"slices"
	"testing"

	"github.com/jtcaraball/goexpdt/query"
)

var (
	b query.FeatV = query.BOT
	o query.FeatV = query.ONE
	z query.FeatV = query.ZERO
)

var test = struct {
	tBytes        []byte
	tree          tree
	nodes         []query.Node
	nodeConsts    []query.QConst
	posLeafConsts []query.QConst
	negLeafConsts []query.QConst
}{
	tBytes: []byte(`
{
	"class_names": ["pos", "neg"],
	"positive": "pos",
	"feature_names": [
		"ft1",
		"ft2",
		"ft3",
		"ft4",
		"ft5",
		"ft6",
		"ft7",
		"ft8",
		"ft9",
		"ft10"
	],
	"nodes": {
		"0": {
		  "id": 0,
		  "type": "internal",
		  "feature_name": "ft2",
		  "feature_index": 6,
		  "threshold": 0.5,
		  "id_left": 1,
		  "id_right": 2
		},
		"1": {
		  "id": 1,
		  "type": "internal",
		  "feature_name": "ft5",
		  "feature_index": 5,
		  "threshold": 0.5,
		  "id_left": 3,
		  "id_right": 4
		},
		"2": {
		  "id": 2,
		  "type": "leaf",
		  "class": "pos"
		},
		"3": {
		  "id": 3,
		  "type": "internal",
		  "feature_name": "ft3",
		  "feature_index": 3,
		  "threshold": 0.5,
		  "id_left": 5,
		  "id_right": 6
		},
		"5": {
		  "id": 5,
		  "type": "leaf",
		  "class": "pos"
		},
		"6": {
		  "id": 6,
		  "type": "leaf",
		  "class": "neg"
		},
		"4": {
		  "id": 4,
		  "type": "internal",
		  "feature_name": "ft7",
		  "feature_index": 7,
		  "threshold": 0.5,
		  "id_left": 7,
		  "id_right": 8
		},
		"7": {
		  "id": 7,
		  "type": "leaf",
		  "class": "neg"
		},
		"8": {
		  "id": 8,
		  "type": "internal",
		  "feature_name": "ft4",
		  "feature_index": 4,
		  "threshold": 0.5,
		  "id_left": 9,
		  "id_right": 10
		},
		"9": {
		  "id": 9,
		  "type": "leaf",
		  "class": "neg"
		},
		"10": {
		  "id": 10,
		  "type": "leaf",
		  "class": "neg"
		}
	}
}
	`),
	nodes: []query.Node{
		{Feat: 6, ZChild: 1, OChild: 2},
		{Feat: 5, ZChild: 3, OChild: 4},
		{Value: true, ZChild: -1, OChild: -1},
		{Feat: 3, ZChild: 5, OChild: 6},
		{Feat: 7, ZChild: 7, OChild: 8},
		{Value: true, ZChild: -1, OChild: -1},
		{Value: false, ZChild: -1, OChild: -1},
		{Value: false, ZChild: -1, OChild: -1},
		{Feat: 4, ZChild: 9, OChild: 10},
		{Value: false, ZChild: -1, OChild: -1},
		{Value: false, ZChild: -1, OChild: -1},
	},
	nodeConsts: []query.QConst{
		{Val: []query.FeatV{b, b, b, b, b, b, b, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, b, z, b, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, b, o, b, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, z, z, b, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, o, z, b, b, b, b}},
		{Val: []query.FeatV{b, b, z, b, z, z, b, b, b, b}},
		{Val: []query.FeatV{b, b, o, b, z, z, b, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, o, z, z, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, o, z, o, b, b, b}},
		{Val: []query.FeatV{b, b, b, z, o, z, o, b, b, b}},
		{Val: []query.FeatV{b, b, b, o, o, z, o, b, b, b}},
	},
	posLeafConsts: []query.QConst{
		{Val: []query.FeatV{b, b, b, b, b, o, b, b, b, b}},
		{Val: []query.FeatV{b, b, z, b, z, z, b, b, b, b}},
	},
	negLeafConsts: []query.QConst{
		{Val: []query.FeatV{b, b, o, b, z, z, b, b, b, b}},
		{Val: []query.FeatV{b, b, b, b, o, z, z, b, b, b}},
		{Val: []query.FeatV{b, b, b, z, o, z, o, b, b, b}},
		{Val: []query.FeatV{b, b, b, o, o, z, o, b, b, b}},
	},
}

func writeNewTree(t *testing.T, treeJSONBytes []byte) (string, error) {
	t.Helper()
	f, err := os.CreateTemp("", "tree_tmp")
	if err != nil {
		return "", err
	}
	if _, err = f.Write(treeJSONBytes); err != nil {
		return "", err
	}
	defer f.Close()
	t.Cleanup(func() {
		os.Remove(f.Name())
	})
	return f.Name(), nil
}

func TestMain(m *testing.M) {
	n10 := node{id: 10, value: false}
	n9 := node{id: 9, value: false}
	n8 := node{id: 8, feat: 4, zeroChild: &n9, oneChild: &n10}
	n7 := node{id: 7, value: false}
	n6 := node{id: 6, value: false}
	n5 := node{id: 5, value: true}
	n4 := node{id: 4, feat: 7, zeroChild: &n7, oneChild: &n8}
	n3 := node{id: 3, feat: 3, zeroChild: &n5, oneChild: &n6}
	n2 := node{id: 2, value: true}
	n1 := node{id: 1, feat: 5, zeroChild: &n3, oneChild: &n4}
	test.tree.root = &node{id: 0, feat: 6, zeroChild: &n1, oneChild: &n2}
	test.tree.nodeCount = 10
	test.tree.featCount = 10
}

func TestLoad_Nodes(t *testing.T) {
	path, err := writeNewTree(t, test.tBytes)
	if err != nil {
		t.Fatalf("Failed to write tree file: %s", err.Error())
	}
	tTree, err := Load(path)
	if err != nil {
		t.Fatalf("Failed to load tree: %s", err.Error())
	}
	if !slices.Equal(test.nodes, tTree.Nodes()) {
		t.Fatalf(
			"Trees not equal.\nExpected %v\nbut got  %v",
			test.nodes,
			tTree.Nodes(),
		)
	}
}

func Test_NodeConsts(t *testing.T) {
	nc := []string{}
	for _, c := range test.tree.NodesConsts() {
		nc = append(nc, c.AsString())
	}
	slices.Sort(nc)

	expnc := []string{}
	for _, c := range test.nodeConsts {
		expnc = append(expnc, c.AsString())
	}

	if !slices.Equal(nc, expnc) {
		t.Errorf("Nodes not equal.\nExpected %s.\nbut got  %s", expnc, nc)
	}
}

func Test_PosLeafConsts(t *testing.T) {
	nc := []string{}
	for _, c := range test.tree.PosLeafsConsts() {
		nc = append(nc, c.AsString())
	}
	slices.Sort(nc)

	expnc := []string{}
	for _, c := range test.posLeafConsts {
		expnc = append(expnc, c.AsString())
	}

	if !slices.Equal(nc, expnc) {
		t.Errorf("Pos leafs not equal.\nExpected %s.\nbut got  %s", expnc, nc)
	}
}

func Test_NegLeafConsts(t *testing.T) {
	nc := []string{}
	for _, c := range test.tree.NegLeafsConsts() {
		nc = append(nc, c.AsString())
	}
	slices.Sort(nc)

	expnc := []string{}
	for _, c := range test.negLeafConsts {
		expnc = append(expnc, c.AsString())
	}

	if !slices.Equal(nc, expnc) {
		t.Errorf("Neg leafs not equal.\nExpected %s.\nbut got  %s", expnc, nc)
	}
}
