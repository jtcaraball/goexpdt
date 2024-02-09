package components

import (
	"fmt"
	"errors"
)

func ValidateConstsDim(
	compName string,
	ctx *Context,
	consts ...Const,
) error {
	for i, c := range consts {
		if len(c) != ctx.Dimension {
			return errors.New(
				fmt.Sprintf(
					"%s -> constant%d: wrong dim %d (%d feats in context)",
					compName,
					i + 1,
					len(c),
					ctx.Dimension,
				),
			)
		}
	}
	return nil
}
