package components

import (
	"context"
	"stratifoiled/cnf"
)

type Component interface {
	Contextualize(context context.Context)
	Encode() *cnf.CNF
	GetChildren() []Component
	Simplify()
}
