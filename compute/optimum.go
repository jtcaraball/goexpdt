package compute

import (
	"os"

	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/logop"
)

// OptOutput corresponds to the output of an optimization computation alongside
// additional information about its execution.
type OptOutput struct {
	// Found is true if a partial instance was computed.
	Found bool
	// Value of the partial instance computed.
	Value query.QConst
	// Calls made to the SAT solver during the computation process.
	Calls int
}

type (
	// VCOrder is a constructor for an encodable formula representing a partial
	// strict order between instances represented as a query variable and a
	// query constant.
	// If the formula contains a "For All Guarded" operator there is no
	// guarantee that the computation will conclude in a polynomial amount of
	// steps and if the order is not strict then computation WILL NOT end at
	// all.
	VCOrder func(v query.QVar, c query.QConst) Encodable
	// SVFormula is a constructor for a single variable formula.
	SVFormula func(v query.QVar) Encodable
)

// ComputeOptim computes the optimal instance that satisfies the formula
// generated by fg according to the order generated by og.
// solverPath indicates the path of the SAT solver executable to be used for
// the computation.
// A query variable v must be provided in order to ensure there are no
// collisions with variables used by the formulas generated by fg and og.
func ComputeOptim(
	fg SVFormula,
	og VCOrder,
	v query.QVar,
	ctx query.QContext,
	solverPath string,
) (OptOutput, error) {
	// Note: the exit code can only be 10 or 20 if step does not return an
	// error so bm will always be evaluated.
	var bm query.QConst

	tmpfp, err := os.CreateTemp("", "tmp.cnf")
	if err != nil {
		return OptOutput{}, nil
	}
	defer func() {
		tmpfp.Close()
		os.Remove(tmpfp.Name())
	}()

	exitcode, out, err := Step(
		logop.WithVar{I: v, Q: fg(v)},
		ctx,
		solverPath,
		tmpfp.Name(),
	)
	if exitcode == 20 { // 20 is the standard unsat code used by solvers.
		return OptOutput{false, query.QConst{}, 1}, nil
	}
	if err != nil {
		return OptOutput{false, query.QConst{}, 0}, err
	}

	steps := 1

	for exitcode == 10 { // 10 is the standard sat code used by solvers.
		bm, err = GetValueFromBytes(out, v, ctx)
		if err != nil {
			return OptOutput{false, query.QConst{}, 0}, err
		}

		ctx.Reset()

		exitcode, out, err = Step(
			logop.WithVar{
				I: v,
				Q: logop.And{
					Q1: fg(v),
					Q2: og(v, bm),
				},
			},
			ctx,
			solverPath,
			tmpfp.Name(),
		)
		if err != nil {
			return OptOutput{false, query.QConst{}, 0}, err
		}

		steps += 1
	}

	return OptOutput{true, bm, steps}, nil
}
