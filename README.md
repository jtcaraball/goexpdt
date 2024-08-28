# Go-ExplainDT

[![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/jtcaraball/goexpdt)
[![Paper](http://img.shields.io/badge/paper-arXiv-red.svg?style=flat-square)](https://arxiv.org/abs/2310.11636)

A practical implementation of the work in "A Uniform Language to Explain
Decision Trees".

- [Goexpdt](#goexpdt)
- [Usage](#usage)
- [Examples](#examples)
- [Testing](#testing)

## Goexpdt

Goexpdt is a Go module that allows to freely write, evaluate and compute fixed
explanation queries over arbitrary Decision Trees, represented as formulas in
the Opt-DT-Foil or Q-DT-Foil logics leveraging the power of modern SAT solvers.

The DT-Foil logics and by extension Goexpdt provide the guarantee that for any
query a valid explanation will be computed (or the lack of one existing
informed) making only a polynomial number of calls, over the model's size, to a
SAT solver.

Goexpdt is capable of evaluating and computing a substantial number of the
explanation notions studied in the literature in a serviceable amount of time
over industry sized models, as can be seeing in the following figures:

<p align="center">
    <img src="https://github.com/jtcaraball/goexpdt-experiments/blob/ebb00e2c7ca3552fb8a1291b6cbd770c3dccc0ff/results_figure.png" alt="Results figures."/>
</p>

The repository containing the source code with which this results where
computed can be found
[here](https://github.com/jtcaraball/goexpdt-experiments).

## Usage

In order to not enforce any encoding Goexpdt does not provide a way to
represent Decision Trees and instead relies on the `query.Model` interface to
interact with them through `query.QContext`. The following snippet shows how to
create a context for a simple model implementation:

```go
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

var ctx = query.BasicQContext(
    DTree{
        // Dimension and nodes definition.
    },
)
```

To represent a query, for example:

> There exists a partial instance x such that all its completions are positive.

It is first necessary to define it as a DT-Foil formula with a free variable
using `logop.WithVar`:

```go
var x = query.QVar("x")

var qry = logop.WithVar{
    I: x,
    Q: logop.And{
        Q1: logop.Not{Q: full.Var{I: x}},
        Q2: allcomp.Var{
            I:               x,
            LeafValue:       false,
            ReachNodeVarGen: func(x query.QVar) query.QVar {
                return query.QVar("r" + string(x))
            },
        },
    },
}
```

> [!NOTE]
> An encoding function is passed to `allcomp.Var` in order to ensure that if
> more than one component of a query encodes the notion of 'reachability' all
> use the same variables to do so.

If a query requires more than one variable this can be expressed nesting the
`query.WithVar` structure:

```go
var (
    x = query.QVar("x")
    y = query.QVar("y")
)

var qry = logop.WithVar{
    I: x,
    Q: logop.WithVar{
        I: y,
        // Rest of the query.
    }
}
```

Finally the query can be evaluated by wrapping it in a Q-DT-Foil atom and
making use of a SAT solver through the `Solver` interface as follows:

```go
var boolComb = compute.QDTFAtom{Query: qry}

var (
    solver compute.Solver
    result bool
    err    error
)

solver, err = compute.BinSolver("path_to_solver_binary")
if err != nil {
    panic(err)
}

result, err = boolComb.Sat(ctx, solver)
if err != nil {
    panic(err)
}

fmt.Println(result)
```

> [!IMPORTANT]
> Goexpdt is agnostic to the SAT solver being used as long as it complies with
> the standard DIMACS CNF format.

If the value of `x` is of importance it can be retrieve through the `solver`
using the method `Values`.

> [!NOTE]
> This was an example of evaluating a Q-DT-Foil formula, to see how to compute
> instances for Opt-DT-Foil formulas, that is, instances that satisfy a
> property and are optimal over a particular order, an example is available in
> the [examples](#examples) below.

## Examples

Complete examples of how to use Goexpdt to evaluate and compute explanations
can be found in the [examples](https://github.com/jtcaraball/goexpdt/tree/78e3710b9cbbf46a863343b5f9b0137197fa3d5f/examples)
directory.

## Testing

To build the test suite and execute it, ensure that
[Docker](https://docs.docker.com/engine/install/) and
[make](<https://en.wikipedia.org/wiki/Make_(software)>) are installed and in the
root of the project run `make test`.
