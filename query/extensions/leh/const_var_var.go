package leh

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstVarVar is the constant-variable-variable version of the LEH extension.
type ConstVarVar struct {
	I1 query.QConst
	I2 query.QVar
	I3 query.QVar
	// HammingDistanceVarGen returns a variable generated from v1 and v2 that
	// will be used to encode the hamming distance between v1 and v2. The
	// resulting query variable should be the same regardless of the order in
	// which v1 and v2 are passed.
	HammingDistanceVarGen func(v1, v2 query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query constant l.I1
// and query variables l.I2 and l.I3 are full and the distance hamming distance
// between l.I1 and l.I2 is smaller that the hamming distance between l.I1 and
// l.I3.
func (l ConstVarVar) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}
	if l.HammingDistanceVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	dim := ctx.Dim()
	sc, _ := ctx.ScopeConst(l.I1)
	sv1 := ctx.ScopeVar(l.I2)
	sv2 := ctx.ScopeVar(l.I3)
	v1chd := l.HammingDistanceVarGen(sv1, query.QVar(sc.AsString()))
	v2chd := l.HammingDistanceVarGen(sv2, query.QVar(sc.AsString()))

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
	v2chdClauses, err := hammingDistVC(sv2, sc, v2chd, ctx)
	if err != nil {
		return cnf.CNF{}, err
	}

	ncnf = ncnf.AppendConsistency(v1chdClauses...)
	ncnf = ncnf.AppendConsistency(v2chdClauses...)

	for i := 1; i <= dim; i++ {
		for j := 0; j < i; j++ {
			ncnf = ncnf.AppendSemantics(
				cnf.Clause{
					-ctx.CNFVar(v1chd, dim-1, i),
					-ctx.CNFVar(v2chd, dim-1, j),
				},
			)
		}
	}

	return ncnf, nil
}
