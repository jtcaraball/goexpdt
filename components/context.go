package components

import (
	"fmt"
	"stratifoiled/trees"
)

// =========================== //
//          CONSTANTS          //
// =========================== //

const (
	ZERO uint8 = 0
	ONE uint8 = 1
	BOT uint8 = 2
)

var Symbols []uint8 = []uint8{ZERO, ONE, BOT}

// =========================== //
//           STRUCTS           //
// =========================== //

type Context struct {
	Dimension int
	Tree *trees.Tree
	TopV int
	vars map[contextVar]int
}

type contextVar struct {
	name string
	idx int
	value uint8
}

// =========================== //
//           METHODS           //
// =========================== //

func NewContext(dim int, tree *trees.Tree) *Context {
	ctx := &Context{Dimension: dim, Tree: tree}
	ctx.vars = make(map[contextVar]int)
	return ctx
}

func (c *Context) AddVar(name string, idx int, value uint8) {
	varS := contextVar{name: name, idx: idx, value: value}
	if c.vars[varS] != 0 {
		return
	}
	c.TopV += 1
	c.vars[varS] = c.TopV
}

func (c *Context) VarExists(name string, idx int, value uint8) bool {
	varS := contextVar{name: name, idx: idx, value: value}
	return c.vars[varS] != 0
}

func (c* Context) VarVal(name string, idx int, value uint8) int {
	varS := contextVar{name: name, idx: idx, value: value}
	varValue := c.vars[varS]
	if varValue == 0 {
		panic(fmt.Sprintf("Var (%s, %d, %d) does not exist", name, idx, value))
	}
	return varValue
}