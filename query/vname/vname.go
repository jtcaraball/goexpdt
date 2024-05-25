// vname defines a set of useful constants and functions for naming variables
// safely and avoid collisions when assigning values to cnf variables.
//
// When declaring custom variable names one should avoid using the value of
// the constants declared in this package indirectly.
package vname

import "strings"

const (
	// Counter variable prefix.
	CounterPrefix string = "ctr"
	// Reachable node variable prefix.
	ReachablePrefix string = "rch"
)

const (
	// Connector for scope variable name generation.
	SConnector rune = 35
	// Connecctor for cnf variable name generation.
	CVConntector rune = 36
)

// SName returns scope formated name generated from p.
func SName(p ...string) string {
	return strings.Join(p, string(SConnector))
}

// CVName returns cnf formated name generated from p.
func CVName(p ...string) string {
	return strings.Join(p, string(CVConntector))
}
