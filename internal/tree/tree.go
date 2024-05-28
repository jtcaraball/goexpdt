package tree

import (
	"fmt"
	"os"

	"github.com/jtcaraball/goexpdt/query"
)

// Node represents a node in a decision tree. Holds both the information
// for leafs and internal nodes.
type node struct {
	id        uint32
	feat      uint32
	value     bool
	zeroChild *node
	oneChild  *node
	parent    *node
}

// Tree represents a decision tree.
type Tree struct {
	root          *node
	nodeCount     int
	featCount     uint32
	nodes         []query.Node
	nodeConsts    []query.Const
	posLeafConsts []query.Const
	negLeafConsts []query.Const
}

// Load returns the tree encoded in a json files passed by path.
func Load(path string) (Tree, error) {
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return Tree{}, err
	}
	treeJSON, err := unmarhsalTree(jsonBytes)
	if err != nil {
		return Tree{}, err
	}
	tree := Tree{}
	if err = tree.populateTree(treeJSON); err != nil {
		return tree, err
	}
	return tree, nil
}

type visitElem struct {
	ID       int
	Parent   *node
	OneChild bool
}

func (t *Tree) populateTree(treeJSON *treeJSON) error {
	t.featCount = uint32(len(treeJSON.Features))
	t.nodeCount = len(treeJSON.Nodes)

	toVisit := []visitElem{{ID: 0}}

	for len(toVisit) > 0 {
		nInfo := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		nodeJSON := treeJSON.Nodes[nInfo.ID]
		if nodeJSON == nil {
			return fmt.Errorf(
				"Tree parsing error: node with id '%d' does not exist",
				nInfo.ID,
			)
		}

		node := &node{id: uint32(nInfo.ID), parent: nInfo.Parent}

		if nInfo.Parent == nil {
			t.root = node
		} else if nInfo.OneChild {
			node.parent.oneChild = node
		} else {
			node.parent.zeroChild = node
		}

		if nodeJSON.Type == "leaf" {
			node.value = nodeJSON.Class == treeJSON.Positive
			continue
		}

		node.feat = uint32(nodeJSON.FeatIdx)

		toVisit = append(
			toVisit,
			visitElem{ID: int(nodeJSON.LeftID), Parent: node, OneChild: false},
			visitElem{ID: int(nodeJSON.RightID), Parent: node, OneChild: true},
		)
	}

	return nil
}

// Nodes returns a slices of the query.Node(s) that compose the tree. Returns
// an empty slice if t is nil.
func (t *Tree) Nodes() []query.Node {
	if t.nodes != nil {
		return t.nodes
	}
	nodes := make([]query.Node, t.nodeCount)
	t.root.appendSubtree(&nodes)
	return nodes
}

func (n node) appendSubtree(nodes *[]query.Node) {
	(*nodes)[n.id] = query.Node{
		Value:  n.value,
		Feat:   uint(n.feat),
		ZChild: -1,
		OChild: -1,
	}

	if n.zeroChild != nil {
		(*nodes)[n.id].ZChild = int(n.zeroChild.id)
		n.zeroChild.appendSubtree(nodes)
	}
	if n.oneChild != nil {
		(*nodes)[n.id].OChild = int(n.oneChild.id)
		n.oneChild.appendSubtree(nodes)
	}
}

// Dim returns the number of features the tree could decide on. Return 0 if t
// is nil.
func (t *Tree) Dim() uint {
	return uint(t.featCount)
}

// nodeElem is an auxiliary structure for iterating over nodes when generating
// their query.Const representation.
type nodeElem struct {
	n *node         // Node
	v []query.FeatV // Node's const value representation
}

// NodeConsts returns a slice of query.Const representing the nodes that
// compose the tree. Returns an error if t is nill or there is an underlying
// error with the tree.
func (t *Tree) NodesConsts() []query.Const {
	if t == nil {
		return nil
	}

	if t.nodeConsts != nil {
		return t.nodeConsts
	}

	var (
		n         *node
		v, zv, ov []query.FeatV
		ninfo     nodeElem
	)

	next := 0
	nconsts := make([]query.Const, t.nodeCount)
	nstack := []nodeElem{{n: t.root, v: make([]query.FeatV, t.featCount)}}

	for len(nstack) > 0 {
		ninfo, nstack = nstack[len(nstack)-1], nstack[:len(nstack)-1]
		n, v = ninfo.n, ninfo.v

		next += 1
		nconsts[next] = query.Const{Val: v}

		if n.zeroChild == nil || n.oneChild == nil {
			continue
		}

		zv = make([]query.FeatV, t.featCount)
		ov = make([]query.FeatV, t.featCount)
		copy(zv, v)
		copy(ov, v)
		zv[n.feat] = query.ZERO
		ov[n.feat] = query.ONE

		nstack = append(nstack, nodeElem{n, zv}, nodeElem{n, zv})
	}

	t.nodeConsts = nconsts
	return t.nodeConsts
}

// PosLeafConsts returns a slice of query.Const representing the tree's
// positive leafs.
func (t *Tree) PosLeafsConsts() []query.Const {
	if t == nil {
		return nil
	}

	if t.posLeafConsts != nil {
		return t.posLeafConsts
	}

	t.computeLeafs()

	return t.posLeafConsts
}

// NegLeafConsts returns a slice of query.Const representing the tree's
// negative leafs.
func (t *Tree) NegLeafsConsts() []query.Const {
	if t == nil {
		return nil
	}

	if t.negLeafConsts != nil {
		return t.negLeafConsts
	}

	t.computeLeafs()

	return t.negLeafConsts
}

// computeLeafs sets posLeafConsts and negLeafConts to a slice of query.Const
// representing the tree's positive and legative leaf respective. Does not
// check if t is nil.
func (t *Tree) computeLeafs() {
	var (
		n         *node
		v, zv, ov []query.FeatV
		ninfo     nodeElem
	)

	pconsts := []query.Const{}
	nconsts := []query.Const{}
	nstack := []nodeElem{{n: t.root, v: make([]query.FeatV, t.featCount)}}

	for len(nstack) > 0 {
		ninfo, nstack = nstack[len(nstack)-1], nstack[:len(nstack)-1]
		n, v = ninfo.n, ninfo.v

		if n.zeroChild == nil || n.oneChild == nil {
			if n.value {
				pconsts = append(pconsts, query.Const{Val: ninfo.v})
				continue
			}
			nconsts = append(nconsts, query.Const{Val: ninfo.v})
			continue
		}

		zv = make([]query.FeatV, t.featCount)
		ov = make([]query.FeatV, t.featCount)
		copy(zv, v)
		copy(ov, v)
		zv[n.feat] = query.ZERO
		ov[n.feat] = query.ONE

		nstack = append(nstack, nodeElem{n, zv}, nodeElem{n, zv})
	}

	t.posLeafConsts = pconsts
	t.negLeafConsts = nconsts
}