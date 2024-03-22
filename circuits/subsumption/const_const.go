package subsumption

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
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


// Return constConst subsumption.
func ConstConst(constInst1, constInst2 base.ConstInstance) *constConst {
	return &constConst{constInst1: constInst1, constInst2: constInst2}
}

// Return CNF encoding of component.
func (s *constConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst1, err := s.constInst1.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst2, err := s.constInst2.Scoped(ctx)
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
	return s.buildEncoding(scpConst1, scpConst2, ctx)
}

// Generate cnf encoding.
func (s *constConst) buildEncoding(
	constInst1, constInst2 base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	for i, ft := range constInst1 {
		if ft != base.BOT && ft != constInst2[i] {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
	}
	return &cnf.CNF{}, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (s *constConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst1, err := s.constInst1.Scoped(ctx)
	if err != nil {
		return s, nil
	}
	scpConst2, err := s.constInst2.Scoped(ctx)
	if err != nil {
		return s, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst1,
		scpConst2,
	); err != nil {
		return nil, err
	}
	return s.buildSimplification(scpConst1, scpConst2, ctx)
}

// Generate simplified component.
func (s *constConst) buildSimplification(
	constInst1 base.Const,
	constInst2 base.Const,
	ctx *base.Context,
) (base.Component, error) {
	for i, ft := range constInst1 {
		if ft != base.BOT && ft != constInst2[i] {
			return base.NewTrivial(false), nil
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (s *constConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *constConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
