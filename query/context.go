package query

// Scoper manages scope for guarded quantifiers.
type Scoper interface {
	// ScopeVar returns an instance variable equal to the original variable
	// plus a suffix representing the stack of scopes at the moment its called.
	ScopeVar(v Var) Var
	// ScopeConst returns an instance constant with its value field set match
	// the value of the scope targeting c at the moment its called if it
	// exists. If the constant is scoped then ok==true else ok==false and the a
	// zero value Const is returned.
	ScopeConst(c Const) (scopedConst Const, ok bool)
	// AddScope adds a scope with target==tgt to the stack.
	AddScope(tgt string)
	// PopScope removes the last scope in the stack.
	PopScope()
	// SetScope adds the value and corresponding index the last scope in the
	// stack. Returns an error if there are no scopes or the last scope is
	// already set.
	SetScope(vIdx int, val []FeatV) error
	// Reset removes all guards in the scope
	Reset()
}

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

// Model provides functionality to access properties of model.
type Model interface {
	// Dim returns the dimension of the model.
	Dim()
	// NodesConsts returns all the model's nodes as constants. The method can
	// fail if the underlying model is inconsistent.
	NodesConsts() ([]Const, error)
	// PosLeafsConsts returns all the model's positive leafs as constants. The
	// method can fail if the underlying model is inconsistent.
	PosLeafsConsts() ([]Const, error)
	// NegLeafsConsts returns all the model's negative leafs as constants. The
	// method can fail if the underlying model is inconsistent.
	NegLeafsConsts() ([]Const, error)
}

// QContext provides a unified interface for passing the Scoper, VarManager and
// Model interfaces to a query components.
type QContext interface {
	Scoper
	VarManager
	Model
	// Reset call the reset method of the embedded interfaces.
	Reset()
}

type baseQContext struct {
	baseScoper
	baseVarManager
	Model
}

// A BasicQContext implements a context with and incremental variable
// assignment and concatenation based instance variable and constant scoping.
func BasicQContext(model Model) QContext {
	return &baseQContext{
		baseScoper:     baseScoper{},
		baseVarManager: baseVarManager{},
		Model:          model,
	}
}

func (qc *baseQContext) Reset() {
	qc.baseScoper.Reset()
	qc.baseVarManager.Reset()
}
