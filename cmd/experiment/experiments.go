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
		"optim:rand:stats:dfs-ll",
		"Optimum (Stats, Random Instances) - DFS under Lesser Level Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newRandStatsDriver(DFS_LL_C),
	),
	newRandOptimExp(
		"optim:rand:stats:sr-ll",
		"Optimum (Stats, Random Instances) - SR under Lesser Level Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newRandStatsDriver(SR_LL_C),
	),
	newRandOptimExp(
		"optim:rand:stats:sr-ss",
		"Optimum (Stats, Random Instances) - SR under Strict Subsumption"+
			"Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newRandStatsDriver(SR_SS_C),
	),
	newRandOptimExp(
		"optim:rand:stats:cr-lh",
		"Optimum (Stats, Random Instances) - CR under Lesser Hamming"+
			"Distance Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newRandStatsDriver(CR_LH_C),
	),
	newRandOptimExp(
		"optim:rand:stats:ca-gh",
		"Optimum (Stats, Random Instances) - CA under Greater Hamming"+
			"Distance Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newRandStatsDriver(CA_GH_C),
	),
	newRandOptimExp(
		"optim:rand:val:dfs-ll",
		"Optimum (Value, Random Instances) - DFS under Lesser Level Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newRandValDriver(DFS_LL_C),
	),
	newOptimExp(
		"optim:val:dfs-ll",
		"Optimum (Value) - DFS under Lesser Level Order.\n"+
			"Arguments:\n"+
			"  - n (instances per input\n"+
			"  - List of <tree_file_inputs>.",
		newValDriver(DFS_LL_O),
	),
	newOptimExp(
		"optim:val:sr-ll",
		"Optimum (Value) - SR under Lesser Level Order.\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		newValDriver(SR_LL_O),
	),
	newOptimExp(
		"optim:val:sr-ss",
		"Optimum (Value) - SR under Strict Subsumption Order.\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		newValDriver(SR_SS_O),
	),
	newOptimExp(
		"optim:val:cr-lh",
		"Optimum (Value) - CR under Less Hamming Distance Order.\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		newValDriver(CR_LH_O),
	),
	newOptimExp(
		"optim:val:ca-gh",
		"Optimum (Value) - CA under Greater Hamming Distance Order.\n"+
			"Arguments:\n"+
			"  - List of <optim_file_input>.",
		newValDriver(CA_GH_O),
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
