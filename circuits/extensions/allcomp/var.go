package allcomp

import (
	"errors"
	"goexpdt/cnf"
	"goexpdt/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type acVar struct {
	varInst base.Var
	leafValue bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Return var allComp.
func Var(varInst base.Var, leafValue bool) *acVar {
	return &acVar{varInst: varInst, leafValue: leafValue}
}

// Return CNF encoding of component.
func (ac *acVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar := ac.varInst.Scoped(ctx)
	return ac.buildEncoding(scpVar, ctx)
}

// Generate cnf encoding.
func (ac *acVar) buildEncoding(
	varInst base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	if ctx.Tree == nil || ctx.Tree.Root == nil {
		return nil, errors.New("Tree or it's root is nil")
	}
	nCNF := &cnf.CNF{}
	rVarName := "r" + string(varInst)
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
				-ctx.Var(string(varInst), node.Feat, base.ZERO.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.LChild.ID, 0),
			},
			{
				-ctx.Var(string(varInst), node.Feat, base.ONE.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.RChild.ID, 0),
			},
			{
				-ctx.Var(string(varInst), node.Feat, base.BOT.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.LChild.ID, 0),
			},
			{
				-ctx.Var(string(varInst), node.Feat, base.BOT.Val()),
				-ctx.IVar(rVarName, node.ID, 0),
				ctx.IVar(rVarName, node.RChild.ID, 0),
			},
			{
				-ctx.IVar(rVarName, node.RChild.ID, 0),
				ctx.IVar(rVarName, node.ID, 0),
			},
			{
				-ctx.IVar(rVarName, node.RChild.ID, 0),
				ctx.Var(string(varInst), node.Feat, base.ONE.Val()),
				ctx.Var(string(varInst), node.Feat, base.BOT.Val()),
			},
			{
				-ctx.IVar(rVarName, node.LChild.ID, 0),
				ctx.IVar(rVarName, node.ID, 0),
			},
			{
				-ctx.IVar(rVarName, node.LChild.ID, 0),
				ctx.Var(string(varInst), node.Feat, base.ZERO.Val()),
				ctx.Var(string(varInst), node.Feat, base.BOT.Val()),
			},
		})
	}
	return nCNF, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (ac *acVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return ac, nil
}

// Return slice of pointers to component's children.
func (ac *acVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (ac *acVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
