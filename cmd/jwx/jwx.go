package main

import (
	"cmp"
	"fmt"
	"io"
	"os"
	"slices"

	"github.com/lestrrat-go/jwx/v4/jwk"
	"github.com/urfave/cli/v2"
)

var topLevelCommands []*cli.Command

type dummyWriteCloser struct {
	io.Writer
}

func (*dummyWriteCloser) Close() error { _ = "STUB: not implemented"; return nil }

func outputFlag() cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func keyFlag(use string) cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func keyFormatFlag() cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func main() {
	var app cli.App
	app.Commands = topLevelCommands
	app.Usage = "Tools for various JWE/JWK/JWS/JWT operations"

	slices.SortFunc(app.Commands, func(a, b *cli.Command) int {
		return cmp.Compare(a.Name, b.Name)
	})

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func dumpJSON(dst io.Writer, v any) error { _ = "STUB: not implemented"; return nil }

func getSource(filename string) (io.ReadCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser), nil
}

func getOutput(filename string) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// Output may be private key material (jwk generate), decrypted
// plaintext (jwe decrypt), or an extracted JWS payload (jws verify).
// Use 0600 and always truncate so a shorter re-run cannot leak
// tail bytes left over from a previous invocation.

func getKeyFile(keyfile, format string) (jwk.Set, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Set), nil
}
