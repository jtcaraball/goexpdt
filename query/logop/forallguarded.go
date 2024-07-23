package logop

import (
	"errors"
	"fmt"

	"github.com/jtcaraball/goexpdt/cnf"
	"github.com/jtcaraball/goexpdt/query"
)

// ForAllGuarded represents a FOR ALL guarded quantifier.
type ForAllGuarded struct {
	// I corresponds to the constant that will be used to materialize the
	// instances that correspond to the ctx's model's nodes. Its ID will be
	// used for scope setting.
	I query.QConst
	// Q corresponds to a sub-query that implements the LogOpQ interface and
	// that is expected to make use of I.
	Q LogOpQ
}

// walkInfo holds information of a step in walk through a model. Used by
// nodeAsValuesIter.
type walkInfo struct {
	prev    int
	visited bool
}

// nodeAsValuesIter allows to iterate over a model's nodes representing them
// as a slice of features stored in t.
type nodeAsValuesIter struct {
	t     *[]query.FeatV
	nodes []query.Node
	walk  []walkInfo
	cp    int
}

// newNodeAsValuesIter initializes a nodeAsValuesIter and returns it alongside
// a pointer to the slice where it will write values to. Returns an error if
// nodes is an empty slice.
func newNodeAsValuesIter(
	dim int,
	nodes []query.Node,
) (*[]query.FeatV, nodeAsValuesIter, error) {
	iter := nodeAsValuesIter{}

	if len(nodes) == 0 {
		return nil, iter, errors.New("Invalid iteration over empty model")
	}

	t := make([]query.FeatV, dim)
	iter.t = &t
	iter.nodes = nodes
	iter.walk = make([]walkInfo, len(nodes))

	return &t, iter, nil
}

// Next attempts to moves the iterator to the next node, updating the value of
// t. Returns true if there is a next node and false otherwise.
func (i *nodeAsValuesIter) Next() bool {
	var pp, zp, op int

	for {
		if i.walk[i.cp].visited {
			// If its a leaf then we attempt to back up.
			if i.nodes[i.cp].IsLeaf() {
				goto backup
			}

			zp = i.nodes[i.cp].ZChild
			if !i.walk[zp].visited {
				(*i.t)[i.nodes[i.cp].Feat] = query.ZERO

				pp = i.cp
				i.cp = zp
				i.walk[i.cp].prev = pp

				continue
			}

			op = i.nodes[i.cp].OChild
			if !i.walk[op].visited {
				(*i.t)[i.nodes[i.cp].Feat] = query.ONE

				pp = i.cp
				i.cp = op
				i.walk[i.cp].prev = pp

				continue
			}

		backup:
			if i.cp == 0 {
				return false
			}

			i.cp = i.walk[i.cp].prev
			(*i.t)[i.nodes[i.cp].Feat] = query.BOT

			continue
		}

		i.walk[i.cp].visited = true

		return true
	}
}

// Encoding returns the CNF formula equivalent to the conjunction all the
// possible CNF formulas of f.Q resulting from instantiating every value of f.I
// in the ctx's model.
func (f ForAllGuarded) Encoding(ctx query.QContext) (ncnf cnf.CNF, err error) {
	defer func() {
		if err != nil {
			err = fmt.Errorf("ForAllGuarded: %w", err)
		}
	}()

	if f.Q == nil {
		return cnf.CNF{}, errors.New("Invalid encoding of nil child")
	}
	if ctx == nil {
		return cnf.CNF{}, errors.New("Invalid encoding with nil ctx")
	}

	ncnf, err = f.buildEncoding(ctx)

	return ncnf, err
}

func (f ForAllGuarded) buildEncoding(
	ctx query.QContext,
) (ncnf cnf.CNF, err error) {
	var cv *[]query.FeatV
	var iter nodeAsValuesIter
	var icnf cnf.CNF

	ctx.AddScope(f.I.ID)
	defer func() {
		if perr := ctx.PopScope(); perr != nil {
			err = errors.Join(err, perr)
		}
	}()

	i := 0
	cv, iter, err = newNodeAsValuesIter(ctx.Dim(), ctx.Nodes())
	if err != nil {
		return ncnf, err
	}

	for iter.Next() {
		if err = ctx.SetScope(i, *cv); err != nil {
			return cnf.CNF{}, err
		}

		icnf, err = f.Q.Encoding(ctx)
		if err != nil {
			return cnf.CNF{}, err
		}

		ncnf = ncnf.Conjunction(icnf)

		i += 1
	}

	return ncnf, err
}
