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

type WithVar struct {
	instance instances.Var
	child components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return CNF encoding of component.
func (wv *WithVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
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
func (wv *WithVar) Simplified() components.Component {
	simpleChild := wv.child.Simplified()
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return components.NewTrivial(value)
	}
	wv.child = simpleChild
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

// Add bread crumbs to error
func withVarErr(err error) error {
	return errors.New(fmt.Sprintf("WithVarErr -> %s", err.Error()))
}
