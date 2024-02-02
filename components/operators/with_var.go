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
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return components.NewTrivial(value)
	}
	return wv
}

// Return slice of pointers to component's children.
func (wv *WithVar) GetChildren() []components.Component {
	return []components.Component{wv.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (wv *WithVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
