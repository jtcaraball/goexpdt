package query

import (
	"errors"
	"strconv"
	"strings"

	"github.com/jtcaraball/goexpdt/query/vname"
)

// Scoper manages scope for guarded quantifiers.
type Scoper interface {
	// ScopeVar returns an instance variable equal to the original variable
	// plus a suffix representing the stack of scopes at the moment its called.
	ScopeVar(v Var) Var
	// ScopeConst returns an instance constant with its value field set match
	// the value of the scope targeting c at the moment its called if it
	// exists. If the constant is scoped then ok==true else ok==false and c
	// is returned.
	ScopeConst(c Const) (scopedConst Const, ok bool)
	// AddScope adds a scope with target==tgt to the stack.
	AddScope(tgt string)
	// PopScope removes the last scope in the stack.
	PopScope() error
	// SetScope adds the value and corresponding index the last scope in the
	// stack. Returns an error if there are no scopes or the last scope is
	// already set.
	SetScope(vIdx int, val []FeatV) error
	// Reset removes all guards in the scope
	Reset()
}

type baseScoper []scope

type scope struct {
	target string  // Constant target of the corresponding guarded quantifier.
	value  []FeatV // The value to be assigned target.
	vIdx   int     // The index of the value to be assigned.
}

func (s *baseScoper) ScopeVar(v Var) Var {
	var sb strings.Builder

	sb.WriteString(string(v))

	for _, g := range *s {
		sb.WriteString(string(vname.SConnector))
		sb.WriteString(vname.SName(g.target, strconv.Itoa(g.vIdx)))
	}

	return Var(sb.String())
}

func (s *baseScoper) ScopeConst(c Const) (Const, bool) {
	for _, g := range *s {
		if g.target == c.ID {
			return Const{c.ID, g.value}, true
		}
	}
	return c, false
}

func (s *baseScoper) AddScope(tgt string) {
	*s = append(*s, scope{target: tgt})
}

func (s *baseScoper) PopScope() error {
	if len(*s) == 0 {
		return errors.New("Invalid scope removal in empty scoper")
	}
	*s = (*s)[:len(*s)-1]
	return nil
}

func (s *baseScoper) SetScope(vIdx int, val []FeatV) error {
	slen := len(*s)
	if slen == 0 {
		return errors.New("Invalid setting of scope in empty scoper")
	}
	(*s)[slen-1].vIdx = vIdx
	(*s)[slen-1].value = val
	return nil
}

func (s *baseScoper) Reset() {
	*s = make(baseScoper, 0)
}
