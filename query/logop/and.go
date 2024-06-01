package logop

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// And represents the logical operator AND.
type And struct {
	// Q1 corresponds to a sub-query that implements the LogOpQ
	// interface.
	Q1 LogOpQ
	// Q2 corresponds to a sub-query that implements the LogOpQ
	// interface.
	Q2 LogOpQ
}

// Encoding returns a CNF formula equivalent to the conjunction of the CNF
// formulas corresponding to its children.
func (a And) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("And: %w", err)
		}
	}()

	if a.Q1 == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child (1)")
	}
	if a.Q2 == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child (2)")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = a.buildEncoding(ctx)

	return ncnf, err
}

func (a And) buildEncoding(ctx query.QContext) (cnf.CNF, error) {
	cnf1, err := a.Q1.Encoding(ctx)
	if err != nil {
		return cnf.CNF{}, fmt.Errorf("Q 1: %w", err)
	}

	cnf2, err := a.Q2.Encoding(ctx)
	if err != nil {
		return cnf.CNF{}, fmt.Errorf("Q 2: %w", err)
	}

	return cnf1.Conjunction(cnf2), nil
}
