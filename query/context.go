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
// assignment and concatenation based instance variable and constant scoping. A
// BaiscQContext uses concatenation based scoping with the unit serparator
// caracater (ascii 31) as seperator and so it should be avoided when naming
// elements.
func BasicQContext(model Model) QContext {
	return &baseQContext{
		baseScoper: baseScoper{},
		baseVarManager: baseVarManager{
			vars: make(map[varRep]int),
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
