package compute

import (
	"errors"

	"github.com/jtcaraball/goexpdt/query"
)

// QDTFCombinator represents a query in the third level of the Q-DT-Foil logic,
// that is, a boolean combination of sub-queries in the logic that do not share
// common variables.
type QDTFCombinator interface {
	// Sat returns true if and only if the query.
	Sat(ctx query.QContext, solver Solver) (bool, error)
}

// AndCombinator is satisfiable if and only if all its elements are.
type AndCombinator []QDTFCombinator

func (ac AndCombinator) Sat(ctx query.QContext, solver Solver) (bool, error) {
	for _, elem := range ac {
		sat, err := elem.Sat(ctx, solver)
		if err != nil {
			return false, err
		}

		if !sat {
			return false, nil
		}
	}

	return true, nil
}

// OrCombinator is satisfiable if any one of its elements is.
type OrCombinator []QDTFCombinator

func (oc OrCombinator) Sat(ctx query.QContext, solver Solver) (bool, error) {
	for _, elem := range oc {
		sat, err := elem.Sat(ctx, solver)
		if err != nil {
			return false, err
		}

		if sat {
			return true, nil
		}
	}

	return false, nil
}

// QDTFAtom represents a single atomic query in the third level of the
// Q-DT-Foil logic. As SAT solvers can not directly answer over queries using
// universal quantifiers it is necessary to negate them and re-interpret the
// solver's answer.
type QDTFAtom struct {
	// Query in the third level of the Q-DT-Foil.
	Query Encodable
	// Negated is true if the original query used universal quantifiers and
	// thus the answer of Sat should be negated.
	Negated bool
}

func (a QDTFAtom) Sat(
	ctx query.QContext,
	solver Solver,
) (bool, error) {
	err := solver.Step(a.Query, ctx)
	if err != nil {
		if errors.Is(err, UnsatError) {
			return a.Negated, nil
		}

		return false, err
	}

	return !a.Negated, nil
}
