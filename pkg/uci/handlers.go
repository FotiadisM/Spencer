package uci

import "fmt"

const startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func uciHandler(e Engine, ei EngineInfo, out chan string) {
	out <- fmt.Sprintf("id name %v@%v\n", ei.Name, ei.Version)
	for _, a := range ei.Authros {
		out <- fmt.Sprintf("id author %v\n", a)
	}

	for _, o := range ei.Options {
		switch o := o.(type) {
		case UCIOptionCheck:
			out <- fmt.Sprintf("option name %v type %v default %v\n", o.Name(), o.Type(), o.Default())
		case UCIOptionSpin:
			out <- fmt.Sprintf("option name %v type %v default %v min %v max %v\n", o.Name(), o.Type(), o.Default(), o.Min(), o.Max())
		case UCIOptionCombo:
			out <- fmt.Sprintf("option name %v type %v default %v\n", o.Name(), o.Type(), o.Default())
		case UCIOptionButton:
			out <- fmt.Sprintf("option name %v type %v\n", o.Name(), o.Type())
		case UCIOptionString:
			out <- fmt.Sprintf("option name %v type %v default %v\n", o.Name(), o.Type(), o.Default())
		}
	}

	out <- "uciok\n"
}

func debugHandler(e Engine, str []string, out chan string) {
	if len(str) != 2 {
		out <- "info string error invalid command\n"
		return
	}
	switch str[1] {
	case "on":
		e.SetDebug(true, out)
	case "off":
		e.SetDebug(false, out)
	default:
		out <- "info string error invalid command\n"
		return
	}
}

func isReadyHandler(e Engine, out chan string) {
	out <- "readyok"
}

func setOptionHandler(e Engine, str []string, out chan string) {
	// TODO: Implement
}

func registerHandler(e Engine, str []string, out chan string) {
	// TODO: Implement
}

func uciNewGameHandler(e Engine, out chan string) {
	e.NewGame(out)
}

func positionHandler(e Engine, str []string, out chan string) {
	if len(str) == 1 {
		out <- "info string error invalid command\n"
		return
	}
	str = str[1:]

	switch str[0] {
	case "startpos":
		e.SetPosition(startFen, out)
		str = str[1:]
	case "fen":
		if len(str) == 1 {
			out <- "info string error invalid command\n"
			return
		}
		str = str[1:]
		e.SetPosition(str[0], out)
		str = str[1:]
	default:
		return
	}

	if len(str) != 0 {
		if str[0] != "moves" {
			out <- "info string error invalid command\n"
			return
		}
		str = str[1:]
		for _, m := range str {
			e.ApplyMove(m, out)
		}
	}
}

func goHandler(e Engine, str []string, out chan string) {
	// TODO: parse str
	esl := EngineSearchLimits{}
	go e.Search(esl, out)
}

func stopHandler(e Engine, out chan string) {
	bm, po := e.Stop()
	if po == "" {
		out <- fmt.Sprintf("bestmove %v\n", bm)
		return
	}
	out <- fmt.Sprintf("bestmove %v ponder %v\n", bm, po)
}

func ponderHitHandler(e Engine, out chan string) {
	// TODO: Implement
}
