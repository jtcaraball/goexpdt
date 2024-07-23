package full

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Const is the constant version of the Full extension.
type Const struct {
	I query.QConst
}

// Encoding returns a CNF that is true if and only if the query constant f.I
// has no features with bottom value.
func (f Const) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc, _ := ctx.ScopeConst(f.I)
	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	for _, ft := range sc.Val {
		if ft == query.BOT {
			return cnf.FalseCNF, nil
		}
	}

	return cnf.TrueCNF, nil
}
