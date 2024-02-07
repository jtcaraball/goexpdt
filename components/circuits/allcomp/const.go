package allcomp

import (
	"fmt"
	"errors"
	"stratifoiled/cnf"
	"stratifoiled/trees"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type acConst struct {
	constInst instances.Const
	leafValue bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func Const(constInst instances.Const, leafValue bool) *acConst {
	return &acConst{constInst: constInst, leafValue: leafValue}
}

// Return CNF encoding of component.
func (ac *acConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := ac.validateInstances(ctx); err != nil {
		return nil, err
	}
	if ctx.Tree == nil || ctx.Tree.Root == nil {
		return nil, errors.New("Tree or it's root is nil")
	}
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
		if ac.constInst[node.Feat] == instances.BOT {
			nStack = append(nStack, node.LChild, node.RChild)
			continue
		}
		if ac.constInst[node.Feat] == instances.ONE {
			nStack = append(nStack, node.RChild)
			continue
		}
		if ac.constInst[node.Feat] == instances.ZERO {
			nStack = append(nStack, node.LChild)
			continue
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (ac *acConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := ac.validateInstances(ctx); err != nil {
		return nil, err
	}
	if ctx.Tree == nil || ctx.Tree.Root == nil {
		return nil, errors.New("Tree or it's root is nil")
	}
	var node *trees.Node
	var nStack  = []*trees.Node{ctx.Tree.Root}
	for len(nStack) > 0 {
		node, nStack = nStack[len(nStack) - 1], nStack[:len(nStack) - 1]
		if node.IsLeaf() {
			if node.Value != ac.leafValue {
				return components.NewTrivial(false), nil
			}
			continue
		}
		if ac.constInst[node.Feat] == instances.BOT {
			nStack = append(nStack, node.LChild, node.RChild)
			continue
		}
		if ac.constInst[node.Feat] == instances.ONE {
			nStack = append(nStack, node.RChild)
			continue
		}
		if ac.constInst[node.Feat] == instances.ZERO {
			nStack = append(nStack, node.LChild)
			continue
		}
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (ac *acConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (ac *acConst) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Return error if constInst value is invalid.
func (ac *acConst) validateInstances(ctx *components.Context) error {
	if len(ac.constInst) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				"allComp.Const -> constant: wrong dim %d (%d feats in context)",
				len(ac.constInst),
				ctx.Dimension,
			),
		)
	}
	return nil
}
