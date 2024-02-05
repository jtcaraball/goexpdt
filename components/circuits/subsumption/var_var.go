package subsumption

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varvar struct {
	varInst1 instances.Var
	varInst2 instances.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varvar subsumption.
func VarVar(varInst1, varInst2 instances.Var) *varvar {
	return &varvar{varInst1: varInst1, varInst2: varInst2}
}

// Return CNF encoding of component.
func (s *varvar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.Var(string(s.varInst1), i, instances.ONE.Val()),
				ctx.Var(string(s.varInst2), i, instances.ONE.Val()),
			},
			[]int{
				-ctx.Var(string(s.varInst1), i, instances.ZERO.Val()),
				ctx.Var(string(s.varInst2), i, instances.ZERO.Val()),
			},
		)
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (s *varvar) Simplified() (components.Component, error) {
	return s, nil
}

// Return slice of pointers to component's children.
func (s *varvar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *varvar) IsTrivial() (yes bool, value bool) {
	return false, false
}
