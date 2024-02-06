package lel

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varVar struct {
	varInst1 instances.Var
	varInst2 instances.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func VarVar(varInst1, varInst2 instances.Var) *varVar {
	return &varVar{varInst1: varInst1, varInst2: varInst2}
}

// Return CNF encoding of component.
func (l *varVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	cnf := &cnf.CNF{}
	cnf.ExtendConsistency(genCountClauses(string(l.varInst1), ctx))
	cnf.ExtendConsistency(genCountClauses(string(l.varInst2), ctx))
	// If we see a number of bots in x then we must see less or equal on y
	var i, j int
	cVarName1 := "c" + string(l.varInst1)
	cVarName2 := "c" + string(l.varInst2)
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
	ctx *components.Context,
) (components.Component, error) {
	return l, nil
}

// Return slice of pointers to component's children.
func (l *varVar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
