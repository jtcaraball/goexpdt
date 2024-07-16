package cons

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarVar is the variable-variable version of the Consistent predicate.
type VarVar struct {
	I1 query.QVar
	I2 query.QVar
}

// Encoding returns a CNF that is true if and only if the query variables c.I1
// and c.I2 have the same values in every feature where both are different than
// bottom.
func (c VarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv1 := ctx.ScopeVar(c.I1)
	sv2 := ctx.ScopeVar(c.I2)

	clauses := []cnf.Clause{}

	for i := 0; i < ctx.Dim(); i++ {
		clauses = append(
			clauses,
			cnf.Clause{
				-ctx.CNFVar(sv1, i, int(query.ONE)),
				-ctx.CNFVar(sv2, i, int(query.ZERO)),
			},
			cnf.Clause{
				-ctx.CNFVar(sv1, i, int(query.ZERO)),
				-ctx.CNFVar(sv2, i, int(query.ONE)),
			},
		)
	}

	return cnf.FromClauses(clauses), nil
}
