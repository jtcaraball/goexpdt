package components

import "stratifoiled/trees"

// =========================== //
//           STRUCTS           //
// =========================== //

type Context struct {
	Dimension int
	Tree *trees.Tree
	TopV int
	Guards []Guard
	vars map[contextVar]int
}

type Guard struct {
	Target string
	InScope []string
	Value Const
}

type contextVar struct {
	name string
	idx int
	value int
	inter bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Generate new context according to passed arguments.
func NewContext(dim int, tree *trees.Tree) *Context {
	ctx := &Context{Dimension: dim, Tree: tree}
	ctx.vars = make(map[contextVar]int)
	return ctx
}

// Return assigned value to variable. If it does not exist it is added.
func (c* Context) Var(name string, idx int, value int) int {
	varS := contextVar{name: name, idx: idx, value: value}
	varValue := c.vars[varS]
	if varValue == 0 {
		c.TopV += 1
		c.vars[varS] = c.TopV
		return c.TopV
	}
	return varValue
}

// Return true if variable exits in context. False otherwise.
func (c *Context) VarExists(name string, idx int, value int) bool {
	varS := contextVar{name: name, idx: idx, value: value}
	return c.vars[varS] != 0
}

// Return assigned value to internal variable. If it does not exist it is added.
func (c* Context) IVar(name string, idx int, value int) int {
	varS := contextVar{name: name, idx: idx, value: value, inter: true}
	varValue := c.vars[varS]
	if varValue == 0 {
		c.TopV += 1
		c.vars[varS] = c.TopV
		return c.TopV
	}
	return varValue
}

// Return true if internal variable exits in context. False otherwise.
func (c *Context) IVarExists(name string, idx int, value int) bool {
	varS := contextVar{name: name, idx: idx, value: value, inter: true}
	return c.vars[varS] != 0
}

// Set context's TopV to the max between the current value and value passed.
// Returns true if the value was updated.
func (c* Context) MaxUpdateTopV(topv int) bool {
	if topv > c.TopV {
		c.TopV = topv
		return true
	}
	return false
}
