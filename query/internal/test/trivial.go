package test

import (
	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

type Trivial bool

func (t Trivial) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if t {
		return cnf.TrueCNF, nil
	}
	return cnf.FalseCNF, nil
}
