package main

import (
	"encoding/csv"
	"goexpdt/base"
	"goexpdt/circuits/extensions/dft"
	"goexpdt/circuits/predicates/lel"
	"goexpdt/compute/orderoptimum"
	"goexpdt/operators"
	"os"
	"path"
	"strconv"
	"time"
)

// DFT LEL order minimum.
type orderOptimExp struct {
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
) *orderOptimExp {
	return &orderOptimExp{name: name, desc: desc, formula: formula, order: order}
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
	outFP, tmpFP := e.fileNames()

	outputFile, err := os.Create(outFP)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	outputWriter := csv.NewWriter(outputFile)

	for _, treeFileName := range args {
		ctx, err := genContext(path.Join(INPUTDIR, treeFileName))
		if err != nil {
			return err
		}

		t := time.Now()

		_, out, err := orderoptimum.Compute(
			e.formula,
			e.order,
			base.Var("x"),
			ctx,
			SOLVER,
			tmpFP,
		)
		if err != nil {
			return err
		}

		if err = e.writeOut(outputWriter, treeFileName, t, out); err != nil {
			return err
		}
		outputWriter.Flush() // Experiments are long. Save outputs often.
	}

	if err = os.Remove(tmpFP); err != nil {
		return err
	}
	return nil
}

// Return output and temporal file names.
func (e *orderOptimExp) fileNames() (string, string) {
	expTS := time.Now().String()
	return path.Join(OUTPUTDIR, "oderoptim_"+expTS), "tmp_" + expTS
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
	timeString := strconv.Itoa(int(time.Second * time.Since(t)))
	return w.Write([]string{tp, timeString, outString})
}

// =========================== //
//         FORMULA GEN         //
// =========================== //

func dftFGen(v base.Var) base.Component {
	return dft.Var(v)
}

// =========================== //
//          ORDER GEN          //
// =========================== //

func lelOGen(v base.Var, c base.Const) base.Component {
	return operators.And(
		lel.VarConst(v, c),
		operators.Not(lel.ConstVar(c, v)),
	)
}
