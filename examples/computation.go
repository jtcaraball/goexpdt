package main

// In this example we compute a minimal sufficient reason for the instance
// c = (0, 1, 1) over the Decision Tree:
//
// T: f0 -0-> f1 -0-> true
//      |       |
//      |        -1-> false
//      |
//       -1-> f1 -0-> f2 -0-> true
//              |       |
//              |        -1-> true
//              |
//               -1-> true
//
// A sufficient reason of c over T is defined as a partial instance x such that
// x is subsumed by c and all possible completions of x have the same
// classification as c.
//
// A minimal partial instance is defined as a partial instance x such that
// there is no other partial instance y where x != y and y is subsumed by x.
//
// Thus a minimal sufficient reason of c is a partial instance x such that
// x is a sufficient reason of c and there are no other partial instance y
// where y != x, y is subsumed by x and y is a sufficient reason of c.
//
// In Opt-DT-Foil this query is understood as follows: What is an instance x
// that satisfies the property 'sufficient reason of c' that is optimal in the
// subsumption order (minimal). Note that it can be expressed in the logic
// because:
//
// 1. Both the property and the order can be written in the second and first
//    layer of the DT-Foil logic respectively.
// 2. The subsumption order is a strict partial order.
//
// The following code computes such an instance.

import (
	"fmt"
	"strings"

	"github.com/jtcaraball/goexpdt/compute"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/allcomp"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

// First we define a decision tree struct that implements the query.Model
// interface.
type OptDTree struct {
	dim   int
	nodes []query.Node
}

func (t OptDTree) Dim() int {
	return t.dim
}

func (t OptDTree) Nodes() []query.Node {
	return t.nodes
}

// Then we define a function to generate, given a variable v, a new variable
// used to encode the notion of reachability in the model.
func reachableVarGen(v query.QVar) query.QVar {
	return query.QVar("r" + string(v))
}

func main() {
	// We instantiate a decision tree and create a basic context from it.
	tree := OptDTree{
		dim: 3,
		nodes: []query.Node{
			{Feat: 0, ZChild: 1, OChild: 6},
			{Feat: 1, ZChild: 2, OChild: 5},
			{Feat: 2, ZChild: 3, OChild: 4},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Feat: 2, ZChild: 7, OChild: 8},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
			{Feat: 1, ZChild: 9, OChild: 10},
			{Value: true, ZChild: query.NoChild, OChild: query.NoChild},
			{Value: false, ZChild: query.NoChild, OChild: query.NoChild},
		},
	}
	ctx := query.BasicQContext(tree)

	// We define the partial instance c = (0, 1, 1).
	c := query.QConst{
		Val: []query.FeatV{query.ZERO, query.ONE, query.ONE},
	}

	// We define two generator function that return queries representing the
	// property of sufficient reason and a strict partial order under
	// subsumption (minimality).
	// These will be used by the optimization algorithm.
	sufficientReason := func(v query.QVar) compute.Encodable {
		return logop.WithVar{
			I: v,
			Q: logop.And{
				Q1: subsumption.VarConst{I1: v, I2: c},
				Q2: logop.And{
					Q1: logop.Or{
						Q1: logop.Not{
							Q: allcomp.Const{I: c, LeafValue: true},
						},
						Q2: allcomp.Var{
							I:               v,
							LeafValue:       true,
							ReachNodeVarGen: reachableVarGen,
						},
					},
					Q2: logop.Or{
						Q1: logop.Not{
							Q: allcomp.Const{I: c, LeafValue: false},
						},
						Q2: allcomp.Var{
							I:               v,
							LeafValue:       false,
							ReachNodeVarGen: reachableVarGen,
						},
					},
				},
			},
		}
	}

	strictSubsumption := func(v query.QVar, u query.QConst) compute.Encodable {
		return logop.And{
			Q1: subsumption.VarConst{I1: v, I2: u},
			Q2: logop.Not{Q: subsumption.ConstVar{I1: u, I2: v}},
		}
	}

	// Now we can compute a minimal sufficient reason for c using a sat solver
	// specified by the path stored in the environment variable
	// "path_to_solver_binary".

	solver, err := compute.NewBinSolver("path_to_solver_binary")
	if err != nil {
		panic(err)
	}

	output, err := compute.ComputeOptim(
		sufficientReason,
		strictSubsumption,
		query.QVar("x"),
		ctx,
		solver,
	)
	if err != nil {
		panic(err)
	}

	if output.Found {
		fmt.Printf(
			"Value found: %s\n",
			strings.Split(output.Value.AsString(), ""),
		)
	} else {
		fmt.Println("No value exists.")
	}
}
