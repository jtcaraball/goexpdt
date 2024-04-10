package main

const SOLVER = "./kissat"

type experiment interface {
	Name() string
	Description() string
	Exec(args ...string) error
}

// Slice of impelmented experiments.
var experiments []experiment = []experiment{
	newOrderOptimExp(
		"dft:min-lel",
		"DFT - Optimum LEL Order",
		dftFGen,
		lelOGen,
	),
}

// Return map of implemented experiments with their name as key.
func expMap() map[string]experiment {
	exps := make(map[string]experiment)
	for _, exp := range experiments {
		exps[exp.Name()] = exp
	}
	return exps
}
