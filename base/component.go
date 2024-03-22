package base

import "goexpdt/cnf"

type Component interface {
	Encoding(ctx *Context) (*cnf.CNF, error)
	Simplified(ctx *Context) (Component, error)
	GetChildren() []Component
	IsTrivial() (bool, bool)
}
