package compute_test

import (
	"slices"
	"testing"

	"github.com/jtcaraball/goexpdt/compute"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/logop"
	"github.com/jtcaraball/goexpdt/query/predicates/lel"
	"github.com/jtcaraball/goexpdt/query/predicates/subsumption"
)

func varGenBotCount(v query.QVar) query.QVar {
	return query.QVar("botc" + string(rune(30)) + string(v))
}

func lelFormula(v query.QVar) compute.Encodable {
	upperEq := query.QConst{
		Val: []query.FeatV{query.ONE, query.ONE, query.ONE},
	}
	lowerStrict := query.QConst{
		Val: []query.FeatV{query.BOT, query.BOT, query.BOT},
	}
	return logop.And{
		Q1: subsumption.VarConst{I1: v, I2: upperEq},
		Q2: logop.Not{Q: subsumption.VarConst{I1: v, I2: lowerStrict}},
	}
}

func lelOrder(v query.QVar, c query.QConst) compute.Encodable {
	return logop.And{
		Q1: lel.VarConst{I1: v, I2: c, CountVarGen: varGenBotCount},
		Q2: logop.Not{
			Q: lel.ConstVar{I1: c, I2: v, CountVarGen: varGenBotCount},
		},
	}
}

func TestCompute_LEL(t *testing.T) {
	v := query.QVar("x")
	validOuts := []query.QConst{
		{Val: []query.FeatV{query.BOT, query.BOT, query.ONE}},
		{Val: []query.FeatV{query.BOT, query.ONE, query.BOT}},
		{Val: []query.FeatV{query.ONE, query.BOT, query.BOT}},
	}
	ctx := query.BasicQContext(mockModel{dim: 3})

	solver, _ := compute.NewBinSolver(SOLVER)

	out, err := compute.ComputeOptim(lelFormula, lelOrder, v, ctx, solver)
	if err != nil {
		t.Error(err.Error())
		return
	}

	if !out.Found {
		t.Error("Invalid output. Value could not be found.")
		return
	}
	if !slices.ContainsFunc(
		validOuts,
		func(c query.QConst) bool {
			return c.EqualValue(out.Value)
		},
	) {
		t.Errorf("Invalid output. %d", out.Value.Val)
		return
	}
}
