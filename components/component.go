package components

import (
	"stratifoiled/cnf"
)

type Component interface {
	Encoding(ctx *Context) *cnf.CNF
	GetChildren() []Component
	Simplified() Component
	IsTrivial() bool
}
