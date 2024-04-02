package leh

import (
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConstConst struct {
	constInst1 base.ConstInstance
	constInst2 base.ConstInstance
	varInst base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConstConst leh.
func VarConstConst(
	varInst base.Var,
	constInst1, constInst2 base.ConstInstance,
) *varConstConst {
	return &varConstConst{
		constInst1: constInst1,
		constInst2: constInst2,
		varInst: varInst,
	}
}

// Return CNF encoding of component.
func (l *varConstConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
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
func (l *varConstConst) buildEncoding(
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
	// Force fullness in var
	nCNF.ExtendSemantics(varFullClauses(varInst, ctx))
	// Generate hamming distance clauses
	distC1Clauses, err := hammingDistVC(varInst, constInst1, ctx)
	if err != nil {
		return nil, err
	}
	distC2Clauses, err := hammingDistVC(varInst, constInst2, ctx)
	if err != nil {
		return nil, err
	}
	nCNF.ExtendConsistency(distC1Clauses)
	nCNF.ExtendConsistency(distC2Clauses)
	// Add distance restriction clauses
	resClauses := [][]int{}
	dvn1 := distVarName(string(varInst), constName(constInst1))
	dvn2 := distVarName(string(varInst), constName(constInst2))
	for i := 1; i <= ctx.Dimension; i++ {
		for j := 0; j < i; j++ {
			resClauses = append(
				resClauses,
				[]int{
					-ctx.IVar(dvn1, ctx.Dimension - 1, i),
					-ctx.IVar(dvn2, ctx.Dimension - 1, j),
				},
			)
		}
	}
	nCNF.ExtendSemantics(resClauses)
	return nCNF, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *varConstConst) Simplified(
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
	return l.buildSimplification(scpConst1, scpConst2)
}

// Generate simplified component.
func (l *varConstConst) buildSimplification(
	constInst1, constInst2 base.Const,
) (base.Component, error) {
	if !constInst1.IsFull() || !constInst2.IsFull() {
		return base.NewTrivial(false), nil
	}
	return l, nil
}

// Return slice of pointers to component's children.
func (l *varConstConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varConstConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
