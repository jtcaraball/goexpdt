package query

// VarManager manages the creation and usage of cnf variables. Variables are
// denoted by the triple (QVar, int, int) and are assigned positive integer
// values. All variables must be assigned a value equal or greater than 1 as 0
// is reserved by the DIMACS standard.
type VarManager interface {
	// TopV returns the value of largest unsigned integer assigned to a
	// variable. This value is expected to be always positive.
	TopV() int
	// UpdateTopV updates the value of the unsigned the value of the largest
	// integer to be assigned. Returns true if the update is valid and took
	// effect.
	UpdateTopV(v int) bool
	// CNFVar return the unsigned integer value assigned to the variable given.
	// If the variable does not exists it first assigns a value to it. This
	// value is expected to be always positive.
	CNFVar(v QVar, idx, val int) int
	// CNFVarExists return true if the passed variable exists.
	CNFVarExists(v QVar, idx, val int) bool
	// Reset removes all variables and sets the next variable to be assigned
	// to 1.
	Reset()
}

type baseVarManager struct {
	topv int            // Value of the largest variable assigned.
	vars map[varRep]int // User variables map.
}

type varRep struct {
	v   QVar
	idx int
	val int
}

func (vm *baseVarManager) TopV() int {
	return vm.topv
}

func (vm *baseVarManager) UpdateTopV(v int) bool {
	if vm.topv > v {
		return false
	}
	vm.topv = v
	return true
}

func (vm *baseVarManager) CNFVar(v QVar, idx, value int) int {
	vr := varRep{v, idx, value}
	if val, ok := vm.vars[vr]; ok {
		return val
	}
	vm.topv += 1
	vm.vars[vr] = vm.topv
	return vm.topv
}

func (vm *baseVarManager) CNFVarExists(v QVar, idx, value int) bool {
	_, ok := vm.vars[varRep{v, idx, value}]
	return ok
}

func (vm *baseVarManager) Reset() {
	vm.topv = 0
	vm.vars = nil
}
