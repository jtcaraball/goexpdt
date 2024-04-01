package leh

import (
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constConstConst struct {
	constInst1 base.ConstInstance
	constInst2 base.ConstInstance
	constInst3 base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constConst lel.
func ConstConstConst(
	constInst1, constInst2, constInst3 base.ConstInstance,
) *constConstConst {
	return &constConstConst{
		constInst1: constInst1,
		constInst2: constInst2,
		constInst3: constInst3,
	}
}

// Return CNF encoding of component.
func (l *constConstConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst1, err := l.constInst1.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst2, err := l.constInst2.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst3, err := l.constInst3.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst1,
		scpConst2,
		scpConst3,
	); err != nil {
		return nil, err
	}
	return l.buildEncoding(scpConst1, scpConst2, scpConst3)
}

// Generate cnf encoding.
func (l *constConstConst) buildEncoding(
	constInst1, constInst2, constInst3 base.Const,
) (*cnf.CNF, error) {
	if !constInst1.IsFull() || !constInst2.IsFull() || !constInst3.IsFull() {
		return cnf.CNFFromClauses([][]int{{}}), nil
	}
	hDistC1C2, err := hammingDist(constInst1, constInst2)
	if err != nil {
		return nil, err
	}
	hDistC1C3, err := hammingDist(constInst1, constInst3)
	if err != nil {
		return nil, err
	}
	if hDistC1C2 > hDistC1C3 {
		return cnf.CNFFromClauses([][]int{{}}), nil
	}
	return &cnf.CNF{}, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *constConstConst) Simplified(
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
	scpConst3, err := l.constInst3.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst1,
		scpConst2,
		scpConst3,
	); err != nil {
		return nil, err
	}
	return l.buildSimplification(scpConst1, scpConst2, scpConst3)
}

// Generate simplified component.
func (l *constConstConst) buildSimplification(
	constInst1, constInst2, constInst3 base.Const,
) (base.Component, error) {
	if !constInst1.IsFull() || !constInst2.IsFull() || !constInst3.IsFull() {
		return base.NewTrivial(false), nil
	}
	hDistC1C2, err := hammingDist(constInst1, constInst2)
	if err != nil {
		return nil, err
	}
	hDistC1C3, err := hammingDist(constInst1, constInst3)
	if err != nil {
		return nil, err
	}
	if hDistC1C2 > hDistC1C3 {
		return base.NewTrivial(false), nil
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *constConstConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constConstConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
