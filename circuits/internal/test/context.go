package test

import (
	"goexpdt/base"
	"slices"
	"testing"
)

func OnlyFeatVariables(t *testing.T, ctx *base.Context, varNames ...string) {
	for key := range ctx.GetFeatVars() {
		if !slices.Contains[[]string](varNames, key.Name) {
			t.Errorf("Unexpected variable name: %s", key.Name)
		}
	}
	for key := range ctx.GetInterVars() {
		if slices.Contains[[]string](varNames, key.Name) {
			t.Errorf("Unexpected variable name: %s", key.Name)
		}
	}
}
