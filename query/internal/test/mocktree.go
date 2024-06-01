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

func (t mockTree) NodesConsts() []query.Const {
	r := make([]query.Const, len(t.nodes))
	r[0] = query.AllBotConst(t.dim)

	for i := 1; i < len(t.nodes); i++ {
		if t.nodes[i].IsLeaf() {
			continue
		}

		zv := (make([]query.FeatV, t.dim))
		ov := (make([]query.FeatV, t.dim))

		copy(zv, r[i].Val)
		copy(ov, r[i].Val)

		zv[t.nodes[i].Feat] = query.ZERO
		ov[t.nodes[i].Feat] = query.ONE

		r[t.nodes[i].ZChild] = query.Const{Val: zv}
		r[t.nodes[i].OChild] = query.Const{Val: ov}
	}

	return r
}

func (t mockTree) PosLeafsConsts() []query.Const {
	var r []query.Const

	nc := t.NodesConsts()
	for i, n := range t.nodes {
		if n.IsLeaf() && n.Value {
			r = append(r, nc[i])
		}
	}

	return r
}

func (t mockTree) NegLeafsConsts() []query.Const {
	var r []query.Const

	nc := t.NodesConsts()
	for i, n := range t.nodes {
		if n.IsLeaf() && !n.Value {
			r = append(r, nc[i])
		}
	}

	return r
}
