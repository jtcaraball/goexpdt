package subsumption

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constVar struct {
	varInst base.Var
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVar subsumption.
func ConstVar(
	constInst base.ConstInstance,
	varInst base.Var,
) *constVar {
	return &constVar{constInst: constInst, varInst: varInst}
}

// Return CNF encoding of component.
func (s *constVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
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
	return s.buildEncoding(scpConst, scpVar, ctx)
}

// Return CNF encoding of component.
func (s *constVar) buildEncoding(
	constInst base.Const,
	varInst base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i, ft := range constInst {
		if ft != base.BOT {
			clauses = append(
				clauses,
				[]int{ctx.Var(string(varInst), i, ft.Val())},
			)
		}
	}
	return cnf.CNFFromClauses(clauses), nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (s *constVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := s.constInst.Scoped(ctx)
	if err != nil {
		return s, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return s.buildSimplification(scpConst, ctx)
}

// Generate simplified component.
func (s *constVar) buildSimplification(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	for _, ft := range constInst {
		if ft != base.BOT {
			return s, nil
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (s *constVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *constVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
