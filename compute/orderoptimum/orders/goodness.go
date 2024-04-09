package orders

import (
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type goodness struct {
	varInst   base.Var
	constInst base.Const
}

// =========================== //
//           HELPERS           //
// =========================== //

// =========================== //
//           METHODS           //
// =========================== //

// Return goodness order instance.
func Goodness(varInst base.Var, constInst base.Const) *goodness {
	return &goodness{varInst: varInst, constInst: constInst}
}

// TODO!
// Return CNF encoding order.
func (g *goodness) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	return nil, nil
}

// TODO!
// Return pointer to simplified equivalent order which might be itself.
func (g *goodness) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return nil, nil
}

// Return empty slice of components.
func (g *goodness) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (g *goodness) IsTrivial() (yes bool, value bool) {
	return false, false
}
