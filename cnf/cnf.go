// cnf implements the CNF and Clause types for representing formulas in
// conjunctive normal form.
//
// This package makes a distinction between semantic and consistency clauses in
// CNF formulas. Where as semantic clauses can be negated consistency clauses
// are supposed to be fixed, and thus will never be negated, as they represent
// the 'world' in which the truth value of a formula is being evaluated.
//
// For example: say we want to define a formula f that should be true if a
// partial instance has n or fewer features with value one. We define both the
// clauses that represent what it means to have any number m of features be
// equal to one (this would be the consistency clauses) and then the clauses
// whose meaning corresponds to not having more than n features be equal to one
// (semantic clauses). This way if we want to represent the formula g that is
// true if a partial instance has more than n features with value one we can
// simply define it as g=-f as the 'counting' clauses in f are not negated.
package cnf

import (
	"fmt"
	"os"
	"strconv"
)

// Clause represents a cnf formula's clause with integer variables where
// negated variables are negative.
type Clause []int

// CNF represents a logical formula in conjoined normal form.
type CNF struct {
	// Top variable value used
	tv uint32
	// Semantic clauses
	sClauses []Clause
	// Consistency clauses
	cClauses []Clause
}

// NegCluases is a slice of clauses that contains only an empty clauses
// which is always evaluated as false.
var NegClauses []Clause = []Clause{{}}

// Create a new CNF struct from clauses. Clauses will be treated as semantic.
func FromClauses(clauses []Clause) CNF {
	ncnf := CNF{sClauses: clauses}
	ncnf.tv = maxVar(clauses)
	return ncnf
}

// Negative returns a negative cnf formula.
func Negative() CNF {
	return CNF{sClauses: NegClauses}
}

// Negate the CNF semantic clauses. The resulting value of CNF's tv is the
// maximum between topv and the current value. This operation will set the CNF
// to an equivalent negation but it will not be equal to negating the
// underlying formula. Returns new tv value.
func (c CNF) Negate(opt_topv ...uint32) CNF {
	var rcnf CNF
	var topv uint32 = c.tv

	// Why did the go team decide against optional arguments?
	if len(opt_topv) > 0 {
		topv = maxUInt(topv, opt_topv[0])
	}

	// Handle empty CNF case.
	if len(c.sClauses) == 0 {
		// An empty CNF is always SAT so to negate it we set it as an always
		// false CNF with a signle empty clause.
		rcnf = Negative()
		rcnf.tv = topv
		return rcnf
	}

	// Handle empty clause in CNF case.
	if c.HasEmptySemanticClause() {
		// A CNF with an empty clause is never SAT so to negate it we set it as
		// an always true empty CNF.
		rcnf = CNF{}
		rcnf.tv = topv
		return rcnf
	}

	// Apply transformation to CNF semantic clauses.
	rcnf = CNF{
		tv:       topv,
		sClauses: c.sClauses,
		cClauses: c.cClauses,
	}
	rcnf.tseytinNegation(topv)
	return rcnf
}

// Generate negation in place using Tseytin transformation. Does not check if
// c is nil.
func (c *CNF) tseytinNegation(tv uint32) {
	clauses := []Clause{}
	enclits := Clause{}

	for _, clause := range c.sClauses {
		// I would rather this function not return errors so we handle the case
		// in which tseytin's transform is not valid (empty clause) by shoving
		// in the appropriate negation and returning.
		if len(clause) == 0 { // An empty clause results in a false formula.
			c.sClauses = nil
			c.cClauses = nil
			c.tv = tv
			return
		}

		auxv := -clause[0]
		if len(clause) > 1 {
			tv += 1
			auxv = int(tv)
			// Direct implication.
			for _, lit := range clause {
				clauses = append(clauses, Clause{-lit, -auxv})
			}
			// Oposite implication.
			clauses = append(clauses, append(clause, auxv))
		}

		// Literals representing negated clauses.
		enclits = append(enclits, auxv)
	}

	// If no errors were found then update CNF.
	c.cClauses = append(c.cClauses, clauses...)
	c.tv = tv

	// Generate bidirectional implication from new enc literal and enclits.
	if len(enclits) == 1 {
		c.sClauses = []Clause{enclits}
		return
	}

	c.addNegationIFFClauses(enclits)
}

// Add "if and only if" clause for the passed enclits to CNF. Does not check if
// c is nil.
func (c *CNF) addNegationIFFClauses(enclits Clause) {
	c.tv += 1
	auxv := int(c.tv)

	for _, lit := range enclits {
		c.cClauses = append(c.cClauses, Clause{-lit, auxv})
	}

	c.cClauses = append(c.cClauses, append(enclits, -auxv))
	c.sClauses = []Clause{{auxv}}
}

// Conjunction generates extend the methods caller semantic and consistency
// clauses with those from the passed CNF.
func (c CNF) Conjunction(oc CNF) CNF {
	rcnf := c.AppendSemantics(oc.sClauses...)
	rcnf.cClauses = c.cClauses
	rcnf = rcnf.AppendConsistency(oc.cClauses...)
	return rcnf
}

// AppendSpemantics appends semantic clauses to CNF and update tv value.
func (c CNF) AppendSemantics(clauses ...Clause) CNF {
	topv := maxUInt(c.tv, maxVar(clauses))
	return CNF{
		tv:       topv,
		sClauses: append(c.sClauses, clauses...),
		cClauses: c.cClauses,
	}
}

// AppendConsistency appends consistency clauses to CNF and update tv value.
func (c CNF) AppendConsistency(clauses ...Clause) CNF {
	topv := maxUInt(c.tv, maxVar(clauses))
	return CNF{
		tv:       topv,
		sClauses: c.sClauses,
		cClauses: append(c.cClauses, clauses...),
	}
}

// ToBytes returns CNF formula as bytes in DIMACS CNF format.
func (c CNF) ToBytes() []byte {
	bString := fmt.Sprintf(
		"p cnf %d %d\n",
		c.tv,
		len(c.sClauses)+len(c.cClauses),
	)

	for _, clause := range c.sClauses {
		bString += fmt.Sprintf("%s\n", clauseToDIMACS(clause))
	}

	for _, clause := range c.cClauses {
		bString += fmt.Sprintf("%s\n", clauseToDIMACS(clause))
	}

	return []byte(bString)
}

// ToFile writes CNF formula to file in DIMACS CNF format.
func (c CNF) ToFile(path string) error {
	// Open file
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write CNF formula
	if _, err = f.WriteString(
		fmt.Sprintf("p cnf %d %d\n", c.tv, len(c.sClauses)+len(c.cClauses)),
	); err != nil {
		return err
	}

	for _, clause := range c.sClauses {
		if _, err = f.WriteString(
			fmt.Sprintf("%s\n", clauseToDIMACS(clause)),
		); err != nil {
			return err
		}
	}

	for _, clause := range c.cClauses {
		if _, err = f.WriteString(
			fmt.Sprintf("%s\n", clauseToDIMACS(clause)),
		); err != nil {
			return err
		}
	}

	return nil
}

// Clauses return CNF's semantic clauses and consistency clauses.
func (c CNF) Clauses() ([]Clause, []Clause) {
	return c.sClauses, c.cClauses
}

// HasEmptySemanticClauses returns true if the CNF has an empty semantic
// clause.
func (c CNF) HasEmptySemanticClause() bool {
	for _, clause := range c.sClauses {
		if len(clause) == 0 {
			return true
		}
	}
	return false
}

// Return clause in DIMACS CNF format including trailing 0.
func clauseToDIMACS(clause Clause) string {
	cString := ""

	for _, v := range clause {
		cString += strconv.Itoa(v) + " "
	}
	cString += "0"

	return cString
}

// maxVar returns the absolute value of the largest variable in a slice of
// clauses.
func maxVar(clauses []Clause) uint32 {
	var topv uint32 = 0
	for _, cl := range clauses {
		for _, v := range cl {
			absV := absInt(v)
			if absV > topv {
				topv = absV
			}
		}
	}
	return topv
}

// absInt returns the absolute value of an integer.
func absInt(v int) uint32 {
	if v > 0 {
		return uint32(v)
	}
	return uint32(-v)
}

// maxUInt returns the maximum between two unsinged integers.
func maxUInt(v1, v2 uint32) uint32 {
	if v1 > v2 {
		return v1
	}
	return v2
}
