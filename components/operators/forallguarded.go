package operators

import (
	"errors"
	"fmt"
	"stratifoiled/cnf"
	"stratifoiled/components"
	"strconv"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type forAllGuarded struct {
	instance components.GuardedConst
	child components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return forAllGuarded operator.
func ForAllGuarded(
	inst components.GuardedConst,
	child components.Component,
) *forAllGuarded {
	return &forAllGuarded{instance: inst, child: child}
}

// Return CNF encoding of component.
func (fag *forAllGuarded) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	nCNF := &cnf.CNF{}
	gIdx := len(ctx.Guards)
	ctx.Guards = append(
		ctx.Guards,
		components.Guard{Target: string(fag.instance)},
	)
	// Remove guard after encoding.
	defer func() {
		ctx.Guards = ctx.Guards[:gIdx]
	}()
	// Get trees nodes from context.
	nodeConsts, err := ctx.NodesAsConsts()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(nodeConsts); i++ {
		// Update guard value to current node.
		ctx.Guards[gIdx].Value = nodeConsts[i]
		ctx.Guards[gIdx].Rep = strconv.Itoa(i)
		// Encode.
		iCNF, err := fag.child.Encoding(ctx)
		if err != nil {
			return nil, forAllGuardedErr(err)
		}
		// Early exit check.
		if iCNF.HasEmptySemanticClause() {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
		nCNF.Conjunction(iCNF)
	}
	return nCNF, nil
}

// Return pointer to simplified equivalent component.
func (fag *forAllGuarded) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	simpleChild, err := fag.child.Simplified(ctx)
	if err != nil {
		return nil, forAllGuardedErr(err)
	}
	if trivial, _ := simpleChild.IsTrivial(); trivial {
		return simpleChild, nil
	}
	return &forAllGuarded{instance: fag.instance, child: simpleChild}, nil
}

// Return slice of pointers to component's children.
func (fag *forAllGuarded) GetChildren() []components.Component {
	return []components.Component{fag.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (fag *forAllGuarded) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Returns error if any of the children are nil.
func (fag *forAllGuarded) nonNilChildren() error {
	if fag.child == nil {
		return forAllGuardedErr(errors.New("child is nil"))
	}
	return nil
}

// Add bread crumbs to error.
func forAllGuardedErr(err error) error {
	return errors.New(fmt.Sprintf("forAllGuarded -> %s", err.Error()))
}
