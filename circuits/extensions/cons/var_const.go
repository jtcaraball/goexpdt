package cons

import (
	"goexpdt/cnf"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConst struct {
	varInst base.Var
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst cons.
func VarConst(
	constInst base.ConstInstance,
	varInst base.Var,
) *varConst {
	return &varConst{constInst: constInst, varInst: varInst}
}

// Return CNF encoding of component.
func (c *varConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst, err := c.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := c.varInst.Scoped(ctx)
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return c.buildEncoding(scpConst, scpVar, ctx), nil
}

// Return CNF encoding of component.
func (c *varConst) buildEncoding(
	constInst base.Const,
	varInst base.Var,
	ctx *base.Context,
) *cnf.CNF {
	clauses := [][]int{}
	for i, ft := range constInst {
		if ft == base.BOT {
			continue
		}
		val := base.ZERO.Val()
		if ft == base.ZERO {
			val = base.ONE.Val()
		}
		clauses = append(
			clauses,
			[]int{-ctx.Var(string(varInst), i, val)},
		)
	}
	if len(clauses) == 0 {
		return &cnf.CNF{}
	}
	return cnf.CNFFromClauses(clauses)
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (c *varConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := c.constInst.Scoped(ctx)
	if err != nil {
		return c, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return c.buildSimplification(scpConst, ctx)
}

// Generate simplified component.
func (c *varConst) buildSimplification(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	return c, nil
}

// Return slice of pointers to component's children.
func (c *varConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (c *varConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
