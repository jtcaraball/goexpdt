// vname defines set of useful constants and functions for naming variables
// safely and avoid collisions when assigning values to cnf variables.
//
// When declaring custom variable names one should avoid using the value of
// the constants declared in this package.
package vname

import "strings"

const (
	CounterPrefix   string = "ctr" // Counter variable prefix.
	ReachablePrefix string = "rch" // Reachable node variable prefix.
)

const (
	SConnector   rune = 35 // Connector for scope variable name generation.
	CVConntector rune = 36 // Connecctor for cnf variable name generation.
)

// SName returns scope formated name generated from p.
func SName(p ...string) string {
	return strings.Join(p, string(SConnector))
}

// CVName returns cnf formated name generated from p.
func CVName(p ...string) string {
	return strings.Join(p, string(CVConntector))
}
