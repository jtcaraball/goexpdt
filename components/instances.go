package components

import (
	"errors"
	"fmt"
	"slices"
	"stratifoiled/cnf"
)

// =========================== //
//         INTERFACES          //
// =========================== //

type ConstInstance interface {
	Scoped(ctx *Context) (Const, error)
}

// =========================== //
//           STRUCTS           //
// =========================== //

type Var string

type GuardedConst string

type Const []featV

type featV struct {
	val int
}

// =========================== //
//          VARIABLES          //
// =========================== //

var (
	ZERO = featV{val: 0}
	ONE = featV{val: 1}
	BOT = featV{val: 2}
	FeatValues = []featV{ZERO, ONE, BOT}
)

// =========================== //
//           METHODS           //
// =========================== //

// Return new var instance with name equal to string passed.
func NewVar(name string) Var {
	var v Var = Var(name)
	return v
}

// Encode v's consistency clauses to cnf and add necesary variables to context.
func (v Var) Encoding(ctx *Context) *cnf.CNF {
	nCNF := &cnf.CNF{}
	// Add consistency clauses
	// Every feature must have at least one value
	reqAllFeats := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		clause := []int{}
		for _, s := range FeatValues {
			clause = append(clause, ctx.Var(string(v), i, s.Val()))
		}
		reqAllFeats = append(reqAllFeats, clause)
	}
	nCNF.ExtendConsistency(reqAllFeats)
	// Every feature must have one and only one value
	for i := 0; i < ctx.Dimension; i++ {
		reqOnePerFeat := [][]int{
			{
				-ctx.Var(string(v), i, ZERO.Val()),
				-ctx.Var(string(v), i, ONE.Val()),
			},
			{
				-ctx.Var(string(v), i, ZERO.Val()),
				-ctx.Var(string(v), i, BOT.Val()),
			},
			{
				-ctx.Var(string(v), i, ONE.Val()),
				-ctx.Var(string(v), i, BOT.Val()),
			},
		}
		nCNF.ExtendConsistency(reqOnePerFeat)
	}
	return nCNF
}

// Return corresponding scoped var. If no scope is found in guard returns the
// method caller.
func (v Var) Scoped(ctx *Context) Var {
	rVar := ""
	for _, guard := range ctx.Guards {
		if slices.Contains[[]string](guard.InScope, string(v)) {
			rVar += guard.Rep
		}
	}
	if rVar != "" {
		return Var(string(v) + rVar)
	}
	return v
}

// Return caller always.
func (c Const) Scoped(ctx *Context) (Const, error) {
	return c, nil
}

// Return corresponding const from list of guards.
func (gc GuardedConst) Scoped(ctx *Context) (Const, error) {
	for _, guard := range ctx.Guards {
		if guard.Target == string(gc) {
			return guard.Value, nil
		}
	}
	return nil, errors.New(
		fmt.Sprintf("No guard with target '%s'", string(gc)),
	)
}

// Return featV value.
func (f featV) Val() int {
	return f.val
}
