package lel

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constConst struct {
	constInst1 components.ConstInstance
	constInst2 components.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //


// Return constConst lel.
func ConstConst(constInst1, constInst2 components.ConstInstance) *constConst {
	return &constConst{constInst1: constInst1, constInst2: constInst2}
}

// Return CNF encoding of component.
func (s *constConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	scpConst1, err := s.constInst1.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst2, err := s.constInst2.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = components.ValidateConstsDim(
		"lel.ConstConst",
		ctx,
		scpConst1,
		scpConst2,
	); err != nil {
		return nil, err
	}
	return s.buildEncoding(scpConst1, scpConst2, ctx)
}

// Generate cnf encoding.
func (l *constConst) buildEncoding(
	constInst1, constInst2 components.Const,
	ctx *components.Context,
) (*cnf.CNF, error) {
	const1Bots := 0
	const2Bots := 0
	for i := 0; i < ctx.Dimension; i++ {
		if constInst1[i] == components.BOT {
			const1Bots += 1
		}
		if constInst2[i] == components.BOT {
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
	ctx *components.Context,
) (components.Component, error) {
	scpConst1, err := l.constInst1.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	scpConst2, err := l.constInst2.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	if err = components.ValidateConstsDim(
		"lel.ConstConst",
		ctx,
		scpConst1,
		scpConst2,
	); err != nil {
		return nil, err
	}
	return l.buildSimplification(scpConst1, scpConst2, ctx)
}

// Generate simplified component.
func (l *constConst) buildSimplification(
	constInst1, constInst2 components.Const,
	ctx *components.Context,
) (components.Component, error) {
	const1Bots := 0
	const2Bots := 0
	for i := 0; i < ctx.Dimension; i++ {
		if constInst1[i] == components.BOT {
			const1Bots += 1
		}
		if constInst2[i] == components.BOT {
			const2Bots += 1
		}
	}
	if const1Bots < const2Bots {
		return components.NewTrivial(false), nil
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *constConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
