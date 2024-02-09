package lel

import (
	"fmt"
	"errors"
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constVar struct {
	varInst components.Var
	constInst components.Const
}

// =========================== //
//           METHODS           //
// =========================== //

// Return constVar lel.
func ConstVar(constInst components.Const, varInst components.Var) *constVar {
	return &constVar{constInst: constInst, varInst: varInst}
}

// Return CNF encoding of component.
func (l *constVar) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := l.validateInstances(ctx); err != nil {
		return nil, err
	}
	cnf := &cnf.CNF{}
	cnf.ExtendConsistency(genCountClauses(string(l.varInst), ctx))
	// Count amount of bots in constant.
	botsInConst := 0
	for _, f := range l.constInst {
		if f == components.BOT {
			botsInConst += 1
		}
	}
	// Ask for var to not have more bots.
	cVarName := "c" + string(l.varInst)
	for i := botsInConst + 1; i < ctx.Dimension + 1; i++ {
		cnf.AppendSemantics([]int{-ctx.IVar(cVarName, ctx.Dimension - 1, i)})
	}
	return cnf, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (l *constVar) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := l.validateInstances(ctx); err != nil {
		return nil, err
	}
	// If const has only bottoms then this predicate is trivialy true.
	for _, f := range l.constInst {
		if f != components.BOT {
			return l, nil
		}
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *constVar) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constVar) IsTrivial() (yes bool, value bool) {
	return false, false
}

func (l *constVar) validateInstances(ctx *components.Context) error {
	if len(l.constInst) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				"lel.constVar -> constant: wrong dim %d (%d feats in context)",
				len(l.constInst),
				ctx.Dimension,
			),
		)
	}
	return nil
}
