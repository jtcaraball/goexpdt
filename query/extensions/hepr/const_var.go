package hepr

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstVar is the constant-variable version of the Higher or Equal Preference
// Rank extension.
type ConstVar struct {
	I1         query.QConst
	I2         query.QVar
	Preference []int
	// FeaturePreferenceRankVarGen returns a variable used to encode the
	// comparative ranking between query variables v1 and v2 under the
	// preference encoded in p.
	FeaturePreferenceRankVarGen func(p, v1, v2 query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query constant d.I1
// has an equal or higher rank than the query variable d.I2 according to the
// feature preference d.Preference.
func (d ConstVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	if d.FeaturePreferenceRankVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	if err := validatePref(d.Preference, ctx.Dim()); err != nil {
		return cnf.CNF{}, err
	}

	dim := ctx.Dim()
	sc, _ := ctx.ScopeConst(d.I1)
	sv := ctx.ScopeVar(d.I2)
	cvp := d.FeaturePreferenceRankVarGen(
		prefVar(d.Preference),
		query.QVar(sc.AsString()),
		sv,
	)

	if err := query.ValidateConstsDim(dim, sc); err != nil {
		return cnf.CNF{}, err
	}

	ncnf := cnf.CNF{}
	plen := len(d.Preference)
	idx := d.Preference[plen-1]

	if sc.Val[idx] != query.BOT {
		ncnf = ncnf.AppendConsistency(cnf.Clause{ctx.CNFVar(cvp, 0, plen-1)})
	} else {
		ncnf = ncnf.AppendConsistency(
			cnf.Clause{
				-ctx.CNFVar(cvp, 0, plen-1),
				ctx.CNFVar(sv, idx, int(query.BOT)),
			},
			cnf.Clause{
				-ctx.CNFVar(sv, idx, int(query.BOT)),
				ctx.CNFVar(cvp, 0, plen-1),
			},
		)
	}

	for i := plen - 2; i >= 0; i-- {
		idx = d.Preference[i]

		if sc.Val[idx] != query.BOT {
			ncnf = ncnf.AppendConsistency(
				cnf.Clause{
					-ctx.CNFVar(cvp, 0, i),
					ctx.CNFVar(sv, idx, int(query.BOT)),
					ctx.CNFVar(cvp, 0, i+1),
				},
				cnf.Clause{
					-ctx.CNFVar(sv, idx, int(query.BOT)),
					ctx.CNFVar(cvp, 0, i),
				},
				cnf.Clause{
					-ctx.CNFVar(cvp, 0, i+1),
					ctx.CNFVar(cvp, 0, i),
				},
			)
			continue
		}

		ncnf = ncnf.AppendConsistency(
			cnf.Clause{
				-ctx.CNFVar(sv, idx, int(query.BOT)),
				-ctx.CNFVar(cvp, 0, i+1),
				ctx.CNFVar(cvp, 0, i),
			},
			cnf.Clause{
				-ctx.CNFVar(cvp, 0, i),
				ctx.CNFVar(sv, idx, int(query.BOT)),
			},
			cnf.Clause{
				-ctx.CNFVar(cvp, 0, i),
				ctx.CNFVar(cvp, 0, i+1),
			},
		)
	}

	return ncnf.AppendSemantics(cnf.Clause{ctx.CNFVar(cvp, 0, 0)}), nil
}