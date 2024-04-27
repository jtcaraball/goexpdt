package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/compute/orderoptimum"
	"os"
	"path"
	"strconv"
	"time"
)

// =========================== //
//          EXPERIMENT         //
// =========================== //

// Method to run on each tree input.
type oeEval func(inf, tpf string, w *csv.Writer) error

// Order minimum 'get stats' experiment.
type optimExp struct {
	name      string
	desc      string
	evaluator oeEval
}

// Return new instance of experiment.
func newOptimExp(name, desc string, evaluator oeEval) *optimExp {
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

	for _, inputFP := range inputPaths {
		if err = e.evaluator(
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
//          EVALUATORS         //
// =========================== //

func valEvalGen(queryGF openOptimQueryGenFactory) oeEval {
	return func(inf, tpf string, w *csv.Writer) error {
		instances, ctx, err := parseInput(inf)
		if err != nil {
			return err
		}

		v := base.Var("x")

		for _, inst := range instances {
			t := time.Now()

			fg, og, err := queryGF(ctx, inst)
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
}

// =========================== //
//            UTILS            //
// =========================== //

func parseInput(inf string) ([]base.Const, *base.Context, error) {
	treeFP, instBytes, err := scanFile(inf)
	if err != nil {
		return nil, nil, err
	}

	ctx, err := genContext(path.Join(INPUTDIR, treeFP))
	if err != nil {
		return nil, nil, err
	}

	instances := make([]base.Const, len(instBytes))
	for i, cb := range instBytes {
		instances[i] = base.AllBotConst(ctx.Dimension)
		err := bToC(cb, instances[i])
		if err != nil {
			return nil, nil, err
		}
	}

	return instances, ctx, nil
}

func scanFile(path string) (string, [][]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return "", nil, errors.New("Empty input file.")
	}
	treeFP := scanner.Text()

	instBytes := [][]byte{}
	for scanner.Scan() {
		instBytes = append(instBytes, scanner.Bytes())
	}
	if len(instBytes) == 0 {
		return "", nil, errors.New("No instances in input file.")
	}

	if scanner.Err() != nil {
		return "", nil, err
	}

	return treeFP, instBytes, nil
}
