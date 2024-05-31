package logop

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Or represents the logical operator OR.
type Or struct {
	// Child1 corresponds to a sub-query that implements the LogOpChild
	// interface.
	Child1 LogOpChild
	// Child2 corresponds to a sub-query that implements the LogOpChild
	// interface.
	Child2 LogOpChild
}

// Encoding returns a CNF formula equivalent to the disjunction of the CNF
// formulas corresponding to its children.
func (o Or) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("Or: %w", err)
		}
	}()

	if o.Child1 == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child (1)")
	}
	if o.Child2 == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child (2)")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = o.buildEncoding(ctx)

	return ncnf, err
}

func (o Or) buildEncoding(ctx query.QContext) (cnf.CNF, error) {
	// De Morgan's law
	// Encode both children
	cnf1, err := o.Child1.Encoding(ctx)
	if err != nil {
		return cnf.CNF{}, fmt.Errorf("Child 1: %w", err)
	}
	cnf2, err := o.Child2.Encoding(ctx)
	if err != nil {
		return cnf.CNF{}, fmt.Errorf("Child 2: %w", err)
	}

	// Negate both children
	cnf1 = cnf1.Negate(ctx.TopV())
	tv := cnf1.TopV()
	ctx.UpdateTopV(tv) // TopV could have increases while negating

	cnf2 = cnf2.Negate(ctx.TopV())
	tv = cnf2.TopV()
	ctx.UpdateTopV(tv) // TopV could have increases while negating

	// Logical and & final negation
	ncnf := cnf1.Conjunction(cnf2).Negate(ctx.TopV())
	ctx.UpdateTopV(ncnf.TopV()) // TopV could have increases while negating

	return ncnf, nil
}
