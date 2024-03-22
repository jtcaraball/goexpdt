package subsumption

import (
	"goexpdt/cnf"
	"goexpdt/base"
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

// Return varVar subsumption.
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
func (s *varVar) buildEncoding(
	varInst1, varInst2 base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.Var(string(varInst1), i, base.ONE.Val()),
				ctx.Var(string(varInst2), i, base.ONE.Val()),
			},
			[]int{
				-ctx.Var(string(varInst1), i, base.ZERO.Val()),
				ctx.Var(string(varInst2), i, base.ZERO.Val()),
			},
		)
	}
	return cnf.CNFFromClauses(clauses), nil
}


// Return pointer to simplified equivalent component which might be itself.
func (s *varVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return s, nil
}

// Return slice of pointers to component's children.
func (s *varVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *varVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
