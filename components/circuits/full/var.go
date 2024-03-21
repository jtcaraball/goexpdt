package full

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type fVar struct {
	varInst components.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func Var(varInst components.Var) *fVar {
	return &fVar{varInst: varInst}
}

// Return CNF encoding of component.
func (f *fVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	scpVar := f.varInst.Scoped(ctx)
	return f.buildEncoding(scpVar, ctx)
}

// Generate cnf encoding.
func (f *fVar) buildEncoding(
	varInst components.Var,
	ctx *components.Context,
) (*cnf.CNF, error) {
	no_bots_clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		no_bots_clauses = append(
			no_bots_clauses,
			[]int{-ctx.Var(string(varInst), i, components.BOT.Val())},
		)
	}
	return cnf.CNFFromClauses(no_bots_clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
func (f *fVar) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	return f, nil
}

// Return slice of pointers to component's children.
func (f *fVar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (f *fVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
