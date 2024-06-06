package isnode

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Var is the variable version of the Node predicate.
type Var struct {
	I query.QVar
	// ReachNodeVarGen returns a variable generated from v that will be used to
	// encode what nodes in the model v is able (and not able) to reach.
	ReachNodeVarGen func(v query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable n.I
// corresponds to a node in the model.
func (f Var) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	nodes := ctx.Nodes()
	if len(nodes) == 0 {
		return cnf.CNF{}, errors.New("Invalid encoding on empty model")
	}

	dim := ctx.Dim()
	ncnf := cnf.CNF{}
	sv := ctx.ScopeVar(f.I)
	svReach := f.ReachNodeVarGen(sv)

	// We will encode that sv reaches a node with index i as the triple
	// (svReach, 0, i).

	// Every feature that is not BOT must be decided on
	fcl := []cnf.Clause{}

	for i := 0; i < dim; i++ {
		fcl = append(fcl, cnf.Clause{ctx.CNFVar(sv, i, int(query.BOT))})
	}
	for i, node := range nodes {
		if node.IsLeaf() {
			continue
		}

		fcl[node.Feat] = append(fcl[node.Feat], ctx.CNFVar(svReach, 0, i))
	}

	ncnf = ncnf.AppendSemantics(fcl...)

	// Root is always reachable
	ncnf = ncnf.AppendConsistency(cnf.Clause{ctx.CNFVar(svReach, 0, 0)})

	// Add non-bot progapation clauses
	for i, node := range nodes {
		if node.IsLeaf() {
			continue
		}
		ncnf = ncnf.AppendConsistency(
			cnf.Clause{
				-ctx.CNFVar(sv, node.Feat, int(query.ZERO)),
				-ctx.CNFVar(svReach, 0, i),
				ctx.CNFVar(svReach, 0, node.ZChild),
			},
			cnf.Clause{
				-ctx.CNFVar(sv, node.Feat, int(query.ONE)),
				-ctx.CNFVar(svReach, 0, i),
				ctx.CNFVar(svReach, 0, node.OChild),
			},
			cnf.Clause{
				-ctx.CNFVar(svReach, 0, node.OChild),
				ctx.CNFVar(svReach, 0, i),
			},
			cnf.Clause{
				-ctx.CNFVar(svReach, 0, node.OChild),
				ctx.CNFVar(sv, node.Feat, int(query.ONE)),
			},
			cnf.Clause{
				-ctx.CNFVar(svReach, 0, node.ZChild),
				ctx.CNFVar(svReach, 0, i),
			},
			cnf.Clause{
				-ctx.CNFVar(svReach, 0, node.ZChild),
				ctx.CNFVar(sv, node.Feat, int(query.ZERO)),
			},
		)
	}

	return ncnf, nil
}
