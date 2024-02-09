package subsumption

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConst struct {
	varInst components.Var
	constInst components.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst subsumption.
func VarConst(varInst components.Var, constInst components.Const) *varConst {
	return &varConst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (s *varConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	scpConst, err := s.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := s.varInst.Scoped(ctx)
	if err = components.ValidateConstsDim(
		"subsumption.VarConst",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return s.buildEncoding(scpVar, scpConst, ctx)
}

// Generate cnf encoding.
func (s *varConst) buildEncoding(
	varInst components.Var,
	constInst components.Const,
	ctx *components.Context,
) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i, f := range constInst {
		if f == components.ONE {
			clauses = append(
				clauses,
				[]int{-ctx.Var(string(varInst), i, components.ZERO.Val())},
			)
			continue
		}
		if f == components.ZERO {
			clauses = append(
				clauses,
				[]int{-ctx.Var(string(varInst), i, components.ONE.Val())},
			)
			continue
		}
		clauses = append(
			clauses,
			[]int{-ctx.Var(string(varInst), i, components.ONE.Val())},
			[]int{-ctx.Var(string(varInst), i, components.ZERO.Val())},
		)
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
func (s *varConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
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
