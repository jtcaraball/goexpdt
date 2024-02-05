package subsumption

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"stratifoiled/components"
	"stratifoiled/components/instances"
	"stratifoiled/components/operators"
	"stratifoiled/sfdtest"
	"testing"
)

const varvarSUFIX = "subsumtpion.varvar"

// =========================== //
//           HELPERS           //
// =========================== //

func runSubsumptionVarVar(
	t *testing.T,
	id int,
	c1, c2 instances.Const,
	expCode int,
) {
	x := instances.NewVar("x")
	y := instances.NewVar("y")
	context := components.NewContext(DIM, nil)
	formula := operators.WithVar(
		x,
		operators.WithVar(
			y,
			operators.And(
				operators.And(VarConst(x, c1), ConstVar(c1, x)),
				operators.And(
					operators.And(VarConst(y, c2), ConstVar(c2, y)),
					VarVar(x, y),
				),
			),
		),
	)
	cnf, err := formula.Encoding(context)
	if err != nil {
		t.Errorf("Encoding error: %s", err.Error())
		return
	}
	filePath := fmt.Sprintf("%s.%s.%d", CNFPATH, varvarSUFIX, id)
	cnf.ToFile(filePath)
	cmd := exec.Command(SOLVER, filePath)
	retCode, err := sfdtest.RunSolver(t, cmd)
	if err != nil {
		t.Errorf("Solver execution error: %s", err.Error())
		return
	}
	if retCode != expCode {
		t.Errorf(
			"Wrong answer: expected exit code %d but got %d",
			expCode,
			retCode,
		)
	}
}

// =========================== //
//            TESTS            //
// =========================== //

func TestVarVar_Encoding(t *testing.T) {
	// Cleanup
	t.Cleanup(
		func() {
			files, err := filepath.Glob(
				fmt.Sprintf("%s.%s*", CNFPATH, varvarSUFIX),
			)
			if err != nil {
				t.Errorf(fmt.Sprintf("Error in cleanup: %s", err.Error()))
				return // No real handling we can do here.
			}
			for _, file := range files {
				os.Remove(file)
			}
		},
	)
	// Run tests
	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			runSubsumptionVarVar(t, i, tc.val1, tc.val2, tc.expCode)
		})
	}
}