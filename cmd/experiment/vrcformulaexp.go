package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/compute/utils"
	"os"
	"path"
	"strconv"
	"time"
)

// Var explicit constant formula experiment.
type vrcFormulaExp struct {
	name    string
	desc    string
	formula vcFormula
}

// Return new instance of experiment.
func newVRCFormulaExp(name, desc string, formula vcFormula) experiment {
	return &vrcFormulaExp{
		name:    name,
		desc:    desc,
		formula: formula,
	}
}

// Return experiment name.
func (e *vrcFormulaExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *vrcFormulaExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *vrcFormulaExp) Exec(args ...string) error {
	if len(args) < 2 {
		return errors.New("Missing arguments.")
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

	nConsts, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("Invalid const count '%s'", args[0])
	}

	for _, treeFP := range args[1:] {
		ctx, err := genContext(path.Join(treeFP))
		if err != nil {
			return err
		}

		if err = e.evalOnTree(
			outputWriter,
			nConsts,
			ctx,
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
func (e *vrcFormulaExp) fileNames() (string, string) {
	expTS := time.Now().String()
	return path.Join(OUTPUTDIR, "vrcformula_"+expTS), "tmp_" + expTS
}

// Run formula over tree and constants.
func (e *vrcFormulaExp) evalOnTree(
	w *csv.Writer,
	n int,
	ctx *base.Context,
	tf, tpf string,
) error {
	v := base.Var("x")
	c := make(base.Const, ctx.Dimension)
	outLine := []string{tf}

	for i := 0; i < n; i++ {
		randConst(c)

		t := time.Now()

		_, _, err := utils.Step(e.formula(v, c), ctx, SOLVER, tpf)
		if err != nil {
			return err
		}

		ctx.Reset()

		outLine = append(outLine, strconv.Itoa(int(time.Since(t)*time.Second)))
	}

	return w.Write(outLine)
}
