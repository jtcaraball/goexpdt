package dft

import (
	"errors"
	"goexpdt/cnf"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type dftVar struct {
	varInst base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return var dft.
func Var(varInst base.Var) *dftVar {
	return &dftVar{varInst: varInst}
}

// Return CNF encoding of component.
func (d *dftVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar := d.varInst.Scoped(ctx)
	return d.buildEncoding(scpVar, ctx)
}

// Generate cnf encoding.
func (d *dftVar) buildEncoding(
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
			nwClause, err := d.notAWitnessClause(varInst, pnConst, nnConst, ctx)
			if err != nil {
				return nil, err
			}
			clauses = append(clauses, nwClause)
		}
	}
	return cnf.CNFFromClauses(clauses), nil
}

// Generate not a witness clause.
func (d *dftVar) notAWitnessClause(
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

// Return pointer to simplified equivalent component which might be itself.
func (d *dftVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return d, nil
}

// Return slice of pointers to component's children.
func (d *dftVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (d *dftVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
