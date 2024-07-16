package allcomp

import (
	"errors"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Var is the variable version of the Node predicate.
type Var struct {
	I query.QVar
	// LeafValue represents the leaf truth value. That is, if LeafValue = true
	// then the extension takes the meaning of AllPos and if LeafValue = false
	// the meaning of AllNeg.
	LeafValue bool
	// ReachNodeVarGen returns a variable generated from v that will be used to
	// encode what nodes in the model v is able (and not able) to reach.
	ReachNodeVarGen func(v query.QVar) query.QVar
}

// Encoding returns a CNF that is true if and only if the query variable ac.I
// is such that all its completions are evaluated as ac.LeafValue by the model.
func (ac Var) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	nodes := ctx.Nodes()
	if len(nodes) == 0 {
		return cnf.CNF{}, errors.New("Invalid encoding on empty model")
	}

	sv := ctx.ScopeVar(ac.I)
	svReach := ac.ReachNodeVarGen(sv)
	ncnf := cnf.CNF{}

	// We will encode that sv reaches a node with index i as the triple
	// (svReach, 0, i).

	ncnf = ncnf.AppendConsistency(cnf.Clause{ctx.CNFVar(svReach, 0, 0)})

	for i, node := range nodes {
		if node.IsLeaf() {
			if node.Value != ac.LeafValue {
				ncnf = ncnf.AppendSemantics(
					cnf.Clause{-ctx.CNFVar(svReach, 0, i)},
				)
			}
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
				-ctx.CNFVar(sv, node.Feat, int(query.BOT)),
				-ctx.CNFVar(svReach, 0, i),
				ctx.CNFVar(svReach, 0, node.ZChild),
			},
			cnf.Clause{
				-ctx.CNFVar(sv, node.Feat, int(query.BOT)),
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
				ctx.CNFVar(sv, node.Feat, int(query.BOT)),
			},
			cnf.Clause{
				-ctx.CNFVar(svReach, 0, node.ZChild),
				ctx.CNFVar(svReach, 0, i),
			},
			cnf.Clause{
				-ctx.CNFVar(svReach, 0, node.ZChild),
				ctx.CNFVar(sv, node.Feat, int(query.ZERO)),
				ctx.CNFVar(sv, node.Feat, int(query.BOT)),
			},
		)
	}

	return ncnf, nil
}
