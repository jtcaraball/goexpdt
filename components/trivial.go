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
func (t *Trivial) Encoding(ctx *Context) *cnf.CNF {
	if bool(*t) {
		return &cnf.CNF{}
	}
	return cnf.CNFFromClauses([][]int{{}})
}

// Return pointer to simplified equivalent component which might be itself.
func (t *Trivial) Simplified() Component {
	return t
}

// Return slice of pointers to component's children.
func (t *Trivial) GetChildren() []Component {
	return nil
}

// yes is true if struct is trivial and value represents its truthiness.
func (t *Trivial) IsTrivial() (yes bool, value bool) {
	return true, bool(*t)
}
