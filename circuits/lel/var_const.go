package lel

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConst struct {
	varInst base.Var
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst lel.
func VarConst(
	varInst base.Var,
	constInst base.ConstInstance,
) *varConst {
	return &varConst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (l *varConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst, err := l.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := l.varInst.Scoped(ctx)
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return l.buildEncoding(scpVar, scpConst, ctx)
}

// Return CNF encoding of component.
func (l *varConst) buildEncoding(
	varInst base.Var,
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	cnf := &cnf.CNF{}
	cnf.ExtendConsistency(genCountClauses(string(varInst), ctx))
	// Count amount of bots in constant.
	botsInConst := 0
	for _, ft := range constInst {
		if ft == base.BOT {
			botsInConst += 1
		}
	}
	// Ask for var to not have fewer bots.
	cVarName := "c" + string(varInst)
	for i := 0; i < botsInConst; i++ {
		cnf.AppendSemantics([]int{-ctx.IVar(cVarName, ctx.Dimension - 1, i)})
	}
	return cnf, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *varConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := l.constInst.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return l.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (l *varConst) buildSimplified(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	for _, ft := range constInst {
		if ft == base.BOT {
			return l, nil
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *varConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
