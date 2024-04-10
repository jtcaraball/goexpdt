package main

import (
	"goexpdt/base"
	"goexpdt/trees"
)

func genContext(treePath string) (*base.Context, error) {
	expT, err := trees.LoadTree(treePath)
	if err != nil {
		return nil, err
	}
	ctx := base.NewContext(expT.FeatCount, expT)
	return ctx, nil
}
