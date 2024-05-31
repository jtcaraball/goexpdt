// The logop package defines the logical operators (AND, NOT and OR) for the
// constructing queries.
//
// Additionally the WithVar and ForAllGuarded types are defined to handle
// variable declarations.
package logop

import (
	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// A LogOpChild allows for the encoding of its meaning into a CNF formula.
type LogOpChild interface {
	// Encoding returns takes in a ctx representing the state of a query or
	// sub-query and returns its CNF encoding along side an error.
	Encoding(ctx query.QContext) (cnf.CNF, error)
}
