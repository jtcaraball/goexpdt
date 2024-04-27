package main

// experiment interface.
type experiment interface {
	Name() string
	Description() string
	Exec(args ...string) error
}

// Slice of impelmented experiments.
var experiments []experiment = []experiment{
	newEvalExp(
		"eval",
		"Evaluate instances in given input files."+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
	),
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
		"Optimum:Value: CR - Less Hamming Distance\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		valEvalGen(CR_LH_O),
	),
	newOptimExp(
		"optim:val:ca-gh",
		"Optimum:Value: CA - Greater Hamming Distance\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		valEvalGen(CA_GH_O),
	),
	newOptimExp(
		"optim:val:sr-ll",
		"Optimum:Value: SR - Less Level\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		valEvalGen(SR_LL_O),
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
