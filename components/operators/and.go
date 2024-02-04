package operators

import (
	"errors"
	"fmt"
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
func (a *And) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	cnf1, err := a.child1.Encoding(ctx)
	if err != nil {
		return nil, andErr(err, 1)
	}
	cnf2, err := a.child2.Encoding(ctx)
	if err != nil {
		return nil, andErr(err, 2)
	}
	cnf1.Conjunction(cnf2)
	return cnf1, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (a *And) Simplified() (components.Component, error) {
	simpleChild1, err := a.child1.Simplified()
	if err != nil {
		return nil, andErr(err, 1)
	}
	simpleChild2, err := a.child2.Simplified()
	if err != nil {
		return nil, andErr(err, 2)
	}
	trivial1, value1 := simpleChild1.IsTrivial()
	trivial2, value2 := simpleChild2.IsTrivial()
	// If any of the two children are trivial and false so to is And.
	if trivial1 && !value1 {
		return components.NewTrivial(false), nil
	}
	if trivial2 && !value2 {
		return components.NewTrivial(false), nil
	}
	// If child1 and child2 are trivial and true then And is true.
	if trivial1 && trivial2 && value1 && value2 {
		return components.NewTrivial(true), nil
	}
	// If child1 is trivial and true the value of And is equal to child2.
	if trivial1 && value1 {
		return simpleChild2, nil
	}
	// If child2 is trivial and true the value of And is equal to child1.
	if trivial2 && value2 {
		return simpleChild1, nil
	}
	a.child1 = simpleChild1
	a.child2 = simpleChild2
	return a, nil
}

// Return slice of pointers to component's children.
func (a *And) GetChildren() []components.Component {
	return []components.Component{a.child1, a.child2}
}

// yes is true if struct is trivial and value represents its truthiness.
func (a *And) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Add bread crumbs to error
func andErr(err error, childIdx uint8) error {
	return errors.New(
		fmt.Sprintf("And:child%d -> %s", childIdx, err.Error()),
	)
}