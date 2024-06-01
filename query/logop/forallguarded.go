package logop

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ForAllGuarded represents a FOR ALL guarded quantifier.
type ForAllGuarded struct {
	// I corresponds to the constant that will be used to materialize
	// the instances that correspond to the ctx's model's nodes. Its ID will be
	// used for scope setting.
	I query.Const
	// Q corresponds to a sub-query that implements the LogOpQ and that
	// is expected to make use of I.
	Q LogOpQ
}

// Encoding returns the CNF formula equivalent to the conjunction all the
// possible CNF formulas of its Q resulting from instantiating every value
// of I in the ctx's model.
func (f ForAllGuarded) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ForAllGuarded: %w", err)
		}
	}()

	if f.Q == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = f.buildEncoding(ctx)

	return ncnf, err
}

func (f ForAllGuarded) buildEncoding(ctx query.QContext) (cnf.CNF, error) {
	ncnf := cnf.CNF{}

	ctx.AddScope(f.I.ID)

	ncs := ctx.NodesConsts()

	for i, nc := range ncs {
		if err := ctx.SetScope(i, nc.Val); err != nil {
			return cnf.CNF{}, err
		}

		icnf, err := f.Q.Encoding(ctx)
		if err != nil {
			return cnf.CNF{}, err
		}

		ncnf = ncnf.Conjunction(icnf)
	}

	if err := ctx.PopScope(); err != nil {
		return cnf.CNF{}, err
	}

	return ncnf, nil
}
