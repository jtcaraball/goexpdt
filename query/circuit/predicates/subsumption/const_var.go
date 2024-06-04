package subsumption

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstVar is the constant-variable version of the Subsumption predicate.
type ConstVar struct {
	I1 query.QConst
	I2 query.QVar
}

// Encoding returns a CNF that is true if and only if the query constant I1 is
// subsumed by the query variable I2.
func (s ConstVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(s.I2)
	sc, _ := ctx.ScopeConst(s.I1)

	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	clauses := []cnf.Clause{}
	for i, ft := range sc.Val {
		if ft != query.BOT {
			clauses = append(
				clauses,
				cnf.Clause{ctx.CNFVar(sv, i, int(ft))},
			)
		}
	}

	return cnf.FromClauses(clauses), nil
}
