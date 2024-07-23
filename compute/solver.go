package compute

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

type Encodable interface {
	Encoding(ctx query.QContext) (cnf.CNF, error)
}

// Steps returns the exitcode and stdout returned by the SAT solver specified
// by solverPath after running it over the encoding of the formula f. The exit
// code will always be 10 (sat) or 20 (unsat) if error == nil. filePath will be
// used to save the encoding of the formula.
func Step(
	f Encodable,
	ctx query.QContext,
	solverPath, filePath string,
) (int, []byte, error) {
	// Write formula as cnf to filePath
	cnf, err := f.Encoding(ctx)
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

// runSolver takes in the solver as an os/exec command and returns exit code
// and stdout as bytes.
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

// GetValueFromBytes returns the value of variable v from a solvers output
// passed as a byte slice.
func GetValueFromBytes(
	out []byte,
	v query.QVar,
	ctx query.QContext,
) (query.QConst, error) {
	var c query.QConst
	vval := make(map[int]bool)
	vfilter := variableFilter(v, ctx)

	for _, line := range bytes.Split(out, []byte("\n")) {
		if len(line) == 0 || line[0] != byte(118) {
			continue
		}
		for _, vb := range bytes.Split(line[2:], []byte(" ")) {
			val, err := strconv.Atoi(string(vb))
			if err != nil {
				return query.QConst{}, err
			}
			absVal := iAbs(val)
			if vfilter[absVal] {
				vval[absVal] = val > 0
			}
		}
	}

	for i := 0; i < ctx.Dim(); i++ {
		if vval[ctx.CNFVar(v, i, int(query.BOT))] {
			c.Val = append(c.Val, query.BOT)
			continue
		}
		if vval[ctx.CNFVar(v, i, int(query.ONE))] {
			c.Val = append(c.Val, query.ONE)
			continue
		}
		if vval[ctx.CNFVar(v, i, int(query.ZERO))] {
			c.Val = append(c.Val, query.ZERO)
			continue
		}
	}

	return c, nil
}

// variableFilter returns a map filter for target variable values.
func variableFilter(v query.QVar, ctx query.QContext) map[int]bool {
	f := make(map[int]bool)

	for i := 0; i < ctx.Dim(); i++ {
		f[ctx.CNFVar(v, i, int(query.BOT))] = true
		f[ctx.CNFVar(v, i, int(query.ZERO))] = true
		f[ctx.CNFVar(v, i, int(query.ONE))] = true
	}

	return f
}

// iAbs returns the absolute value of the integer v.
func iAbs(v int) int {
	if v > 0 {
		return v
	}
	return -v
}
