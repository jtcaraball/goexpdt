package leh

import (
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constVarVar struct {
	varInst1 base.Var
	varInst2 base.Var
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVarVar leh.
func ConstVarVar(
	constInst base.ConstInstance,
	varInst1, varInst2 base.Var,
) *constVarVar {
	return &constVarVar{
		varInst1: varInst1,
		varInst2: varInst2,
		constInst: constInst,
	}
}

// Return CNF encoding of component.
func (l *constVarVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar1 := l.varInst1.Scoped(ctx)
	scpVar2 := l.varInst2.Scoped(ctx)
	scpConst, err := l.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return l.buildEncoding(scpVar1, scpVar2, scpConst, ctx)
}

// Generate cnf encoding.
func (l *constVarVar) buildEncoding(
	varInst1, varInst2 base.Var,
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	// Check for easy out
	if !constInst.IsFull() {
		return cnf.CNFFromClauses([][]int{{}}), nil
	}
	// Generate cnf
	nCNF := &cnf.CNF{}
	// Force fullness in var
	nCNF.ExtendSemantics(varFullClauses(varInst1, ctx))
	nCNF.ExtendSemantics(varFullClauses(varInst2, ctx))
	// Generate hamming distance clauses
	distV1Clauses, err := hammingDistVC(varInst1, constInst, ctx)
	if err != nil {
		return nil, err
	}
	distV2Clauses, err := hammingDistVC(varInst2, constInst, ctx)
	if err != nil {
		return nil, err
	}
	nCNF.ExtendConsistency(distV1Clauses)
	nCNF.ExtendConsistency(distV2Clauses)
	// Add distance restriction clauses
	resClauses := [][]int{}
	// Consistent the order in which params where passed creating hamming dist.
	dvnVar1Const := distVarName(string(varInst1), constName(constInst))
	dvnVar2Const := distVarName(string(varInst2), constName(constInst))
	for i := 1; i <= ctx.Dimension; i++ {
		for j := 0; j < i; j++ {
			resClauses = append(
				resClauses,
				[]int{
					-ctx.IVar(dvnVar1Const, ctx.Dimension-1, i),
					-ctx.IVar(dvnVar2Const, ctx.Dimension-1, j),
				},
			)
		}
	}
	nCNF.ExtendSemantics(resClauses)
	return nCNF, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *constVarVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := l.constInst.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return l, nil
}

// Return slice of pointers to component's children.
func (l *constVarVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constVarVar) IsTrivial() (yes bool, value bool) {
	return false, false
}

