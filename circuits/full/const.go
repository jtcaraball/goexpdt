package full

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type fConst struct {
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return const full.
func Const(constInst base.ConstInstance) *fConst {
	return &fConst{constInst: constInst}
}

// Return CNF encoding of component.
func (f *fConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst, err := f.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return f.buildEncoding(scpConst, ctx)
}

// Generate cnf encoding.
func (f *fConst) buildEncoding(
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	for _, ft := range constInst {
		if ft == base.BOT {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
func (f *fConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := f.constInst.Scoped(ctx)
	if err != nil {
		return f, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return f.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (f *fConst) buildSimplified(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	for _, ft := range constInst {
		if ft == base.BOT {
			return base.NewTrivial(false), nil
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (f *fConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (f *fConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
