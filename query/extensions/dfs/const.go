package dfs

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Const is the constant version of the DFS extension.
type Const struct {
	I query.QConst
}

// Encoding returns a CNF that is true if and only if the query constant d.I
// is such that all its completions have the same Model evaluation.
func (d Const) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc, _ := ctx.ScopeConst(d.I)
	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	pleaf, nleaf, err := leafsAsConsts(ctx.Dim(), ctx.Nodes())
	if err != nil {
		return cnf.CNF{}, err
	}

	if len(pleaf) == 0 && len(nleaf) == 0 {
		return cnf.CNF{}, errors.New("Invalid encoding on empty model")
	}

	for _, p := range pleaf {
		for _, n := range nleaf {
			cons, err := leafCompare(sc, p, n)
			if err != nil {
				return cnf.CNF{}, err
			}

			if !cons {
				return cnf.FalseCNF, nil
			}
		}
	}

	return cnf.TrueCNF, nil
}

func leafCompare(c, l1, l2 query.QConst) (bool, error) {
	if len(c.Val) != len(l1.Val) || len(c.Val) != len(l2.Val) {
		return false,
			errors.New("Invalid comparison of constants of different dim")
	}

	for i, ft := range c.Val {
		if ft != query.BOT &&
			l1.Val[i] != query.BOT &&
			l2.Val[i] != query.BOT &&
			l1.Val[i] != l2.Val[i] {
			return true, nil
		}
	}

	return false, nil
}
