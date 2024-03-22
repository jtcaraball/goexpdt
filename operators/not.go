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

type not struct {
	child base.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return not operator.
func Not(child base.Component) *not {
	return &not{child: child}
}

// Return CNF encoding of component.
func (n *not) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	if err := n.nonNilChildren(); err != nil {
		return nil, err
	}
	cnf, err := n.child.Encoding(ctx)
	if err != nil {
		return nil, notErr(err)
	}
	tv := cnf.Negate(ctx.TopV)
	ctx.MaxUpdateTopV(tv)
	return cnf, nil
}

// Return pointer to simplified equivalent component.
func (n *not) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	if err := n.nonNilChildren(); err != nil {
		return nil, err
	}
	simpleChild, err := n.child.Simplified(ctx)
	if err != nil {
		return nil, notErr(err)
	}
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return base.NewTrivial(!value), nil
	}
	return &not{child: simpleChild}, nil
}

// Return slice of pointers to component's children.
func (n *not) GetChildren() []base.Component {
	return []base.Component{n.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (n *not) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Returns error if any of the children are nil.
func (n *not) nonNilChildren() error {
	if n.child == nil {
		return notErr(errors.New("child is nil"))
	}
	return nil
}

// Add bread crumbs to error.
func notErr(err error) error {
	return errors.New(fmt.Sprintf("not -> %s", err.Error()))
}
