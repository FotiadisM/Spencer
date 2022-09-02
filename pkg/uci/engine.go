package uci

type UCIOptionType int

const (
	Check UCIOptionType = iota
	Spin
	Combo
	Button
	String
)

type EngineOption interface {
	Name() string
	Type() UCIOptionType
}

type UCIOptionCheck interface {
	EngineOption
	Default() bool
	Set(bool)
}

type UCIOptionSpin interface {
	EngineOption
	Default() int
	Min() int
	Max() int
	St(int)
}

type UCIOptionCombo interface {
	EngineOption
	Default() string
	Set(string)
}

type UCIOptionButton interface {
	EngineOption
	Set()
}

type UCIOptionString interface {
	EngineOption
	Default() string
	Set(string)
}

type EngineInfo struct {
	Name    string
	Version string
	Authros []string
	Options []EngineOption
}

type EngineSearchLimits struct {
	SearchMoves []string
	Ponder      bool
	WTime       int
	BTime       int
	WInc        int
	BInc        int
	MovesToGo   int
	Depth       int
	Nodes       int
	Mate        int
	MoveTime    int
	Infinite    bool
}

type Engine interface {
	SetDebug(b bool, out chan string)
	NewGame(out chan string)
	SetPosition(fen string, out chan string)
	ApplyMove(mv string, out chan string)
	Search(esl EngineSearchLimits, out chan string)
	Stop() (bm string, po string)
}
