package operators

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type And struct {
	child1 components.Component
	child2 components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return CNF encoding of component.
func (a *And) Encoding(ctx *components.Context) *cnf.CNF {
	cnf := a.child1.Encoding(ctx)
	cnf.Conjunction(a.child2.Encoding(ctx))
	return cnf
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (a *And) Simplified() components.Component {
	simpleChild1 := a.child1.Simplified()
	simpleChild2 := a.child2.Simplified()
	trivial1, value1 := simpleChild1.IsTrivial()
	trivial2, value2 := simpleChild2.IsTrivial()
	// If any of the two children are trivial and false to is And.
	if trivial1 && !value1 {
		return components.NewTrivial(false)
	}
	if trivial2 && !value2 {
		return components.NewTrivial(false)
	}
	// If child1 and child2 are trivial and true then And is true.
	if trivial1 && trivial2 && value1 && value2 {
		return components.NewTrivial(true)
	}
	// If child1 is trivial and true the value of And is equal to child2.
	if trivial1 && value1 {
		return simpleChild2
	}
	// If child2 is trivial and true the value of And is equal to child1.
	if trivial2 && value2 {
		return simpleChild1
	}
	a.child1 = simpleChild1
	a.child2 = simpleChild2
	return a
}

// Return slice of pointers to component's children.
func (a *And) GetChildren() []components.Component {
	return []components.Component{a.child1, a.child2}
}

// yes is true if struct is trivial and value represents its truthiness.
func (a *And) IsTrivial() (yes bool, value bool) {
	return false, false
}
