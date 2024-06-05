package lel

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstConst is the constant-constant version of the Less or Equal Level
// predicate.
type ConstConst struct {
	I1 query.QConst
	I2 query.QConst
}

// Ecoding returns a CNF that is true if and only if the query constant l.I1
// has more or equal amount of BOT valued features than the query constant
// l.I2.
func (l ConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc1, _ := ctx.ScopeConst(l.I1)
	sc2, _ := ctx.ScopeConst(l.I2)

	if err := query.ValidateConstsDim(
		ctx.Dim(),
		sc1,
		sc2,
	); err != nil {
		return cnf.CNF{}, err
	}

	if sc1.BotCount() < sc2.BotCount() {
		return cnf.FalseCNF, nil
	}

	return cnf.TrueCNF, nil
}
