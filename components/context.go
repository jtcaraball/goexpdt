package components

import (
	"slices"
	"stratifoiled/trees"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type Context struct {
	Dimension int
	Tree *trees.Tree
	TopV int
	Guards []Guard
	vars map[ContextVar]int
}

type Guard struct {
	Target string
	InScope []string
	Value Const
	Rep string
}

type ContextVar struct {
	Name string
	Idx int
	Value int
	Inter bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Generate new context according to passed arguments.
func NewContext(dim int, tree *trees.Tree) *Context {
	ctx := &Context{Dimension: dim, Tree: tree}
	ctx.vars = make(map[ContextVar]int)
	return ctx
}

// Return assigned value to variable. If it does not exist it is added.
func (c* Context) Var(name string, idx int, value int) int {
	varS := ContextVar{Name: name, Idx: idx, Value: value}
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
	varS := ContextVar{Name: name, Idx: idx, Value: value}
	return c.vars[varS] != 0
}

// Return assigned value to internal variable. If it does not exist it is added.
func (c* Context) IVar(name string, idx int, value int) int {
	varS := ContextVar{Name: name, Idx: idx, Value: value, Inter: true}
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
	varS := ContextVar{Name: name, Idx: idx, Value: value, Inter: true}
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

// Add var name to guard's scopes.
func (c *Context) AddVarToScope(varInst Var) {
	// The amount of vars in a formula should tend to be small so slices.Contain
	// is more than good enough.
	for i := 0; i < len(c.Guards); i++ {
		if !slices.Contains[[]string](c.Guards[i].InScope, string(varInst)) {
			c.Guards[i].InScope = append(c.Guards[i].InScope, string(varInst))
		}
	}
}

// Return context's vars.
func (c *Context) GetVars() map[ContextVar]int {
	return c.vars
}
