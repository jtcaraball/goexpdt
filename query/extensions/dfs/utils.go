package dfs

import (
	"fmt"

	"github.com/jtcaraball/goexpdt/query"
)

// leafsAsConsts returns, as query constants, the positive and negative leafs
// (in that order) of the model represented in nodes.
func leafsAsConsts(
	dim int,
	nodes []query.Node,
) ([]query.QConst, []query.QConst, error) {
	for _, n := range nodes {
		if err := validateNode(n, dim, len(nodes)); err != nil {
			return nil, nil, err
		}
	}

	pleaf := []query.QConst{}
	nleaf := []query.QConst{}

	var tmp []query.FeatV
	cv := make([]query.FeatV, dim)

	var cp, pp, zp, op int
	walk := make(
		[]struct {
			prev    int
			visited bool
		},
		len(nodes),
	)

	for {
		if walk[cp].visited {
			// If its a leaf then we attempt to back up.
			if nodes[cp].IsLeaf() {
				goto backup
			}

			zp = nodes[cp].ZChild
			if !walk[zp].visited {
				cv[nodes[cp].Feat] = query.ZERO

				pp = cp
				cp = zp
				walk[cp].prev = pp

				continue
			}

			op = nodes[cp].OChild
			if !walk[op].visited {
				cv[nodes[cp].Feat] = query.ONE

				pp = cp
				cp = op
				walk[cp].prev = pp

				continue
			}

		backup:
			if cp == 0 {
				break
			}

			cp = walk[cp].prev
			cv[nodes[cp].Feat] = query.BOT

			continue
		}

		walk[cp].visited = true

		if !nodes[cp].IsLeaf() {
			continue
		}

		tmp = make([]query.FeatV, dim)
		copy(tmp, cv)

		if nodes[cp].Value {
			pleaf = append(pleaf, query.QConst{Val: tmp})
		} else {
			nleaf = append(nleaf, query.QConst{Val: tmp})
		}
	}

	return pleaf, nleaf, nil
}

// validateNode returns a non nill error if the node n is invalid for the the
// given dimension and number of nodes.
func validateNode(n query.Node, dim, nlen int) error {
	if n.IsLeaf() {
		return nil
	}

	if n.Feat < 0 || n.Feat >= dim {
		return fmt.Errorf(
			"Node's feature %d is out of range [0, %d]",
			n.Feat,
			dim-1,
		)
	}

	if n.ZChild < 0 || n.ZChild >= nlen {
		return fmt.Errorf("Node's ZChild out of bounds %d", n.ZChild)
	}

	if n.OChild < 0 || n.OChild >= nlen {
		return fmt.Errorf("Node's OChild out of bounds %d", n.OChild)
	}

	return nil
}
