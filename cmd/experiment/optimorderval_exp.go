package main

import (
	"encoding/csv"
	"errors"
	"goexpdt/base"
	"goexpdt/compute/orderoptimum"
	"os"
	"strconv"
	"time"
)

// Order minimum 'get value' experiment.
type orderOptimValueExp struct {
	name        string
	desc        string
	formulaGF optimFormulaGenFactory
	orderGF   optimOrderGenFactory
}

// Return new instance of experiment.
func newOrderOptimValueExp(
	name, desc string,
	formulaGF optimFormulaGenFactory,
	orderGF optimOrderGenFactory,
) *orderOptimValueExp {
	return &orderOptimValueExp{
		name:        name,
		desc:        desc,
		formulaGF: formulaGF,
		orderGF:   orderGF,
	}
}

// Return experiment name.
func (e *orderOptimValueExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *orderOptimValueExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *orderOptimValueExp) Exec(args ...string) error {
	if len(args) == 0 {
		return errors.New("Missing tree file names.")
	}

	outFP, tmpFP := fileNames("order_optim_val")

	outputFile, err := os.Create(outFP)
	if err != nil {
		return err
	}
	defer func() {
		outputFile.Close()
		os.Remove(tmpFP)
	}()

	outputWriter := csv.NewWriter(outputFile)

	treePaths, err := filesToPaths(args)
	if err != nil {
		return err
	}

	for _, treeFP := range treePaths {
		ctx, err := genContext(treeFP)
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
func (e *orderOptimValueExp) evalOnTree(
	w *csv.Writer,
	ctx *base.Context,
	tf, tpf string,
) error {
	t := time.Now()

	v := base.Var("x")

	fg, err := e.formulaGF()
	if err != nil {
		return err
	}

	og, err := e.orderGF()
	if err != nil {
		return err
	}

	_, out, err := orderoptimum.Compute(fg, og, v, ctx, SOLVER, tpf)
	if err != nil {
		return err
	}

	if err = e.writeOut(w, tf, t, ctx, out); err != nil {
		return err
	}
	w.Flush() // Experiments are long. Save outputs often.

	return nil
}

// Write compute output to csv writer.
func (e *orderOptimValueExp) writeOut(
	w *csv.Writer,
	tp string,
	t time.Time,
	ctx *base.Context,
	out base.Const,
) error {
	outString := "-"
	if out != nil {
		outString = out.AsString()
	}
	timeString := strconv.Itoa(int(time.Since(t)))
	return w.Write(
		[]string{
			tp,
			strconv.Itoa(ctx.Dimension),
			strconv.Itoa(ctx.Tree.NodeCount),
			timeString,
			outString,
		},
	)
}
