package main

// In this example we evaluate the query Q(x): x is partial instance such that
// x is a DFS (Determinant Feature Set) of a Decision Tree T and for every
// instance y if y has less defined features than x then y is not a DFS of T.
// In particular we will evaluate Q(c) where c = (_, 0, _) over T defined as:
//
// T: f0 -0-> f1 -0-> f2 -0-> false
//      |       |       |
//      |       |        -1-> false
//      |       |
//      |        -1-> f2 -0-> true
//      |               |
//      |                -1-> true
//      |
//       -1-> f2 -0-> f1 -0-> false
//              |       |
//              |        -1-> true
//              |
//               -1-> f1 -0-> false
//                      |
//                       -1-> true
//
// A DFS of a Decision Tree T is defined as a partial instance x such all
// possible completions of x have the same classification. Thus Q(x) is true
// if and only if x is the one of the partial instances with the least defined
// features that are a DFS of T.
//
// In Q-DT-Foil this query is understood as follows: Given an partial instance
// c determine the truth value of Q(c). Note that it can be expressed in the
// logic because it is boolean combination of formulas with a single
// non-guarded quantifier.
//
// The following code evaluates Q(c).

import (
	"fmt"

	"github.com/jtcaraball/goexpdt/compute"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/dfs"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/lel"
)

// First we define a decision tree struct that implements the query.Model
// interface.
type evalDTree struct {
	dim   int
	nodes []query.Node
}

func (t evalDTree) Dim() int {
	return t.dim
}

func (t evalDTree) Nodes() []query.Node {
	return t.nodes
}

// Then we define a function to generate, given a variable v, a new variable
// used to encode the count of bottom variables.
func botCountVarGen(v query.QVar) query.QVar {
	return query.QVar("cbot" + string(v))
}

func EvaluationExample() {
	// We instantiate a decision tree and create a basic context from it.
	tree := evalDTree{
		dim: 3,
		nodes: []query.Node{
			{Feat: 0, ZChild: 1, OChild: 8},
			{Feat: 1, ZChild: 2, OChild: 5},
			{Feat: 2, ZChild: 3, OChild: 4},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			{Feat: 2, ZChild: 6, OChild: 7},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Feat: 2, ZChild: 9, OChild: 12},
			{Feat: 1, ZChild: 10, OChild: 11},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Feat: 1, ZChild: 13, OChild: 14},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
		},
	}
	ctx := query.BasicQContext(tree)

	// We define the partial instance c = (_, 0, _) and variable y.
	c := query.QConst{
		Val: []query.FeatV{query.BOT, query.ZERO, query.BOT},
	}
	y := query.QVar("y")

	// Note that Q(c) uses a universal quantifier while SAT solvers, as the
	// name implies, can only answer in regards to the existance of a single
	// evaluation that makes a formula true. Because of this in order to answer
	// Q(c) we must define and solve for Q'(c) = -Q(c). For a formal definition
	// of DFS(x) please refer to section 5.1 of the paper.
	qry := logop.Or{
		Q1: logop.Not{Q: dfs.Const{I: c}},
		Q2: logop.WithVar{
			I: y,
			Q: logop.And{
				Q1: dfs.Var{I: y},
				Q2: logop.And{
					Q1: lel.VarConst{
						I1:          y,
						I2:          c,
						CountVarGen: botCountVarGen,
					},
					Q2: logop.Not{
						Q: lel.ConstVar{
							I1:          c,
							I2:          y,
							CountVarGen: botCountVarGen,
						},
					},
				},
			},
		},
	}

	// Now we can evaluate if Q'(c) is satsfiable or not using a sat solver
	// specified by the path stored in the environment variable
	// "path_to_solver_binary".

	solver, err := compute.BinSolver("path_to_solver_binary")
	if err != nil {
		panic(err)
	}

	boolComb := compute.QDTFAtom{
		Query:   qry,
		Negated: true,
	}

	result, err := boolComb.Sat(ctx, solver)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
