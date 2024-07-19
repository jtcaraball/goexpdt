package lel

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstVar is the constant-variable version of the Less or Equal Level
// predicate.
type ConstVar struct {
	I1 query.QConst
	I2 query.QVar
	// CountVarGen returns a variable generated from v that will be used to
	// encode the amount of features equal to bot in v.
	CountVarGen func(v query.QVar) query.QVar
}

// Ecoding returns a CNF that is true if and only if the query constant l.I1
// has more or equal amount of BOT valued features than the query variable
// l.I2.
func (l ConstVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv := ctx.ScopeVar(l.I2)
	sc, _ := ctx.ScopeConst(l.I1)
	svCount := l.CountVarGen(sv)

	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	ncnf := cnf.CNF{}
	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv, svCount, ctx)...)

	for i := sc.BotCount() + 1; i < ctx.Dim()+1; i++ {
		ncnf = ncnf.AppendSemantics(
			cnf.Clause{-ctx.CNFVar(svCount, ctx.Dim()-1, i)},
		)
	}

	return ncnf, nil
}
