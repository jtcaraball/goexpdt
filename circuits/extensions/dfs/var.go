package dfs

import (
	"errors"
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type dfsVar struct {
	varInst base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return var dfs.
func Var(varInst base.Var) *dfsVar {
	return &dfsVar{varInst: varInst}
}

// Return CNF encoding of component.
func (d *dfsVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar := d.varInst.Scoped(ctx)
	return d.buildEncoding(scpVar, ctx)
}

// Generate cnf encoding.
func (d *dfsVar) buildEncoding(
	varInst base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	pnConsts, nnConsts, err := ctx.LeafsAsConsts()
	if err != nil {
		return nil, err
	}
	clauses := [][]int{}
	for _, pnConst := range pnConsts {
		for _, nnConst := range nnConsts {
			clauses = append(
				clauses,
				d.unsafeNotAWitnessClause(varInst, pnConst, nnConst, ctx),
			)
		}
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Generate not a witness clause.
func (d *dfsVar) notAWitnessClause(
	varInst base.Var,
	lConst1, lConst2 base.Const,
	ctx *base.Context,
) ([]int, error) {
	if len(lConst1) != ctx.Dimension || len(lConst2) != ctx.Dimension {
		return nil, errors.New("Constant and leaf have different dimensions")
	}
	clause := []int{}
	for i := 0; i < ctx.Dimension; i++ {
		if lConst1[i] != base.BOT &&
			lConst2[i] != base.BOT &&
			lConst1[i] != lConst2[i] {
			clause = append(
				clause,
				-ctx.Var(string(varInst), i, base.BOT.Val()),
			)
		}
	}
	return clause, nil
}

// Generate not a witness clause.
// Does not check for an index out of bound error.
func (d *dfsVar) unsafeNotAWitnessClause(
	varInst base.Var,
	lConst1, lConst2 base.Const,
	ctx *base.Context,
) []int {
	clause := []int{}
	for i := 0; i < ctx.Dimension; i++ {
		if lConst1[i] != base.BOT &&
			lConst2[i] != base.BOT &&
			lConst1[i] != lConst2[i] {
			clause = append(
				clause,
				-ctx.Var(string(varInst), i, base.BOT.Val()),
			)
		}
	}
	return clause
}

// Return pointer to simplified equivalent component which might be itself.
func (d *dfsVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return d, nil
}

// Return slice of pointers to component's children.
func (d *dfsVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (d *dfsVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
