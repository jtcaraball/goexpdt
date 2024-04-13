package orders

import (
	"errors"
	"fmt"
	"goexpdt/base"
	"goexpdt/cnf"
	"strconv"
)

// =========================== //
//           STRUCTS           //
// =========================== //

type ftPrefLL struct {
	varInst   base.Var
	constInst base.Const
	pref      []int
}

// =========================== //
//           METHODS           //
// =========================== //

// Return ftPrefLL order instance.
func FTPrefLL(varInst base.Var, constInst base.Const, pref []int) *ftPrefLL {
	return &ftPrefLL{varInst: varInst, constInst: constInst, pref: pref}
}

// Return CNF encoding of component.
func (o *ftPrefLL) Encoding(ctx *base.Context) (*cnf.CNF, error) {
	var err error

	if err = validPref(o.pref, ctx.Dimension); err != nil {
		return nil, err
	}

	scpVar := o.varInst.Scoped(ctx)

	if err = base.ValidateConstsDim(
		ctx.Dimension,
		o.constInst,
	); err != nil {
		return nil, err
	}

	return o.buildEncoding(scpVar, o.constInst, ctx)
}

// Return CNF encoding order.
func (o *ftPrefLL) buildEncoding(
	varInst base.Var,
	constInst base.Const,
	ctx *base.Context,
) (*cnf.CNF, error) {
	nCNF := &cnf.CNF{}
	cVName := "c$" + string(varInst)
	llVName := "ll$" + string(varInst) + "$" + constInst.AsString()
	prefVName := "fp$" + o.priorityAsString() + "$" + string(varInst) +
		"$" + constInst.AsString()

	botN := 0
	for _, ft := range constInst {
		if ft == base.BOT {
			botN += 1
		}
	}

	nCNF.ExtendConsistency(genCountClauses(string(varInst), cVName, ctx))
	nCNF.ExtendConsistency(equivLLVarClauses(cVName, llVName, botN, ctx))
	nCNF.ExtendConsistency(
		genFTPrefClauses(
			string(varInst),
			prefVName,
			constInst,
			o.pref,
			ctx,
		),
	)

	// Var is on a lower level or they are equal and var wins on feature
	// preference.
	nCNF.ExtendSemantics(
		[][]int{
			{
				ctx.IVar(llVName, 0, 0),
				ctx.IVar(cVName, ctx.Dimension-1, botN),
			},
			{
				ctx.IVar(llVName, 0, 0),
				ctx.IVar(prefVName, 0, 0),
			},
		},
	)

	return nCNF, nil
}

func (o *ftPrefLL) priorityAsString() string {
	r := ""
	for _, ft := range o.pref {
		r += strconv.Itoa(ft)
	}
	return r
}

// TODO!
// Return pointer to simplified equivalent order which mio ht be itself.
func (o *ftPrefLL) Simplified(
	ctx *base.Context,
) (base.Component, error) {
	return o, nil
}

// Return empty slice of components.
func (o *ftPrefLL) GetChildren() []base.Component {
	return []base.Component{}
}

// yes is true if struct is trivial and value represents its truthiness.
func (o *ftPrefLL) IsTrivial() (yes bool, value bool) {
	return false, false
}

// =========================== //
//           HELPERS           //
// =========================== //

func validPref(pref []int, ftCount int) error {
	if len(pref) == 0 {
		return errors.New("Preference must have at least one feature.")
	}
	seen := make([]bool, ftCount)
	for _, i := range pref {
		if i < 0 || i >= ftCount {
			return fmt.Errorf("Feature index out of range: '%d'.", i)
		}
		if seen[i] {
			return errors.New("Duplicated preference.")
		}
		seen[i] = true
	}
	return nil
}

func genCountClauses(varName, cName string, ctx *base.Context) [][]int {
	var i, j int
	clauses := [][]int{}
	for i = 1; i < ctx.Dimension; i++ {
		clauses = append(
			clauses,
			[]int{
				-ctx.IVar(cName, i, 0),
				-ctx.Var(varName, i, base.BOT.Val()),
			},
			[]int{
				-ctx.IVar(cName, i, 0),
				ctx.IVar(cName, i-1, 0),
			},
			[]int{
				-ctx.IVar(cName, i-1, 0),
				ctx.Var(varName, i, base.BOT.Val()),
				ctx.IVar(cName, i, 0),
			},
		)
		for j = 1; j < i+2; j++ {
			clauses = append(
				clauses,
				[]int{
					-ctx.IVar(cName, i, j),
					ctx.IVar(cName, i-1, j-1),
					-ctx.Var(varName, i, base.BOT.Val()),
				},
				[]int{
					ctx.IVar(cName, i, j),
					-ctx.IVar(cName, i-1, j-1),
					-ctx.Var(varName, i, base.BOT.Val()),
				},
				[]int{
					-ctx.IVar(cName, i, j),
					ctx.IVar(cName, i-1, j),
					ctx.Var(varName, i, base.BOT.Val()),
				},
				[]int{
					ctx.IVar(cName, i, j),
					-ctx.IVar(cName, i-1, j),
					ctx.Var(varName, i, base.BOT.Val()),
				},
			)
		}
	}
	for i = 0; i < ctx.Dimension; i++ {
		for j = i + 2; j <= ctx.Dimension; j++ {
			clauses = append(clauses, []int{-ctx.IVar(cName, i, j)})
		}
	}
	clauses = append(
		clauses,
		[]int{
			-ctx.IVar(cName, 0, 1),
			ctx.Var(varName, 0, base.BOT.Val()),
		},
		[]int{
			-ctx.Var(varName, 0, base.BOT.Val()),
			ctx.IVar(cName, 0, 1),
		},
		[]int{
			-ctx.IVar(cName, 0, 0),
			-ctx.Var(varName, 0, base.BOT.Val()),
		},
		[]int{
			ctx.Var(varName, 0, base.BOT.Val()),
			ctx.IVar(cName, 0, 0),
		},
	)
	return clauses
}

func equivLLVarClauses(cName, llName string, l int, ctx *base.Context) [][]int {
	clauses := [][]int{}

	// b = n
	// -c1 ^ -c2 ^ ... ^ -cn <-> ll
	// c1 v c2 v .... v cn v ll
	// (-ll v -c1) ^ (-ll v -c2) ^ ... ^ (-ll v -cn)

	countThenAux := []int{ctx.IVar(llName, 0, 0)}
	for i := 0; i <= l; i++ {
		countThenAux = append(countThenAux, ctx.IVar(cName, ctx.Dimension-1, i))
		clauses = append(clauses, []int{
			-ctx.IVar(cName, ctx.Dimension-1, i),
			-ctx.IVar(llName, 0, 0),
		})
	}
	clauses = append(clauses, countThenAux)

	return clauses
}

func genFTPrefClauses(
	vName, fpName string,
	c base.Const,
	pref []int,
	ctx *base.Context,
) [][]int {
	clauses := [][]int{}

	ftIdx := pref[len(pref)-1]

	// fp_n <-> v_n != bot ^ c_n == bot
	if c[ftIdx] != base.BOT {
		// fp_n = false
		clauses = append(clauses, []int{-ctx.IVar(fpName, len(pref)-1, 0)})
	} else {
		// fp_n <-> -(v_n == bot)
		clauses = append(
			clauses,
			[]int{
				ctx.Var(vName, ftIdx, base.BOT.Val()),
				ctx.IVar(fpName, len(pref)-1, 0),
			},
			[]int{
				-ctx.IVar(fpName, len(pref)-1, 0),
				-ctx.Var(vName, ftIdx, base.BOT.Val()),
			},
		)
	}

	for i := len(pref) - 2; i >= 0; i-- {
		ftIdx = pref[i]
		// fp_i <-> (v_i != bot ^ c_i == bot) v
		//          (
		//             [(v_i == bot ^ c_i == bot) v (v_i != bot ^ c_i != bot)]
		//             ^
		//             fp_{i+1}
        //          )
		if c[ftIdx] != base.BOT {
			// fp_i <-> (-(v_i == bot) ^ fp_{i+1})
			// (v_i == bot v -fp_{i+1} v fp_i)
			// (-fp_i v -(v_i == bot)) ^ (-fp_i v fp_{i+1})
			clauses = append(
				clauses,
				[]int{
					ctx.Var(vName, ftIdx, base.BOT.Val()),
					-ctx.IVar(fpName, i+1, 0),
					ctx.IVar(fpName, i, 0),
				},
				[]int{
					-ctx.IVar(fpName, i, 0),
					-ctx.Var(vName, ftIdx, base.BOT.Val()),
				},
				[]int{
					-ctx.IVar(fpName, i, 0),
					ctx.IVar(fpName, i+1, 0),
				},
			)
			continue
		}
		// fp_i <-> (v_i != bot) v (v_i == bot ^ fp_{i+1})
		// fp_i <-> -(v_i == bot) v fp_{i+1}
		// -(v_i == bot) v fp_{i+1} v -fp_i
		// (fp_i v (v_i == bot)) ^ (fp_i v -fp_{i+1})
		clauses = append(
			clauses,
			[]int{
				-ctx.Var(vName, ftIdx, base.BOT.Val()),
				ctx.IVar(fpName, i+1, 0),
				-ctx.IVar(fpName, i, 0),
			},
			[]int{
				ctx.IVar(fpName, i, 0),
				ctx.Var(vName, ftIdx, base.BOT.Val()),
			},
			[]int{
				ctx.IVar(fpName, i, 0),
				-ctx.IVar(fpName, i+1, 0),
			},
		)
	}

	return clauses
}
