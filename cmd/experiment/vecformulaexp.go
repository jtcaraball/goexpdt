package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"goexpdt/base"
	"goexpdt/compute/utils"
	"os"
	"path"
	"strconv"
	"time"
)

// Var explicit constant formula experiment.
type vecFormulaExp struct {
	name    string
	desc    string
	formula vcFormula
}

// Return new instance of experiment.
func newVECFormulaExp(name, desc string, formula vcFormula) experiment {
	return &vecFormulaExp{
		name:    name,
		desc:    desc,
		formula: formula,
	}
}

// Return experiment name.
func (e *vecFormulaExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *vecFormulaExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *vecFormulaExp) Exec(args ...string) error {
	if len(args)%2 != 0 {
		return errors.New(
			"Arguments must be pairs of type <treeFile> <constFile>.",
		)
	}

	outFP, tmpFP := e.fileNames()

	outputFile, err := os.Create(outFP)
	if err != nil {
		return err
	}
	defer func() {
		outputFile.Close()
		os.Remove(tmpFP)
	}()

	outputWriter := csv.NewWriter(outputFile)

	for i := 0; i < len(args)/2; i++ {
		treeFP := path.Join(INPUTDIR, args[2*i])
		constFP := path.Join(INPUTDIR, args[2*i+1])

		ctx, err := genContext(treeFP)
		if err != nil {
			return err
		}

		if err = e.evalOnTree(
			outputWriter,
			ctx,
			constFP,
			treeFP,
			tmpFP,
		); err != nil {
			return err
		}

		outputWriter.Flush()
	}

	return nil
}

// Return output and temporal file names.
func (e *vecFormulaExp) fileNames() (string, string) {
	expTS := time.Now().String()
	return path.Join(OUTPUTDIR, "vecformula_"+expTS), "tmp_" + expTS
}

// Run formula over tree and constants.
func (e *vecFormulaExp) evalOnTree(
	w *csv.Writer,
	ctx *base.Context,
	cf, tf, tpf string,
) error {
	f, err := os.Open(cf)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := bufio.NewScanner(f)

	v := base.Var("x")
	c := make(base.Const, ctx.Dimension)
	outLine := []string{tf}

	for reader.Scan() {
		t := time.Now()

		if err = btoC(reader.Bytes(), c); err != nil {
			return err
		}

		_, _, err = utils.Step(e.formula(v, c), ctx, SOLVER, tpf)
		if err != nil {
			return err
		}

		ctx.Reset()

		outLine = append(outLine, strconv.Itoa(int(time.Since(t)*time.Second)))
	}

	return w.Write(outLine)
}
