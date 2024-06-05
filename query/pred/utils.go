package pred

import (
	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// varBotCountClauses generates the counting clauses used to encode the number
// of features with value equal to bot in v. To do this an auxiliary Var cv is
// used.
func varBotCountClauses(
	v query.QVar,
	cv query.QVar,
	ctx query.QContext,
) []cnf.Clause {
	var i, j int

	dim := ctx.Dim()
	clauses := []cnf.Clause{}

	for i = 1; i < dim; i++ {
		clauses = append(
			clauses,
			cnf.Clause{
				-ctx.CNFVar(cv, i, 0),
				-ctx.CNFVar(v, i, int(query.BOT)),
			},
			cnf.Clause{
				-ctx.CNFVar(cv, i, 0),
				ctx.CNFVar(cv, i-1, 0),
			},
			cnf.Clause{
				-ctx.CNFVar(cv, i-1, 0),
				ctx.CNFVar(v, i, int(query.BOT)),
				ctx.CNFVar(cv, i, 0),
			},
		)

		for j = 1; j < i+2; j++ {
			clauses = append(
				clauses,
				cnf.Clause{
					// -ctx.V[('c', var.name, i, j)],
					// ctx.V[('c', var.name, i - 1, j - 1)],
					// -ctx.V[(var.name, i, Symbol.BOT)]
					-ctx.CNFVar(cv, i, j),
					ctx.CNFVar(cv, i-1, j-1),
					-ctx.CNFVar(v, i, int(query.BOT)),
				},
				cnf.Clause{
					// ctx.V[('c', var.name, i, j)],
					// -ctx.V[('c', var.name, i - 1, j - 1)],
					// -ctx.V[(var.name, i, Symbol.BOT)]
					ctx.CNFVar(cv, i, j),
					-ctx.CNFVar(cv, i-1, j-1),
					-ctx.CNFVar(v, i, int(query.BOT)),
				},
				cnf.Clause{
					// -ctx.V[('c', var.name, i, j)],
					// ctx.V[('c', var.name, i - 1, j)],
					// ctx.V[(var.name, i, Symbol.BOT)]
					-ctx.CNFVar(cv, i, j),
					ctx.CNFVar(cv, i-1, j),
					ctx.CNFVar(v, i, int(query.BOT)),
				},
				cnf.Clause{
					// ctx.V[('c', var.name, i, j)],
					// -ctx.V[('c', var.name, i - 1, j)],
					// ctx.V[(var.name, i, Symbol.BOT)]
					ctx.CNFVar(cv, i, j),
					-ctx.CNFVar(cv, i-1, j),
					ctx.CNFVar(v, i, int(query.BOT)),
				},
			)
		}
	}

	for i = 0; i < dim; i++ {
		for j = i + 2; j <= dim; j++ {
			clauses = append(clauses, cnf.Clause{-ctx.CNFVar(cv, i, j)})
		}
	}

	clauses = append(
		clauses,
		cnf.Clause{
			// -ctx.V[('c', var.name, 0, 1)],
			// ctx.V[(var.name, 0, Symbol.BOT)]
			-ctx.CNFVar(cv, 0, 1),
			ctx.CNFVar(v, 0, int(query.BOT)),
		},
		cnf.Clause{
			// -ctx.V[(var.name, 0, Symbol.BOT)],
			// ctx.V[('c', var.name, 0, 1)]
			-ctx.CNFVar(v, 0, int(query.BOT)),
			ctx.CNFVar(cv, 0, 1),
		},
		cnf.Clause{
			// -ctx.V[('c', var.name, 0, 0)],
			// -ctx.V[(var.name, 0, Symbol.BOT)]
			-ctx.CNFVar(cv, 0, 0),
			-ctx.CNFVar(v, 0, int(query.BOT)),
		},
		cnf.Clause{
			// ctx.V[(var.name, 0, Symbol.BOT)],
			// ctx.V[('c', var.name, 0, 0)]
			ctx.CNFVar(v, 0, int(query.BOT)),
			ctx.CNFVar(cv, 0, 0),
		},
	)

	return clauses
}
