package logop

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// WithVar represents a sub-query that makes use of an instance variable.
type WithVar struct {
	// Instance corresponds to a partial instance variable.
	Instance query.Var
	// Child corresponds to a sub-query that implements the LogOpChild and that
	// is expected to make use of Instance.
	Child LogOpChild
}

// Encoding returns the CNF formula encoding the consistency clauses of its
// Instance variable and its Child CNF formula.
func (w WithVar) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("WithVar: %w", err)
		}
	}()

	if w.Child == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = w.buildEncoding(ctx)
	if err != nil {
		return cnf.CNF{}, err
	}
	return ncnf, nil
}

func (w WithVar) buildEncoding(
	ctx query.QContext,
) (cnf.CNF, error) {
	icnf := w.encodeInstance(ctx)

	ccnf, err := w.Child.Encoding(ctx)
	if err != nil {
		return cnf.CNF{}, fmt.Errorf("WithVar: %w", err)
	}

	return icnf.Conjunction(ccnf), nil
}

// encodeInstance returns a CNF with consistency clauses enforcing a consistent
// valuation of w.Instance.
func (w WithVar) encodeInstance(ctx query.QContext) cnf.CNF {
	ncnf := cnf.CNF{}

	v := ctx.ScopeVar(w.Instance)

	// Lets not add consistency clauses twice.
	if ctx.VarExists(string(v), 0, int(query.BOT)) {
		return ncnf
	}

	// Every feature must have at least one value
	reqAllFeats := []cnf.Clause{}
	for i := 0; i < ctx.Dim(); i++ {
		reqAllFeats = append(
			reqAllFeats,
			cnf.Clause{
				ctx.Var(string(v), i, int(query.ZERO)),
				ctx.Var(string(v), i, int(query.ONE)),
				ctx.Var(string(v), i, int(query.BOT)),
			},
		)
	}

	ncnf = ncnf.AppendConsistency(reqAllFeats...)
	// Every feature must have one and only one value
	for i := 0; i < ctx.Dim(); i++ {
		reqOnePerFeat := []cnf.Clause{
			{
				-ctx.Var(string(v), i, int(query.ZERO)),
				-ctx.Var(string(v), i, int(query.ONE)),
			},
			{
				-ctx.Var(string(v), i, int(query.ZERO)),
				-ctx.Var(string(v), i, int(query.BOT)),
			},
			{
				-ctx.Var(string(v), i, int(query.ONE)),
				-ctx.Var(string(v), i, int(query.BOT)),
			},
		}
		ncnf = ncnf.AppendConsistency(reqOnePerFeat...)
	}

	return ncnf
}
