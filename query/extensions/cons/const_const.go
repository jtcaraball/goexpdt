package cons

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstConst is the constant-constant version of the Consistent extension.
type ConstConst struct {
	I1 query.QConst
	I2 query.QConst
}

// Encoding returns a CNF that is true if and only if the query constants c.I1
// and c.I2 have the same values in every feature where both are different than
// bottom.
func (c ConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc1, _ := ctx.ScopeConst(c.I1)
	sc2, _ := ctx.ScopeConst(c.I2)

	if err := query.ValidateConstsDim(
		ctx.Dim(),
		sc1,
		sc2,
	); err != nil {
		return cnf.CNF{}, err
	}

	for i := 0; i < ctx.Dim(); i++ {
		if sc1.Val[i] != query.BOT &&
			sc2.Val[i] != query.BOT &&
			sc1.Val[i] != sc2.Val[i] {
			return cnf.FalseCNF, nil
		}
	}

	return cnf.TrueCNF, nil
}
