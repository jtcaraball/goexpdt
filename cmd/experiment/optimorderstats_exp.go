package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/compute/orderoptimum"
	"os"
	"strconv"
	"time"
)

// Order minimum 'get stats' experiment.
type orderOptimStatsExp struct {
	name      string
	desc      string
	formulaGF optimFormulaGenFactory
	orderGF   optimOrderGenFactory
}

// Return new instance of experiment.
func newOrderOptimStatsExp(
	name, desc string,
	formulaGF optimFormulaGenFactory,
	orderGF optimOrderGenFactory,
) *orderOptimStatsExp {
	return &orderOptimStatsExp{
		name:      name,
		desc:      desc,
		formulaGF: formulaGF,
		orderGF:   orderGF,
	}
}

// Return experiment name.
func (e *orderOptimStatsExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *orderOptimStatsExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *orderOptimStatsExp) Exec(args ...string) error {
	if len(args) < 2 {
		return errors.New("Missing arguments.")
	}

	outFP, tmpFP := fileNames("order_optim_stats")

	outputFile, err := os.Create(outFP)
	if err != nil {
		return err
	}
	defer func() {
		outputFile.Close()
		os.Remove(tmpFP)
	}()

	outputWriter := csv.NewWriter(outputFile)

	mult, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid mult '%s'", args[0])
	}

	treePaths, err := filesToPaths(args[1:])
	if err != nil {
		return err
	}

	for _, treeFP := range treePaths {
		ctx, err := genContext(treeFP)
		if err != nil {
			return err
		}

		if err = e.evalOnTree(
			outputWriter,
			mult,
			ctx,
			treeFP,
			tmpFP,
		); err != nil {
			return err
		}

	}

	return nil
}

// Compute value on tree.
func (e *orderOptimStatsExp) evalOnTree(
	w *csv.Writer,
	n int,
	ctx *base.Context,
	tf, tpf string,
) error {
	count := 0
	var (
		min      time.Duration = time.Duration(1<<63 - 1)
		max, avg time.Duration
	)

	c := base.AllBotConst(ctx.Dimension)

	og, err := e.orderGF()
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		randConst(c, true)

		fg, err := e.formulaGF(c)
		if err != nil {
			return err
		}

		t := time.Now()

		_, _, err = orderoptimum.Compute(
			fg,
			og,
			base.Var("x"),
			ctx,
			SOLVER,
			tpf,
		)
		if err != nil {
			return fmt.Errorf("Compute error: %s", err.Error())
		}

		ctx.Reset()

		ts := time.Since(t)
		min = dMin(min, ts)
		max = dMax(max, ts)
		avg += ts
		count += 1
	}

	if err := w.Write(
		[]string{
			tf,
			strconv.Itoa(ctx.Dimension),
			strconv.Itoa(ctx.Tree.NodeCount),
			strconv.Itoa(int(min)),
			strconv.Itoa(int(max)),
			strconv.Itoa(int(avg) / count),
		},
	); err != nil {
		return err
	}

	w.Flush() // Experiments are long. Save outputs often.

	return nil
}
