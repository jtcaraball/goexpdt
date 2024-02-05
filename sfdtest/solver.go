package sfdtest

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"stratifoiled/components"
	"testing"
)

const (
	CNFPATH = "tmpCNF"
	SOLVER = "/kissat"
	DELSTAR = "%s.%s.%t"
)

// Run solver command and return exit code.
func runSolver(t *testing.T, cmd *exec.Cmd) (int, error) {
	// Capturing stderr to debug kissat errors.
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		if exitCode != 10 && exitCode != 20 {
			t.Error(stderr.String())
		}
		return exitErr.ExitCode(), nil
	} else if err != nil {
		return 0, err
	}
	return 0, errors.New("Solver exit code could not be recovered")
}

// Encode formula and add error if results is not equal to the expected one.
func RunFormulaTest(
	t *testing.T,
	id, expCode int,
	formula components.Component,
	context *components.Context,
	cnfPath string,
) {
	cnf, err := formula.Encoding(context)
	if err = cnf.ToFile(cnfPath); err != nil {
		t.Errorf("CNF writing error: %s", err.Error())
		return
	}
	cmd := exec.Command(SOLVER, cnfPath)
	retCode, err := runSolver(t, cmd)
	if err != nil {
		t.Errorf("Solver execution error: %s", err.Error())
		return
	}
	if retCode != expCode {
		t.Errorf(
			"Wrong answer: expected exit code %d but got %d",
			expCode,
			retCode,
		)
	}
}

// Return cnf file name
func CNFName(sufix string, testID int, simplified bool) string {
	return fmt.Sprintf("%s.%s.%t.%d", CNFPATH, sufix, simplified, testID)
}

// Add cleanup for cnf files
func AddCleanup(t *testing.T, sufix string, simplified bool) {
	t.Cleanup(
		func() {
			files, err := filepath.Glob(
				fmt.Sprintf(DELSTAR, CNFPATH, sufix, simplified),
			)
			if err != nil {
				t.Errorf(fmt.Sprintf("Error in cleanup: %s", err.Error()))
				return
			}
			for _, file := range files {
				os.Remove(file)
			}
		},
	)
}
