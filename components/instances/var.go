package instances

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type Var string

// =========================== //
//           METHODS           //
// =========================== //

// Return new var instance with name equal to string passed.
func NewVar(name string) Var {
	var v Var = Var(name)
	return v
}

// Encode v's consistency clauses to cnf and add necesary variables to context.
func (v Var) Encoding(ctx *components.Context) *cnf.CNF {
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
