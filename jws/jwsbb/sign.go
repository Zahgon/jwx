package jwsbb

import (
	"crypto"
	"io"
)

// Sign generates a JWS signature using the specified key and algorithm.
//
// This function loads the signer registered in the jwsbb package _ONLY_.
// It does not support custom signers that the user might have registered.
//
// rr is an io.Reader that provides randomness for signing. If rr is nil, it defaults to rand.Reader.
// Not all algorithms require this parameter, but it is included for consistency.
// 99% of the time, you can pass nil for rr, and it will work fine.
//
// Deprecated in spirit: in the next major release of jwx (v5), the
// signature of Sign will change to match [SignWithOpts], i.e. it will
// accept an additional [crypto.SignerOpts] parameter immediately before
// rr. Callers that need to pass per-call options today should use
// [SignWithOpts]; callers that do not can keep using Sign and migrate
// when v5 ships by threading a nil opts argument through at the call
// site.
func Sign(key any, alg string, payload []byte, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignWithOpts is like [Sign] but threads an optional
// [crypto.SignerOpts] through to the underlying dsig signer. The
// canonical use case is composite ML-DSA signatures, where a per-call
// domain-separation context (`*mldsa.Options`) must reach
// `filippo.io/mldsa` via the `dsig.SignerWithOpts` interface
// implemented by `github.com/jwx-go/mldsa/v4`. For built-in families
// (HMAC, RSA, ECDSA, EdDSA) the opts argument is ignored.
//
// This function exists as a transitional API. In the next major release
// of jwx (v5) it will be removed and its signature will become the
// canonical shape of [Sign]. Code that uses SignWithOpts today will
// need a mechanical rename to Sign (and nothing else) when v5 ships.
func SignWithOpts(key any, alg string, payload []byte, opts crypto.SignerOpts, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// For custom algorithms registered with dsig, JWS name = dsig name

// Get dsig algorithm info to determine key conversion strategy

func dispatchHMACSign(key any, dsigAlg string, payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func dispatchRSASign(key any, dsigAlg string, payload []byte, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	// Try crypto.Signer first (dsig can handle it directly)
	return nil, nil
}

// Verify it's an RSA key

// Fall back to concrete key types

func dispatchECDSASign(key any, dsigAlg string, payload []byte, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	// Try crypto.Signer first (dsig can handle it directly)
	return nil, nil
}

// Verify it's an ECDSA key

// Fall back to concrete key types

func dispatchEdDSASign(key any, jwsAlg, dsigAlg string, payload []byte, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	// Note: Extension algorithms (e.g. Ed448) are registered as dsig.Custom family,
	// so they take the dsig.Custom branch in Sign() and never reach this function.
	return nil, nil
}

// Try crypto.Signer first (dsig can handle it directly)

// Verify it's an EdDSA key

// Fall back to concrete key types
