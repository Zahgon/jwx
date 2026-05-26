package jwxtest

import (
	"context"
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/rsa"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/lestrrat-go/jwx/v4/jwa"
	"github.com/lestrrat-go/jwx/v4/jwk"
)

// JKUFetcher is a minimal jwk.Fetcher for tests that exercise
// jws.WithVerifyAuto / jwt.WithVerifyAuto. It mirrors enough of the
// jwkfetch.Client contract to cover the jku verification paths without
// pulling the jwkfetch companion into the core module's dependency
// graph.
//
// Allow is consulted on every Fetch call before any network request.
// A nil Allow permits every URL, matching jwkfetch.Client's default.
// Set Allow to a restrictive predicate to simulate whitelist-driven
// rejection.
type JKUFetcher struct {
	Client *http.Client
	Allow  func(url string) bool
}

// Fetch implements jwk.Fetcher. It applies Allow, GETs url via
// Client, and parses the response body as a jwk.Set. No body-size cap,
// no redirect policy, no parse options — tests that need richer
// behavior should declare their own fetcher.
func (f *JKUFetcher) Fetch(ctx context.Context, url string) (jwk.Set, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Set), nil
}

func GenerateRsaKey() (*rsa.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func GenerateRsaJwk() (jwk.Key, error) { _ = "STUB: not implemented"; return *new(jwk.Key), nil }

func GenerateRsaPublicJwk() (jwk.Key, error) { _ = "STUB: not implemented"; return *new(jwk.Key), nil }

func GenerateEcdsaKey(alg jwa.EllipticCurveAlgorithm) (*ecdsa.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GenerateEcdsaJwk() (jwk.Key, error) { _ = "STUB: not implemented"; return *new(jwk.Key), nil }

func GenerateEcdsaPublicJwk() (jwk.Key, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Key), nil
}

func GenerateSymmetricKey() []byte { _ = "STUB: not implemented"; return nil }

func GenerateSymmetricJwk() (jwk.Key, error) { _ = "STUB: not implemented"; return *new(jwk.Key), nil }

func GenerateEd25519Key() (ed25519.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(ed25519.PrivateKey), nil
}

func GenerateEd25519Jwk() (jwk.Key, error) { _ = "STUB: not implemented"; return *new(jwk.Key), nil }

func GenerateX25519Key() (*ecdh.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

func GenerateX25519Jwk() (jwk.Key, error) { _ = "STUB: not implemented"; return *new(jwk.Key), nil }

func WriteFile(dir, template string, src io.Reader) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func WriteJSONFile(dir, template string, v any) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func DumpFile(t *testing.T, file string) { _ = "STUB: not implemented"; return }

// Looks like a JSON-like thing. Dump that in a formatted manner, and
// be done with it

// If the contents do not look like JSON, then we attempt to parse each content
// based on heuristics (from its file name) and do our best

// cross our fingers our jwe implementation works

func CreateTempFile(dir, template string) (*os.File, func(), error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func ReadFile(file string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func ParseJwkFile(_ context.Context, file string) (jwk.Key, error) {
	_ = "STUB: not implemented"
	return *new(jwk.Key), nil
}

func DecryptJweFile(ctx context.Context, file string, alg jwa.KeyEncryptionAlgorithm, jwkfile string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncryptJweFile(ctx context.Context, dir string, payload []byte, keyalg jwa.KeyEncryptionAlgorithm, keyfile string, contentalg jwa.ContentEncryptionAlgorithm, compressalg jwa.CompressionAlgorithm) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// The latchset jose CLI rejects PBES2 tokens with p2c above 32768.
// jwx's own defaults follow OWASP 2023 (six-digit counts) which jose
// refuses. Downgrade explicitly here so interop roundtrips pass;
// real-world producers should not use this helper.

func VerifyJwsFile(ctx context.Context, file string, alg jwa.SignatureAlgorithm, jwkfile string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func SignJwsFile(ctx context.Context, dir string, payload []byte, alg jwa.SignatureAlgorithm, keyfile string) (string, func(), error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}
