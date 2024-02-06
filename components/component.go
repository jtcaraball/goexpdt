package components

import (
	"stratifoiled/cnf"
)

type Component interface {
	Encoding(ctx *Context) (*cnf.CNF, error)
	Simplified() (Component, error)
	GetChildren() []Component
	IsTrivial() (bool, bool)
}
