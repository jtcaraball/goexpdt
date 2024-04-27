package main

import (
	"bufio"
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/compute/utils"
	"goexpdt/trees"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

const (
	SOLVER    = "./kissat"
	INPUTDIR  = "inputs"
	OUTPUTDIR = "outputs"
)

var globPath = regexp.MustCompile(`\*`)

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
func bToC(b []byte, c base.Const) error {
	if len(b) != len(c) {
		return fmt.Errorf(
			"Invalid bytes length %d expected %d.",
			len(b),
			len(c),
		)
	}
	for i, ch := range b {
		switch ch {
		case 48: // 0
			c[i] = base.ZERO
		case 49: // 1
			c[i] = base.ONE
		case 50: // 2
			c[i] = base.BOT
		default:
			return fmt.Errorf("Invalid const feature value in index %d", i)
		}
	}
	return nil
}

// Set the values of c to constant represented in bytes b.
func sToC(s string, c base.Const) error {
	if len(s) != len(c) {
		return fmt.Errorf(
			"Invalid string length %d expected %d.",
			len(s),
			len(c),
		)
	}
	for i, ch := range s {
		switch ch {
		case 48: // 0
			c[i] = base.ZERO
		case 49: // 1
			c[i] = base.ONE
		case 50: // 2
			c[i] = base.BOT
		default:
			return fmt.Errorf("Invalid const feature value in index %d", i)
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

// Write random constant to c with truth value equal to tVal.
func randValConst(c base.Const, tVal bool, tree *trees.Tree) error {
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
func evalConst(c base.Const, tree *trees.Tree) (bool, error) {
	node := tree.Root
	for !node.IsLeaf() {
		if node.Feat < 0 || node.Feat >= len(c) {
			return false, errors.New("Node feature out of index.")
		}
		switch c[node.Feat] {
		case base.ONE:
			node = node.RChild
		case base.ZERO:
			node = node.LChild
		default:
			return false, errors.New("Constant is not full")
		}
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

// Return instances and context represented in the tree-instance input file
// passed by path.
func parseTIInput(inf string) ([]base.Const, *base.Context, error) {
	treeFP, instStrings, err := scanTIFile(inf)
	if err != nil {
		return nil, nil, err
	}

	ctx, err := genContext(path.Join(INPUTDIR, treeFP))
	if err != nil {
		return nil, nil, err
	}

	instances := make([]base.Const, len(instStrings))
	for i, cb := range instStrings {
		instances[i] = base.AllBotConst(ctx.Dimension)
		err := sToC(cb, instances[i])
		if err != nil {
			return nil, nil, err
		}
	}

	return instances, ctx, nil
}

// Scan a tree/instance input file by path.
func scanTIFile(path string) (string, []string, error) {
	file, err := os.Open(path)
	if err != nil {
		return "", nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		return "", nil, errors.New("Empty input file.")
	}
	treeFP := scanner.Text()

	instStrings := []string{}
	for scanner.Scan() {
		instStrings = append(instStrings, scanner.Text())
	}
	if len(instStrings) == 0 {
		return "", nil, errors.New("No instances in input file.")
	}

	if scanner.Err() != nil {
		return "", nil, err
	}

	return treeFP, instStrings, nil
}
