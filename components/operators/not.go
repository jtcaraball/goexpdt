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

type not struct {
	child components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return not operator.
func Not(child components.Component) *not {
	return &not{child: child}
}

// Return CNF encoding of component.
func (n *not) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	cnf, err := n.child.Encoding(ctx)
	if err != nil {
		return nil, notErr(err)
	}
	tv := cnf.Negate(ctx.TopV)
	ctx.MaxUpdateTopV(tv)
	return cnf, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (n *not) Simplified(ctx *components.Context) (components.Component, error) {
	simpleChild, err := n.child.Simplified(ctx)
	if err != nil {
		return nil, notErr(err)
	}
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return components.NewTrivial(!value), nil
	}
	n.child = simpleChild
	return n, nil
}

// Return slice of pointers to component's children.
func (n *not) GetChildren() []components.Component {
	return []components.Component{n.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (n *not) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Add bread crumbs to error
func notErr(err error) error {
	return errors.New(fmt.Sprintf("not -> %s", err.Error()))
}
