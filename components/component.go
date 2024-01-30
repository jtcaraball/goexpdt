package components

import "context"

type Component interface {
	Contextualize(context context.Context)
	Encode() CNF
	GetChildren() []Component
	Simplify()
}
