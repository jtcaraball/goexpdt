package context

import (
	"goexpdt/base"
	"slices"
	"testing"
)

func OnlyFeatVariables(t *testing.T, ctx *base.Context, varNames ...string) {
	for key := range ctx.GetFeatVars() {
		if !slices.Contains(varNames, key.Name) {
			t.Errorf("Unexpected variable name in featVars: %s", key.Name)
		}
	}
	for key := range ctx.GetInterVars() {
		if slices.Contains(varNames, key.Name) {
			t.Errorf("Unexpected variable name in interVars: %s", key.Name)
		}
	}
}
