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

type withVar struct {
	instance components.Var
	child components.Component
}

// =========================== //
//           METHODS           //
// =========================== //

// Return withVar operator.
func WithVar(inst components.Var, child components.Component) *withVar {
	return &withVar{instance: inst, child: child}
}

// Return CNF encoding of component.
func (wv *withVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := wv.nonNilChildren(); err != nil {
		return nil, err
	}
	ctx.AddVarToScope(wv.instance)
	scopedInst := wv.instance.Scoped(ctx)
	return wv.buildEncoding(scopedInst, wv.child, ctx)
}

// Generate CNF encoding.
func (wv *withVar) buildEncoding(
	varInstance components.Var,
	child components.Component,
	ctx *components.Context,
) (*cnf.CNF, error) {
	iCNF := wv.instance.Encoding(ctx)
	cCNF, err := wv.child.Encoding(ctx)
	if err != nil {
		return nil, withVarErr(err)
	}
	iCNF.Conjunction(cCNF)
	return iCNF, nil
}

// Return pointer to simplified equivalent component.
func (wv *withVar) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := wv.nonNilChildren(); err != nil {
		return nil, err
	}
	simpleChild, err := wv.child.Simplified(ctx)
	if err != nil {
		return nil, withVarErr(err)
	}
	trivial, value := simpleChild.IsTrivial()
	if trivial {
		return components.NewTrivial(value), nil
	}
	return &withVar{instance: wv.instance, child: simpleChild}, nil
}

// Return slice of pointers to component's children.
func (wv *withVar) GetChildren() []components.Component {
	return []components.Component{wv.child}
}

// yes is true if struct is trivial and value represents its truthiness.
func (wv *withVar) IsTrivial() (yes bool, value bool) {
	return false, false
}

// Returns error if any of the children are nil.
func (wv *withVar) nonNilChildren() error {
	if wv.child == nil {
		return withVarErr(errors.New("child is nil"))
	}
	return nil
}

// Add bread crumbs to error.
func withVarErr(err error) error {
	return errors.New(fmt.Sprintf("withVar -> %s", err.Error()))
}
