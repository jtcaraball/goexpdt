package lel

import (
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConst struct {
	varInst components.Var
	constInst components.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst lel.
func VarConst(
	varInst components.Var,
	constInst components.ConstInstance,
) *varConst {
	return &varConst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (l *varConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	scpConst, err := l.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	scpVar := l.varInst.Scoped(ctx)
	if err = components.ValidateConstsDim(
		"lel.VarConst",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return l.buildEncoding(scpVar, scpConst, ctx)
}

// Return CNF encoding of component.
func (l *varConst) buildEncoding(
	varInst components.Var,
	constInst components.Const,
	ctx *components.Context,
) (*cnf.CNF, error) {
	cnf := &cnf.CNF{}
	cnf.ExtendConsistency(genCountClauses(string(varInst), ctx))
	// Count amount of bots in constant.
	botsInConst := 0
	for _, f := range constInst {
		if f == components.BOT {
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
	ctx *components.Context,
) (components.Component, error) {
	scpConst, err := l.constInst.Scoped(ctx)
	if err != nil {
		return l, nil
	}
	if err = components.ValidateConstsDim(
		"lel.VarConst",
		ctx,
		scpConst,
	); err != nil {
		return nil, err
	}
	return l.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (l *varConst) buildSimplified(
	constInst components.Const,
	ctx *components.Context,
) (components.Component, error) {
	for _, f := range constInst {
		if f == components.BOT {
			return l, nil
		}
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *varConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *varConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
