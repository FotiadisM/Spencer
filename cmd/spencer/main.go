package main

import (
	"fmt"
	"os"

	"github.com/FotiadisM/spencer/pkg/engine"
	"github.com/FotiadisM/spencer/pkg/uci"
)

func main() {
	e := engine.Engine{}
	ei := uci.EngineInfo{
		Name:    "Spencer",
		Version: "alpha",
		Authros: []string{"Michalis Fotiadis"},
		Options: []uci.EngineOption{},
	}

	if err := uci.Start(os.Stdin, os.Stdout, e, ei); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
