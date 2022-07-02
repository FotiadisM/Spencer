package uci

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Start(r io.Reader, w io.Writer, e Engine, ei EngineInfo) error {
	out := make(chan string)
	defer close(out)

	go func() {
		for s := range out {
			fmt.Fprint(w, s)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))

		str := strings.Fields(line)
		if len(str) == 0 {
			continue
		}

		switch str[0] {
		case "uci":
			uciHandler(e, ei, out)
		case "debug":
			debugHandler(e, str, out)
		case "isready":
			isReadyHandler(e, out)
		case "setoption":
			setOptionHandler(e, str, out)
		case "register":
			registerHandler(e, str, out)
		case "ucinewgame":
			uciNewGameHandler(e, out)
		case "position":
			positionHandler(e, str, out)
		case "go":
			goHandler(e, str, out)
		case "stop":
			stopHandler(e, out)
		case "ponderhit":
			ponderHitHandler(e, out)
		case "quit", "q":
			return nil
		default:
			out <- "info string error invalid command\n"
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
