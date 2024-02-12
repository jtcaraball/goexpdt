package subsumption

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constVar struct {
	varInst components.Var
	constInst components.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVar subsumption.
func ConstVar(
	constInst components.ConstInstance,
	varInst components.Var,
) *constVar {
	return &constVar{constInst: constInst, varInst: varInst}
}

// Return CNF encoding of component.
func (s *constVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	scpConst, err := s.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := s.varInst.Scoped(ctx)
	if err = components.ValidateConstsDim(
		"subsumption.ConstVar",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return s.buildEncoding(scpConst, scpVar, ctx)
}

// Return CNF encoding of component.
func (s *constVar) buildEncoding(
	constInst components.Const,
	varInst components.Var,
	ctx *components.Context,
) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i, ft := range constInst {
		if ft != components.BOT {
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
	ctx *components.Context,
) (components.Component, error) {
	scpConst, err := s.constInst.Scoped(ctx)
	if err != nil {
		return s, nil
	}
	if err = components.ValidateConstsDim(
		"subsumption.ConstVar",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return s.buildSimplification(scpConst, ctx)
}

// Generate simplified component.
func (s *constVar) buildSimplification(
	constInst components.Const,
	ctx *components.Context,
) (components.Component, error) {
	for _, f := range constInst {
		if f != components.BOT {
			return s, nil
		}
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (s *constVar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *constVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
