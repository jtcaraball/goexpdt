package main

import (
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/circuits/extensions/allcomp"
	"goexpdt/circuits/extensions/dft"
	"goexpdt/circuits/extensions/full"
	"goexpdt/circuits/extensions/leh"
	"goexpdt/circuits/predicates/lel"
	"goexpdt/circuits/predicates/subsumption"
	"goexpdt/compute/orderoptimum"
	"goexpdt/compute/utils"
	"goexpdt/operators"
	"goexpdt/trees"
	"math/rand"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type (
	optimFormulaGenFactory func(cs ...base.Const) (orderoptimum.VFormula, error)
	optimOrderGenFactory   func(cs ...base.Const) (orderoptimum.VCOrder, error)
)

const (
	SOLVER    = "./kissat"
	INPUTDIR  = "inputs"
	OUTPUTDIR = "outputs"
)

var globPath = regexp.MustCompile(`\*`)

// =========================== //
//      VAR FORMULA GEN        //
// =========================== //

func dftFGF(cs ...base.Const) (orderoptimum.VFormula, error) {
	return func(v base.Var) base.Component {
		return dft.Var(v)
	}, nil
}

func srFGF(cs ...base.Const) (orderoptimum.VFormula, error) {
	if len(cs) == 0 {
		return nil, errors.New("Missing constant in order generation factory")
	}

	c := cs[0]

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
	}, nil
}

func crFGF(cs ...base.Const) (orderoptimum.VFormula, error) {
	if len(cs) == 0 {
		return nil, errors.New("Missing constant in order generation factory")
	}

	c := cs[0]

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
	}, nil
}

// =========================== //
//          ORDER GEN          //
// =========================== //

func llOGF(cs ...base.Const) (orderoptimum.VCOrder, error) {
	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			lel.VarConst(v, c),
			operators.Not(lel.ConstVar(c, v)),
		)
	}, nil
}

func ssOGF(cs ...base.Const) (orderoptimum.VCOrder, error) {
	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			subsumption.VarConst(v, c),
			operators.Not(subsumption.ConstVar(c, v)),
		)
	}, nil
}

func lhOGF(cs ...base.Const) (orderoptimum.VCOrder, error) {
	if len(cs) == 0 {
		return nil, errors.New("Missing constant in order generation factory")
	}

	cp := cs[0]

	return func(v base.Var, c base.Const) base.Component {
		return operators.And(
			leh.ConstVarConst(cp, v, c),
			operators.Not(leh.ConstConstVar(cp, c, v)),
		)
	}, nil
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
func randConst(c base.Const, full bool) {
	limit := 3
	if full {
		limit = 2
	}

	for i := 0; i < len(c); i++ {
		r := rand.Intn(limit)
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

// Wrtie random constant to c with truth value equal to tVal.
func randValConst(c base.Const, tVal bool, tree trees.Tree) error {
	match := false
	for !match {
		randConst(c, true)
		val, err := evalConst(c, tree)
		if err != nil {
			return err
		}
		match = val == tVal
	}
	return nil
}

// Return valuation of constant in tree. C must be full.
func evalConst(c base.Const, tree trees.Tree) (bool, error) {
	node := tree.Root
	for !node.IsLeaf() {
		if node.Feat < 0 || node.Feat >= len(c) {
			return false, errors.New("Node feature out of index.")
		}
		if c[node.Feat] == base.ONE {
			node = node.RChild
			continue
		} else if c[node.Feat] == base.ZERO {
			node = node.LChild
			continue
		}
		return false, errors.New("Constant is not full")
	}
	return node.Value, nil
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
	return path.Join(OUTPUTDIR, tag+"_"+expId+".csv"), tag + "tmp_" + expId
}

// Format input paths to be used in experiment. Expands any gblob paths.
func filesToPaths(filenames []string) ([]string, error) {
	paths := []string{}

	for _, f := range filenames {
		fp := filepath.Join(INPUTDIR, f)

		if globPath.MatchString(fp) {
			expand, err := filepath.Glob(fp)
			if err != nil {
				return nil, err
			}
			paths = append(paths, expand...)
			continue
		}

		paths = append(paths, fp)
	}
	return paths, nil
}

// Return minimum duration.
func dMin(t1, t2 time.Duration) time.Duration {
	if t1 >= t2 {
		return t2
	}
	return t1
}

// Return maximum duration.
func dMax(t1, t2 time.Duration) time.Duration {
	if t1 <= t2 {
		return t2
	}
	return t1
}
