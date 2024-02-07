package trees

import (
	"errors"
	"fmt"
	"os"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type Node struct {
	ID int
	Feat int
	Value bool
	LChild *Node
	RChild *Node
}

type Tree struct {
	Root *Node
	NodeCount int
	FeatCount int
	PosLeafs []*Node
	NegLeafs []*Node
}

type VisitElem struct {
	ID int
	Parent *Node
	Right bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Return tree from json file.
func LoadTree(path string) (*Tree, error) {
	jsonBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	treeJSON, err := unmarhsalTree(jsonBytes)
	if err != nil {
		return nil, err
	}
	tree := &Tree{}
	if err = tree.populateTree(treeJSON); err != nil {
		return nil, err
	}
	return tree, nil
}

// Populate tree from tree info parsed from json.
func (t *Tree) populateTree(treeJSON *TreeInJSON) error {
	t.FeatCount = int(len(treeJSON.Features))
	t.NodeCount = int(len(treeJSON.Nodes))

	toVisit := []VisitElem{{ID: 0}}
	for len(toVisit) > 0 {
		nInfo := toVisit[len(toVisit) - 1]
		toVisit = toVisit[:len(toVisit) - 1]

		nodeJSON := treeJSON.Nodes[nInfo.ID]
		if nodeJSON == nil {
			return errors.New(
				fmt.Sprintf(
					"Tree parsing error: node with id '%d' does not exist",
					nInfo.ID,
				),
			)
		}

		node := &Node{ID: nInfo.ID}
		if nInfo.Parent == nil {
			t.Root = node
		} else if nInfo.Right {
			nInfo.Parent.RChild = node
		} else {
			nInfo.Parent.LChild = node
		}

		if nodeJSON.Type == "leaf" {
			node.Value = nodeJSON.Class == treeJSON.Positive
			if node.Value {
				t.PosLeafs = append(t.PosLeafs, node)
			} else {
				t.NegLeafs = append(t.NegLeafs, node)
			}
			continue
		}

		node.Feat = nodeJSON.FeatIdx
		toVisit = append(
			toVisit,
			VisitElem{ID: int(nodeJSON.LeftID), Parent: node, Right: false},
			VisitElem{ID: int(nodeJSON.RightID), Parent: node, Right: true},
		)
	}

	return nil
}

// Return slice of nodes in tree.
func (t *Tree) Nodes() []*Node {
	nodes := []*Node{}
	t.Root.AppendSubtree(&nodes)
	return nodes
}

// Return slice of nodes in node subtree.
func (n *Node) AppendSubtree(nodes *[]*Node) {
	*nodes = append(*nodes, n)
	if n.LChild != nil {
		n.LChild.AppendSubtree(nodes)
	}
	if n.RChild != nil {
		n.RChild.AppendSubtree(nodes)
	}
}

// Return true if node is a leaf.
func (n *Node) IsLeaf() bool {
	return n.LChild == nil || n.RChild == nil
}
