// A basic implementation of a Decision tree that implements the Model
// interface using a custom tree json encoding based on scikit-learn
// DecisiontreeClassifier encoding.
package tree

import (
	"fmt"
	"os"

	"github.com/jtcaraball/goexpdt/query"
)

// Node represents a node in a decision tree. Holds both the information
// for leafs and internal nodes.
type node struct {
	id        int
	feat      int
	value     bool
	zeroChild *node
	oneChild  *node
	parent    *node
}

// tree represents a decision tree.
type tree struct {
	root          *node
	nodeCount     int
	featCount     int
	nodes         []query.Node
	nodeConsts    []query.QConst
	posLeafConsts []query.QConst
	negLeafConsts []query.QConst
}

// Load returns the tree encoded in a json files passed by path.
func Load(path string) (tree, error) {
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return tree{}, err
	}
	treeJSON, err := unmarhsalTree(jsonBytes)
	if err != nil {
		return tree{}, err
	}
	tree := tree{}
	if err = tree.populatetree(treeJSON); err != nil {
		return tree, err
	}
	return tree, nil
}

type visitElem struct {
	ID       int
	Parent   *node
	OneChild bool
}

func (t *tree) populatetree(treeJSON *treeJSON) error {
	t.featCount = len(treeJSON.Features)
	t.nodeCount = len(treeJSON.Nodes)

	toVisit := []visitElem{{ID: 0}}

	for len(toVisit) > 0 {
		nInfo := toVisit[len(toVisit)-1]
		toVisit = toVisit[:len(toVisit)-1]

		nodeJSON := treeJSON.Nodes[nInfo.ID]
		if nodeJSON == nil {
			return fmt.Errorf(
				"tree parsing error: node with id '%d' does not exist",
				nInfo.ID,
			)
		}

		node := &node{id: nInfo.ID, parent: nInfo.Parent}

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

		node.feat = nodeJSON.FeatIdx

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
func (t *tree) Nodes() []query.Node {
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
		Feat:   n.feat,
		ZChild: query.NoChild,
		OChild: query.NoChild,
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
func (t *tree) Dim() int {
	if t == nil {
		return 0
	}
	return t.featCount
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
func (t *tree) NodesConsts() []query.QConst {
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
	nconsts := make([]query.QConst, t.nodeCount)
	nstack := []nodeElem{{n: t.root, v: make([]query.FeatV, t.featCount)}}

	for len(nstack) > 0 {
		ninfo, nstack = nstack[len(nstack)-1], nstack[:len(nstack)-1]
		n, v = ninfo.n, ninfo.v

		next += 1
		nconsts[next] = query.QConst{Val: v}

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
func (t *tree) PosLeafsConsts() []query.QConst {
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
func (t *tree) NegLeafsConsts() []query.QConst {
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
func (t *tree) computeLeafs() {
	var (
		n         *node
		v, zv, ov []query.FeatV
		ninfo     nodeElem
	)

	pconsts := []query.QConst{}
	nconsts := []query.QConst{}
	nstack := []nodeElem{{n: t.root, v: make([]query.FeatV, t.featCount)}}

	for len(nstack) > 0 {
		ninfo, nstack = nstack[len(nstack)-1], nstack[:len(nstack)-1]
		n, v = ninfo.n, ninfo.v

		if n.zeroChild == nil || n.oneChild == nil {
			if n.value {
				pconsts = append(pconsts, query.QConst{Val: ninfo.v})
				continue
			}
			nconsts = append(nconsts, query.QConst{Val: ninfo.v})
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
