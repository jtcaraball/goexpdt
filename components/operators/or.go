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

type Or struct {
	child1 components.Component
	child2 components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return CNF encoding of component.
func (o *Or) Encoding(ctx *components.Context) (*cnf.CNF, error) {
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

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (o *Or) Simplified() (components.Component, error) {
	simpleChild1, err := o.child1.Simplified()
	if err != nil {
		return nil, orErr(err, 1)
	}
	simpleChild2, err := o.child2.Simplified()
	if err != nil {
		return nil, orErr(err, 1)
	}
	trivial1, value1 := simpleChild1.IsTrivial()
	trivial2, value2 := simpleChild2.IsTrivial()
	// If child1 true then so is Or.
	if trivial1 && value1 {
		return simpleChild1, nil
	}
	// If child2 true then so is Or.
	if trivial2 && value2 {
		return simpleChild2, nil
	}
	// If both children are trivial but none are true then Or must be false.
	if trivial1 && trivial2 {
		return components.NewTrivial(false), nil
	}
	// If child1 is false and child2 is not Or's value is equal to child2.
	if trivial1 && !value1 {
		return simpleChild2, nil
	}
	// If child2 is false and child1 is not Or's value is equal to child1.
	if trivial2 && !value2 {
		return simpleChild1, nil
	}
	// If no trivial statements are recovered we update children to their
	// simplified form and return pointer to self.
	o.child1 = simpleChild1
	o.child2 = simpleChild2
	return o, nil
}

// Return slice of pointers to component's children.
func (o *Or) GetChildren() []components.Component {
	return []components.Component{o.child1, o.child2}
}

// yes is true if struct is trivial and value represents its truthiness.
func (o *Or) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Add bread crumbs to error
func orErr(err error, childIdx uint8) error {
	return errors.New(
		fmt.Sprintf("Or:child%d -> %s", childIdx, err.Error()),
	)
}