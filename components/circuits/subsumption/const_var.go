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

type constVar struct {
	varInst instances.Var
	constInst instances.Const
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVar subsumption.
func ConstVar(constInst instances.Const, varInst instances.Var) *constVar {
	return &constVar{constInst: constInst, varInst: varInst}
}

// Return CNF encoding of component.
func (s *constVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if len(s.constInst) != ctx.Dimension {
		return nil, errors.New(
			fmt.Sprintf(
				`subsumption.constVar -> constant: wrong dim %d
				(%d feats in context)`,
				len(s.constInst),
				ctx.Dimension,
			),
		)
	}
	clauses := [][]int{}
	for i, f := range s.constInst {
		if f != instances.BOT {
			clauses = append(
				clauses,
				[]int{ctx.Var(string(s.varInst), i, f.Val())},
			)
		}
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (s *constVar) Simplified() (components.Component, error) {
	return s, nil
}

// Return slice of pointers to component's children.
func (s *constVar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *constVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
