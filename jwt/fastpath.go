package jwt

import (
	"github.com/lestrrat-go/jwx/v4/jwa"
)

// fastPathKidSafe reports whether kid can be concatenated into a
// hand-built JSON header literal without escaping. Callers that
// receive false fall through to jws.Sign where encoding/json handles
// the escaping.
func fastPathKidSafe(kid string) bool { _ = "STUB: not implemented"; return false }

// signFast reinvents the wheel a bit to avoid the overhead of
// going through the entire jws.Sign() machinery.
func signFast(t Token, alg jwa.SignatureAlgorithm, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil,

		// Unlike kid, an unsafe alg name cannot silently fall back to
		// jws.Sign: a caller that registered an algorithm with a name
		// containing JSON-special bytes is misconfigured, and slow-path
		// encoding would still produce output under that rogue alg.
		nil
}

// Setup headers
// {"alg":"","typ":"JWT"}
// 1234567890123456789012

// also, if kid != "", we need to add "kid":"$kid"

// "kid":""
// 12345689

// setup the buffer to sign with

// Reuse the combined buffer (base64(hdr).base64(payload)) and append .base64(sig)
// instead of re-encoding hdr and payload from scratch via JoinCompact
