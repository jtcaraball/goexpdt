package orderoptimum

import (
	"goexpdt/base"
	"goexpdt/compute/utils"
	"goexpdt/operators"
)

type (
	// Constructor for component representing a partial strict order between
	// instances represented as variable and constant.
	// Note: If component has a "ForAllGuarded" children there is no guarantee
	// this order will be of polynomial length and if the order is not strict
	// then computation WILL NOT end at all.
	VCOrder func(v base.Var, c base.Const) base.Component
	// Constructor for component representing a single variable formula.
	VFormula func(v base.Var) base.Component
)

// Compute optimal instance that satisfies formula according to order.
func Compute(
	formula VFormula,
	order VCOrder,
	variable base.Var,
	ctx *base.Context,
	solverPath, filePath string,
) (bool, base.Const, error) {
	// Note: the exit code can only be 10 or 20 if step does not return an
	// error so bm will always be evaluated.
	var bm base.Const

	exitcode, out, err := utils.Step(
		operators.WithVar(variable, formula(variable)),
		ctx,
		solverPath,
		filePath,
	)
	if exitcode == 20 { // 20 is the standard unsat code used by solvers.
		return false, nil, nil
	}
	if err != nil {
		return false, nil, err
	}

	for exitcode == 10 { // 10 is the standard sat code used by solvers.
		bm, err = utils.GetValueFromBytes(out, variable, ctx)
		if err != nil {
			return false, nil, err
		}
		ctx.Reset()
		exitcode, out, err = utils.Step(
			operators.WithVar(
				variable,
				operators.And(formula(variable), order(variable, bm)),
			),
			ctx,
			solverPath,
			filePath,
		)
		if err != nil {
			return false, nil, err
		}
	}

	return true, bm, nil
}
