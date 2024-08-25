package compute_test

import (
	"fmt"
	"testing"

	"github.com/jtcaraball/goexpdt/compute"
	"github.com/jtcaraball/goexpdt/query"
	"github.com/jtcaraball/goexpdt/query/predicates/lel"
)

var (
	satQDTFAttom = compute.QDTFAtom{
		Query: lel.ConstConst{
			I1: query.QConst{Val: []query.FeatV{query.BOT}},
			I2: query.QConst{Val: []query.FeatV{query.ONE}},
		},
	}
	unsatQDTFAttom = compute.QDTFAtom{
		Query: lel.ConstConst{
			I1: query.QConst{Val: []query.FeatV{query.ONE}},
			I2: query.QConst{Val: []query.FeatV{query.BOT}},
		},
	}
)

func TestAndCombinator_True(t *testing.T) {
	comb := compute.AndCombinator{satQDTFAttom, satQDTFAttom}
	ctx := query.BasicQContext(mockModel{dim: 1})
	solver, _ := compute.NewBinSolver(SOLVER)

	sat, err := comb.Sat(ctx, solver)
	if err != nil {
		t.Error(fmt.Errorf("Unexpected error: %w", err))
	}

	if !sat {
		t.Error("Invalid output. Expected SAT but got UNSAT")
	}
}

func TestAndCombinator_False(t *testing.T) {
	comb := compute.AndCombinator{satQDTFAttom, unsatQDTFAttom}
	ctx := query.BasicQContext(mockModel{dim: 1})
	solver, _ := compute.NewBinSolver(SOLVER)

	sat, err := comb.Sat(ctx, solver)
	if err != nil {
		t.Error(fmt.Errorf("Unexpected error: %w", err))
	}

	if sat {
		t.Error("Invalid output. Expected UNSAT but got SAT")
	}
}

func TestOrCombinator_True(t *testing.T) {
	comb := compute.OrCombinator{satQDTFAttom, unsatQDTFAttom}
	ctx := query.BasicQContext(mockModel{dim: 1})
	solver, _ := compute.NewBinSolver(SOLVER)

	sat, err := comb.Sat(ctx, solver)
	if err != nil {
		t.Error(fmt.Errorf("Unexpected error: %w", err))
	}

	if !sat {
		t.Error("Invalid output. Expected SAT but got UNSAT")
	}
}

func TestOrCombinator_False(t *testing.T) {
	comb := compute.OrCombinator{unsatQDTFAttom, unsatQDTFAttom}
	ctx := query.BasicQContext(mockModel{dim: 1})
	solver, _ := compute.NewBinSolver(SOLVER)

	sat, err := comb.Sat(ctx, solver)
	if err != nil {
		t.Error(fmt.Errorf("Unexpected error: %w", err))
	}

	if sat {
		t.Error("Invalid output. Expected UNSAT but got SAT")
	}
}

func TestAtom_Positive(t *testing.T) {
	comb := compute.QDTFAtom{
		Query: lel.ConstConst{
			I1: query.QConst{Val: []query.FeatV{query.BOT}},
			I2: query.QConst{Val: []query.FeatV{query.ONE}},
		},
	}
	ctx := query.BasicQContext(mockModel{dim: 1})
	solver, _ := compute.NewBinSolver(SOLVER)

	sat, err := comb.Sat(ctx, solver)
	if err != nil {
		t.Error(fmt.Errorf("Unexpected error: %w", err))
	}

	if !sat {
		t.Error("Invalid output. Expected SAT but got UNSAT")
	}
}

func TestAtom_Negative(t *testing.T) {
	comb := compute.QDTFAtom{
		Query: lel.ConstConst{
			I1: query.QConst{Val: []query.FeatV{query.BOT}},
			I2: query.QConst{Val: []query.FeatV{query.ONE}},
		},
		Negated: true,
	}
	ctx := query.BasicQContext(mockModel{dim: 1})
	solver, _ := compute.NewBinSolver(SOLVER)

	sat, err := comb.Sat(ctx, solver)
	if err != nil {
		t.Error(fmt.Errorf("Unexpected error: %w", err))
	}

	if sat {
		t.Error("Invalid output. Expected UNSAT but got SAT")
	}
}
