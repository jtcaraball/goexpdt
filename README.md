# Go-ExplainDT

A dependency free GO implementation of the work in the paper "A Uniform
Language to Explain Decision Trees".

## Example
The following program computes the minimal sufficient reason of an arbitrary
constant and decision tree.

```go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jtcaraball/goexpdt/compute"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/extensions/allcomp"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

// First we define a decision tree struct that implements the query.Model
// interface.
type DTree struct {
	dim   int
	nodes []query.Node
}

func (t DTree) Dim() int {
	return t.dim
}

func (t DTree) Nodes() []query.Node {
	return t.nodes
}

// Then we define a function to generate, given a variable v, a new variable
// used to encode the notion of reachability in the model.
func reachableVarGen(v query.QVar) query.QVar {
	return query.QVar("r" + string(v))
}

func main() {
	// We instantiate a decision tree and create a basic context from it.
	tree := DTree{
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

	// We define the constant we want to compute a minimal sufficient reason
	// for.
	c := query.QConst{
		Val: []query.FeatV{query.ZERO, query.ONE, query.ONE},
	}

	// Finally we define two generator function that return queries
	// representing the property of sufficient reason and a strict partial
	// order under subsumption (minimality). These will be used by the
	// optimisation algorithm.
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
	// which path is stored in the environment variable "SAT_SOLVER_EXEC_PATH".

	output, err := compute.ComputeOptim(
		sufficientReason,
		strictSubsumption,
		query.QVar("x"),
		ctx,
		os.Getenv("SAT_SOLVER_EXEC_PATH"),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if output.Found {
		fmt.Printf(
			"Value found: %s\n",
			strings.Split(output.Value.AsString(), ""),
		)
	} else {
		fmt.Println("No value exists.")
	}

	os.Exit(0)
}
```

## Tests

To build the test suite and execute it, ensure that
[Docker](https://docs.docker.com/engine/install/) and
[make](https://en.wikipedia.org/wiki/Make_(software)) are installed and run the
following command:

```
make test
```
