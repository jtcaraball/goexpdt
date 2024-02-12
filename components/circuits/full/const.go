package full

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type fConst struct {
	constInst components.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return const full.
func Const(constInst components.ConstInstance) *fConst {
	return &fConst{constInst: constInst}
}

// Return CNF encoding of component.
func (f *fConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	scpConst, err := f.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = components.ValidateConstsDim(
		"full.Const",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return f.buildEncoding(scpConst, ctx)
}

// Generate cnf encoding.
func (f *fConst) buildEncoding(
	constInst components.Const,
	ctx *components.Context,
) (*cnf.CNF, error) {
	for _, ft := range constInst {
		if ft == components.BOT {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
func (f *fConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	scpConst, err := f.constInst.Scoped(ctx)
	if err != nil {
		return f, nil
	}
	if err = components.ValidateConstsDim(
		"full.Const",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return f.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (f *fConst) buildSimplified(
	constInst components.Const,
	ctx *components.Context,
) (components.Component, error) {
	for _, ft := range constInst {
		if ft == components.BOT {
			return components.NewTrivial(false), nil
		}
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (f *fConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (f *fConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
