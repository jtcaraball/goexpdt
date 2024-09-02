package test

import "github.com/jtcaraball/goexpdt/query"

const sep string = string(rune(30))

// VarGenBotCount returns a variable with value equal to v with the addition of
// the prefix "botc" separated with the record separator character (ascii 30).
func VarGenBotCount(v query.QVar) query.QVar {
	return query.QVar("botc" + sep + string(v))
}

// VarGenBotCount returns a variable with value equal to v with the addition of
// the prefix "reach" separated with the record separator character (ascii 30).
func VarGenNodeReach(v query.QVar) query.QVar {
	return query.QVar("reach" + sep + string(v))
}

// VarGenHammingDistance returns a variable equal to the sorted concatenation
// of variables v1 and v2 with the addition of the prefix "hdist" separated
// using the record separator character (ascii 30).
func VarGenHammingDistance(v1, v2 query.QVar) query.QVar {
	if string(v1) < string(v2) {
		return query.QVar("hdist" + sep + string(v1) + sep + string(v2))
	}
	return query.QVar("hdist" + sep + string(v2) + sep + string(v1))
}

// VarGenEqualFeat returns a variable equal to the sorted concatenation of
// variables v1 and v2 with the addition of the prefix "eqf" separated using
// the record separator character (ascii 30).
func VarGenEqualFeat(v1, v2 query.QVar) query.QVar {
	if string(v1) < string(v2) {
		return query.QVar("eqf" + sep + string(v1) + sep + string(v2))
	}
	return query.QVar("eqf" + sep + string(v2) + sep + string(v1))
}

// VarGenFeaturePreference returns a variable equal to the concatenation of
// variable v1 and v2 with the addition of the prefix "fp" + string(p)
// separated using the record separator character (ascii 30).
func VarGenFeaturePreference(p, v1, v2 query.QVar) query.QVar {
	return query.QVar(
		"fp" + sep + string(p) + sep + string(v1) + sep + string(v2),
	)
}
