package lel

import (
	"stratifoiled/components"
	"stratifoiled/components/instances"
)

func genCountClauses(varName string, ctx *components.Context) [][]int {
	var i, j int
	clauses := [][]int{}
	cVarName := "c" + varName
	for i = 1; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.IVar(cVarName, i, 0),
				-ctx.Var(varName, i, instances.BOT.Val()),
			},
			[]int{
				-ctx.IVar(cVarName, i, 0),
				ctx.IVar(cVarName, i - 1, 0),
			},
			[]int{
				-ctx.IVar(cVarName, i - 1, 0),
				ctx.Var(varName, i, instances.BOT.Val()),
				ctx.IVar(cVarName, i, 0),
			},
		)
		for j = 1; j < i + 2; j++ {
			clauses = append(
				clauses,
				[]int{
					// -context.V[('c', var.name, i, j)],
					// context.V[('c', var.name, i - 1, j - 1)],
					// -context.V[(var.name, i, Symbol.BOT)]
					-ctx.IVar(cVarName, i, j),
					ctx.IVar(cVarName, i - 1, j - 1),
					-ctx.Var(varName, i, instances.BOT.Val()),
				},
				[]int {
					// context.V[('c', var.name, i, j)],
					// -context.V[('c', var.name, i - 1, j - 1)],
					// -context.V[(var.name, i, Symbol.BOT)]
					ctx.IVar(cVarName, i, j),
					-ctx.IVar(cVarName, i - 1, j - 1),
					-ctx.Var(varName, i, instances.BOT.Val()),
				},
				[]int{
					// -context.V[('c', var.name, i, j)],
					// context.V[('c', var.name, i - 1, j)],
					// context.V[(var.name, i, Symbol.BOT)]
					-ctx.IVar(cVarName, i, j),
					ctx.IVar(cVarName, i - 1, j),
					ctx.Var(varName, i, instances.Bot().Val()),
				},
				[]int {
					// context.V[('c', var.name, i, j)],
					// -context.V[('c', var.name, i - 1, j)],
					// context.V[(var.name, i, Symbol.BOT)]
					ctx.IVar(cVarName, i, j),
					-ctx.IVar(cVarName, i - 1, j),
					ctx.Var(varName, i, instances.BOT.Val()),
				},
			)
		}
	}
	for i = 0; i < ctx.Dimension; i++ {
		for j = i + 2; j <= ctx.Dimension; j++ {
			clauses = append(clauses, []int{-ctx.IVar(cVarName, i, j)})
		}
	}
	clauses = append(
		clauses,
		[]int{
			// -context.V[('c', var.name, 0, 1)],
			// context.V[(var.name, 0, Symbol.BOT)]
			-ctx.IVar(cVarName, 0, 1),
			ctx.Var(varName, 0, instances.BOT.Val()),
		},
		[]int{
			// -context.V[(var.name, 0, Symbol.BOT)],
			// context.V[('c', var.name, 0, 1)]
			-ctx.Var(varName, 0, instances.BOT.Val()),
			ctx.IVar(cVarName, 0, 1),
		},
		[]int{
			// -context.V[('c', var.name, 0, 0)],
			// -context.V[(var.name, 0, Symbol.BOT)]
			-ctx.IVar(cVarName, 0, 0),
			-ctx.Var(varName, 0, instances.BOT.Val()),
		},
		[]int{
			// context.V[(var.name, 0, Symbol.BOT)],
			// context.V[('c', var.name, 0, 0)]
			ctx.Var(varName, 0, instances.BOT.Val()),
			ctx.IVar(cVarName, 0, 0),
		},
	)
	return clauses
}
