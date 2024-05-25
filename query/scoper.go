package query

import (
	"errors"
	"slices"
	"strconv"
	"strings"

	"github.com/jtcaraball/goexpdt/query/vname"
)

// baseScoper is a basic scope variable and constant manager.
type baseScoper []Guard

// A Guard is the representation of a guarded quantifier in the scope.
type Guard struct {
	Target  string  // Constant target of the corresponding guarded quantifier.
	InScope []Var   // Variables declared inside the scope.
	Value   []FeatV // The value to be assigned target.
	VIdx    int     // The index of the value to be assigned.
}

func (s *baseScoper) ScopeVar(v Var) Var {
	var sb strings.Builder
	for _, g := range *s {
		if slices.Contains(g.InScope, v) {
			sb.WriteString(vname.SName(g.Target, strconv.Itoa(g.VIdx)))
		}
	}
	if sb.Len() == 0 {
		return v
	}
	return Var(vname.SName(string(v), sb.String()))
}

func (s *baseScoper) ScopeConst(c Const) (Const, bool) {
	for _, g := range *s {
		if g.Target == c.ID {
			return Const{c.ID, g.Value}, true
		}
	}
	return Const{}, false
}

func (s *baseScoper) AddGuard(tgt string) {
	*s = append(*s, Guard{Target: tgt})
}

func (s *baseScoper) PopGuard() {
	*s = (*s)[:len(*s)-1]
}

func (s *baseScoper) SetGuard(vIdx int, val []FeatV) error {
	slen := len(*s)
	if slen == 0 || (*s)[slen].Target != "" {
		return errors.New("Invalid guard setting")
	}
	(*s)[slen-1].VIdx = vIdx
	(*s)[slen-1].Value = val
	return nil
}

func (s *baseScoper) Reset() { *s = nil }
