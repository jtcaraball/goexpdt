package allcomp

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// Const is the constant version of the AllComp extension.
type Const struct {
	I query.QConst
	// LeafValue represents the leaf truth value. That is, if LeafValue = true
	// then the extension takes the meaning of AllPos and if LeafValue = false
	// the meaning of AllNeg.
	LeafValue bool
}

// Encoding returns a CNF that is true if and only if the query constant ac.I
// is such that all its completions are evaluated as ac.LeafValue by the model.
func (ac Const) Encoding(ctx query.QContext) (cnf.CNF, error) {
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	sc, _ := ctx.ScopeConst(ac.I)
	if err := query.ValidateConstsDim(ctx.Dim(), sc); err != nil {
		return cnf.CNF{}, err
	}

	nodes := ctx.Nodes()
	if len(nodes) == 0 {
		return cnf.CNF{}, errors.New("Invalid encoding on empty model")
	}

	var n query.Node
	dim := ctx.Dim()
	nvisited := make(map[int]struct{})
	nstack := []query.Node{nodes[0]}

	for len(nstack) > 0 {
		n, nstack = nstack[len(nstack)-1], nstack[:len(nstack)-1]

		if n.IsLeaf() {
			if n.Value != ac.LeafValue {
				return cnf.FalseCNF, nil
			}
			continue
		}

		if n.Feat < 0 || n.Feat >= dim {
			return cnf.CNF{}, fmt.Errorf(
				"Node's feature %d is out of range [0, %d]",
				n.Feat,
				dim-1,
			)
		}

		switch sc.Val[n.Feat] {
		case query.BOT:
			if _, ok := nvisited[n.ZChild]; !ok {
				nvisited[n.ZChild] = struct{}{}
				nstack = append(nstack, nodes[n.ZChild])
			}
			if _, ok := nvisited[n.OChild]; !ok {
				nvisited[n.OChild] = struct{}{}
				nstack = append(nstack, nodes[n.OChild])
			}
		case query.ZERO:
			if _, ok := nvisited[n.ZChild]; !ok {
				nvisited[n.ZChild] = struct{}{}
				nstack = append(nstack, nodes[n.ZChild])
			}
		case query.ONE:
			if _, ok := nvisited[n.OChild]; !ok {
				nvisited[n.OChild] = struct{}{}
				nstack = append(nstack, nodes[n.OChild])
			}
		}
	}

	return cnf.TrueCNF, nil
}
