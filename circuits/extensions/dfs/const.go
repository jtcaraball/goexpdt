package dfs

import (
	"errors"
	"goexpdt/base"
	"goexpdt/cnf"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type dfsConst struct {
	constInst base.ConstInstance
}

// =========================== //
//           METHODS           //
// =========================== //

// Return const dfs.
func Const(constInst base.ConstInstance) *dfsConst {
	return &dfsConst{constInst: constInst}
}

// Return CNF encoding of component.
func (d *dfsConst) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	scpConst, err := d.constInst.Scoped(ctx)
	if err != nil {
		return nil, err
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return d.buildEncoding(scpConst, ctx)
}

// Generate cnf encoding.
func (d *dfsConst) buildEncoding(
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	pnConsts, nnConsts, err := ctx.LeafsAsConsts()
	if err != nil {
		return nil, err
	}
	for _, pnConst := range pnConsts {
		for _, nnConst := range nnConsts {
			if !d.unsafeCompareConsToLeafs(constInst, pnConst, nnConst) {
				return cnf.CNFFromClauses([][]int{{}}), nil
			}
		}
	}
	return &cnf.CNF{}, nil
}

// Return pointer to simplified equivalent component which might be itself.
func (d *dfsConst) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	scpConst, err := d.constInst.Scoped(ctx)
	if err != nil {
		return d, nil
	}
	if err = base.ValidateConstsDim(
		ctx.Dimension,
		scpConst,
	); err != nil {
		return nil, err
	}
	return d.buildSimplified(scpConst, ctx)
}

// Generate simplified component.
func (d *dfsConst) buildSimplified(
	constInst base.Const,
	ctx *base.Context,
) (base.Component, error) {
	pnConsts, nnConsts, err := ctx.LeafsAsConsts()
	if err != nil {
		return nil, err
	}
	for _, pnConst := range pnConsts {
		for _, nnConst := range nnConsts {
			if !d.unsafeCompareConsToLeafs(constInst, pnConst, nnConst) {
				return base.NewTrivial(false), nil
			}
		}
	}
	return base.NewTrivial(true), nil
}

// Return true if constant is not a ancestor of the leafs passed.
func (d *dfsConst) compareConsToLeafs(
	constInst, lConst1, lConst2 base.Const,
) (bool, error) {
	if len(constInst) != len(lConst1) || len(constInst) != len(lConst2) {
		return false, errors.New("Constant and leaf have different dimensions")
	}
	for i, ft := range constInst {
		if ft != base.BOT &&
			lConst1[i] != base.BOT &&
			lConst2[i] != base.BOT &&
			lConst1[i] != lConst2[i] {
			return true, nil
		}
	}
	return false, nil
}

// Return true if constant is not a ancestor of the leafs passed.
// Does not check for an index out of bound error.
func (d *dfsConst) unsafeCompareConsToLeafs(
	constInst, lConst1, lConst2 base.Const,
) bool {
	for i, ft := range constInst {
		if ft != base.BOT &&
			lConst1[i] != base.BOT &&
			lConst2[i] != base.BOT &&
			lConst1[i] != lConst2[i] {
			return true
		}
	}
	return false
}

// Return slice of pointers to component's children.
func (d *dfsConst) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (d *dfsConst) IsTrivial() (yes bool, value bool) {
	return false, false
}
