package components

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type CNF struct {
	nv int              // Next variable value
	mClauses [][]int    // Meaning clauses
	cClauses [][]int    // Consistency clauses
	negated bool
}

// =========================== //
//           METHODS           //
// =========================== //

// Create a new CNF struct from clauses. Clauses will be treated as meaning.
func cnfFromClauses(clauses [][]int) *CNF {
	var clause []int
	var variable, absVariable int
	newCNF := &CNF{mClauses: clauses}
	for _, clause = range clauses {
		for _, variable = range clause {
			absVariable = absInt(variable)
			if absVariable > newCNF.nv {
				newCNF.nv = absVariable
			}
		}
	}
	return newCNF
}

// Negate the CNF meaning clauses. The resulting value of CNF's nv is the
// maximum between topv and the current value. This operation will set the CNF
// to an equivalent negation but it will not be equal to negating the
// underlying formula.
func (c *CNF) Negate(opt_topv ...int) error {
	topv := 0
	if len(opt_topv) > 0 {
		topv = opt_topv[0]
	}
	// Handle empty CNF case.
	if len(c.mClauses) == 0 {
		// An empty CNF is always SAT so to negate it we set it as an always
		// false CNF with a signle empty clause.
		c.mClauses = append(c.mClauses, []int{})
		return nil
	}
	// Handle empty clause in CNF case.
	if c.hasEmptyMeaningClause() {
		// A CNF with an empty clause is never SAT so to negate it we set it as
		// an always true empty CNF.
		c.mClauses = nil
		c.cClauses = nil
		return nil
	}
	// If CNF is negated then negation is simple: negate single literal.
	if c.negated {
		if len(c.mClauses) != 1 || len(c.mClauses[0]) != 1 {
			return errors.New(
				"Invalid CNF: Negated CNF with more than one clause or" +
				"literal.",
			)
		}
		c.mClauses[0][0] = -c.mClauses[0][0]
		c.negated = false
	}
	// Apply transformation to CNF meaning clauses.
	newNV := maxInt(topv, c.nv)
	if err := c.generateNegation(newNV); err != nil {
		return err
	}
	return nil
}

// Generate a new CNF corresponding to the conjunction of method caller and
// passed CNF.
func (c *CNF) Conjunction(oc *CNF) *CNF {
	newCNF := &CNF{}
	newCNF.mClauses = append(c.mClauses, oc.mClauses...)
	newCNF.cClauses = append(c.cClauses, oc.cClauses...)
	newCNF.nv = maxInt(c.nv, oc.nv)
	return nil
}

// Append a meaning clause to CNF and update nv value.
func (c *CNF) AppendMeaning(clause []int) {
	for _, v := range clause {
		absV := absInt(v)
		if absV > c.nv {
			c.nv = absV
		}
	}
	c.mClauses = append(c.mClauses, clause)
}

// Append a consistency clause to CNF and update nv value.
func (c *CNF) AppendConsistency(clause []int) {
	for _, v := range clause {
		absV := absInt(v)
		if absV > c.nv {
			c.nv = absV
		}
	}
	c.cClauses = append(c.cClauses, clause)
}

// Extend menaing clauses in CNF and update nv value.
func (c *CNF) ExtendMeaning(clauses [][]int) {
	for _, clause := range clauses {
		c.AppendMeaning(clause)
	}
}

// Extend consistency clauses in CNF and update nv value.
func (c *CNF) ExtendConsistency(clauses [][]int) {
	for _, clause := range clauses {
		c.AppendConsistency(clause)
	}
}

// Convert CNF formula to bytes in DIMACS CNF format.
func (c *CNF) ToBytes() []byte {
	bString := fmt.Sprintf(
		"p CNF %d %d\n",
		c.nv,
		len(c.mClauses) + len(c.cClauses),
	)
	for _, clause := range c.mClauses {
		bString += fmt.Sprintf("%s\n", clauseToDIMACS(clause))
	}
	for _, clause := range c.cClauses {
		bString += fmt.Sprintf("%s\n", clauseToDIMACS(clause))
	}
	return []byte(bString)
}

// Write CNF formula to file in DIMACS CNF format.
func (c *CNF) ToFile(path string) error {
	// Open file
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// Write CNF formula
	f.WriteString(
		fmt.Sprintf("p CNF %d %d\n", c.nv, len(c.mClauses) + len(c.cClauses)),
	)
	for _, clause := range c.mClauses {
		f.WriteString(fmt.Sprintf("%s\n", clauseToDIMACS(clause)))
	}
	for _, clause := range c.cClauses {
		f.WriteString(fmt.Sprintf("%s\n", clauseToDIMACS(clause)))
	}
	return nil
}

// Return CNF's meaning clauses and consistency clauses.
func (c *CNF) Clauses() ([][]int, [][]int) {
	return c.mClauses, c.cClauses
}

// =========================== //
//            UTILS            //
// =========================== //

// Returns true if the CNF has an empty meaning clause.
func (c *CNF) hasEmptyMeaningClause() bool {
	for _, clause := range c.mClauses {
		if len(clause) == 0 {
			return true
		}
	}
	return false
}

// Generate negation in place using Tseytin transformation.
func (c *CNF) generateNegation(nv int) error {
	clauses := [][]int{}
	enclits := []int{}
	for _, clause := range c.mClauses {
		if len(clause) == 0 {
			return errors.New(
				"Invalid CNF: Tseytin transformation can not be applied to" +
				"empty clause CNF.",
			)
		}
		auxv := -clause[0]
		if len(clause) > 1 {
			nv += 1
			auxv = nv
			// Direct implication.
			for _, lit := range clause {
				clauses = append(clauses, []int{-lit, -auxv})
			}
			// Oposite implication.
			clauses = append(clauses, append(clause, auxv))
		}
		// Literals representing negated clauses.
		enclits = append(enclits, auxv)
	}
	// If no errors were found then update CNF.
	c.cClauses = append(c.cClauses, clauses...)
	c.nv = nv
	// Generate bidirectional implication from new enc literal and enclits.
	if len(enclits) == 1 {
		c.mClauses = [][]int{enclits}
		return nil
	}
	c.addNegationIFFClauses(enclits)
	return nil
}

// Add "if and only if" clause for the passed enclits to CNF.
func (c *CNF) addNegationIFFClauses(enclits []int) {
	c.nv += 1
	auxv := c.nv
	for _, lit := range enclits {
		c.cClauses = append(c.cClauses, []int{-lit, auxv})
	}
	c.cClauses = append(c.cClauses, append(enclits, -auxv))
	c.mClauses = [][]int{{auxv}}
}

// Return clause in DIMACS CNF format including trailing 0.
func clauseToDIMACS(clause []int) string {
	cString := ""
	for _, v := range clause {
		cString += strconv.Itoa(v) + " "
	}
	cString += "0"
	return cString
}
