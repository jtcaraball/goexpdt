package cons

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstVar is the variable-variable version of the Consistent extension.
type ConstVar struct {
	I1 query.QConst
	I2 query.QVar
}

// Encoding returns a CNF that is true if and only if the query constants c.I1
// and query variable c.I2 have the same values in every feature where both are
// different than bottom.
func (c ConstVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc, _ := ctx.ScopeConst(c.I1)
	sv := ctx.ScopeVar(c.I2)

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
