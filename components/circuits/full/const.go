package full

import (
	"fmt"
	"errors"
	"stratifoiled/cnf"
	"stratifoiled/components"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type fConst struct {
	constInst components.Const
}

// =========================== //
//           METHODS           //
// =========================== //

// Return const full.
func Const(constInst components.Const) *fConst {
	return &fConst{constInst: constInst}
}

// Return CNF encoding of component.
func (f *fConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if err := f.validateInstances(ctx); err != nil {
		return nil, err
	}
	for _, ft := range f.constInst {
		if ft == components.BOT {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (f *fConst) Simplified(
	ctx *components.Context,
) (components.Component, error) {
	if err := f.validateInstances(ctx); err != nil {
		return nil, err
	}
	for _, ft := range f.constInst {
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

func (f *fConst) validateInstances(ctx *components.Context) error {
	if len(f.constInst) != ctx.Dimension {
		return errors.New(
			fmt.Sprintf(
				"full -> constant: wrong dim %d (%d feats in context)",
				len(f.constInst),
				ctx.Dimension,
			),
		)
	}
	return nil
}
