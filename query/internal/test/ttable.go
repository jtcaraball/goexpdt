// Test cases tables for pred package.
package test

import "github.com/jtcaraball/goexpdt/query"

// OTRecord represents a test case that uses a single value.
type OTRecord struct {
	Dim     int
	Name    string
	Val     []query.FeatV
	ExpCode int
}

// BTRecord represents a test case that uses two values.
type BTRecord struct {
	Dim     int
	Name    string
	Val1    []query.FeatV
	Val2    []query.FeatV
	ExpCode int
}

// TTRecord represents a test case that uses three values.
type TTRecord struct {
	Dim     int
	Name    string
	Val1    []query.FeatV
	Val2    []query.FeatV
	Val3    []query.FeatV
	ExpCode int
}
