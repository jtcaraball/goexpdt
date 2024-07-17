package leh

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarVarVar is the variable-variable-variable version of the LEH extension.
type VarVarVar struct {
	I1 query.QVar
	I2 query.QVar
	I3 query.QVar
	// HammingDistanceVarGen returns a variable generated from v1 and v2 that
	// will be used to encode the hamming distance between v1 and v2. The
	// resulting query variable should be the same regardless of the order in
	// which v1 and v2 are passed.
	HammingDistanceVarGen func(v1, v2 query.QVar) query.QVar
	// EqualFeatureVarGen returns a variable generated from v1 and v2 that will
	// be used to encode equality between the features of v1 and v2. The
	// resulting query variable should be the same regardless of the order in
	// which v1 and v2 are passed.
	EqualFeatureVarGen func(v1, v2 query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query variables l.I1,
// l.I2 and l.I3 are full and the distance hamming distance between  l.I1 and
// l.I2 is smaller that the hamming distance between l.I1 and l.I3.
func (l VarVarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}
	if l.HammingDistanceVarGen == nil || l.EqualFeatureVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	dim := ctx.Dim()
	sv1 := ctx.ScopeVar(l.I1)
	sv2 := ctx.ScopeVar(l.I2)
	sv3 := ctx.ScopeVar(l.I3)
	v1v2dh := l.HammingDistanceVarGen(sv1, sv2)
	v1v3dh := l.HammingDistanceVarGen(sv1, sv3)
	v1v2ef := l.EqualFeatureVarGen(sv1, sv2)
	v1v3ef := l.EqualFeatureVarGen(sv1, sv3)

	ncnf := cnf.CNF{}

	ncnf = ncnf.AppendSemantics(varFullClauses(sv1, ctx)...)
	ncnf = ncnf.AppendSemantics(varFullClauses(sv2, ctx)...)
	ncnf = ncnf.AppendSemantics(varFullClauses(sv3, ctx)...)

	ncnf = ncnf.AppendConsistency(fullVarEqualClauses(sv1, sv2, v1v2ef, ctx)...)
	ncnf = ncnf.AppendConsistency(fullVarEqualClauses(sv1, sv3, v1v3ef, ctx)...)

	ncnf = ncnf.AppendConsistency(
		hammingDistVV(sv1, sv2, v1v2dh, v1v2ef, ctx)...,
	)
	ncnf = ncnf.AppendConsistency(
		hammingDistVV(sv2, sv3, v1v3dh, v1v3ef, ctx)...,
	)

	leqhd := []cnf.Clause{}
	for i := 1; i <= dim; i++ {
		for j := 0; j < i; j++ {
			leqhd = append(
				leqhd,
				cnf.Clause{
					-ctx.CNFVar(v1v2dh, dim-1, i),
					-ctx.CNFVar(v1v3dh, dim-1, j),
				},
			)
		}
	}

	return ncnf.AppendSemantics(leqhd...), nil
}
