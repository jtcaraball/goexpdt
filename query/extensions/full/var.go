package full

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Var is the variable version of the Full extension.
type Var struct {
	I query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable f.I
// has no features with bottom value.
func (f Var) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(f.I)
	clauses := []cnf.Clause{}

	for i := 0; i < ctx.Dim(); i++ {
		clauses = append(
			clauses,
			cnf.Clause{-ctx.CNFVar(sv, i, int(query.BOT))},
		)
	}

	return cnf.FromClauses(clauses), nil
}
