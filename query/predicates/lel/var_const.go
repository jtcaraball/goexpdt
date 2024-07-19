package lel

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarConst is the variable-constant version of the Less or Equal Level
// predicate.
type VarConst struct {
	I1 query.QVar
	I2 query.QConst
	// CountVarGen returns a variable generated from v that will be used to
	// encode the amount of features equal to bot in v.
	CountVarGen func(v query.QVar) query.QVar
}

// Ecoding returns a CNF that is true if and only if the query variable l.I1
// has more or equal amount of BOT valued features than the query constant
// l.I2.
func (l VarConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(l.I1)
	sc, _ := ctx.ScopeConst(l.I2)
	svCount := l.CountVarGen(sv)

	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	ncnf := cnf.CNF{}
	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv, svCount, ctx)...)

	for i := 0; i < sc.BotCount(); i++ {
		ncnf = ncnf.AppendSemantics(
			cnf.Clause{-ctx.CNFVar(svCount, ctx.Dim()-1, i)},
		)
	}

	return ncnf, nil
}
