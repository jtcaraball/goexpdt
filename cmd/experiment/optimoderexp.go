package main

import (
	"goexpdt/base"
	"goexpdt/compute/orderoptimum"
	"goexpdt/trees"
	"time"
)

// DFT LEL order minimum.
type oderOptimExp struct {
	name    string
	desc    string
	formula orderoptimum.VFormula
	order   orderoptimum.VCOrder
}

// Return new instance of experiment
func newOrderOptimExp(
	name, desc string,
	formula orderoptimum.VFormula,
	order orderoptimum.VCOrder,
) *oderOptimExp {
	return &oderOptimExp{name: name, desc: desc, formula: formula, order: order}
}

// Return experiment name.
func (e *oderOptimExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *oderOptimExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *oderOptimExp) Exec(args ...string) error {
	v := base.Var("x")
	for _, treePath := range args {
		expT, err := trees.LoadTree(treePath)
		if err != nil {
			return err
		}
		ctx := base.NewContext(expT.FeatCount, expT)

		t := time.Now()
		_, out, err := orderoptimum.Compute(
			e.formula,
			e.order,
			v,
			ctx,
			SOLVER,
			"dftminlel",
		)
		if err != nil {
			return err
		}

		if err = recordOutput(treePath, time.Since(t), out); err != nil {
			return err
		}
	}
	return nil
}

// Record experiment.
func recordOutput(fp string, t time.Duration, val base.Const) error {
	return nil
}
