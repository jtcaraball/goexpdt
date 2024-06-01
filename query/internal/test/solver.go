package test

import (
	"bytes"
	"errors"
	"os/exec"
	"path"
	"strconv"
	"testing"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

const solverPath string = "/kissat"

type Encodable interface {
	Encoding(ctx query.QContext) (cnf.CNF, error)
}

func EncodeAndRun(t *testing.T, f Encodable, ctx query.QContext, id, code int) {
	d := t.TempDir()
	p := path.Join(d, strconv.Itoa(id)+".cnf")

	ncnf, err := f.Encoding(ctx)
	if err != nil {
		t.Fatalf("Formula encoding error: %s", err.Error())
	}
	if err = ncnf.ToFile(p); err != nil {
		t.Fatalf("CNF writing error: %s", err.Error())
	}

	runFormulaTest(t, code, p)
}

func runFormulaTest(t *testing.T, code int, path string) {
	cmd := exec.Command(solverPath, path)

	ec, err := runSolver(cmd)
	if err != nil {
		t.Fatalf("Solver exec error: %s", err.Error())
	}

	if ec != code {
		t.Errorf("Wrong answer. Expected exit code %d but got %d", code, ec)
	}
}

func runSolver(cmd *exec.Cmd) (int, error) {
	// Caputre stderr to debug kissat errors.
	var s bytes.Buffer
	cmd.Stderr = &s

	err := cmd.Run()

	// Attempt to cast error to exit error in order to read exit code
	if e, ok := err.(*exec.ExitError); ok {
		ec := e.ExitCode()
		// If exit code is not 10 or 20 then the solver failed to run and an
		// error has to have been thrown.
		if ec != 10 && ec != 20 {
			return 0, errors.New(s.String())
		}
		// If it is then we have our answer.
		return ec, nil
	} else if err != nil {
		return 0, err
	}

	return 0, errors.New("Solver exit code could not be recovered")
}
