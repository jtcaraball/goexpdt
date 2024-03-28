package full

import (
	"goexpdt/cnf"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type fVar struct {
	varInst base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return var full.
func Var(varInst base.Var) *fVar {
	return &fVar{varInst: varInst}
}

// Return CNF encoding of component.
func (f *fVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar := f.varInst.Scoped(ctx)
	return f.buildEncoding(scpVar, ctx)
}

// Generate cnf encoding.
func (f *fVar) buildEncoding(
	varInst base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	no_bots_clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		no_bots_clauses = append(
			no_bots_clauses,
			[]int{-ctx.Var(string(varInst), i, base.BOT.Val())},
		)
	}
	return cnf.CNFFromClauses(no_bots_clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
func (f *fVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return f, nil
}

// Return slice of pointers to component's children.
func (f *fVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (f *fVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
