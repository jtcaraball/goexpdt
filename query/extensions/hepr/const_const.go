package hepr

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstConst is the constant-constant version of the Higher or Equal
// Preference Rank extension.
type ConstConst struct {
	I1         query.QConst
	I2         query.QConst
	Preference []int
}

// Encoding returns a CNF that is true if and only if the query constants d.I1
// has an equal or higher rank than d.I2 according to the feature preference
// d.Preference.
func (d ConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	if err := validatePref(d.Preference, ctx.Dim()); err != nil {
		return cnf.CNF{}, err
	}

	sc1, _ := ctx.ScopeConst(d.I1)
	sc2, _ := ctx.ScopeConst(d.I2)
	if err := query.ValidateConstsDim(ctx.Dim(), sc1, sc2); err != nil {
		return cnf.CNF{}, err
	}

	if len(d.Preference) == 0 {
		return cnf.TrueCNF, nil
	}

	for _, idx := range d.Preference {
		if sc1.Val[idx] != query.BOT && sc2.Val[idx] == query.BOT {
			return cnf.TrueCNF, nil
		}
		if sc1.Val[idx] == query.BOT && sc2.Val[idx] != query.BOT {
			return cnf.FalseCNF, nil
		}
	}

	return cnf.TrueCNF, nil
}
