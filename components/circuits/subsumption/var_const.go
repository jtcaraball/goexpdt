package subsumption

import (
	"errors"
	"fmt"
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varconst struct {
	varInst instances.Var
	constInst instances.Const
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varvar subsumption.
func VarConst(varInst instances.Var, constInst instances.Const) *varconst {
	return &varconst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (s *varconst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if len(s.constInst) != ctx.Dimension {
		return nil, errors.New(
			fmt.Sprintf(
				`subsumption.varconst -> constant: wrong dim %d
				(%d feats in context)`,
				len(s.constInst),
				ctx.Dimension,
			),
		)
	}
	clauses := [][]int{}
	for i, f := range s.constInst {
		if f == instances.ONE {
			clauses = append(
				clauses,
				[]int{-ctx.Var(string(s.varInst), i, instances.ZERO.Val())},
			)
			continue
		}
		if f == instances.ZERO {
			clauses = append(
				clauses,
				[]int{-ctx.Var(string(s.varInst), i, instances.ONE.Val())},
			)
			continue
		}
		clauses = append(
			clauses,
			[]int{-ctx.Var(string(s.varInst), i, instances.ONE.Val())},
			[]int{-ctx.Var(string(s.varInst), i, instances.ZERO.Val())},
		)
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (s *varconst) Simplified() (components.Component, error) {
	return s, nil
}

// Return slice of pointers to component's children.
func (s *varconst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *varconst) IsTrivial() (yes bool, value bool) {
	return false, false
}
