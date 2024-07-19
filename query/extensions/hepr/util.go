package hepr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jtcaraball/goexpdt/query"
)

// validatePref returns a non nill error if pref contains feature indexes out
// of range of the given dimension dim or has any duplicates.
func validatePref(pref []int, dim int) error {
	seen := make([]bool, dim)
	for _, i := range pref {
		if i < 0 || i >= dim {
			return fmt.Errorf("Preference feature index out of range '%d'", i)
		}
		if seen[i] {
			return fmt.Errorf("Invalid duplicate preference '%d'", i)
		}
		seen[i] = true
	}
	return nil
}

// prefVar returns a query variable corresponding to pref.
func prefVar(pref []int) query.QVar {
	var sb strings.Builder

	for _, idx := range pref {
		sb.WriteString(strconv.Itoa(idx))
	}

	return query.QVar(sb.String())
}
