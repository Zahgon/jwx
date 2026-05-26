package main

import (
	"github.com/urfave/cli/v2"
)

func init() {
	topLevelCommands = append(topLevelCommands, makeJwaCmd())
}

func makeJwaCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

// should not reach here
