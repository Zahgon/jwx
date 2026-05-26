package main

import (
	"github.com/urfave/cli/v2"
)

func init() {
	topLevelCommands = append(topLevelCommands, makeJweCmd())
}

func makeJweCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

func keyEncryptionFlag(required bool) cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func makeJweEncryptCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

func makeJweDecryptCmd() *cli.Command { _ = "STUB: not implemented"; return nil }
