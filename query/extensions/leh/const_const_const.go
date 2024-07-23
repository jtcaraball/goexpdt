package leh

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ConstConstConst is the constant-constant-constant version of the LEH
// extension.
type ConstConstConst struct {
	I1 query.QConst
	I2 query.QConst
	I3 query.QConst
}

// Encoding returns a CNF that is true if and only if the query constants l.I1,
// l.I2 and l.I3 are full and the distance hamming distance between  l.I1 and
// l.I2 is smaller that the hamming distance between l.I1 and l.I3.
func (l ConstConstConst) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc1, _ := ctx.ScopeConst(l.I1)
	sc2, _ := ctx.ScopeConst(l.I2)
	sc3, _ := ctx.ScopeConst(l.I3)

	if err := query.ValidateConstsDim(
		ctx.Dim(),
		sc1,
		sc2,
		sc3,
	); err != nil {
		return cnf.CNF{}, err
	}

	if !sc1.IsFull() || !sc2.IsFull() || !sc3.IsFull() {
		return cnf.FalseCNF, nil
	}

	hd12, err := hammingDistCC(sc1, sc2)
	if err != nil {
		return cnf.CNF{}, nil
	}
	hd13, err := hammingDistCC(sc1, sc3)
	if err != nil {
		return cnf.CNF{}, nil
	}

	if hd12 > hd13 {
		return cnf.FalseCNF, nil
	}

	return cnf.TrueCNF, nil
}
