package operators

import (
	"errors"
	"fmt"
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type withVar struct {
	instance instances.Var
	child components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return withVar operator.
func WithVar(inst instances.Var, child components.Component) *withVar {
	return &withVar{instance: inst, child: child}
}

// Return CNF encoding of component.
func (wv *withVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	iCNF := wv.instance.Encoding(ctx)
	cCNF, err := wv.child.Encoding(ctx)
	if err != nil {
		return nil, withVarErr(err)
	}
	iCNF.Conjunction(cCNF)
	return iCNF, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (wv *withVar) Simplified() (components.Component, error) {
	simpleChild, err := wv.child.Simplified()
	if err != nil {
		return nil, withVarErr(err)
	}
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return components.NewTrivial(value), nil
	}
	wv.child = simpleChild
	return wv, nil
}

// Return slice of pointers to component's children.
func (wv *withVar) GetChildren() []components.Component {
	return []components.Component{wv.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (wv *withVar) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Add bread crumbs to error
func withVarErr(err error) error {
	return errors.New(fmt.Sprintf("withVar -> %s", err.Error()))
}
