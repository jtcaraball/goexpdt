package orderoptimum

import (
	"goexpdt/base"
	"goexpdt/circuits/predicates/lel"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/operators"
	"os"
	"slices"
	"testing"
)

// =========================== //
//           HELPERS           //
// =========================== //

const (
	FILEPATH = "tmpCNF"
	SOLVER   = "/kissat"
)

func equalToConstFunc(c1 base.Const) func(c base.Const) bool {
	return func(c2 base.Const) bool {
		if len(c1) != len(c2) {
			return false
		}
		for i, v := range c1 {
			if v != c2[i] {
				return false
			}
		}
		return true
	}
}

func lelFormula(v base.Var) base.Component {
	upperEq := base.Const{base.ONE, base.ONE, base.ONE}
	lowerStrict := base.Const{base.BOT, base.BOT, base.BOT}
	return operators.And(
		subsumption.VarConst(v, upperEq),
		operators.Not(subsumption.VarConst(v, lowerStrict)),
	)
}

func lelOrder(v base.Var, c base.Const) base.Component {
	return operators.And(
		lel.VarConst(v, c),
		operators.Not(lel.ConstVar(c, v)),
	)
}

// =========================== //
//            TESTS            //
// =========================== //

func TestCompute_LEL(t *testing.T) {
	t.Cleanup(
		func() {
			os.Remove(FILEPATH)
		},
	)
	variable := base.Var("x")
	validOuts := []base.Const{
		{base.BOT, base.BOT, base.ONE},
		{base.BOT, base.ONE, base.BOT},
		{base.ONE, base.BOT, base.BOT},
	}
	found, out, err := Compute(
		lelFormula,
		lelOrder,
		variable,
		base.NewContext(3, nil),
		SOLVER,
		FILEPATH,
	)
	if err != nil {
		t.Error(err.Error())
		return
	}
	if !found {
		t.Error("Invalid output. Value could not be found.")
		return
	}
	if !slices.ContainsFunc(validOuts, equalToConstFunc(out)) {
		t.Errorf("Invalid output. %d", out)
		return
	}
}
