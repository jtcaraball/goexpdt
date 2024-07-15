package lel

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarVar is the variable-variable version of the Less or Equal Level
// predicate.
type VarVar struct {
	I1 query.QVar
	I2 query.QVar
	// CountVarGen returns a variable generated from v that will be used to
	// encode the amount of features equal to bot in v.
	CountVarGen func(v query.QVar) query.QVar
}

// Ecoding returns a CNF that is true if and only if the query variable l.I1
// has more or equal amount of BOT valued features than the query variable
// l.I2.
func (l VarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sv1 := ctx.ScopeVar(l.I1)
	sv2 := ctx.ScopeVar(l.I2)
	svCount1 := l.CountVarGen(sv1)
	svCount2 := l.CountVarGen(sv2)

	ncnf := cnf.CNF{}

	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv1, svCount1, ctx)...)
	ncnf = ncnf.AppendConsistency(varBotCountClauses(sv2, svCount2, ctx)...)

	// If we see a number of bots in x then we must see less or equal on y
	lcl := []cnf.Clause{}

	var i, j int
	for i = 0; i < ctx.Dim(); i++ {
		cl := cnf.Clause{-ctx.CNFVar(svCount1, ctx.Dim()-1, i)}
		for j = 0; j <= i; j++ {
			cl = append(
				cl,
				ctx.CNFVar(svCount2, ctx.Dim()-1, j),
			)
		}
		lcl = append(lcl, cl)
	}

	ncnf = ncnf.AppendSemantics(lcl...)

	return ncnf, nil
}
