package lel

import (
	"errors"
	"fmt"
	"stratifoiled/cnf"
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type constConst struct {
	constInst1 instances.Const
	constInst2 instances.Const
}

// =========================== //
//           METHODS           //
// =========================== //


// Return constConst lel.
func ConstConst(constInst1, constInst2 instances.Const) *constConst {
	return &constConst{constInst1: constInst1, constInst2: constInst2}
}

// Return CNF encoding of component.
func (l *constConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := l.validateInstances(ctx); err != nil {
		return nil, err
	}
	const1Bots := 0
	const2Bots := 0
	for i := 0; i < ctx.Dimension; i++ {
		if l.constInst1[i] == instances.BOT {
			const1Bots += 1
		}
		if l.constInst2[i] == instances.BOT {
			const2Bots += 1
		}
	}
	if const1Bots < const2Bots {
		return cnf.CNFFromClauses([][]int{{}}), nil
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (l *constConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := l.validateInstances(ctx); err != nil {
		return nil, err
	}
	const1Bots := 0
	const2Bots := 0
	for i := 0; i < ctx.Dimension; i++ {
		if l.constInst1[i] == instances.BOT {
			const1Bots += 1
		}
		if l.constInst2[i] == instances.BOT {
			const2Bots += 1
		}
	}
	if const1Bots < const2Bots {
		return components.NewTrivial(false), nil
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (l *constConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (l *constConst) IsTrivial() (yes bool, value bool) {
	return false, false
}

func (l *constConst) validateInstances(ctx *components.Context) error {
	if len(l.constInst1) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				`subsumption.constConst -> constant%d: wrong dim %d
				(%d feats in context)`,
				1,
				len(l.constInst1),
				ctx.Dimension,
			),
		)
	}
	if len(l.constInst2) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				`lel.constConst -> constant%d: wrong dim %d
				(%d feats in context)`,
				2,
				len(l.constInst2),
				ctx.Dimension,
			),
		)
	}
	return nil
}
