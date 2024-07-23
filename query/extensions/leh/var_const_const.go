package leh

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// VarConstConst is the variable-constant-constant version of the LEH
// extension.
type VarConstConst struct {
	I1 query.QVar
	I2 query.QConst
	I3 query.QConst
	// HammingDistanceVarGen returns a variable generated from v1 and v2 that
	// will be used to encode the hamming distance between v1 and v2. The
	// resulting query variable should be the same regardless of the order in
	// which v1 and v2 are passed.
	HammingDistanceVarGen func(v1, v2 query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query constants l.I2
// and l.I3 and the query variable l.I1 are full and the distance hamming
// distance between  l.I1 and l.I2 is smaller that the hamming distance between
// l.I1 and l.I3.
func (l VarConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}
	if l.HammingDistanceVarGen == nil {
		return cnf.CNF{}, errors.New("Invalid nil var generation function")
	}

	dim := ctx.Dim()
	sc1, _ := ctx.ScopeConst(l.I2)
	sc2, _ := ctx.ScopeConst(l.I3)
	sv := ctx.ScopeVar(l.I1)
	vc1hd := l.HammingDistanceVarGen(sv, query.QVar(sc1.AsString()))
	vc2hd := l.HammingDistanceVarGen(sv, query.QVar(sc2.AsString()))

	if err := query.ValidateConstsDim(dim, sc1, sc2); err != nil {
		return cnf.CNF{}, err
	}

	if !sc1.IsFull() || !sc2.IsFull() {
		return cnf.FalseCNF, nil
	}

	ncnf := cnf.FromClauses(varFullClauses(sv, ctx))

	hd1Clauses, err := hammingDistVC(sv, sc1, vc1hd, ctx)
	if err != nil {
		return cnf.CNF{}, err
	}
	hd2Clauses, err := hammingDistVC(sv, sc2, vc2hd, ctx)
	if err != nil {
		return cnf.CNF{}, err
	}

	ncnf = ncnf.AppendConsistency(hd1Clauses...)
	ncnf = ncnf.AppendConsistency(hd2Clauses...)

	for i := 1; i <= dim; i++ {
		for j := 0; j < i; j++ {
			ncnf = ncnf.AppendSemantics(
				cnf.Clause{
					-ctx.CNFVar(vc1hd, dim-1, i),
					-ctx.CNFVar(vc2hd, dim-1, j),
				},
			)
		}
	}

	return ncnf, nil
}
