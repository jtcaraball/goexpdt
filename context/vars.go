package context

// =========================== //
//          CONSTANTS          //
// =========================== //

const (
	ZERO uint8 = 0
	ONE uint8 = 1
	BOT uint8 = 2
)

// =========================== //
//           STRUCTS           //
// =========================== //

type Var struct {
	name string
	feat int
	value uint8
}

// =========================== //
//           METHODS           //
// =========================== //

func newVar(name string, feat int, value uint8) *Var{
	return &Var{name: name, feat: feat, value: value}
}
