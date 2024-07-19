package hepr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jtcaraball/goexpdt/query"
)

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

func prefVar(pref []int) query.QVar {
	var sb strings.Builder

	for _, idx := range pref {
		sb.WriteString(strconv.Itoa(idx))
	}

	return query.QVar(sb.String())
}
