package engine

import "github.com/FotiadisM/spencer/pkg/uci"

type Engine struct{}

func (e Engine) SetDebug(b bool, out chan string) {
}

func (e Engine) NewGame(out chan string) {
}

func (e Engine) SetPosition(fen string, out chan string) {
}

func (e Engine) ApplyMove(mv string, out chan string) {
}

func (e Engine) Search(esl uci.EngineSearchLimits, out chan string) {
}

func (e Engine) Stop() (bm string, po string) {
	return "", ""
}
