package subsumption

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


// Return constConst subsumption.
func ConstConst(constInst1, constInst2 instances.Const) *constConst {
	return &constConst{constInst1: constInst1, constInst2: constInst2}
}

// Return CNF encoding of component.
func (s *constConst) Encoding(ctx *components.Context) (*cnf.CNF, error) {
	if len(s.constInst1) != ctx.Dimension {
		return nil, errors.New(
			fmt.Sprintf(
				`subsumption.constConst -> constant%d: wrong dim %d
				(%d feats in context)`,
				1,
				len(s.constInst1),
				ctx.Dimension,
			),
		)
	}
	if len(s.constInst2) != ctx.Dimension {
		return nil, errors.New(
			fmt.Sprintf(
				`subsumption.constConst -> constant%d: wrong dim %d
				(%d feats in context)`,
				2,
				len(s.constInst2),
				ctx.Dimension,
			),
		)
	}
	for i, f := range s.constInst1 {
		if f != instances.BOT && f != s.constInst2[i] {
			return cnf.CNFFromClauses([][]int{{}}), nil
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
// This method may change the state of the caller.
func (s *constConst) Simplified() (components.Component, error) {
	if len(s.constInst1) != len(s.constInst2) {
		return nil, errors.New(
			fmt.Sprintf(
				`subsumption.constConst -> constants: different length
				(%d != %d)`,
				len(s.constInst1),
				len(s.constInst2),
			),
		)
	}
	for i, f := range s.constInst1 {
		if f != instances.BOT && f != s.constInst2[i] {
			return components.NewTrivial(false), nil
		}
	}
	return components.NewTrivial(true), nil
}

// Return slice of pointers to component's children.
func (s *constConst) GetChildren() []components.Component {
	return []components.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (s *constConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
