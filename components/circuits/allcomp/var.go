package allcomp

import (
	"errors"
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type acVar struct {
	varInst instances.Var
	leafValue bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varVar lel.
func Var(varInst instances.Var, leafValue bool) *acVar {
	return &acVar{varInst: varInst, leafValue: leafValue}
}

// Return CNF encoding of component.
func (ac *acVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if ctx.Tree == nil || ctx.Tree.Root == nil {
		return nil, errors.New("Tree or it's root is nil")
	}
	nCNF := &cnf.CNF{}
	rVarName := "r" + string(ac.varInst)
	// Root is always reachable
	nCNF.AppendConsistency([]int{ctx.IVar(rVarName, ctx.Tree.Root.ID, 0)})
	// Add progapation clauses
	for _, node := range ctx.Tree.Nodes() {
		if node.IsLeaf() {
			if node.Value != ac.leafValue {
				nCNF.AppendSemantics([]int{-ctx.IVar(rVarName, node.ID, 0)})
			}
			continue
		}
		nCNF.ExtendConsistency([][]int{
			{
				-ctx.Var(string(ac.varInst), node.Feat, instances.ZERO.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.LChild.ID, 0),
			},
			{
				-ctx.Var(string(ac.varInst), node.Feat, instances.ONE.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.RChild.ID, 0),
			},
			{
				-ctx.Var(string(ac.varInst), node.Feat, instances.BOT.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.LChild.ID, 0),
			},
			{
				-ctx.Var(string(ac.varInst), node.Feat, instances.BOT.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.RChild.ID, 0),
			},
			{
				-ctx.IVar(rVarName, node.RChild.ID, 0),
				ctx.IVar(rVarName, node.ID, 0),
			},
			{
				-ctx.Var(rVarName, node.RChild.ID, 0),
				ctx.Var(string(ac.varInst), node.Feat, instances.ONE.Val()),
				ctx.Var(string(ac.varInst), node.Feat, instances.BOT.Val()),
			},
			{
				-ctx.IVar(rVarName, node.LChild.ID, 0),
				ctx.IVar(rVarName, node.ID, 0),
			},
			{
				-ctx.Var(rVarName, node.LChild.ID, 0),
				ctx.Var(string(ac.varInst), node.Feat, instances.ZERO.Val()),
				ctx.Var(string(ac.varInst), node.Feat, instances.BOT.Val()),
			},
		})
	}
	return nCNF, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (ac *acVar) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	return ac, nil
}

// Return slice of pointers to component's children.
func (ac *acVar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (ac *acVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
