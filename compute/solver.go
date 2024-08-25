package compute

import (
	"bytes"
	"errors"
	"os"
	"os/exec"
	"strconv"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// UnsatError is used as a sentinel for query satisfiability computation.
var UnsatError = errors.New("Formula is not satisfiable")

// Encodable query that can be evaluated by a Solver.
type Encodable interface {
	// Encoding of caller as a cnf formula.
	Encoding(ctx query.QContext) (cnf.CNF, error)
}

// Solver allows to determine if a valuation that makes a query true exists and
// subsequently retrieve particular assignation in it based on query variables.
type Solver interface {
	// Values assigned to query variables in the valuation returned by the last
	// successful call to Step. ctx should be the same as the one passed in the
	// Step call.
	Values(vars []query.QVar, ctx query.QContext) ([]query.QConst, error)
	// Step determines if a valuation that makes q true, in ctx, exists and if
	// so records it to be accessed by calling Values. If q is not satisfiable
	// then error == UnsatError.
	Step(q Encodable, ctx query.QContext) error
}

// BinSolver holds a path to a SAT solver's executable binary and its outputs,
// both of which can be accessed making use of the Solver interface.
type BinSolver struct {
	output     []byte
	solverPath string
}

// NewBinSolver returns a new BinSolver. An error is returned if the path
// provided does not correspond to a valid SAT solver executable.
func NewBinSolver(solverPath string) (*BinSolver, error) {
	// TODO: Look into how to determine if the path corresponds to a SAT solver
	// executable.
	if _, err := os.Stat(solverPath); err != nil {
		return nil, err
	}

	return &BinSolver{solverPath: solverPath}, nil
}

func (s *BinSolver) Step(
	q Encodable,
	ctx query.QContext,
) error {
	// Create tmp encoding file
	encFile, err := os.CreateTemp("", "tmp.cnf")
	if err != nil {
		return err
	}
	defer func() {
		encFile.Close()
		os.Remove(encFile.Name())
	}()

	// Write formula as cnf to filePath
	cnf, err := q.Encoding(ctx)
	if err != nil {
		return err
	}
	if err = cnf.ToFile(encFile.Name()); err != nil {
		return err
	}

	// Create solver command and run it
	cmd := exec.Command(s.solverPath, encFile.Name())

	code, output, err := runSolver(cmd)
	if err != nil {
		return err
	}
	if code == 20 {
		return UnsatError
	}
	s.output = output

	return nil
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

func (s *BinSolver) Values(
	vars []query.QVar,
	ctx query.QContext,
) ([]query.QConst, error) {
	out := make([]query.QConst, len(vars))

	for i, v := range vars {
		c, err := varValueFromBytes(s.output, v, ctx)
		if err != nil {
			return nil, err
		}
		out[i] = c
	}

	return out, nil
}

// varValueFromBytes returns the value of variable v from a solvers output
// passed as a byte slice.
func varValueFromBytes(
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
