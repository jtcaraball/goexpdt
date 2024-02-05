package sfdtest

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const (
	CNFPATH = "tmpCNF"
	SOLVER = "/kissat"
	DELSTAR = "%s.%s.%t*"
)

// Run solver command and return exit code.
func runSolver(t *testing.T, cmd *exec.Cmd) (int, error) {
	// Capturing stderr to debug kissat errors.
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	// Attpemt to cast error to exit error to read exit code
	if exitErr, ok := err.(*exec.ExitError); ok {
		exitCode := exitErr.ExitCode()
		// If exit code is not 10 or 20 then solver failed to run and an error
		// has to be thrown
		if exitCode != 10 && exitCode != 20 {
			return 0, errors.New(stderr.String())
		}
		// If it is just return it
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
	cnfPath string,
) {
	cmd := exec.Command(SOLVER, cnfPath)
	retCode, err := runSolver(t, cmd)
	if err != nil {
		t.Errorf("Solver execution error. %s", err.Error())
		return
	}
	if retCode != expCode {
		t.Errorf(
			"Wrong answer. expected exit code %d but got %d",
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
				t.Errorf(fmt.Sprintf("Error in cleanup. %s", err.Error()))
				return
			}
			for _, file := range files {
				os.Remove(file)
			}
		},
	)
}
