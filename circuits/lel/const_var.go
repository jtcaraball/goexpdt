package lel

import (
	"stratifoiled/cnf"
	"stratifoiled/base"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constVar struct {
	varInst base.Var
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVar lel.
func ConstVar(
	constInst base.ConstInstance,
	varInst base.Var,
) *constVar {
	return &constVar{constInst: constInst, varInst: varInst}
}

// Return CNF encoding of component.
func (l *constVar) Encoding(ctx *base.Context) (*cnf.CNF, error) {
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
	return l.buildEncoding(scpConst, scpVar, ctx)
}

// Return CNF encoding of component.
func (l *constVar) buildEncoding(
	constInst base.Const,
	varInst base.Var,
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
	// Ask for var to not have more bots.
	cVarName := "c" + string(varInst)
	for i := botsInConst + 1; i < ctx.Dimension + 1; i++ {
		cnf.AppendSemantics([]int{-ctx.IVar(cVarName, ctx.Dimension - 1, i)})
	}
	return cnf, nil
}

// TODO: Add correct simplification for guarded const.
// Return pointer to simplified equivalent component which might be itself.
func (l *constVar) Simplified(
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
	return l.buildSimplification(scpConst, ctx)
}

// Generate simplified component.
func (l *constVar) buildSimplification(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	// If const has only bottoms then this predicate is trivialy true.
	for _, ft := range constInst {
		if ft != base.BOT {
			return l, nil
		}
	}
	return base.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *constVar) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constVar) IsTrivial() (yes bool, value bool) {
	return false, false
}
