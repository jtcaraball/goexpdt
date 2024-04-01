package leh

import (
	"fmt"
	"goexpdt/base"
)

// Return hamming distance between to constant instances. Return error if
// constants have different dimensions.
func hammingDist(constInst1, constInst2 base.Const) (int, error) {
	hDist := 0
	if len(constInst1) != len(constInst2) {
		return 0, fmt.Errorf(
			"Mismatched constant dimensions: %d - %d",
			len(constInst1),
			len(constInst2),
		)
	}
	for i, ft := range(constInst1) {
		if ft != constInst2[i] {
			hDist += 1
		}
	}
	return hDist, nil
}
