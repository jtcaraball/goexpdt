package sfdtest

import (
	"testing"
	"bytes"
	"os/exec"
	"errors"
)

// Run solver command and return exit code.
func RunSolver(t *testing.T, cmd *exec.Cmd) (int, error) {
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
