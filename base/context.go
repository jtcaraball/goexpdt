package base

import (
	"fmt"
	"errors"
	"goexpdt/trees"
	"slices"
	"strconv"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type Context struct {
	Dimension int
	Tree *trees.Tree
	TopV int
	Guards []Guard
	nodeConsts []Const
	featVars map[ContextVar]int
	interVars map[ContextVar]int
}

type Guard struct {
	Target string
	InScope []string
	Value Const
	Idx int
}

type ContextVar struct {
	Name string
	Idx int
	Value int
}

// =========================== //
//           METHODS           //
// =========================== //

// Generate new context according to passed arguments.
func NewContext(dim int, tree *trees.Tree) *Context {
	ctx := &Context{Dimension: dim, Tree: tree}
	ctx.featVars = make(map[ContextVar]int)
	ctx.interVars = make(map[ContextVar]int)
	return ctx
}

// Return assigned value to variable. If it does not exist it is added.
func (c* Context) Var(name string, idx int, value int) int {
	varS := ContextVar{Name: name, Idx: idx, Value: value}
	varValue := c.featVars[varS]
	if varValue == 0 {
		c.TopV += 1
		c.featVars[varS] = c.TopV
		return c.TopV
	}
	return varValue
}

// Return true if variable exits in context. False otherwise.
func (c *Context) VarExists(name string, idx int, value int) bool {
	varS := ContextVar{Name: name, Idx: idx, Value: value}
	return c.featVars[varS] != 0
}

// Return assigned value to internal variable. If it does not exist it is added.
func (c* Context) IVar(name string, idx int, value int) int {
	varS := ContextVar{Name: name, Idx: idx, Value: value}
	varValue := c.interVars[varS]
	if varValue == 0 {
		c.TopV += 1
		c.interVars[varS] = c.TopV
		return c.TopV
	}
	return varValue
}

// Return true if internal variable exits in context. False otherwise.
func (c *Context) IVarExists(name string, idx int, value int) bool {
	varS := ContextVar{Name: name, Idx: idx, Value: value}
	return c.interVars[varS] != 0
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
func (c *Context) GetFeatVars() map[ContextVar]int {
	return c.featVars
}

// Return context's vars.
func (c *Context) GetInterVars() map[ContextVar]int {
	return c.interVars
}

// Add guard with the given target.
func (c *Context) AddGuard(target string) {
	c.Guards = append(c.Guards, Guard{Target: target})
}

// Remove last guard.
func (c *Context) PopGuard() {
	c.Guards = c.Guards[:len(c.Guards)]
}

// Set guard iterable values.
func (c *Context) SetGuard(gIdx, vIdx int, value Const) error {
	if gIdx > len(c.Guards) {
		return errors.New("Guard index out of range")
	}
	c.Guards[gIdx].Idx = vIdx
	c.Guards[gIdx].Value = value
	return nil
}

// Return variable scope suffix.
func (c *Context) ScopeSuffix(vName string) string {
	suffix := ""
	for _, guard := range c.Guards {
		if slices.Contains[[]string](guard.InScope, vName) {
			suffix += "#" + guard.Rep()
		}
	}
	return suffix
}

// Return matching target guard's value.
func (c *Context) GuardValueByTarget(target string) (Const, error) {
	for _, guard := range c.Guards {
		if guard.Target == target {
			return guard.Value, nil
		}
	}
	return nil, errors.New(
		fmt.Sprintf("No guard with target '%s'", target),
	)
}

// Return all trees nodes as slice of constants.
func (c *Context) NodesAsConsts() ([]Const, error) {
	if c.nodeConsts != nil {
		return c.nodeConsts, nil
	}
	var node *trees.Node
	var nConst, lnConst, rnConst Const
	var nStack = []*trees.Node{c.Tree.Root}
	var ncStack = []Const{AllBotConst(c.Dimension)}
	var nConsts = []Const{}
	for len(nStack) > 0 {
		node, nStack = nStack[len(nStack) - 1], nStack[:len(nStack) - 1]
		nConst, ncStack = ncStack[len(ncStack) - 1], ncStack[:len(ncStack) - 1]
		// Check for valid indexing.
		if node.Feat >= c.Dimension {
			return nil, errors.New("Node with invalid feature index.")
		}
		// Add node const to slice.
		nConsts = append(nConsts, nConst)
		if node.IsLeaf() {
			continue
		}
		// Add left node and const to stack.
		nStack = append(nStack, node.LChild)
		lnConst = AllBotConst(c.Dimension)
		copy(lnConst, nConst)
		lnConst[node.Feat] = ZERO
		ncStack = append(ncStack, lnConst)
		// Add right node and const to stack.
		nStack = append(nStack, node.RChild)
		rnConst = AllBotConst(c.Dimension)
		copy(rnConst, nConst)
		rnConst[node.Feat] = ONE
		ncStack = append(ncStack, rnConst)
	}
	c.nodeConsts = nConsts
	return c.nodeConsts, nil
}

// Return guard's representation.
func (g Guard) Rep() string {
	return g.Target + "#" + strconv.Itoa(g.Idx)
}
