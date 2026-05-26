package main

import (
	"io"

	"github.com/lestrrat-go/jwx/v4/jwk"
	"github.com/urfave/cli/v2"
)

func init() {
	topLevelCommands = append(topLevelCommands, makeJwkCmd())
}

func jwkSetFlag() cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func jwkOutputFormatFlag() cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func publicKeyFlag() cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func makeJwkCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

func dumpJWKSet(dst io.Writer, keyset jwk.Set, format string, preserve bool) error {
	_ = "STUB: not implemented"
	return nil
}

func makeJwkGenerateCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

// If the caller is about to dump private key material to a
// terminal — the default when -o is omitted and --public-key
// is not set — emit a stderr warning. The key still goes to
// stdout, so pipes (`| jq`, `| tee key.json`) and shell
// redirections continue to work; only the interactive
// "dump-to-scrollback" footgun gets a signal.

func makeJwkFormatCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

// jwx jwk format <file>
