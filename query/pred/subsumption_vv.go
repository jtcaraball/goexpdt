package pred

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// SubsumptionVarVar is the variable-variable version of the Subsumption
// predicate.
type SubsumptionVarVar struct {
	I1 query.QVar
	I2 query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable s.I1
// is subsumed by the query variable s.I2.
func (s SubsumptionVarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv1 := ctx.ScopeVar(s.I1)
	sv2 := ctx.ScopeVar(s.I2)

	clauses := []cnf.Clause{}

	for i := 0; i < ctx.Dim(); i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.CNFVar(sv1, i, int(query.ONE)),
				ctx.CNFVar(sv2, i, int(query.ONE)),
			},
			[]int{
				-ctx.CNFVar(sv1, i, int(query.ZERO)),
				ctx.CNFVar(sv2, i, int(query.ZERO)),
			},
		)
	}

	return cnf.FromClauses(clauses), nil
}
