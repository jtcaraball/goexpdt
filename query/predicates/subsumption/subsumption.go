// Subsumption predicate explanation here.
//
// Given a model M and partial instances e1 and e2 we say that M satisfies
// Subsumption(e1, e2) if and only if for every feature i it holds that
// if e1[i] != bottom then e1[i] == e2[i].
package subsumption