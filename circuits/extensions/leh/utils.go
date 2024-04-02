package leh

import (
	"fmt"
	"goexpdt/base"
	"strconv"
)

// Return hamming distance between to constant instances. Return error if
// constants have different dimensions.
func hammingDistCC(constInst1, constInst2 base.Const) (int, error) {
	hDist := 0
	if len(constInst1) != len(constInst2) {
		return 0, fmt.Errorf(
			"Mismatched constant dimensions: %d - %d",
			len(constInst1),
			len(constInst2),
		)
	}
	for i, ft := range constInst1 {
		if ft != constInst2[i] {
			hDist += 1
		}
	}
	return hDist, nil
}

// Return constant as a string.
func constName(c base.Const) string {
	name := ""
	for _, ft := range c {
		name += strconv.Itoa(ft.Val())
	}
	return name
}

// Return hamming distance variable name.
func distVarName(name1, name2 string) string {
	return "d$" + name1 + "$" + name2
}

// Return the clauses encoding the hamming distance between a variable and a
// constant.
func hammingDistVC(
	varInst base.Var,
	constInst base.Const,
	ctx *base.Context,
) ([][]int, error) {
	if err := base.ValidateConstsDim(ctx.Dimension, constInst); err != nil {
		return nil, err
	}
	dvName := distVarName(string(varInst), constName(constInst))
	clauses := [][]int{
		{
			-ctx.IVar(dvName, 0, 0),
			ctx.Var(string(varInst), 0, constInst[0].Val()),
		},
		{
			-ctx.Var(string(varInst), 0, constInst[0].Val()),
			ctx.IVar(dvName, 0, 0),
		},
		{
			-ctx.IVar(dvName, 0, 1),
			-ctx.Var(string(varInst), 0, constInst[0].Val()),
		},
		{
			ctx.Var(string(varInst), 0, constInst[0].Val()),
			ctx.IVar(dvName, 0, 1),
		},
	}
	for i := 1; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.IVar(dvName, i, 0),
				ctx.IVar(dvName, i-1, 0),
			},
			[]int{
				-ctx.IVar(dvName, i, 0),
				ctx.Var(string(varInst), i, constInst[i].Val()),
			},
			[]int{
				-ctx.IVar(dvName, i-1, 0),
				-ctx.Var(string(varInst), i, constInst[i].Val()),
				ctx.IVar(dvName, i, 0),
			},
		)
		for j := 1; j <= i+1; j++ {
			clauses = append(
				clauses,
				[]int{
					-ctx.IVar(dvName, i, j),
					ctx.IVar(dvName, i-1, j-1),
					ctx.Var(string(varInst), i, constInst[i].Val()),
				},
				[]int{
					-ctx.IVar(dvName, i, j),
					ctx.IVar(dvName, i-1, j),
					-ctx.Var(string(varInst), i, constInst[i].Val()),
				},
				[]int{
					ctx.IVar(dvName, i, j),
					-ctx.IVar(dvName, i-1, j-1),
					ctx.Var(string(varInst), i, constInst[i].Val()),
				},
				[]int{
					ctx.IVar(dvName, i, j),
					-ctx.IVar(dvName, i-1, j),
					-ctx.Var(string(varInst), i, constInst[i].Val()),
				},
			)
		}
	}
	for i := 0; i < ctx.Dimension; i++ {
		for j := i + 2; j <= ctx.Dimension; j++ {
			clauses = append(clauses, []int{-ctx.IVar(dvName, i, j)})
		}
	}
	return clauses, nil
}

// Return the clauses encoding the hamming distance between two variables.
func hammingDistVV(varInst1, varInst2 base.Var) ([][]int, error) {
	return nil, nil
}

// Return the clauses encoding that variable is full.
func varFullClauses(varInst base.Var, ctx *base.Context) [][]int {
	clauses := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{-ctx.Var(string(varInst), i, base.BOT.Val())},
		)
	}
	return clauses
}
