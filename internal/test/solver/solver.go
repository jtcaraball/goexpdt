package solver

import (
	"bytes"
	"errors"
	"fmt"
	"goexpdt/base"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

const (
	cnfPath = "tmpCNF"
	solver  = "/kissat"
	delStar = "%s.%s.%t*"
)

func EncodeAndRun(
	t *testing.T,
	formula base.Component,
	context *base.Context,
	filePath string,
	id, expCode int,
	simplify bool,
) {
	var err error
	if simplify {
		formula, err = formula.Simplified(context)
		if err != nil {
			t.Errorf("Formula simplification error. %s", err.Error())
			return
		}
	}
	cnf, err := formula.Encoding(context)
	if err != nil {
		t.Errorf("Formula encoding error. %s", err.Error())
		return
	}
	if err = cnf.ToFile(filePath); err != nil {
		t.Errorf("CNF writing error. %s", err.Error())
		return
	}
	runFormulaTest(t, expCode, filePath)
}

// Encode formula and add error if results is not equal to the expected one.
func runFormulaTest(
	t *testing.T,
	expCode int,
	cnfPath string,
) {
	cmd := exec.Command(solver, cnfPath)
	retCode, err := runSolver(cmd)
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

// Run solver command and return exit code.
func runSolver(cmd *exec.Cmd) (int, error) {
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

// Return cnf file name
func CNFName(sufix string, testID int, simplified bool) string {
	return fmt.Sprintf("%s.%s.%t.%d", cnfPath, sufix, simplified, testID)
}

// Add cleanup for cnf files
func AddCleanup(t *testing.T, sufix string, simplified bool) {
	t.Cleanup(
		func() {
			files, err := filepath.Glob(
				fmt.Sprintf(delStar, cnfPath, sufix, simplified),
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
