package pred

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// SubsumptionConstConst is the constant-constant version of the Subsumption
// predicate.
type SubsumptionConstConst struct {
	I1 query.QConst
	I2 query.QConst
}

// Encoding returns a CNF that is true if and only if the query constant s.I1
// is subsumed by the query constant s.I2.
func (s SubsumptionConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc1, _ := ctx.ScopeConst(s.I1)
	sc2, _ := ctx.ScopeConst(s.I2)

	if err := query.ValidateConstsDim(
		ctx.Dim(),
		sc1,
		sc2,
	); err != nil {
		return cnf.CNF{}, err
	}

	for i, ft := range sc1.Val {
		if ft != query.BOT && ft != sc2.Val[i] {
			return cnf.FalseCNF, nil
		}
	}

	return cnf.TrueCNF, nil
}
