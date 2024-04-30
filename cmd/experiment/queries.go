package main

import (
	"errors"
	"goexpdt/base"
	"goexpdt/circuits/extensions/allcomp"
	"goexpdt/circuits/extensions/dft"
	"goexpdt/circuits/extensions/full"
	"goexpdt/circuits/extensions/leh"
	"goexpdt/circuits/predicates/lel"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/compute/orderoptimum"
	"goexpdt/operators"
)

type (
	closeOptimQueryGenFactory func(ctx *base.Context) (
		orderoptimum.VFormula,
		orderoptimum.VCOrder,
		error,
	)
	openOptimQueryGenFactory func(ctx *base.Context, cs ...base.Const) (
		orderoptimum.VFormula,
		orderoptimum.VCOrder,
		error,
	)
)

// =========================== //
//        CLOSE QUERIES        //
// =========================== //

func SR_LL_C(ctx *base.Context) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	c := base.AllBotConst(ctx.Dimension)
	err := randValConst(c, true, ctx.Tree)
	if err != nil {
		return nil, nil, err
	}
	return srFGF(c), llOGF(), nil
}

func SR_SS_C(ctx *base.Context) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	c := base.AllBotConst(ctx.Dimension)
	err := randValConst(c, true, ctx.Tree)
	if err != nil {
		return nil, nil, err
	}
	return srFGF(c), ssOGF(), nil
}

func DFS_LL_C(ctx *base.Context) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	return dfsFGF(), llOGF(), nil
}

func CR_LH_C(ctx *base.Context) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	c := base.AllBotConst(ctx.Dimension)
	err := randValConst(c, true, ctx.Tree)
	if err != nil {
		return nil, nil, err
	}
	return crFGF(c), lhOGF(c), nil
}

func CA_GH_C(ctx *base.Context) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	c := base.AllBotConst(ctx.Dimension)
	err := randValConst(c, true, ctx.Tree)
	if err != nil {
		return nil, nil, err
	}
	return caFGF(c), ghOGF(c), nil
}

// =========================== //
//         OPEN QUERIES        //
// =========================== //

func CR_LH_O(ctx *base.Context, cs ...base.Const) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	if len(cs) == 0 {
		return nil, nil, errors.New("Missing constant in query factory.")
	}
	return crFGF(cs[0]), lhOGF(cs[0]), nil
}

func CA_GH_O(ctx *base.Context, cs ...base.Const) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	if len(cs) == 0 {
		return nil, nil, errors.New("Missing constant in query factory.")
	}
	return caFGF(cs[0]), ghOGF(cs[0]), nil
}

func SR_LL_O(ctx *base.Context, cs ...base.Const) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	return srFGF(cs[0]), llOGF(), nil
}

func SR_SS_O(ctx *base.Context, cs ...base.Const) (
	orderoptimum.VFormula,
	orderoptimum.VCOrder,
	error,
) {
	return srFGF(cs[0]), ssOGF(), nil
}

// =========================== //
//      VAR FORMULA GEN        //
// =========================== //

func dfsFGF() orderoptimum.VFormula {
	return func(v base.Var) base.Component {
		return dft.Var(v)
	}
}

func srFGF(c base.Const) orderoptimum.VFormula {
	return func(v base.Var) base.Component {
		return operators.WithVar(
			v,
			operators.And(
				subsumption.VarConst(v, c),
				operators.And(
					operators.Or(
						operators.Not(allcomp.Const(c, true)),
						allcomp.Var(v, true),
					),
					operators.Or(
						operators.Not(allcomp.Const(c, false)),
						allcomp.Var(v, false),
					),
				),
			),
		)
	}
}

func crFGF(c base.Const) orderoptimum.VFormula {
	return func(v base.Var) base.Component {
		return operators.WithVar(
			v,
			operators.And(
				full.Var(v),
				operators.And(
					full.Const(c),
					operators.Or(
						operators.And(
							allcomp.Var(v, true),
							operators.Not(allcomp.Const(c, true)),
						),
						operators.And(
							allcomp.Const(c, true),
							operators.Not(allcomp.Var(v, true)),
						),
					),
				),
			),
		)
	}
}

func caFGF(c base.Const) orderoptimum.VFormula {
	return func(v base.Var) base.Component {
		return operators.WithVar(
			v,
			operators.And(
				full.Var(v),
				operators.And(
					full.Const(c),
					operators.And(
						operators.Or(
							operators.Not(allcomp.Var(v, true)),
							allcomp.Const(c, true),
						),
						operators.Or(
							operators.Not(allcomp.Const(c, true)),
							allcomp.Var(v, true),
						),
					),
				),
			),
		)
	}
}

// =========================== //
//          ORDER GEN          //
// =========================== //

func llOGF() orderoptimum.VCOrder {
	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			lel.VarConst(v, c),
			operators.Not(lel.ConstVar(c, v)),
		)
	}
}

func ssOGF() orderoptimum.VCOrder {
	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			subsumption.VarConst(v, c),
			operators.Not(subsumption.ConstVar(c, v)),
		)
	}
}

func lhOGF(cp base.Const) orderoptimum.VCOrder {
	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			leh.ConstVarConst(cp, v, c),
			operators.Not(leh.ConstConstVar(cp, c, v)),
		)
	}
}

func ghOGF(cp base.Const) orderoptimum.VCOrder {
	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			leh.ConstConstVar(cp, c, v),
			operators.Not(leh.ConstVarConst(cp, v, c)),
		)
	}
}
