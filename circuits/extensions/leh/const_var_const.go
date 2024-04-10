package leh

import (
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constVarConst struct {
	constInst1 base.ConstInstance
	constInst2 base.ConstInstance
	varInst    base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVarConst leh.
func ConstVarConst(
	constInst1 base.ConstInstance,
	varInst base.Var,
	constInst2 base.ConstInstance,
) *constVarConst {
	return &constVarConst{
		constInst1: constInst1,
		constInst2: constInst2,
		varInst:    varInst,
	}
}

// Return CNF encoding of component.
func (l *constVarConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst1, err := l.constInst1.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpConst2, err := l.constInst2.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := l.varInst.Scoped(ctx)
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst1,
		scpConst2,
	); err != nil {
		return nil, err
	}
	return l.buildEncoding(scpConst1, scpConst2, scpVar, ctx)
}

// Generate cnf encoding.
func (l *constVarConst) buildEncoding(
	constInst1, constInst2 base.Const,
	varInst base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	// Check for easy out
	if !constInst1.IsFull() || !constInst2.IsFull() {
		return cnf.CNFFromClauses([][]int{{}}), nil
	}
	// Generate cnf
	nCNF := &cnf.CNF{}
	// Determine hamming distance between constants
	cchd, err := hammingDistCC(constInst1, constInst2)
	if err != nil {
		return nil, err
	}
	if cchd == ctx.Dimension {
		return nCNF, nil
	}
	// Force fullness in var
	nCNF.ExtendSemantics(varFullClauses(varInst, ctx))
	// Generate hamming distance clauses
	distClauses, err := hammingDistVC(varInst, constInst1, ctx)
	if err != nil {
		return nil, err
	}
	nCNF.ExtendConsistency(distClauses)
	// Add distance restriction clauses
	resClauses := [][]int{}
	dvn := distVarName(string(varInst), constInst1.AsString())
	for i := cchd + 1; i <= ctx.Dimension; i++ {
		resClauses = append(
			resClauses,
			[]int{-ctx.IVar(dvn, ctx.Dimension-1, i)},
		)
	}
	nCNF.ExtendSemantics(resClauses)
	return nCNF, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *constVarConst) Simplified(
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
func (l *constVarConst) buildSimplification(
	constInst1, constInst2 base.Const,
	ctx *base.Context,
) (base.Component, error) {
	if !constInst1.IsFull() || !constInst2.IsFull() {
		return base.NewTrivial(false), nil
	}
	cchd, err := hammingDistCC(constInst1, constInst2)
	if err != nil {
		return nil, err
	}
	if cchd == ctx.Dimension {
		return base.NewTrivial(true), nil
	}
	return l, nil
}

// Return slice of pointers to component's children.
func (l *constVarConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constVarConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
