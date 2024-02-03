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

type Not struct {
	child components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return CNF encoding of component.
func (n *Not) Encoding(ctx *components.Context) (*cnf.CNF, error) {
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
func (n *Not) Simplified() components.Component {
	simpleChild := n.child.Simplified()
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return components.NewTrivial(!value)
	}
	n.child = simpleChild
	return n
}

// Return slice of pointers to component's children.
func (n *Not) GetChildren() []components.Component {
	return []components.Component{n.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (n *Not) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Add bread crumbs to error
func notErr(err error) error {
	return errors.New(fmt.Sprintf("Not -> %s", err.Error()))
}
