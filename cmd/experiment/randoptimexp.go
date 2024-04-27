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

// Method to run on each tree input.
type roeEval func(n int, tf, tpf string, ctx *base.Context, w *csv.Writer) error

// Order minimum 'get stats' experiment.
type randOptimExp struct {
	name      string
	desc      string
	evaluator roeEval
}

// Return new instance of experiment.
func newRandOptimExp(name, desc string, evaluator roeEval) *randOptimExp {
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

	for _, treeFP := range treePaths {
		ctx, err := genContext(treeFP)
		if err != nil {
			return err
		}

		if err = e.evaluator(
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

// Compute value on tree.
func randStatsEvalGen(queryGF closeOptimQueryGenFactory) roeEval {
	return func(
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
			fg, og, err := queryGF(ctx)
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
}

// Compute value on tree.
func randValEvalGen(queryGF closeOptimQueryGenFactory) roeEval {
	return func(
		n int,
		tf, tpf string,
		ctx *base.Context,
		w *csv.Writer,
	) error {
		v := base.Var("x")

		for i := 0; i < n; i++ {
			fg, og, err := queryGF(ctx)
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
}
