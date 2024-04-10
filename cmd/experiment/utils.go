package main

import (
	"goexpdt/base"
	"goexpdt/trees"
)

const (
	INPUTDIR  = "input"
	OUTPUTDIR = "output"
)

func genContext(treePath string) (*base.Context, error) {
	expT, err := trees.LoadTree(treePath)
	if err != nil {
		return nil, err
	}
	ctx := base.NewContext(expT.FeatCount, expT)
	return ctx, nil
}
