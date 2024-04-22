package main

// experiment interface.
type experiment interface {
	Name() string
	Description() string
	Exec(args ...string) error
}

// Slice of impelmented experiments.
var experiments []experiment = []experiment{
	newOrderOptimValueExp(
		"optim:value:dft-ll",
		"Optimum:Value: DFT - Lesser Level Order\nArguments:\n"+
			"  - List of <tree file name>.",
		dftFGF,
		llOGF,
	),
	newOrderOptimStatsExp(
		"optim:stats:sr-ll",
		"Optimum:Stats: SR - Lesser Level Order\nArguments:\n"+
			"  - n consts per instance\n  - List of <tree file name>.",
		srFGF,
		llOGF,
	),
	newOrderOptimStatsExp(
		"optim:stats:sr-ss",
		"Optimum:Stats: SR - Strict Subsumption\nArguments:\n"+
			"  - n consts per instance\n  - List of <tree file name>.",
		srFGF,
		ssOGF,
	),
	newOrderOptimStatsExp(
		"optim:stats:cr-lh",
		"Optimum:Stats: CR - Less Hamming Distance\nArguments:\n"+
			"  - n consts per instance\n  - List of <tree file name>.",
		crFGF,
		lhOGF,
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
