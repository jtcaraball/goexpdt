package query

// VarManager manages the creation and usage of cnf variables. Variables are
// divided into user and internal variables and are denoted by the triple
// (name, idx, val).
type VarManager interface {
	// TopV returns the value of largest unsigned integer assigned to a
	// variable.
	TopV() uint
	// UpdateTopV updates the value of the unsigned the value of the largest
	// integer to be assigned. Returns true if the update is valid and took
	// effect.
	UpdateTopV(v uint) bool
	// Var return the unsigned integer value assigned to the user variable
	// given. If the variable does not exists it first assigns a value to it.
	Var(name string, idx, val int) uint
	// VarExists return true if the user variable given exists.
	VarExists(name string, idx, val int) bool
	// Var return the unsigned integer value assigned to the internal variable
	// given. If the variable does not exists it first assigns a value to it.
	IVar(name string, idx, val int) uint
	// VarExists return true if the internal variable given exists.
	IVarExists(name string, idx, val int) bool
	// Reset removes all variables and sets the next variable to be assigned
	// to 0.
	Reset()
}

type baseVarManager struct {
	topv         uint            // Value of the largest variable assigned.
	userVars     map[varRep]uint // User variables map.
	internalVars map[varRep]uint // Internal variables map.
}

type varRep struct {
	Name  string
	Idx   int
	Value int
}

func (vm *baseVarManager) TopV() uint {
	return vm.topv
}

func (vm *baseVarManager) UpdateTopV(v uint) bool {
	if vm.topv > v {
		return false
	}
	vm.topv = v
	return true
}

func (vm *baseVarManager) Var(name string, idx, value int) uint {
	v := varRep{name, idx, value}
	if val, ok := vm.userVars[v]; ok {
		return val
	}
	vm.topv += 1
	vm.userVars[v] = vm.topv
	return vm.topv
}

func (vm *baseVarManager) VarExists(name string, idx, value int) bool {
	_, ok := vm.userVars[varRep{name, idx, value}]
	return ok
}

func (vm *baseVarManager) IVar(name string, idx, value int) uint {
	v := varRep{name, idx, value}
	if val, ok := vm.internalVars[v]; ok {
		return val
	}
	vm.topv += 1
	vm.internalVars[v] = vm.topv
	return vm.topv
}

func (vm *baseVarManager) IVarExists(name string, idx, value int) bool {
	_, ok := vm.internalVars[varRep{name, idx, value}]
	return ok
}

func (vm *baseVarManager) Reset() {
	vm.topv = 0
	vm.userVars = nil
	vm.internalVars = nil
}