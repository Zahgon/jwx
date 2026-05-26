package jwsbb

import (
	"crypto"
)

// Verify verifies a JWS signature using the specified key and algorithm.
//
// This function loads the verifier registered in the jwsbb package _ONLY_.
// It does not support custom verifiers that the user might have registered.
//
// Deprecated in spirit: in the next major release of jwx (v5), the
// signature of Verify will change to match [VerifyWithOpts], i.e. it
// will accept an additional [crypto.SignerOpts] parameter at the end.
// Callers that need to pass per-call options today should use
// [VerifyWithOpts]; callers that do not can keep using Verify and
// migrate when v5 ships by threading a nil opts argument through at
// the call site.
func Verify(key any, alg string, payload, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyWithOpts is like [Verify] but threads an optional
// [crypto.SignerOpts] through to the underlying dsig verifier. See
// [SignWithOpts] for the rationale and the migration story.
func VerifyWithOpts(key any, alg string, payload, signature []byte, opts crypto.SignerOpts) error {
	_ = "STUB: not implemented"
	return nil
}

// For custom algorithms registered with dsig, JWS name = dsig name

// Get dsig algorithm info to determine key conversion strategy

func dispatchHMACVerify(key any, dsigAlg string, payload, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func dispatchRSAVerify(key any, dsigAlg string, payload, signature []byte) error {
	_ = "STUB: not implemented"
	// Try crypto.Signer first (dsig can handle it directly)
	return nil
}

// Verify it's an RSA key

// Fall back to concrete key types

func dispatchECDSAVerify(key any, dsigAlg string, payload, signature []byte) error {
	_ = "STUB: not implemented"
	// Try crypto.Signer first (dsig can handle it directly)
	return nil
}

// Verify it's an ECDSA key

// Fall back to concrete key types

func dispatchEdDSAVerify(key any, jwsAlg, dsigAlg string, payload, signature []byte) error {
	_ = "STUB: not implemented"
	// Note: Extension algorithms (e.g. Ed448) are registered as dsig.Custom family,
	// so they take the dsig.Custom branch in Verify() and never reach this function.
	return nil
}

// Try crypto.Signer first (dsig can handle it directly)

// Verify it's an EdDSA key

// Fall back to concrete key types
