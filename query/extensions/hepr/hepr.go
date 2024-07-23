// Higher or Equal Preference Rank formula in the DT-FOIL logic.
//
// Given a model M, partial instances e1 and e2 and a sequence of feature
// indexes f_1, f_2, ..., f_n we say that M satisfies HEPR(e1, e2) if and only
// if it holds that either all the features of the sequence in e1 and e2 are
// different to bottom or if the first index f_i in the sequence where either
// e1 or e2 have a value equal to bottom is such that e1[f_i] != bottom.
package hepr
