package base

import (
	"fmt"
	"goexpdt/cnf"
	"strconv"
)

// =========================== //
//         INTERFACES          //
// =========================== //

type ConstInstance interface {
	Scoped(ctx *Context) (Const, error)
}

// =========================== //
//           STRUCTS           //
// =========================== //

type Var string

type GuardedConst string

type Const []featV

type featV struct {
	val int
}

// =========================== //
//          VARIABLES          //
// =========================== //

var (
	ZERO = featV{val: 0}
	ONE  = featV{val: 1}
	BOT  = featV{val: 2}
)

// =========================== //
//          VAR METHODS        //
// =========================== //

// Return new var instance with name equal to string passed.
func NewVar(name string) Var {
	var v Var = Var(name)
	return v
}

// Encode v's consistency clauses to cnf and add necesary variables to context.
func (v Var) Encoding(ctx *Context) *cnf.CNF {
	nCNF := &cnf.CNF{}
	// Add consistency clauses
	if ctx.VarExists(string(v), 0, BOT.Val()) {
		return nCNF
	}
	// Every feature must have at least one value
	reqAllFeats := [][]int{}
	for i := 0; i < ctx.Dimension; i++ {
		reqAllFeats = append(
			reqAllFeats,
			[]int{
				ctx.Var(string(v), i, ZERO.Val()),
				ctx.Var(string(v), i, ONE.Val()),
				ctx.Var(string(v), i, BOT.Val()),
			},
		)
	}
	nCNF.ExtendConsistency(reqAllFeats)
	// Every feature must have one and only one value
	for i := 0; i < ctx.Dimension; i++ {
		reqOnePerFeat := [][]int{
			{
				-ctx.Var(string(v), i, ZERO.Val()),
				-ctx.Var(string(v), i, ONE.Val()),
			},
			{
				-ctx.Var(string(v), i, ZERO.Val()),
				-ctx.Var(string(v), i, BOT.Val()),
			},
			{
				-ctx.Var(string(v), i, ONE.Val()),
				-ctx.Var(string(v), i, BOT.Val()),
			},
		}
		nCNF.ExtendConsistency(reqOnePerFeat)
	}
	return nCNF
}

// Return corresponding scoped var. If no scope is found in guard returns the
// method caller.
func (v Var) Scoped(ctx *Context) Var {
	// This can generate a conflict in feature variable names as Var is always
	// saved to ctx.featVars so if any variable outside of scope is named
	// v + scopeName they will collide.
	// I have too much on my life to deal with this rn thou :c
	rVar := ctx.ScopeSuffix(string(v))
	if rVar != "" {
		return Var(string(v) + rVar)
	}
	return v
}

// =========================== //
//        CONST METHODS        //
// =========================== //

// Return all bots const of len dim.
func AllBotConst(dim int) Const {
	feats := []featV{}
	for i := 0; i < dim; i++ {
		feats = append(feats, BOT)
	}
	return feats
}

// Return caller always.
func (c Const) Scoped(ctx *Context) (Const, error) {
	return c, nil
}

// Return true if caller is full.
func (c Const) IsFull() bool {
	for _, ft := range c {
		if ft == BOT {
			return false
		}
	}
	return true
}

// Return caller string representation.
func (c Const) AsString() string {
	name := ""
	for _, ft := range c {
		name += strconv.Itoa(ft.Val())
	}
	return name
}

// Return number of bot features in constant.
func (c Const) BotCount() int {
	count := 0
	for _, ft := range c {
		if ft == BOT {
			count += 1
		}
	}
	return count
}

// Return corresponding const from list of guards.
func (gc GuardedConst) Scoped(ctx *Context) (Const, error) {
	return ctx.GuardValueByTarget(string(gc))
}

// Return featV value.
func (f featV) Val() int {
	return f.val
}

// Returns an error if any of passed consts have lenght different to Dim.
func ValidateConstsDim(
	constDim int,
	consts ...Const,
) error {
	for i, c := range consts {
		if len(c) != constDim {
			return fmt.Errorf(
				"constant%d: wrong dim %d (%d feats in context)",
				i+1,
				len(c),
				constDim,
			)
		}
	}
	return nil
}
