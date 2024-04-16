package main

import (
	"encoding/csv"
	"errors"
	"goexpdt/base"
	"goexpdt/compute/orderoptimum"
	"os"
	"path"
	"strconv"
	"time"
)

// Order minimum experiment.
type orderOptimExp struct {
	name    string
	desc    string
	formula orderoptimum.VFormula
	order   orderoptimum.VCOrder
}

// Return new instance of experiment.
func newOrderOptimExp(
	name, desc string,
	formula orderoptimum.VFormula,
	order orderoptimum.VCOrder,
) *orderOptimExp {
	return &orderOptimExp{
		name:    name,
		desc:    desc,
		formula: formula,
		order:   order,
	}
}

// Return experiment name.
func (e *orderOptimExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *orderOptimExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *orderOptimExp) Exec(args ...string) error {
	if len(args) == 0 {
		return errors.New("Missing tree file names.")
	}

	outFP, tmpFP := fileNames("order_optim_")

	outputFile, err := os.Create(outFP)
	if err != nil {
		return err
	}
	defer func() {
		outputFile.Close()
		os.Remove(tmpFP)
	}()

	outputWriter := csv.NewWriter(outputFile)

	for _, treeFP := range args {
		ctx, err := genContext(path.Join(INPUTDIR, treeFP))
		if err != nil {
			return err
		}

		if err = e.evalOnTree(outputWriter, ctx, treeFP, tmpFP); err != nil {
			return err
		}

	}

	return nil
}

// Compute value on tree.
func (e *orderOptimExp) evalOnTree(
	w *csv.Writer,
	ctx *base.Context,
	tf, tpf string,
) error {
	t := time.Now()

	_, out, err := orderoptimum.Compute(
		e.formula,
		e.order,
		base.Var("x"),
		ctx,
		SOLVER,
		tpf,
	)
	if err != nil {
		return err
	}

	if err = e.writeOut(w, tf, t, out); err != nil {
		return err
	}
	w.Flush() // Experiments are long. Save outputs often.

	return nil
}

// Write compute output to csv writer.
func (e *orderOptimExp) writeOut(
	w *csv.Writer,
	tp string,
	t time.Time,
	out base.Const,
) error {
	outString := "-"
	if out != nil {
		outString = out.AsString()
	}
	timeString := strconv.Itoa(int(time.Since(t)))
	return w.Write([]string{tp, timeString, outString})
}
