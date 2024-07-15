package isnode

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Const is the constant version of the Node predicate.
type Const struct {
	I            query.QConst
	visitedFeats []int
}

// Encoding returns a CNF that is true if and only if the query constant n.I
// corresponds to a node in the model.
func (n Const) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc, _ := ctx.ScopeConst(n.I)
	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	dim := ctx.Dim()
	n.initVisitedFeats(dim, sc)

	nodes := ctx.Nodes()
	if len(nodes) == 0 {
		return cnf.CNF{}, errors.New("Invalid encoding on empty model")
	}

	// Lets walk down the model from the root and see if all the decisions made
	// are over non BOT features.
	node := nodes[0]
	for {
		// If we reach a leaf then we decided on every possible feature.
		if node.IsLeaf() {
			break
		}

		if node.Feat < 0 || node.Feat >= dim {
			return cnf.CNF{}, fmt.Errorf(
				"Node's feature %d is out of range [0, %d]",
				node.Feat,
				dim-1,
			)
		}

		// A constant with a BOT value on a decided feature can not be a node.
		if sc.Val[node.Feat] == query.BOT {
			break
		}

		// Mark the feature we are deciding on as visited and set next
		// iteration node as the current's corresponding children.
		n.visitedFeats[node.Feat] = 1

		if sc.Val[node.Feat] == query.ZERO {
			if node.ZChild < 0 || node.ZChild >= len(nodes) {
				return cnf.CNF{}, fmt.Errorf(
					"Node's ZChild out of bounds %d",
					node.ZChild,
				)
			}
			node = nodes[node.ZChild]
			continue
		}

		if node.OChild < 0 || node.OChild >= len(nodes) {
			return cnf.CNF{}, fmt.Errorf(
				"Node's OChild out of bounds %d",
				node.OChild,
			)
		}
		node = nodes[node.OChild]
	}

	// If all features where visited then the constant must represent a node.
	if n.allFeatsVisited() {
		return cnf.TrueCNF, nil
	}

	return cnf.FalseCNF, nil
}

// initVisitedFeatures marks every BOT valued feature as "visited" and all
// others as "not visited".
func (n *Const) initVisitedFeats(dim int, c query.QConst) {
	// As n is passed on as value on the Encoding method this is always
	// evaluating to true. Someday ill think about a way of not allocating a
	// slice every time its called, but that day is not today.
	if n.visitedFeats == nil {
		n.visitedFeats = make([]int, dim)
	}
	for i, ft := range c.Val {
		if ft == query.BOT {
			n.visitedFeats[i] = 1
			continue
		}
		n.visitedFeats[i] = 0
	}
}

// allVisited returns true if every feature is marked as "visited".
func (n Const) allFeatsVisited() bool {
	sum := 0
	for _, val := range n.visitedFeats {
		sum += val
	}
	return sum == len(n.visitedFeats)
}
