package main

import (
	"github.com/lestrrat-go/jwx/v4/jwa"
	"github.com/urfave/cli/v2"
)

func init() {
	topLevelCommands = append(topLevelCommands, makeJwsCmd())
}

// resolveSignatureAlgorithm looks up a JWS signature algorithm by name and
// refuses "none". Accepting alg=none at a CLI boundary is the canonical JWT
// footgun (RFC 7518 §3.6): sign would emit an unsigned blob, verify would
// accept any unsigned token. There is no legitimate CLI use case, so we
// reject unconditionally.
func resolveSignatureAlgorithm(name string) (jwa.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return *new(jwa.SignatureAlgorithm), nil
}

func jwsAlgorithmFlag(use string) cli.Flag { _ = "STUB: not implemented"; return *new(cli.Flag) }

func makeJwsCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

func makeJwsParseCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

// jwx jws parse <file>

func makeJwsVerifyCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

// jwx jws verify <file>

func makeJwsSignCmd() *cli.Command { _ = "STUB: not implemented"; return nil }

// jwx jws verify <file>

// headers must go to WithKeySuboptions
