package lel

import (
	"fmt"
	"errors"
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type varConst struct {
	varInst instances.Var
	constInst instances.Const
}

// =========================== //
//           METHODS           //
// =========================== //

// Return varConst lel.
func VarConst(varInst instances.Var, constInst instances.Const) *varConst {
	return &varConst{varInst: varInst, constInst: constInst}
}

// Return CNF encoding of component.
func (l *varConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := l.validateInstances(ctx); err != nil {
		return nil, err
	}
	cnf := &cnf.CNF{}
	cnf.ExtendConsistency(genCountClauses(string(l.varInst), ctx))
	// Count amount of bots in constant.
	botsInConst := 0
	for _, f := range l.constInst {
		if f == instances.BOT {
			botsInConst += 1
		}
	}
	// Ask for var to not have fewer bots.
	cVarName := "c" + string(l.varInst)
	for i := 0; i < botsInConst; i++ {
		cnf.AppendSemantics([]int{-ctx.IVar(cVarName, ctx.Dimension - 1, i)})
	}
	return cnf, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (l *varConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := l.validateInstances(ctx); err != nil {
		return nil, err
	}
	for _, f := range l.constInst {
		if f == instances.BOT {
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

func (l *varConst) validateInstances(ctx *components.Context) error {
	if len(l.constInst) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				"lel.varConst -> constant: wrong dim %d (%d feats in context)",
				len(l.constInst),
				ctx.Dimension,
			),
		)
	}
	return nil
}
