package query

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
		baseScoper: baseScoper{},
		baseVarManager: baseVarManager{
			userVars:     make(map[varRep]uint),
			internalVars: make(map[varRep]uint),
		},
		Model: model,
	}
}

// Reset the context's scoper and variable manager to their original state.
func (qc *baseQContext) Reset() {
	if qc == nil {
		return
	}
	qc.baseScoper.Reset()
	qc.baseVarManager.Reset()
}
