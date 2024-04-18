package main

import (
	"fmt"
	"goexpdt/base"
	"goexpdt/circuits/extensions/allcomp"
	"goexpdt/circuits/extensions/dft"
	"goexpdt/circuits/predicates/lel"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/compute/utils"
	"goexpdt/operators"
	"goexpdt/trees"
	"math/rand"
	"path"
	"strings"
	"time"
)

type vcFormula func(v base.Var, c base.Const) base.Component

const (
	SOLVER    = "./kissat"
	INPUTDIR  = "inputs"
	OUTPUTDIR = "outputs"
)

// =========================== //
//      VAR FORMULA GEN        //
// =========================== //

func dftFGen(v base.Var) base.Component {
	return dft.Var(v)
}

// =========================== //
//    VAR CONST FORMULA GEN    //
// =========================== //

func srFGen(v base.Var, c base.Const) base.Component {
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

// =========================== //
//          ORDER GEN          //
// =========================== //

func lelOGen(v base.Var, c base.Const) base.Component {
	return operators.And(
		lel.VarConst(v, c),
		operators.Not(lel.ConstVar(c, v)),
	)
}

// =========================== //
//             UTILS           //
// =========================== //

// Solve formula and return ok, const value. ok is false if the formula is
// unsatisfiable.
func solveFormula(
	formula base.Component,
	v base.Var,
	ctx *base.Context,
	solverPath, cnfPath string,
) (bool, base.Const, error) {
	exitcode, out, err := utils.Step(
		formula,
		ctx,
		solverPath,
		cnfPath,
	)
	if err != nil {
		return false, nil, err
	}

	if exitcode == 10 {
		outConst, err := utils.GetValueFromBytes(out, v, ctx)
		if err != nil {
			return false, nil, err
		}
		return true, outConst, nil
	}

	return false, nil, nil
}

// Set the values of c to constant represented in bytes b.
func btoC(b []byte, c base.Const) error {
	for i, ch := range b {
		switch ch {
		case 48: // 0
			c[i] = base.ZERO
		case 49: // 1
			c[i] = base.ONE
		case 50: // 2
			c[i] = base.BOT
		default:
			return fmt.Errorf("Invalid const feature value '%s'.", string(ch))
		}
	}
	return nil
}

// Randomize values of c.
func randConst(c base.Const) {
	for i := 0; i < len(c); i++ {
		r := rand.Intn(3)
		switch r {
		case 0:
			c[i] = base.ZERO
		case 1:
			c[i] = base.ONE
		case 2:
			c[i] = base.BOT
		}
	}
}

// Initialize context from tree file path.
func genContext(treePath string) (*base.Context, error) {
	expT, err := trees.LoadTree(treePath)
	if err != nil {
		return nil, err
	}
	ctx := base.NewContext(expT.FeatCount, expT)
	return ctx, nil
}

// Return output and temporal file names.
func fileNames(tag string) (string, string) {
	tAsString := strings.Split(time.Now().String(), " ")
	expId := strings.Join(tAsString[:2], "_")
	return path.Join(OUTPUTDIR, tag+"_"+expId+".out"), tag + "tmp_" + expId
}
