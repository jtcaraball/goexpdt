package dfs

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Var is the variable version of the DFS extension.
type Var struct {
	I query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable d.I
// is such that all its completions have the same Model evaluation.
func (d Var) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(d.I)
	pleaf := ctx.PosLeafsConsts()
	nleaf := ctx.NegLeafsConsts()

	if len(pleaf) == 0 && len(nleaf) == 0 {
		return cnf.CNF{}, errors.New("Invalid encoding on empty model")
	}

	clauses := []cnf.Clause{}

	for _, p := range pleaf {
		for _, n := range nleaf {
			cl, err := notAWitnessClause(sv, p, n, ctx)
			if err != nil {
				return cnf.CNF{}, err
			}

			clauses = append(clauses, cl)
		}
	}

	return cnf.FromClauses(clauses), nil
}

func notAWitnessClause(
	v query.QVar,
	l1, l2 query.QConst,
	ctx query.QContext,
) (cnf.Clause, error) {
	dim := ctx.Dim()
	clause := cnf.Clause{}

	if dim != len(l1.Val) || dim != len(l2.Val) {
		return clause,
			errors.New("Invalid comparison of constants of different dim")
	}

	for i := 0; i < ctx.Dim(); i++ {
		if l1.Val[i] != query.BOT &&
			l2.Val[i] != query.BOT &&
			l1.Val[i] != l2.Val[i] {
			clause = append(clause, -ctx.CNFVar(v, i, int(query.BOT)))
		}
	}

	return clause, nil
}
