package main

import (
	"errors"
	"fmt"
)

// =========================== //
//          EXPERIMENT         //
// =========================== //

// Order minimum 'get stats' experiment.
type evalExp struct {
	name string
	desc string
}

// Return new instance of experiment.
func newEvalExp(name, desc string) *evalExp {
	return &evalExp{
		name: name,
		desc: desc,
	}
}

// Return experiment name.
func (e *evalExp) Name() string {
	return e.name
}

// Return experiment description.
func (e *evalExp) Description() string {
	return e.desc
}

// Run experiment.
func (e *evalExp) Exec(args ...string) error {
	if len(args) == 0 {
		return errors.New("Missing arguments.")
	}

	inputPaths, err := filesToPaths(args)
	if err != nil {
		return err
	}

	for _, inputFP := range inputPaths {
		instances, ctx, err := parseTIInput(inputFP)
		if err != nil {
			return err
		}

		for _, inst := range instances {
			val, err := evalConst(inst, ctx.Tree)
			if err != nil {
				return err
			}

			fmt.Println(val)
		}

		fmt.Println()
	}

	return nil
}
