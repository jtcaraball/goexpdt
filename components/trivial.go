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

// Return true if the component is trivial. False otherwise.
func (t *Trivial) IsTrivial() bool {
	return true
}
