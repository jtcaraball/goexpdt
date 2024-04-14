package isnode

import (
	"fmt"
	"errors"
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type inConst struct {
	constInst base.ConstInstance
	featReg []int
}

// =========================== //
//           METHODS           //
// =========================== //

// Return const isnode.
func Const(constInst base.ConstInstance) *inConst {
	return &inConst{constInst: constInst}
}

// Return CNF encoding of component.
func (n *inConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst, err := n.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return n.buildEncoding(scpConst, ctx)
}

// Generate cnf encoding.
func (n *inConst) buildEncoding(
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	if n.featReg == nil {
		n.featReg = make([]int, ctx.Dimension)
	}
	if err := n.setFeats(constInst); err != nil {
		return nil, err
	}
	node := ctx.Tree.Root
	for true {
		if node.IsLeaf() {
			break
		}
		if node.Feat < 0 || node.Feat >= ctx.Dimension {
			return nil, errors.New(
				fmt.Sprintf(
					"Node's feature %d is out of range [0, %d]",
					node.Feat,
					ctx.Dimension - 1,
				),
			)
		}
		if constInst[node.Feat] == base.BOT {
			break
		}
		// Mark feature as decided and set corresponding children node.
		n.featReg[node.Feat] = 1
		if constInst[node.Feat] == base.ZERO {
			node = node.LChild
			continue
		}
		node = node.RChild
	}
	if n.validFeatReg() {
		return &cnf.CNF{}, nil
	}
	return cnf.CNFFromClauses([][]int{{}}), nil
}

// Set component's featReg to match the passed constant.
func (n *inConst) setFeats(constInst base.Const) error {
	if err := base.ValidateConstsDim(len(n.featReg), constInst); err != nil {
		return err
	}
	for i, ft := range constInst {
		if ft == base.BOT {
			n.featReg[i] = 1
			continue
		}
		n.featReg[i] = 0
	}
	return nil
}

// Return true if all elements in n's featReg are 1.
func (n *inConst) validFeatReg() bool {
	sum := 0
	for _, val := range n.featReg {
		sum += val
	}
	return sum == len(n.featReg)
}

// Return pointer to simplified equivalent component which might be itself.
func (n *inConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := n.constInst.Scoped(ctx)
	if err != nil {
		return n, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return n.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (n *inConst) buildSimplified(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	if n.featReg == nil {
		n.featReg = make([]int, ctx.Dimension)
	}
	if err := n.setFeats(constInst); err != nil {
		return nil, err
	}
	node := ctx.Tree.Root
	for true {
		if node.IsLeaf() {
			break
		}
		if node.Feat < 0 || node.Feat >= ctx.Dimension {
			return nil, errors.New(
				fmt.Sprintf(
					"Node's feature %d is larger than ctx's dimension %d",
					node.Feat,
					ctx.Dimension,
				),
			)
		}
		if constInst[node.Feat] == base.BOT {
			break
		}
		// Mark feature as decided and set corresponding children node.
		n.featReg[node.Feat] = 1
		if constInst[node.Feat] == base.ZERO {
			node = node.LChild
			continue
		}
		node = node.RChild
	}
	if n.validFeatReg() {
		return base.NewTrivial(true), nil
	}
	return base.NewTrivial(false), nil
}

// Return slice of pointers to component's children.
func (n *inConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (n *inConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
