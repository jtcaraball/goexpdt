package query

// VarManager manages the creation and usage of cnf variables. Variables are
// denoted by the triple (name, idx, val) and are assigned positive integer
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
	// Var return the unsigned integer value assigned to the variable given. If
	// the variable does not exists it first assigns a value to it. This value
	// is expected to be always positive.
	Var(name string, idx, val int) int
	// VarExists return true if the passed variable exists.
	VarExists(name string, idx, val int) bool
	// Reset removes all variables and sets the next variable to be assigned
	// to 1.
	Reset()
}

type baseVarManager struct {
	topv         int            // Value of the largest variable assigned.
	vars     map[varRep]int // User variables map.
}

type varRep struct {
	Name  string
	Idx   int
	Value int
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

func (vm *baseVarManager) Var(name string, idx, value int) int {
	v := varRep{name, idx, value}
	if val, ok := vm.vars[v]; ok {
		return val
	}
	vm.topv += 1
	vm.vars[v] = vm.topv
	return vm.topv
}

func (vm *baseVarManager) VarExists(name string, idx, value int) bool {
	_, ok := vm.vars[varRep{name, idx, value}]
	return ok
}

func (vm *baseVarManager) Reset() {
	vm.topv = 0
	vm.vars = nil
}
