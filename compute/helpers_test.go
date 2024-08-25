package compute_test

import "github.com/jtcaraball/goexpdt/query"

const SOLVER = "/kissat"

type mockModel struct {
	dim int
}

func (m mockModel) Dim() int {
	return m.dim
}

func (m mockModel) Nodes() []query.Node {
	return nil
}

