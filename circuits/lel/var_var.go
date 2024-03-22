package lel

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varVar struct {
	varInst1 base.Var
	varInst2 base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func VarVar(varInst1, varInst2 base.Var) *varVar {
	return &varVar{varInst1: varInst1, varInst2: varInst2}
}

// Return CNF encoding of component.
func (s *varVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar1 := s.varInst1.Scoped(ctx)
	scpVar2 := s.varInst2.Scoped(ctx)
	return s.buildEncoding(scpVar1, scpVar2, ctx)
}

// Generate cnf encoding.
func (l *varVar) buildEncoding(
	varInst1, varInst2 base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	cnf := &cnf.CNF{}
	cnf.ExtendConsistency(genCountClauses(string(varInst1), ctx))
	cnf.ExtendConsistency(genCountClauses(string(varInst2), ctx))
	// If we see a number of bots in x then we must see less or equal on y
	var i, j int
	cVarName1 := "c" + string(varInst1)
	cVarName2 := "c" + string(varInst2)
	for i = 0; i < ctx.Dimension; i++ {
		lel_clauses := []int{-ctx.IVar(cVarName1, ctx.Dimension - 1, i)}
		for j = 0; j <= i; j++ {
			lel_clauses = append(
				lel_clauses,
				ctx.IVar(cVarName2, ctx.Dimension - 1, j),
			)
		}
		cnf.AppendSemantics(lel_clauses)
	}
	return cnf, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (l *varVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return l, nil
}

// Return slice of pointers to component's children.
func (l *varVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
