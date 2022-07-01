package uci

type UCIOptionCheck interface {
	Default() bool
	Set(bool)
}

type UCIOptionSpin interface {
	Default() int
	Min() int
	Max() int
	St(int)
}

type UCIOptionCombo interface {
	Default() string
	Set(string)
}

type UCIOptionButton interface {
	Set()
}

type UCIOptionString interface {
	Default() string
	Set(string)
}

type EngineOption interface {
	Name() string
	Type() string
}

type EngineInfo struct {
	Name    string
	Version string
	Authros []string
	Options []EngineOption
}

type Engine interface {
	SetPosition(pos string)
	ApplyMove(mv string)
}
