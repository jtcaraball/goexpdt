package logop

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Not represents the logical operator NOT.
type Not struct {
	// Q corresponds to a sub-query that implements the LogOpQ
	// interface.
	Q LogOpQ
}

// Encoding returns a CNF formula equivalent to the negations of its children's
// CNF formula.
func (n Not) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Not: %w", err)
		}
	}()

	if n.Q == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = n.buildEncoding(ctx)

	return ncnf, err
}

func (n Not) buildEncoding(ctx query.QContext) (cnf.CNF, error) {
	ncnf, err := n.Q.Encoding(ctx)
	if err != nil {
		return cnf.CNF{}, err
	}

	ncnf = ncnf.Negate(ctx.TopV())
	ctx.UpdateTopV(ncnf.TopV())

	return ncnf, nil
}
