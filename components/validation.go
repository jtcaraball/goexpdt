package components

import (
	"fmt"
	"errors"
)

func ValidateConstsDim(
	compName string,
	dim int,
	consts ...Const,
) error {
	for i, c := range consts {
		if len(c) != dim {
			return errors.New(
				fmt.Sprintf(
					"%s -> constant%d: wrong dim %d (%d feats in context)",
					compName,
					i + 1,
					len(c),
					dim,
				),
			)
		}
	}
	return nil
}
