package isnode

import (
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type inVar struct {
	varInst base.Var
}

// =========================== //
//           METHODS           //
// =========================== //

// Return var isnode.
func Var(varInst base.Var) *inVar {
	return &inVar{varInst: varInst}
}

// Return CNF encoding of component.
func (f *inVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpVar := f.varInst.Scoped(ctx)
	return f.buildEncoding(scpVar, ctx)
}

// Generate cnf encoding.
func (f *inVar) buildEncoding(
	varInst base.Var,
	ctx *base.Context,
) (*cnf.CNF, error) {
	rVarName := "r" + string(varInst)
	nCNF := &cnf.CNF{}
	// Every feature that is not BOT must be decided on
	featClauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		featClauses = append(
			featClauses,
			[]int{ctx.Var(string(varInst), i, base.BOT.Val())},
		)
	}
	for _, node := range ctx.Tree.Nodes() {
		if node.IsLeaf() {
			continue
		}
		if node.Feat < 0 || node.Feat >= ctx.Dimension {
			return nil, errors.New(
				fmt.Sprintf(
					"Node's feature %d is larger than context's dimension %d",
					node.Feat,
					ctx.Dimension,
				),
			)
		}
		featClauses[node.Feat] = append(
			featClauses[node.Feat],
			ctx.IVar(rVarName, node.ID, 0),
		)
	}
	nCNF.ExtendSemantics(featClauses)
	// Root is always reachable
	nCNF.AppendConsistency([]int{ctx.IVar(rVarName, ctx.Tree.Root.ID, 0)})
	// Add non-bot progapation clauses
	for _, node := range ctx.Tree.Nodes() {
		if node.IsLeaf() {
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
				-ctx.IVar(rVarName, node.RChild.ID, 0),
				ctx.IVar(rVarName, node.ID, 0),
			},
			{
				-ctx.IVar(rVarName, node.RChild.ID, 0),
				ctx.Var(string(varInst), node.Feat, base.ONE.Val()),
			},
			{
				-ctx.IVar(rVarName, node.LChild.ID, 0),
				ctx.IVar(rVarName, node.ID, 0),
			},
			{
				-ctx.IVar(rVarName, node.LChild.ID, 0),
				ctx.Var(string(varInst), node.Feat, base.ZERO.Val()),
			},
		})
	}
	return nCNF, nil
}

// Return pointer to simplified equivalent component which might be itself.
func (f *inVar) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return f, nil
}

// Return slice of pointers to component's children.
func (f *inVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (f *inVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
