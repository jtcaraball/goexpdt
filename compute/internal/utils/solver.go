package utils

import (
	"bytes"
	"errors"
	"fmt"
	"goexpdt/base"
	"os/exec"
	"strconv"
)

// Run a compute step. Returns the exitcode and stdout. The exit code will
// always be 10 or 20 if error == nil.
func Step(
	formula base.Component,
	ctx *base.Context,
	solverPath, filePath string,
) (int, []byte, error) {
	// Write formula as cnf to filePath
	cnf, err := formula.Encoding(ctx)
	if err != nil {
		return 0, nil, err
	}
	if err = cnf.ToFile(filePath); err != nil {
		return 0, nil, err
	}
	// Create solver command and run it
	cmd := exec.Command(solverPath, filePath)
	return runSolver(cmd)
}

// Run solver. Takes in the solver an os/exec command and returns exit code and
// stdout as bytes.
func runSolver(cmd *exec.Cmd) (int, []byte, error) {
	// Capturing stderr to debug kissat errors.
	var stderr, stdout bytes.Buffer
	cmd.Stderr = &stderr
	cmd.Stdout = &stdout
	err := cmd.Run()
	// Attempt to cast error to exit error to read exit code
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		// If exit code is not 10 or 20 then solver failed to run and an error
		// has to be thrown
		if exitCode != 10 && exitCode != 20 {
			return exitCode, nil, errors.New(stderr.String())
		}
		// If it is just return it alongside stdout
		return exitErr.ExitCode(), stdout.Bytes(), nil
	} else if err != nil {
		return 0, nil, err
	}
	return 0, nil, errors.New("Solver exit code could not be recovered")
}

// Extract variable v value from output.
func GetValueFromBytes(
	out []byte,
	v base.Var,
	vFilter map[int]bool,
	ctx *base.Context,
) (base.Const, error) {
	varValues := make(map[int]bool)
	for _, line := range bytes.Split(out, []byte("\n")) {
		if len(line) == 0 || line[0] != byte(118) {
			continue
		}
		for _, valAsBytes := range bytes.Split(line[2:], []byte(" ")) {
			val, err := strconv.Atoi(string(valAsBytes))
			if err != nil {
				return nil, err
			}
			absVal := iAbs(val)
			if vFilter[absVal] {
				varValues[absVal] = val > 0
			}
		}
	}
	var evlConst base.Const
	for i := 0; i < ctx.Dimension; i++ {
		if varValues[ctx.Var(string(v), i, base.BOT.Val())] {
			evlConst = append(evlConst, base.BOT)
			continue
		}
		if varValues[ctx.Var(string(v), i, base.ONE.Val())] {
			evlConst = append(evlConst, base.ONE)
			continue
		}
		if varValues[ctx.Var(string(v), i, base.ZERO.Val())] {
			evlConst = append(evlConst, base.ZERO)
			continue
		}
	}
	return evlConst, nil
}

// Generate map filter for target variable values.
func VariableFilter(v base.Var, ctx *base.Context) map[int]bool {
	f := make(map[int]bool)
	for i := 0; i < ctx.Dimension; i++ {
		f[ctx.Var(string(v), i, base.BOT.Val())] = true
		f[ctx.Var(string(v), i, base.ZERO.Val())] = true
		f[ctx.Var(string(v), i, base.ONE.Val())] = true
	}
	return f
}

// Absolute value for integers.
func iAbs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
