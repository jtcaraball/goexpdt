package instances

// =========================== //
//           STRUCTS           //
// =========================== //

type Const []featV

type featV struct {
	val int
}

// =========================== //
//          VARIABLES          //
// =========================== //

var (
	ZERO = Zero()
	ONE = One()
	BOT = Bot()
	FeatValues = []featV{ZERO, ONE, BOT}
)

// =========================== //
//           METHODS           //
// =========================== //

func Zero() featV {
	return featV{val: 0}
}

func One() featV {
	return featV{val: 1}
}

func Bot() featV {
	return featV{val: 2}
}

func (f featV) Val() int {
	return f.val
}
