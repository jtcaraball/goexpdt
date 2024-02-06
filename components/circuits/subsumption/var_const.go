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

type varConst struct {
	varInst instances.Var
	constInst instances.Const
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst subsumption.
func VarConst(varInst instances.Var, constInst instances.Const) *varConst {
	return &varConst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (s *varConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := s.validateInstances(ctx); err != nil {
		return nil, err
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
func (s *varConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := s.validateInstances(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

// Return slice of pointers to component's children.
func (s *varConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *varConst) IsTrivial() (yes bool, value bool) {
	return false, false
}

func (s *varConst) validateInstances(ctx *components.Context) error {
	if len(s.constInst) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				`subsumption.constVar -> constant: wrong dim %d
				(%d feats in context)`,
				len(s.constInst),
				ctx.Dimension,
			),
		)
	}
	return nil
}
