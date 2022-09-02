package engine

import (
	"github.com/FotiadisM/spencer/pkg/uci"
)

// Engine satisfies the interface uci.Engine
type Engine struct {
	position *Position
}

func (e Engine) SetDebug(b bool, out chan string) {
}

func (e Engine) NewGame(out chan string) {
}

func (e Engine) SetPosition(fen string, out chan string) {
	e.position = NewPosition(fen)
	out <- e.position.String()
}

func (e Engine) ApplyMove(mv string, out chan string) {
	out <- mv + "\n"
}

func (e Engine) Search(esl uci.EngineSearchLimits, out chan string) {
}

func (e Engine) Stop() (bm string, po string) {
	return "", ""
}
