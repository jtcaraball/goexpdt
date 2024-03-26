package cons

import (
	"goexpdt/cnf"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constConst struct {
	constInst1 base.ConstInstance
	constInst2 base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constConst cons.
func ConstConst(constInst1, constInst2 base.ConstInstance) *constConst {
	return &constConst{constInst1: constInst1, constInst2: constInst2}
}

// Return CNF encoding of component.
func (c *constConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst1, err := c.constInst1.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst2, err := c.constInst2.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	return c.buildEncoding(scpConst1, scpConst2, ctx)
}

// Generate cnf encoding.
func (c *constConst) buildEncoding(
	constInst1, constInst2 base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	if err := base.ValidateConstsDim(
		ctx.Dimension,
		constInst1,
		constInst2,
	); err != nil {
		return nil, err
	}
	for i := 0; i < ctx.Dimension; i++ {
		if constInst1[i] != base.BOT &&
			constInst2[i] != base.BOT &&
			constInst1[i] != constInst2[i] {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
	}
	return &cnf.CNF{}, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (c *constConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst1, err := c.constInst1.Scoped(ctx)
	if err != nil {
		return c, nil
	}
	scpConst2, err := c.constInst2.Scoped(ctx)
	if err != nil {
		return c, nil
	}
	return c.buildSimplification(scpConst1, scpConst2, ctx)
}

// Generate simplified component.
func (c *constConst) buildSimplification(
	constInst1, constInst2 base.Const,
	ctx *base.Context,
) (base.Component, error) {
	if err := base.ValidateConstsDim(
		ctx.Dimension,
		constInst1,
		constInst2,
	); err != nil {
		return nil, err
	}
	for i := 0; i < ctx.Dimension; i++ {
		if constInst1[i] != base.BOT &&
			constInst2[i] != base.BOT &&
			constInst1[i] != constInst2[i] {
			return base.NewTrivial(false), nil
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (c *constConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (c *constConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
