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

// =========================== //
//          EXPERIMENT         //
// =========================== //

// Random optimisation experiment evaluator interface.
type roeDriver interface {
	WriteHeader(n int, w *csv.Writer) error
	Eval(n int, tf, tpf string, ctx *base.Context, w *csv.Writer) error
}

// Order minimum 'get stats' experiment.
type randOptimExp struct {
	name      string
	desc      string
	evaluator roeDriver
}

// Return new instance of experiment.
func newRandOptimExp(name, desc string, evaluator roeDriver) *randOptimExp {
	return &randOptimExp{
		name:      name,
		desc:      desc,
		evaluator: evaluator,
	}
}

// Return experiment name.
func (e *randOptimExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *randOptimExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *randOptimExp) Exec(args ...string) error {
	if len(args) < 2 {
		return errors.New("Missing arguments.")
	}

	outFP, tmpFP := fileNames(e.Name())

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

	if err = e.evaluator.WriteHeader(mult, outputWriter); err != nil {
		return err
	}

	for _, treeFP := range treePaths {
		ctx, err := genContext(treeFP)
		if err != nil {
			return err
		}

		if err = e.evaluator.Eval(
			mult,
			treeFP,
			tmpFP,
			ctx,
			outputWriter,
		); err != nil {
			return err
		}

	}

	return nil
}

// =========================== //
//          EVALUATORS         //
// =========================== //

type randStatsDriver struct {
	queryGF closeOptimQueryGenFactory
}

// Return rand stats evaluator.
func newRandStatsDriver(queryGF closeOptimQueryGenFactory) randStatsDriver {
	return randStatsDriver{queryGF: queryGF}
}

// Write output header.
func (e randStatsDriver) WriteHeader(
	n int,
	w *csv.Writer,
) error {
	header := []string{"file_name", "tree_dim", "tree_nodes"}
	for i := 0; i < n; i++ {
		header = append(
			header,
			fmt.Sprintf("bots_%d", i),
			fmt.Sprintf("calls_%d", i),
			fmt.Sprintf("time_%d", i),
		)
	}
	return w.Write(header)
}

// Compute value on tree.
func (e randStatsDriver) Eval(
	n int,
	tf, tpf string,
	ctx *base.Context,
	w *csv.Writer,
) error {
	output := []string{
		tf,
		strconv.Itoa(ctx.Dimension),
		strconv.Itoa(ctx.Tree.NodeCount),
	}

	for i := 0; i < n; i++ {
		fg, og, err := e.queryGF(ctx)
		if err != nil {
			return err
		}

		t := time.Now()

		out, err := orderoptimum.Compute(
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
		output = append(
			output,
			strconv.Itoa(out.Value.BotCount()),
			strconv.Itoa(out.Calls),
			strconv.Itoa(int(ts)),
		)
	}

	if err := w.Write(output); err != nil {
		return err
	}

	w.Flush() // Experiments are long. Save outputs often.

	return nil
}

type randValDriver struct {
	queryGF closeOptimQueryGenFactory
}

// Return rand value evaluator.
func newRandValDriver(queryGF closeOptimQueryGenFactory) randValDriver {
	return randValDriver{queryGF: queryGF}
}

// Write output header.
func (e randValDriver) WriteHeader(
	n int,
	w *csv.Writer,
) error {
	header := []string{"file_name", "tree_dim", "tree_nodes", "time", "value"}
	return w.Write(header)
}

// Compute value on tree.
func (e randValDriver) Eval(
	n int,
	tf, tpf string,
	ctx *base.Context,
	w *csv.Writer,
) error {
	v := base.Var("x")

	for i := 0; i < n; i++ {
		fg, og, err := e.queryGF(ctx)
		if err != nil {
			return err
		}

		t := time.Now()

		out, err := orderoptimum.Compute(fg, og, v, ctx, SOLVER, tpf)
		if err != nil {
			return err
		}

		outString := "-"
		if out.Found {
			outString = out.Value.AsString()
		}
		timeString := strconv.Itoa(int(time.Since(t)))

		if err = w.Write(
			[]string{
				tf,
				strconv.Itoa(ctx.Dimension),
				strconv.Itoa(ctx.Tree.NodeCount),
				timeString,
				outString,
			},
		); err != nil {
			return err
		}

		w.Flush() // Experiments are long. Save outputs often.
		ctx.Reset()
	}

	return nil
}
