package cons

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarConst is the variable-constant version of the Consistent extension.
type VarConst struct {
	I1 query.QVar
	I2 query.QConst
}

// Encoding returns a CNF that is true if and only if the query variable c.I1
// and query constant c.I2 have the same values in every feature where both are
// different than bottom.
func (c VarConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(c.I1)
	sc, _ := ctx.ScopeConst(c.I2)

	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	clauses := []cnf.Clause{}

	for i, ft := range sc.Val {
		if ft == query.BOT {
			continue
		}

		val := int(query.ZERO)
		if ft == query.ZERO {
			val = int(query.ONE)
		}

		clauses = append(clauses, cnf.Clause{-ctx.CNFVar(sv, i, val)})
	}

	if len(clauses) == 0 {
		return cnf.TrueCNF, nil
	}

	return cnf.FromClauses(clauses), nil
}
