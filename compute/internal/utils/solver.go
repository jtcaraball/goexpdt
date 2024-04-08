package utils

import (
	"bytes"
	"errors"
	"goexpdt/base"
	"os/exec"
	"slices"
	"strconv"
	"strings"
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
	ctx *base.Context,
) (base.Const, error) {
	return nil, nil
}

// Absolute value for integers.
func iAbs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
