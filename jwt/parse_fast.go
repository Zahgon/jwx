package jwt

import (
	"github.com/lestrrat-go/jwx/v4/jwa"
)

type fastParseCtx struct {
	alg          jwa.SignatureAlgorithm
	key          any
	skipValidate bool
}

// tryFastPath checks whether the fast path can be used.
// The fast path requires:
//  1. One or two options: a WithKey(SignatureAlgorithm, key) with no suboptions,
//     optionally followed by WithValidate(false)
//  2. The data is not JSON (first byte != '{')
//  3. The data has exactly two '.' separators (compact JWS format)
func tryFastPath(ctx *fastParseCtx, data []byte, options []ParseOption) bool {
	_ = "STUB: not implemented"
	return false
}

// First option must be WithKey

// parseCompactFast is the fast path for parsing JWS compact JWTs.
// It bypasses format detection, option conversion, and the nested decode loop.
// Validation is performed unless ctx.skipValidate is true.
func parseCompactFast(data []byte, ctx *fastParseCtx) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}

// VerifyCompactFast refuses crit-bearing messages. jwt.Parse
// must not be laxer than jws.Verify, so fall through to the
// full jws.Verify path which enforces validateCritical with
// the default-strict (empty) WithCritExtension allowlist.

// The fast path uses strict base64url (RFC 7515). On a
// strict-decode failure, surface a diagnosis first ("input
// is not strict RFC 7515 base64url") and only then mention
// the conditional remedy — the failure shape can't
// distinguish a known-non-conforming issuer from genuinely
// malformed / tampered input, so the caller has to make
// that call deliberately.

// parseCompactCritFallback routes a fast-path-eligible input through
// jws.Verify so the full RFC 7515 §4.1.11 "crit" rule set applies.
// Reached only when the protected header actually contains "crit" (or
// fails to split), so the extra cost is limited to adversarial / RFC 7797
// inputs.
func parseCompactCritFallback(data []byte, ctx *fastParseCtx) (Token, error) {
	_ = "STUB: not implemented"
	return *new(Token), nil
}
