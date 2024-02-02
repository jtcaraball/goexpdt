package operators

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type WithVar struct {
	instance instances.Var
	child components.Component
	ctx *components.Context
}

// =========================== //
//           METHODS           //
// =========================== //

// Return CNF encoding of component.
func (wv *WithVar) Encoding(ctx *components.Context) *cnf.CNF {
	nCNF := wv.instance.Encoding(ctx)
	nCNF.Conjunction(wv.child.Encoding(ctx))
	return nCNF
}

// Return pointer to simplified equivalent component which might be itself.
func (wv *WithVar) Simplified() components.Component {
	simpleChild := wv.child.Simplified()
	if simpleChild.IsTrivial() {
		return simpleChild
	}
	return wv
}

// Return slice of pointers to component's children.
func (wv *WithVar) GetChildren() []components.Component {
	return []components.Component{wv.child}
}

// Return true if the component is trivial.
func (wv *WithVar) IsTrivial() bool {
	return false
}
