package main

// experiment interface.
type experiment interface {
	Name() string
	Description() string
	Exec(args ...string) error
}

// Slice of impelmented experiments.
var experiments []experiment = []experiment{
	newRandOptimExp(
		"optim:rand:val:dft-ll",
		"Optimum:Value: DFT - Lesser Level Order (Random Instances):\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		randValEvalGen(DFT_LL_C),
	),
	newRandOptimExp(
		"optim:rand:stats:sr-ll",
		"Optimum:Value: SR - Lesser Level (Random Instances):\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		randStatsEvalGen(SR_LL_C),
	),
	newRandOptimExp(
		"optim:rand:stats:sr-ss",
		"Optimum:Value: SR - Strict Subsumption (Random Instances):\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		randStatsEvalGen(SR_SS_C),
	),
	newOptimExp(
		"optim:val:cr-lh",
		"Optimum:Stats: CR - Less Hamming Distance\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		valEvalGen(CR_LH_O),
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
