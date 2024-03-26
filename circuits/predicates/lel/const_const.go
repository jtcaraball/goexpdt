package lel

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


// Return constConst lel.
func ConstConst(constInst1, constInst2 base.ConstInstance) *constConst {
	return &constConst{constInst1: constInst1, constInst2: constInst2}
}

// Return CNF encoding of component.
func (l *constConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst1, err := l.constInst1.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst2, err := l.constInst2.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst1,
		scpConst2,
	); err != nil {
		return nil, err
	}
	return l.buildEncoding(scpConst1, scpConst2, ctx)
}

// Generate cnf encoding.
func (l *constConst) buildEncoding(
	constInst1, constInst2 base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	const1Bots := 0
	const2Bots := 0
	for i := 0; i < ctx.Dimension; i++ {
		if constInst1[i] == base.BOT {
			const1Bots += 1
		}
		if constInst2[i] == base.BOT {
			const2Bots += 1
		}
	}
	if const1Bots < const2Bots {
		return cnf.CNFFromClauses([][]int{{}}), nil
	}
	return &cnf.CNF{}, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *constConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst1, err := l.constInst1.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	scpConst2, err := l.constInst2.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst1,
		scpConst2,
	); err != nil {
		return nil, err
	}
	return l.buildSimplification(scpConst1, scpConst2, ctx)
}

// Generate simplified component.
func (l *constConst) buildSimplification(
	constInst1, constInst2 base.Const,
	ctx *base.Context,
) (base.Component, error) {
	const1Bots := 0
	const2Bots := 0
	for i := 0; i < ctx.Dimension; i++ {
		if constInst1[i] == base.BOT {
			const1Bots += 1
		}
		if constInst2[i] == base.BOT {
			const2Bots += 1
		}
	}
	if const1Bots < const2Bots {
		return base.NewTrivial(false), nil
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *constConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
