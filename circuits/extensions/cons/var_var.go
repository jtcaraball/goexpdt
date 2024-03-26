package cons

import (
	"goexpdt/cnf"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varVar struct {
	varInst1 base.Var
	varInst2 base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar cons.
func VarVar(varInst1, varInst2 base.Var) *varVar {
	return &varVar{varInst1: varInst1, varInst2: varInst2}
}

// Return CNF encoding of component.
func (c *varVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar1 := c.varInst1.Scoped(ctx)
	scpVar2 := c.varInst2.Scoped(ctx)
	return c.buildEncoding(scpVar1, scpVar2, ctx)
}

// Generate cnf encoding.
func (c *varVar) buildEncoding(
	varInst1, varInst2 base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.Var(string(varInst1), i, base.ONE.Val()),
				-ctx.Var(string(varInst2), i, base.ZERO.Val()),
			},
			[]int{
				-ctx.Var(string(varInst1), i, base.ZERO.Val()),
				-ctx.Var(string(varInst2), i, base.ONE.Val()),
			},
		)
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (c *varVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return c, nil
}

// Return slice of pointers to component's children.
func (c *varVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varVar) IsTrivial() (yes bool, value bool) {
	return false, false
}

