package subsumption

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarConst is the variable-constant version of the Subsumption predicate.
type VarConst struct {
	I1 query.QVar
	I2 query.QConst
}

// Encoding returns a CNF that is true if and only if the query variable s.I1
// is subsumed by the query constant s.I2.
func (s VarConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(s.I1)
	sc, _ := ctx.ScopeConst(s.I2)

	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	clauses := []cnf.Clause{}
	for i, ft := range sc.Val {
		if ft == query.ONE {
			clauses = append(
				clauses,
				cnf.Clause{-ctx.CNFVar(sv, i, int(query.ZERO))},
			)
			continue
		}
		if ft == query.ZERO {
			clauses = append(
				clauses,
				cnf.Clause{-ctx.CNFVar(sv, i, int(query.ONE))},
			)
			continue
		}
		clauses = append(
			clauses,
			cnf.Clause{-ctx.CNFVar(sv, i, int(query.ONE))},
			cnf.Clause{-ctx.CNFVar(sv, i, int(query.ZERO))},
		)
	}

	return cnf.FromClauses(clauses), nil
}
