package query

import (
	"fmt"
	"strings"
)

// Var corresponds to the identifier of a partial instance in a query.
type Var string

// FeatV corresponds to the value of a partial instance.
type FeatV int

// Const corresponds to an explicit partial instance. It may not have an
// assigned value to accomodate for guarded quantifiers.
type Const struct {
	ID  string
	Val []FeatV
}

// Supported feature values.
const (
	// A BOT value represents the bottom feature value of a partial instance.
	BOT FeatV = iota
	// A ZERO value represents the zero feature value of a partial instance.
	ZERO
	// A ONE value represents the one feature value of a partial instance.
	ONE
)

// AllBotConst returns a all bot constant len dim with a zero value id.
func AllBotConst(dim int) Const {
	feats := []FeatV{}
	for i := 0; i < dim; i++ {
		feats = append(feats, BOT)
	}
	return Const{Val: feats}
}

// IsFull return true if constant caller has no features equal to BOT.
func (c Const) IsFull() bool {
	for _, ft := range c.Val {
		if ft == BOT {
			return false
		}
	}
	return true
}

// AsString returns the caller's features represented as a string.
func (c Const) AsString() string {
	var sb strings.Builder
	for _, ft := range c.Val {
		switch ft {
		case BOT:
			sb.WriteRune(95) // "_"
		case ZERO:
			sb.WriteRune(48) // "0"
		case ONE:
			sb.WriteRune(49) // "1"
		}
	}
	return sb.String()
}

// BotCound return number of bot features in constant.
func (c Const) BotCount() int {
	count := 0
	for _, ft := range c.Val {
		if ft == BOT {
			count += 1
		}
	}
	return count
}

// ValidateConstsDim returns an error if any of passed consts have length
// different to d.
func ValidateConstsDim(
	d int,
	consts ...Const,
) error {
	for i, c := range consts {
		if len(c.Val) != d {
			return fmt.Errorf(
				"constant%d: wrong dim %d (%d feats in context)",
				i+1,
				len(c.Val),
				d,
			)
		}
	}
	return nil
}
