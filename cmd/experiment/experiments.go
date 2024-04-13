package main

// experiment interface.
type experiment interface {
	Name() string
	Description() string
	Exec(args ...string) error
}

// Slice of impelmented experiments.
var experiments []experiment = []experiment{
	newOrderOptimExp(
		"optim:dft-lel",
		"Optimum: DFT - LEL Order\nArguments:\n"+
			"  - List of <tree file name>.",
		dftFGen,
		lelOGen,
	),
	newVECFormulaExp(
		"vec:sr",
		"Formula: Variable - Explicit Constants\nArguments:\n"+
			"  - List of pairs <tree file name> <constants file name>.",
		srFGen,
	),
	newVRCFormulaExp(
		"vrc:sr",
		"Formula: Variable - Random Constants\nArguments:\n"+
			"  - List of <tree file name>.",
		srFGen,
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
