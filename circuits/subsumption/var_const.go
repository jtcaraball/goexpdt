package subsumption

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConst struct {
	varInst base.Var
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst subsumption.
func VarConst(
	varInst base.Var,
	constInst base.ConstInstance,
) *varConst {
	return &varConst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (s *varConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst, err := s.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := s.varInst.Scoped(ctx)
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return s.buildEncoding(scpVar, scpConst, ctx)
}

// Generate cnf encoding.
func (s *varConst) buildEncoding(
	varInst base.Var,
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i, ft := range constInst {
		if ft == base.ONE {
			clauses = append(
				clauses,
				[]int{-ctx.Var(string(varInst), i, base.ZERO.Val())},
			)
			continue
		}
		if ft == base.ZERO {
			clauses = append(
				clauses,
				[]int{-ctx.Var(string(varInst), i, base.ONE.Val())},
			)
			continue
		}
		clauses = append(
			clauses,
			[]int{-ctx.Var(string(varInst), i, base.ONE.Val())},
			[]int{-ctx.Var(string(varInst), i, base.ZERO.Val())},
		)
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Return pointer to simplified equivalent component which might be itself.
func (s *varConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return s, nil
}

// Return slice of pointers to component's children.
func (s *varConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *varConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
