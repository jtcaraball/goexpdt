package leh

import (
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varVarVar struct {
	varInst1 base.Var
	varInst2 base.Var
	varInst3 base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVarVar leh.
func VarVarVar(
	varInst1, varInst2, varInst3 base.Var,
) *varVarVar {
	return &varVarVar{
		varInst1: varInst1,
		varInst2: varInst2,
		varInst3: varInst3,
	}
}

// Return CNF encoding of component.
func (l *varVarVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar1 := l.varInst1.Scoped(ctx)
	scpVar2 := l.varInst2.Scoped(ctx)
	scpVar3 := l.varInst3.Scoped(ctx)
	return l.buildEncoding(scpVar1, scpVar2, scpVar3, ctx)
}

// Generate cnf encoding.
func (l *varVarVar) buildEncoding(
	varInst1, varInst2, varInst3 base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	// Generate cnf
	nCNF := &cnf.CNF{}
	// Force fullness in vars
	nCNF.ExtendSemantics(varFullClauses(varInst1, ctx))
	nCNF.ExtendSemantics(varFullClauses(varInst2, ctx))
	nCNF.ExtendSemantics(varFullClauses(varInst3, ctx))
	// Generate var equality clauses
	nCNF.ExtendConsistency(fullVarEqualClauses(varInst1, varInst2, ctx))
	nCNF.ExtendConsistency(fullVarEqualClauses(varInst1, varInst3, ctx))
	// Generate hamming distance clauses
	nCNF.ExtendConsistency(hammingDistVV(varInst1, varInst2, ctx))
	nCNF.ExtendConsistency(hammingDistVV(varInst1, varInst3, ctx))
	// Add restriction clauses
	resClauses := [][]int{}
	dvn11 := distVarName(string(varInst1), string(varInst2))
	dvn13 := distVarName(string(varInst1), string(varInst3))
	for i := 1; i <= ctx.Dimension; i++ {
		for j := 0; j < i; j++ {
			resClauses = append(
				resClauses,
				[]int{
					-ctx.IVar(dvn11, ctx.Dimension-1, i),
					-ctx.IVar(dvn13, ctx.Dimension-1, j),
				},
			)
		}
	}
	nCNF.ExtendSemantics(resClauses)
	return nCNF, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *varVarVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return l, nil
}

// Return slice of pointers to component's children.
func (l *varVarVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varVarVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
