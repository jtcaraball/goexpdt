package query

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
