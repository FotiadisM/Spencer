package uci

import "fmt"

const startFen = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"

func uciHandler(e Engine, ei EngineInfo, out chan string) {
	out <- fmt.Sprintf("id name %v@%v\n", ei.Name, ei.Version)
	for _, a := range ei.Authros {
		out <- fmt.Sprintf("id author %v\n", a)
	}
	// TODO: print available option
}

func debugHandler(e Engine, out chan string) {}

func isReadyHandler(e Engine, out chan string) {}

func setOptionHandler(e Engine, str []string, out chan string) {}

func registerHandler(e Engine, str []string, out chan string) {}

func uciNewGameHandler(e Engine, out chan string) {}

func positionHandler(e Engine, str []string, out chan string) {
	if len(str) == 1 {
		out <- "info string error invalid command\n"
		return
	}
	str = str[1:]

	switch str[0] {
	case "startpos":
		e.SetPosition(startFen)
		str = str[1:]
	case "fen":
		if len(str) == 1 {
			out <- "info string error invalid command\n"
			return
		}
		str = str[1:]
		e.SetPosition(str[0])
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
			e.ApplyMove(m)
		}
	}
}

func goHandler(e Engine, str []string, out chan string) {}

func stopHandler(e Engine, out chan string) {}

func ponderHitHandler(e Engine, out chan string) {}
