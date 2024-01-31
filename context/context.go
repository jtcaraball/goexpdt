package context

import "stratifoiled/trees"

// =========================== //
//           STRUCTS           //
// =========================== //

type Context struct {
	Dimension int
	Tree *trees.Tree
	Vars map[Var]int
	TopV int
}

// =========================== //
//           METHODS           //
// =========================== //

func (c *Context) AddVar(varKey Var) {
	if c.Vars[varKey] != 0 {
		return
	}
	c.TopV += 1
	c.Vars[varKey] = c.TopV
}
