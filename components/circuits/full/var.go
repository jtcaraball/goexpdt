package full

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type fVar struct {
	varInst instances.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func Var(varInst instances.Var) *fVar {
	return &fVar{varInst: varInst}
}

// Return CNF encoding of component.
func (f *fVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	no_bots_clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		no_bots_clauses = append(
			no_bots_clauses,
			[]int{-ctx.Var(string(f.varInst), i, instances.BOT.Val())},
		)
	}
	return cnf.CNFFromClauses(no_bots_clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
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
