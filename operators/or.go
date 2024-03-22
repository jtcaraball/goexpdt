package operators

import (
	"errors"
	"fmt"
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type or struct {
	child1 base.Component
	child2 base.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return or operator.
func Or(child1, child2 base.Component) *or {
	return &or{child1: child1, child2: child2}
}

// Return CNF encoding of component.
func (o *or) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	if err := o.nonNilChildren(); err != nil {
		return nil, err
	}
	// De Morgan's law
	// Encode both children
	cnf1, err := o.child1.Encoding(ctx)
	if err != nil {
		return nil, orErr(err, 1)
	}
	cnf2, err := o.child2.Encoding(ctx)
	if err != nil {
		return nil, orErr(err, 2)
	}
	// Negate both children
	var tv int
	tv = cnf1.Negate(ctx.TopV)
	ctx.MaxUpdateTopV(tv) // TopV could have increases while negating
	tv = cnf2.Negate(ctx.TopV)
	ctx.MaxUpdateTopV(tv) // TopV could have increases while negating
	// Logical and and final negation
	cnf1.Conjunction(cnf2)
	tv = cnf1.Negate(ctx.TopV)
	ctx.MaxUpdateTopV(tv) // TopV could have increases while negating
	return cnf1, nil
}

// Return pointer to simplified equivalent component.
func (o *or) Simplified(ctx *base.Context) (base.Component, error) {
	if err := o.nonNilChildren(); err != nil {
		return nil, err
	}
	simpleChild1, err := o.child1.Simplified(ctx)
	if err != nil {
		return nil, orErr(err, 1)
	}
	simpleChild2, err := o.child2.Simplified(ctx)
	if err != nil {
		return nil, orErr(err, 2)
	}
	trivial1, value1 := simpleChild1.IsTrivial()
	trivial2, value2 := simpleChild2.IsTrivial()
	// If child1 true then so is or.
	if trivial1 && value1 {
		return simpleChild1, nil
	}
	// If child2 true then so is or.
	if trivial2 && value2 {
		return simpleChild2, nil
	}
	// If both children are trivial but none are true then or must be false.
	if trivial1 && trivial2 {
		return base.NewTrivial(false), nil
	}
	// If child1 is false and child2 is not or's value is equal to child2.
	if trivial1 && !value1 {
		return simpleChild2, nil
	}
	// If child2 is false and child1 is not or's value is equal to child1.
	if trivial2 && !value2 {
		return simpleChild1, nil
	}
	return &or{child1: simpleChild1, child2: simpleChild2}, nil
}

// Return slice of pointers to component's children.
func (o *or) GetChildren() []base.Component {
	return []base.Component{o.child1, o.child2}
}

// yes is true if struct is trivial and value represents its truthiness.
func (o *or) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Returns error if any of the children are nil.
func (o *or) nonNilChildren() error {
	if o.child1 == nil {
		return orErr(errors.New("child is nil"), 1)
	}
	if o.child2 == nil {
		return orErr(errors.New("child is nil"), 2)
	}
	return nil
}

// Add bread crumbs to error.
func orErr(err error, childIdx uint8) error {
	return errors.New(
		fmt.Sprintf("or:child%d -> %s", childIdx, err.Error()),
	)

}