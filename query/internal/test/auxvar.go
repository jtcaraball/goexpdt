package test

import "github.com/jtcaraball/goexpdt/query"

const sep string = string(rune(30))

// VarGenBotCount returns a variable with value equal to v with the
// addition of the prefix "botc" separated with the record separator
// character (ascii 30).
func VarGenBotCount(v query.QVar) query.QVar {
	return query.QVar("botc" + sep + string(v))
}
