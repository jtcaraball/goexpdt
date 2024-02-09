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

type and struct {
	child1 components.Component
	child2 components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return and operator.
func And(child1, child2 components.Component) *and {
	return &and{child1: child1, child2: child2}
}

// Return CNF encoding of component.
func (a *and) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := a.nonNilChildren(); err != nil {
		return nil, err
	}
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

// Return pointer to simplified equivalent component.
func (a *and) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := a.nonNilChildren(); err != nil {
		return nil, err
	}
	simpleChild1, err := a.child1.Simplified(ctx)
	if err != nil {
		return nil, andErr(err, 1)
	}
	simpleChild2, err := a.child2.Simplified(ctx)
	if err != nil {
		return nil, andErr(err, 2)
	}
	trivial1, value1 := simpleChild1.IsTrivial()
	trivial2, value2 := simpleChild2.IsTrivial()
	// If any of the two children are trivial and false so to is and.
	if trivial1 && !value1 {
		return components.NewTrivial(false), nil
	}
	if trivial2 && !value2 {
		return components.NewTrivial(false), nil
	}
	// If child1 and child2 are trivial and true then and is true.
	if trivial1 && trivial2 && value1 && value2 {
		return components.NewTrivial(true), nil
	}
	// If child1 is trivial and true the value of and is equal to child2.
	if trivial1 && value1 {
		return simpleChild2, nil
	}
	// If child2 is trivial and true the value of and is equal to child1.
	if trivial2 && value2 {
		return simpleChild1, nil
	}
	return &and{child1: simpleChild1, child2: simpleChild2}, nil
}

// Return slice of pointers to component's children.
func (a *and) GetChildren() []components.Component {
	return []components.Component{a.child1, a.child2}
}

// yes is true if struct is trivial and value represents its truthiness.
func (a *and) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Returns error if any of the children are nil.
func (a *and) nonNilChildren() error {
	if a.child1 == nil {
		return andErr(errors.New("child is nil"), 1)
	}
	if a.child2 == nil {
		return andErr(errors.New("child is nil"), 2)
	}
	return nil
}

// Add bread crumbs to error.
func andErr(err error, childIdx uint8) error {
	return errors.New(
		fmt.Sprintf("and:child%d -> %s", childIdx, err.Error()),
	)
}
