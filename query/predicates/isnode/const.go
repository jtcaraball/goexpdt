package isnode

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Const is the constant version of the Node predicate.
type Const struct {
	I query.QConst
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
	visited := newVisitRecord(dim, sc)

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
		visited[node.Feat] = 1

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
	if allFeatsVisited(visited) {
		return cnf.TrueCNF, nil
	}

	return cnf.FalseCNF, nil
}

// newVisitRecord returns an int slice denoting with a 1 which features the
// query constant c has "visited" and with a 0 those is has not. Initially only
// the features with value equal to bottom are marked as visited.
func newVisitRecord(dim int, c query.QConst) []int {
	vr := make([]int, dim)

	for i, ft := range c.Val {
		if ft == query.BOT {
			vr[i] = 1
			continue
		}
		vr[i] = 0
	}

	return vr
}

// allVisited returns true if every feature is marked as "visited".
func allFeatsVisited(vr []int) bool {
	sum := 0

	for _, val := range vr {
		sum += val
	}

	return sum == len(vr)
}
