package uci

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Interface struct{}

func NewInterface() Interface {
	return Interface{}
}

func (ui Interface) Run(r io.Reader) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.ToLower(strings.TrimSpace(scanner.Text()))

		cmds := strings.Fields(line)
		switch cmds[0] {
		case "uci":
		case "debug":
		case "isready":
		case "setoption":
		case "register":
		case "ucinewgame":
		case "position":
		case "stop":
		case "ponderhit":
		case "quit", "q":
			break
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
