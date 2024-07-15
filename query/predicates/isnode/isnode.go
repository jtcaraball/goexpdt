// Node predicate in the DT-FOIL logic.
//
// Given a decision tree T and a partial instance e we say that T satisfies
// Node(e) if and only if there exists a node u in T such that e = e_u.
// That is, Node(e) is true if and only if there is a node u in T such that
// for every feature that T decides on to reach u, e has the same value as
// the path chosen and for every feature not decided on e has a bottom value.
package isnode
