package jws

import (
	"sync"

	"github.com/lestrrat-go/jwx/v4/jwa"
)

type defaultVerifier struct {
	alg jwa.SignatureAlgorithm
}

func (v defaultVerifier) Verify(key any, payload, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// Verifier is an interface for objects that can verify signatures.
//
// Custom verification algorithms can be registered by implementing this
// interface and calling RegisterVerifier.
type Verifier interface {
	Verify(key any, payload, signature []byte) error
}

// VerifierFunc is a function type that implements the Verifier interface.
type VerifierFunc func(key any, payload, signature []byte) error

func (f VerifierFunc) Verify(key any, payload, signature []byte) error {
	_ = "STUB: not implemented"
	return nil
}

var verifierDB sync.Map // map[jwa.SignatureAlgorithm]Verifier

// VerifierFor returns a Verifier for the given signature algorithm.
//
// If a custom Verifier has been registered for the algorithm, it is returned.
// Otherwise, a default verifier that delegates to jwsbb.Verify is returned.
func VerifierFor(alg jwa.SignatureAlgorithm) (Verifier, error) {
	_ = "STUB: not implemented"
	return *new(Verifier), nil
}

//nolint:forcetypeassert
// always stored as Verifier

// RegisterVerifier registers a custom Verifier for the given algorithm.
//
// This function also calls jwa.RegisterSignatureAlgorithm to register
// the algorithm in this module's algorithm database.
//
// Built-in algorithm identifiers are reserved; attempts to register a Verifier
// using a built-in name with different metadata will fail. Re-registering the
// exact built-in algorithm value is allowed.
func RegisterVerifier(alg jwa.SignatureAlgorithm, v Verifier) error {
	_ = "STUB: not implemented"
	return nil
}

// UnregisterVerifier removes the verifier associated with the given algorithm.
//
// Note that when you call this function, the algorithm itself is
// not automatically unregistered from this module's algorithm database.
// This is because the algorithm may still be required for signing or
// some other operation (however unlikely, it is still possible).
// Therefore, in order to completely remove the algorithm, you must
// call jwa.UnregisterSignatureAlgorithm yourself.
//
// The error return is reserved for future validation (for example,
// refusing to unregister a built-in algorithm) and is always nil
// today. Callers — especially those scripting
// Register/Unregister cycles from init() — should check the
// returned value and propagate on failure to stay
// forward-compatible, matching the convention on [RegisterVerifier].
func UnregisterVerifier(alg jwa.SignatureAlgorithm) error { _ = "STUB: not implemented"; return nil }
