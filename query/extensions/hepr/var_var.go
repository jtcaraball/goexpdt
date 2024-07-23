package hepr

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarVar is the variable-variable version of the Higher or Equal Preference
// Rank extension.
type VarVar struct {
	I1         query.QVar
	I2         query.QVar
	Preference []int
	// FeaturePreferenceRankVarGen returns a variable used to encode the
	// comparative ranking between query variables v1 and v2 under the
	// preference encoded in p.
	FeaturePreferenceRankVarGen func(p, v1, v2 query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable d.I1
// has an equal or higher rank than d.I2 according to the feature preference
// d.Preference.
func (d VarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	if d.FeaturePreferenceRankVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	if err := validatePref(d.Preference, ctx.Dim()); err != nil {
		return cnf.CNF{}, err
	}

	sv1 := ctx.ScopeVar(d.I1)
	sv2 := ctx.ScopeVar(d.I2)
	vvp := d.FeaturePreferenceRankVarGen(prefVar(d.Preference), sv1, sv2)

	ncnf := cnf.CNF{}
	plen := len(d.Preference)
	idx := d.Preference[plen-1]

	ncnf = ncnf.AppendConsistency(
		cnf.Clause{
			-ctx.CNFVar(vvp, 0, plen-1),
			-ctx.CNFVar(sv1, idx, int(query.BOT)),
			ctx.CNFVar(sv2, idx, int(query.BOT)),
		},
		cnf.Clause{
			ctx.CNFVar(sv1, idx, int(query.BOT)),
			ctx.CNFVar(vvp, 0, plen-1),
		},
		cnf.Clause{
			-ctx.CNFVar(sv2, idx, int(query.BOT)),
			ctx.CNFVar(vvp, 0, plen-1),
		},
	)

	for i := plen - 2; i >= 0; i-- {
		idx = d.Preference[i]
		ncnf = ncnf.AppendConsistency(
			// ->
			cnf.Clause{
				-ctx.CNFVar(sv1, idx, int(query.BOT)),
				ctx.CNFVar(sv2, idx, int(query.BOT)),
				-ctx.CNFVar(vvp, 0, i),
			},
			cnf.Clause{
				-ctx.CNFVar(sv1, idx, int(query.BOT)),
				-ctx.CNFVar(vvp, 0, i),
				ctx.CNFVar(vvp, 0, i+1),
			},
			cnf.Clause{
				ctx.CNFVar(sv2, idx, int(query.BOT)),
				-ctx.CNFVar(vvp, 0, i),
				ctx.CNFVar(vvp, 0, i+1),
			},
			// <-
			cnf.Clause{
				ctx.CNFVar(sv1, idx, int(query.BOT)),
				-ctx.CNFVar(sv2, idx, int(query.BOT)),
				ctx.CNFVar(vvp, 0, i),
			},
			cnf.Clause{
				ctx.CNFVar(sv1, idx, int(query.BOT)),
				ctx.CNFVar(vvp, 0, i),
				-ctx.CNFVar(vvp, 0, i+1),
			},
			cnf.Clause{
				-ctx.CNFVar(sv2, idx, int(query.BOT)),
				ctx.CNFVar(vvp, 0, i),
				-ctx.CNFVar(vvp, 0, i+1),
			},
		)
	}

	return ncnf.AppendSemantics(cnf.Clause{ctx.CNFVar(vvp, 0, 0)}), nil
}
