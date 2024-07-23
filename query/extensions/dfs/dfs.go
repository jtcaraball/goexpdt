// Determinant Feature Set formula in the DT-FOIL logic.
//
// Given a model T and partial instance e we say that T satisfies DFS(e) if and
// only if for every partial instance e', such that the set of features with
// value equal to bottom in e is the same to the set of features with value
// equal to bottom in e', it holds that all the completions of e' receive the
// same classification over T.
package dfs
