package leh

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarVarConst is the variable-variable-constant version of the LEH extension.
type VarVarConst struct {
	I1 query.QVar
	I2 query.QVar
	I3 query.QConst
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

// Return CNF encoding of component.
func (l VarVarConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}
	if l.HammingDistanceVarGen == nil || l.EqualFeatureVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	dim := ctx.Dim()
	sv1 := ctx.ScopeVar(l.I1)
	sv2 := ctx.ScopeVar(l.I2)
	sc, _ := ctx.ScopeConst(l.I3)
	v1chd := l.HammingDistanceVarGen(sv1, query.QVar(sc.AsString()))
	vvhd := l.HammingDistanceVarGen(sv1, sv2)
	vvef := l.EqualFeatureVarGen(sv1, sv2)

	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	if !sc.IsFull() {
		return cnf.FalseCNF, nil
	}

	ncnf := cnf.CNF{}

	ncnf = ncnf.AppendSemantics(varFullClauses(sv1, ctx)...)
	ncnf = ncnf.AppendSemantics(varFullClauses(sv2, ctx)...)

	v1chdClauses, err := hammingDistVC(sv1, sc, v1chd, ctx)
	if err != nil {
		return cnf.CNF{}, err
	}
	ncnf = ncnf.AppendConsistency(fullVarEqualClauses(sv1, sv2, vvef, ctx)...)
	ncnf = ncnf.AppendConsistency(hammingDistVV(sv1, sv2, vvhd, vvef, ctx)...)
	ncnf = ncnf.AppendConsistency(v1chdClauses...)

	leqhd := []cnf.Clause{}
	for i := 1; i <= dim; i++ {
		for j := 0; j < i; j++ {
			leqhd = append(
				leqhd,
				cnf.Clause{
					-ctx.CNFVar(vvhd, dim-1, i),
					-ctx.CNFVar(v1chd, dim-1, j),
				},
			)
		}
	}

	return ncnf.AppendSemantics(leqhd...), nil
}
