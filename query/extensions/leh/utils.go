package leh

import (
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// hammingDistCC returns the hamming distance between to constant instances.
// Return error if constants have different dimensions.
func hammingDistCC(c1, c2 query.QConst) (int, error) {
	if len(c1.Val) != len(c2.Val) {
		return 0, fmt.Errorf(
			"Mismatched constant dimensions: %d - %d",
			len(c1.Val),
			len(c2.Val),
		)
	}

	hd := 0

	for i, ft := range c1.Val {
		if ft != c2.Val[i] {
			hd += 1
		}
	}

	return hd, nil
}

// hammingDistVC returns the clauses encoding the hamming distance between the
// query variable v and the constant c. The query variable vcd is used to
// represent the value of this distance.
func hammingDistVC(
	v query.QVar,
	c query.QConst,
	vcd query.QVar,
	ctx query.QContext,
) ([]cnf.Clause, error) {
	dim := ctx.Dim()
	clauses := []cnf.Clause{}

	if err := query.ValidateConstsDim(ctx.Dim(), c); err != nil {
		return clauses, err
	}

	clauses = append(
		clauses,
		cnf.Clause{-ctx.CNFVar(vcd, 0, 0), ctx.CNFVar(v, 0, int(c.Val[0]))},
		cnf.Clause{-ctx.CNFVar(v, 0, int(c.Val[0])), ctx.CNFVar(vcd, 0, 0)},
		cnf.Clause{-ctx.CNFVar(vcd, 0, 1), -ctx.CNFVar(v, 0, int(c.Val[0]))},
		cnf.Clause{ctx.CNFVar(v, 0, int(c.Val[0])), ctx.CNFVar(vcd, 0, 1)},
	)

	for i := 1; i < dim; i++ {
		clauses = append(
			clauses,
			cnf.Clause{
				-ctx.CNFVar(vcd, i, 0),
				ctx.CNFVar(vcd, i-1, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(vcd, i, 0),
				ctx.CNFVar(v, i, int(c.Val[i])),
			},
			cnf.Clause{
				-ctx.CNFVar(vcd, i-1, 0),
				-ctx.CNFVar(v, i, int(c.Val[i])),
				ctx.CNFVar(vcd, i, 0),
			},
		)
		for j := 1; j <= i+1; j++ {
			clauses = append(
				clauses,
				cnf.Clause{
					-ctx.CNFVar(vcd, i, j),
					ctx.CNFVar(vcd, i-1, j-1),
					ctx.CNFVar(v, i, int(c.Val[i])),
				},
				cnf.Clause{
					-ctx.CNFVar(vcd, i, j),
					ctx.CNFVar(vcd, i-1, j),
					-ctx.CNFVar(v, i, int(c.Val[i])),
				},
				cnf.Clause{
					ctx.CNFVar(vcd, i, j),
					-ctx.CNFVar(vcd, i-1, j-1),
					ctx.CNFVar(v, i, int(c.Val[i])),
				},
				cnf.Clause{
					ctx.CNFVar(vcd, i, j),
					-ctx.CNFVar(vcd, i-1, j),
					-ctx.CNFVar(v, i, int(c.Val[i])),
				},
			)
		}
	}

	for i := 0; i < dim; i++ {
		for j := i + 2; j <= dim; j++ {
			clauses = append(clauses, cnf.Clause{-ctx.CNFVar(vcd, i, j)})
		}
	}

	return clauses, nil
}

// hammingDistVV returns the clauses encoding the hamming distance between the
// query variables v1 and v2. The query variable vd is used to represent the
// value of this distance and vfeq is used to encode v1 and v2 having features
// with equal value.
func hammingDistVV(vd, vfeq query.QVar, ctx query.QContext) []cnf.Clause {
	dim := ctx.Dim()
	clauses := []cnf.Clause{}

	clauses = append(
		clauses,
		cnf.Clause{-ctx.CNFVar(vd, 0, 0), ctx.CNFVar(vfeq, 0, 0)},
		cnf.Clause{-ctx.CNFVar(vfeq, 0, 0), ctx.CNFVar(vd, 0, 0)},
		cnf.Clause{-ctx.CNFVar(vd, 0, 1), -ctx.CNFVar(vfeq, 0, 0)},
		cnf.Clause{ctx.CNFVar(vfeq, 0, 0), ctx.CNFVar(vd, 0, 1)},
	)

	for i := 1; i < dim; i++ {
		clauses = append(
			clauses,
			cnf.Clause{
				-ctx.CNFVar(vd, i, 0),
				ctx.CNFVar(vd, i-1, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(vd, i, 0),
				ctx.CNFVar(vfeq, i, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(vd, i-1, 0),
				-ctx.CNFVar(vfeq, i, 0),
				ctx.CNFVar(vd, i, 0),
			},
		)
		for j := 1; j <= i+1; j++ {
			clauses = append(
				clauses,
				cnf.Clause{
					-ctx.CNFVar(vd, i, j),
					ctx.CNFVar(vd, i-1, j-1),
					ctx.CNFVar(vfeq, i, 0),
				},
				cnf.Clause{
					-ctx.CNFVar(vd, i, j),
					ctx.CNFVar(vd, i-1, j),
					-ctx.CNFVar(vfeq, i, 0),
				},
				cnf.Clause{
					ctx.CNFVar(vd, i, j),
					-ctx.CNFVar(vd, i-1, j-1),
					ctx.CNFVar(vfeq, i, 0),
				},
				cnf.Clause{
					ctx.CNFVar(vd, i, j),
					-ctx.CNFVar(vd, i-1, j),
					-ctx.CNFVar(vfeq, i, 0),
				},
			)
		}
	}

	for i := 0; i < dim; i++ {
		for j := i + 2; j <= dim; j++ {
			clauses = append(clauses, cnf.Clause{-ctx.CNFVar(vd, i, j)})
		}
	}

	return clauses
}

// varFullClauses returns the clauses encoding that the query variable v has
// no features with bottom value.
func varFullClauses(v query.QVar, ctx query.QContext) []cnf.Clause {
	clauses := []cnf.Clause{}

	for i := 0; i < ctx.Dim(); i++ {
		clauses = append(
			clauses,
			cnf.Clause{-ctx.CNFVar(v, i, int(query.BOT))},
		)
	}

	return clauses
}

// fullVarEqualClauses returns clauses encoding that the query variables v1 and
// v2 have the same value on every features. Assumes that both v1 and v2 are
// will be forced to be full. The variable vfeq is used to encode v1 and v2
// having features with equal value.
func fullVarEqualClauses(
	v1, v2, vfeq query.QVar,
	ctx query.QContext,
) []cnf.Clause {
	clauses := []cnf.Clause{}

	for i := 0; i < ctx.Dim(); i++ {
		clauses = append(
			clauses,
			cnf.Clause{
				-ctx.CNFVar(v1, i, int(query.ONE)),
				-ctx.CNFVar(v2, i, int(query.ONE)),
				ctx.CNFVar(vfeq, i, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(v1, i, int(query.ONE)),
				-ctx.CNFVar(v2, i, int(query.ZERO)),
				-ctx.CNFVar(vfeq, i, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(v1, i, int(query.ZERO)),
				-ctx.CNFVar(v2, i, int(query.ONE)),
				-ctx.CNFVar(vfeq, i, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(v1, i, int(query.ZERO)),
				-ctx.CNFVar(v2, i, int(query.ZERO)),
				ctx.CNFVar(vfeq, i, 0),
			},
		)
	}

	return clauses
}

// // Return hamming distance variable name.
// func distVarName(name1, name2 string) string {
// 	return "d$" + name1 + "$" + name2
// }

// // Return eq variable name.
// func eqVarName(name1, name2 string) string {
// 	return "eq$" + name1 + "$" + name2
// }

// Return the clauses encoding the hamming distance between a variable and a
// constant.
// func hammingDistVC(
// 	varInst base.Var,
// 	constInst base.Const,
// 	ctx *base.Context,
// ) ([][]int, error) {
// 	if err := base.ValidateConstsDim(ctx.Dimension, constInst); err != nil {
// 		return nil, err
// 	}
// 	dvName := distVarName(string(varInst), constInst.AsString())
// 	clauses := [][]int{
// 		{
// 			-ctx.IVar(dvName, 0, 0),
// 			ctx.Var(string(varInst), 0, constInst[0].Val()),
// 		},
// 		{
// 			-ctx.Var(string(varInst), 0, constInst[0].Val()),
// 			ctx.IVar(dvName, 0, 0),
// 		},
// 		{
// 			-ctx.IVar(dvName, 0, 1),
// 			-ctx.Var(string(varInst), 0, constInst[0].Val()),
// 		},
// 		{
// 			ctx.Var(string(varInst), 0, constInst[0].Val()),
// 			ctx.IVar(dvName, 0, 1),
// 		},
// 	}
// 	for i := 1; i < ctx.Dimension; i++ {
// 		clauses = append(
// 			clauses,
// 			[]int{
// 				-ctx.IVar(dvName, i, 0),
// 				ctx.IVar(dvName, i-1, 0),
// 			},
// 			[]int{
// 				-ctx.IVar(dvName, i, 0),
// 				ctx.Var(string(varInst), i, constInst[i].Val()),
// 			},
// 			[]int{
// 				-ctx.IVar(dvName, i-1, 0),
// 				-ctx.Var(string(varInst), i, constInst[i].Val()),
// 				ctx.IVar(dvName, i, 0),
// 			},
// 		)
// 		for j := 1; j <= i+1; j++ {
// 			clauses = append(
// 				clauses,
// 				[]int{
// 					-ctx.IVar(dvName, i, j),
// 					ctx.IVar(dvName, i-1, j-1),
// 					ctx.Var(string(varInst), i, constInst[i].Val()),
// 				},
// 				[]int{
// 					-ctx.IVar(dvName, i, j),
// 					ctx.IVar(dvName, i-1, j),
// 					-ctx.Var(string(varInst), i, constInst[i].Val()),
// 				},
// 				[]int{
// 					ctx.IVar(dvName, i, j),
// 					-ctx.IVar(dvName, i-1, j-1),
// 					ctx.Var(string(varInst), i, constInst[i].Val()),
// 				},
// 				[]int{
// 					ctx.IVar(dvName, i, j),
// 					-ctx.IVar(dvName, i-1, j),
// 					-ctx.Var(string(varInst), i, constInst[i].Val()),
// 				},
// 			)
// 		}
// 	}
// 	for i := 0; i < ctx.Dimension; i++ {
// 		for j := i + 2; j <= ctx.Dimension; j++ {
// 			clauses = append(clauses, []int{-ctx.IVar(dvName, i, j)})
// 		}
// 	}
// 	return clauses, nil
// }

// Return the clauses encoding the hamming distance between two variables.
// func hammingDistVV(varInst1, varInst2 base.Var, ctx *base.Context) [][]int {
// 	clauses := [][]int{}
// 	dvName := distVarName(string(varInst1), string(varInst2))
// 	evName := eqVarName(string(varInst1), string(varInst2))
// 	clauses = append(
// 		clauses,
// 		[]int{-ctx.IVar(dvName, 0, 0), ctx.IVar(evName, 0, 0)},
// 		[]int{-ctx.IVar(evName, 0, 0), ctx.IVar(dvName, 0, 0)},
// 		[]int{-ctx.IVar(dvName, 0, 1), -ctx.IVar(evName, 0, 0)},
// 		[]int{ctx.IVar(evName, 0, 0), ctx.IVar(dvName, 0, 1)},
// 	)
// 	for i := 1; i < ctx.Dimension; i++ {
// 		clauses = append(
// 			clauses,
// 			[]int{-ctx.IVar(dvName, i, 0), ctx.IVar(dvName, i-1, 0)},
// 			[]int{-ctx.IVar(dvName, i, 0), ctx.IVar(evName, i, 0)},
// 			[]int{
// 				-ctx.IVar(dvName, i-1, 0),
// 				-ctx.IVar(evName, i, 0),
// 				ctx.IVar(dvName, i, 0),
// 			},
// 		)
// 		for j := 1; j <= i+1; j++ {
// 			clauses = append(
// 				clauses,
// 				[]int{
// 					-ctx.IVar(dvName, i, j),
// 					ctx.IVar(dvName, i-1, j-1),
// 					ctx.IVar(evName, i, 0),
// 				},
// 				[]int{
// 					-ctx.IVar(dvName, i, j),
// 					ctx.IVar(dvName, i-1, j),
// 					-ctx.IVar(evName, i, 0),
// 				},
// 				[]int{
// 					ctx.IVar(dvName, i, j),
// 					-ctx.IVar(dvName, i-1, j-1),
// 					ctx.IVar(evName, i, 0),
// 				},
// 				[]int{
// 					ctx.IVar(dvName, i, j),
// 					-ctx.IVar(dvName, i-1, j),
// 					-ctx.IVar(evName, i, 0),
// 				},
// 			)
// 		}
// 	}
// 	for i := 0; i < ctx.Dimension; i++ {
// 		for j := i + 2; j <= ctx.Dimension; j++ {
// 			clauses = append(clauses, []int{-ctx.IVar(dvName, i, j)})
// 		}
// 	}
// 	return clauses
// }

// Return the clauses encoding that variable is full.
// func varFullClauses(varInst base.Var, ctx *base.Context) [][]int {
// 	clauses := [][]int{}
// 	for i := 0; i < ctx.Dimension; i++ {
// 		clauses = append(
// 			clauses,
// 			[]int{-ctx.Var(string(varInst), i, base.BOT.Val())},
// 		)
// 	}
// 	return clauses
// }

// Return the clauses encoding equality in variables features. Assumes both
// variables are full.
// func fullVarEqualClauses(
// 	varInst1, varInst2 base.Var,
// 	ctx *base.Context,
// ) [][]int {
// 	clauses := [][]int{}
// 	eqName := eqVarName(string(varInst1), string(varInst2))
// 	for i := 0; i < ctx.Dimension; i++ {
// 		clauses = append(
// 			clauses,
// 			[]int{
// 				-ctx.Var(string(varInst1), i, base.ONE.Val()),
// 				-ctx.Var(string(varInst2), i, base.ONE.Val()),
// 				ctx.IVar(eqName, i, 0),
// 			},
// 			[]int{
// 				-ctx.Var(string(varInst1), i, base.ONE.Val()),
// 				-ctx.Var(string(varInst2), i, base.ZERO.Val()),
// 				-ctx.IVar(eqName, i, 0),
// 			},
// 			[]int{
// 				-ctx.Var(string(varInst1), i, base.ZERO.Val()),
// 				-ctx.Var(string(varInst2), i, base.ONE.Val()),
// 				-ctx.IVar(eqName, i, 0),
// 			},
// 			[]int{
// 				-ctx.Var(string(varInst1), i, base.ZERO.Val()),
// 				-ctx.Var(string(varInst2), i, base.ZERO.Val()),
// 				ctx.IVar(eqName, i, 0),
// 			},
// 		)
// 	}
// 	return clauses
// }
