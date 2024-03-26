package allcomp

import (
	"errors"
	"goexpdt/cnf"
	"goexpdt/trees"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type acConst struct {
	constInst base.ConstInstance
	leafValue bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func Const(constInst base.ConstInstance, leafValue bool) *acConst {
	return &acConst{constInst: constInst, leafValue: leafValue}
}

// Return CNF encoding of component.
func (ac *acConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	if ctx.Tree == nil || ctx.Tree.Root == nil {
		return nil, errors.New("Tree or it's root is nil")
	}
	scpConst, err := ac.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return ac.buildEncoding(scpConst, ctx)
}

// Generate cnf encoding.
func (ac *acConst) buildEncoding(
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	var node *trees.Node
	var nStack  = []*trees.Node{ctx.Tree.Root}
	for len(nStack) > 0 {
		node, nStack = nStack[len(nStack) - 1], nStack[:len(nStack) - 1]
		if node.IsLeaf() {
			if node.Value != ac.leafValue {
				return cnf.CNFFromClauses([][]int{{}}), nil
			}
			continue
		}
		if constInst[node.Feat] == base.BOT {
			nStack = append(nStack, node.LChild, node.RChild)
			continue
		}
		if constInst[node.Feat] == base.ONE {
			nStack = append(nStack, node.RChild)
			continue
		}
		if constInst[node.Feat] == base.ZERO {
			nStack = append(nStack, node.LChild)
			continue
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
func (ac *acConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	if ctx.Tree == nil || ctx.Tree.Root == nil {
		return nil, errors.New("Tree or it's root is nil")
	}
	scpConst, err := ac.constInst.Scoped(ctx)
	if err != nil {
		return ac, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return ac.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (ac *acConst) buildSimplified(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	var node *trees.Node
	var nStack  = []*trees.Node{ctx.Tree.Root}
	for len(nStack) > 0 {
		node, nStack = nStack[len(nStack) - 1], nStack[:len(nStack) - 1]
		if node.IsLeaf() {
			if node.Value != ac.leafValue {
				return base.NewTrivial(false), nil
			}
			continue
		}
		if constInst[node.Feat] == base.BOT {
			nStack = append(nStack, node.LChild, node.RChild)
			continue
		}
		if constInst[node.Feat] == base.ONE {
			nStack = append(nStack, node.RChild)
			continue
		}
		if constInst[node.Feat] == base.ZERO {
			nStack = append(nStack, node.LChild)
			continue
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (ac *acConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (ac *acConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
