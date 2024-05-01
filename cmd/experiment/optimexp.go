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
type oeDriver interface {
	WriteHeader(w *csv.Writer) error
	Eval(tf, tpf string, w *csv.Writer) error
}

// Order minimum 'get stats' experiment.
type optimExp struct {
	name      string
	desc      string
	evaluator oeDriver
}

// Return new instance of experiment.
func newOptimExp(name, desc string, evaluator oeDriver) *optimExp {
	return &optimExp{
		name:      name,
		desc:      desc,
		evaluator: evaluator,
	}
}

// Return experiment name.
func (e *optimExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *optimExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *optimExp) Exec(args ...string) error {
	if len(args) == 0 {
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

	inputPaths, err := filesToPaths(args)
	if err != nil {
		return err
	}

	if err = e.evaluator.WriteHeader(outputWriter); err != nil {
		return err
	}

	for _, inputFP := range inputPaths {
		if err = e.evaluator.Eval(
			inputFP,
			tmpFP,
			outputWriter,
		); err != nil {
			return err
		}
	}

	return nil
}

// =========================== //
//            DRIVERS          //
// =========================== //

type valDriver struct {
	queryGF openOptimQueryGenFactory
}

// Return rand stats evaluator.
func newValDriver(queryGF openOptimQueryGenFactory) valDriver {
	return valDriver{queryGF: queryGF}
}

// Write output header.
func (e valDriver) WriteHeader(
	w *csv.Writer,
) error {
	return w.Write(
		[]string{"file_name", "tree_dim", "tree_nodes", "time", "value"},
	)
}

// Compute value on tree.
func (e valDriver) Eval(inf, tpf string, w *csv.Writer) error {
	instances, ctx, err := parseTIInput(inf)
	if err != nil {
		return err
	}

	v := base.Var("x")

	for _, inst := range instances {
		t := time.Now()

		fg, og, err := e.queryGF(ctx, inst)
		if err != nil {
			return err
		}

		out, err := orderoptimum.Compute(fg, og, v, ctx, SOLVER, tpf)
		if err != nil {
			return fmt.Errorf("Compute error: %s", err.Error())
		}

		ctx.Reset()

		outString := "-"
		if out.Found {
			outString = out.Value.AsString()
		}
		timeString := strconv.Itoa(int(time.Since(t)))

		if err = w.Write(
			[]string{
				inf,
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
