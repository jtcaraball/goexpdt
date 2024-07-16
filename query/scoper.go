package query

import (
	"errors"
	"strconv"
	"strings"
)

// Scoper manages scope for guarded quantifiers.
type Scoper interface {
	// ScopeVar returns an instance variable that corresponds to v "inside" the
	// stack of scopes at the moment its called.
	// For any two variables u, v and any two stacks s1, s2 it must hold that
	// if u != v then s1:ScopeQVar(u) != s2:ScopeQVar(v).
	ScopeVar(v QVar) QVar
	// ScopeConst returns an instance constant with its value field set match
	// the value of the scope targeting c at the moment its called if it
	// exists. If the constant is scoped then ok==true else ok==false and c
	// is returned.
	ScopeConst(c QConst) (scopedQConst QConst, ok bool)
	// AddScope adds a scope with target==tgt to the stack.
	AddScope(tgt string)
	// PopScope removes the last scope in the stack.
	PopScope() error
	// SetScope adds the value and corresponding index to the last scope in the
	// stack. Returns an error if there are no scopes.
	SetScope(vIdx int, val []FeatV) error
	// Reset removes all guards in the scope
	Reset()
}

type baseScoper struct {
	builder strings.Builder // Scope string builder
	scopes  []scope
}

type scope struct {
	target string  // QConst target of the corresponding guarded quantifier.
	value  []FeatV // The value to be assigned target.
	vIdx   int     // The index of the value to be assigned.
}

func (s *baseScoper) ScopeVar(v QVar) QVar {
	if len(s.scopes) == 0 {
		return v
	}

	defer func() {
		s.builder.Reset()
	}()

	s.builder.WriteString(string(v))

	for _, g := range s.scopes {
		s.builder.WriteRune(31) // Unit separator
		s.builder.WriteString(g.target)
		s.builder.WriteRune(31) // Unit separator
		s.builder.WriteString(strconv.Itoa(g.vIdx))
	}

	return QVar(s.builder.String())
}

func (s *baseScoper) ScopeConst(c QConst) (QConst, bool) {
	for _, g := range s.scopes {
		if g.target == c.ID {
			return QConst{c.ID, g.value}, true
		}
	}
	return c, false
}

func (s *baseScoper) AddScope(tgt string) {
	s.scopes = append(s.scopes, scope{target: tgt})
}

func (s *baseScoper) PopScope() error {
	if len(s.scopes) == 0 {
		return errors.New("Invalid scope removal in empty scoper")
	}
	s.scopes = s.scopes[:len(s.scopes)-1]
	return nil
}

func (s *baseScoper) SetScope(vIdx int, val []FeatV) error {
	slen := len(s.scopes)
	if slen == 0 {
		return errors.New("Invalid setting of scope in empty scoper")
	}
	(s.scopes)[slen-1].vIdx = vIdx
	(s.scopes)[slen-1].value = val
	return nil
}

func (s *baseScoper) Reset() {
	s.scopes = nil
}
