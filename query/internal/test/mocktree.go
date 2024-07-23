package test

import (
	"errors"

	"github.com/jtcaraball/goexpdt/query"
)

type mockTree struct {
	dim   int
	nodes []query.Node
}

func NewMockTree(dim int, nodes []query.Node) (mockTree, error) {
	for _, n := range nodes {
		invF := n.Feat < 0 || n.Feat >= dim
		invC := n.Feat < 0 || n.Feat >= dim
		if invF || invC {
			return mockTree{},
				errors.New("Invalid node at mocktree initialization")
		}
	}
	return mockTree{dim, nodes}, nil
}

func (t mockTree) Dim() int {
	return t.dim
}

func (t mockTree) Nodes() []query.Node {
	return t.nodes
}
