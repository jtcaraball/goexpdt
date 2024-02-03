package components

import (
	"stratifoiled/cnf"
)

// =========================== //
//            TYPE             //
// =========================== //

type Trivial bool

// =========================== //
//           METHODS           //
// =========================== //

// Return pointer to trivial component with truthiness equal to value.
func NewTrivial(value bool) *Trivial {
	var t Trivial = Trivial(value)
	return &t
}

// Return CNF encoding of component.
func (t *Trivial) Encoding(ctx *Context) (*cnf.CNF, error) {
	if bool(*t) {
		return &cnf.CNF{}, nil
	}
	return cnf.CNFFromClauses([][]int{{}}), nil
}

// Return pointer to simplified equivalent component which might be itself.
func (t *Trivial) Simplified() (Component, error) {
	return t, nil
}

// Return slice of pointers to component's children.
func (t *Trivial) GetChildren() []Component {
	return nil
}

// yes is true if struct is trivial and value represents its truthiness.
func (t *Trivial) IsTrivial() (yes bool, value bool) {
	return true, bool(*t)
}
