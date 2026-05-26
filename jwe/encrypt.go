package jwe

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v4/jwa"
	"github.com/lestrrat-go/jwx/v4/jwe/internal/keygen"
	"github.com/lestrrat-go/jwx/v4/jwe/jwebb"
)

// encryptKey dispatches key encryption to the appropriate algorithm-specific
// function. This is a standalone function to avoid allocating an encrypter struct.
func encryptKey(cek []byte, keyalg jwa.KeyEncryptionAlgorithm, ctalg jwa.ContentEncryptionAlgorithm, key any, apu, apv []byte, pbes2Count int) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}

// mlkemEncrypterFromKey converts a user-supplied key value into a
// jwebb.MLKEMKeyEncrypter. ML-KEM support is provided by the
// github.com/jwx-go/mlkem companion module — its init() registers raw
// key importers (for stdlib *mlkem.EncapsulationKey768/1024) and a
// jwk.Export adapter that yields an MLKEMKeyEncrypter wrapper.
func mlkemEncrypterFromKey(key any) (jwebb.MLKEMKeyEncrypter, error) {
	_ = "STUB: not implemented"
	return *new(jwebb.MLKEMKeyEncrypter), nil
}

func requireByteKey(key any, alg string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// validateAlgorithmForKey checks that alg is family-compatible with
// key at the WithKey option boundary, surfacing wrong-shape mismatches
// as crisp `jwe.WithKey: ...` errors instead of nested errors deep in
// the dispatcher (e.g. requireByteKey inside the AESKW path). Mirrors
// jws.validateAlgorithmForKey in spirit but is shaped to JWE's
// per-family key conventions.
//
// Permissive carve-outs (return nil, deferring validation):
//
//   - Untyped/extension algorithms: HPKE and ML-KEM. These are
//     extension-pluggable via Register{HPKE,MLKEM}Algorithm; an
//     extension may accept arbitrary key shapes (e.g. an
//     HPKEKeyEncrypter raw type), so the option-time gate cannot
//     enforce a closed-set rule. The downstream dispatch handles
//     unsupported shapes via a typed error.
//   - jwk.Key: the dispatcher unwraps via jwk.Export to a raw key,
//     so the kty-vs-alg check happens then.
//   - Nil key: legitimate for `dir` (caller provides CEK separately
//     via WithCEK) and for callers exploring the API.
//
// All other built-in algorithm families enforce a concrete key-shape
// expectation here. The error is wrapped by the WithKey site so the
// caller sees `jwe.WithKey: ...` consistently.
func validateAlgorithmForKey(alg jwa.KeyEncryptionAlgorithm, key any) error {
	_ = "STUB: not implemented"
	return nil
}

// jwk.Key wrappers: defer to dispatch-time kty validation.

// Caller-supplied KeyEncrypter / KeyDecrypter implementations
// take responsibility for their own key-shape validation. Defer.

// HPKE raw key types implementing the HPKE key interfaces are
// extension-pluggable; defer.

// Extension-pluggable; key shape is the extension's contract.

// "dir" requires a byte slice (the CEK) or a symmetric jwk.

// Unknown algorithm family: defer to dispatch.

func encryptKeyRSA(cek []byte, alg string, key any, encryptFn func([]byte, string, *rsa.PublicKey) (keygen.ByteSource, error)) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	// Handle rsa.PublicKey by value - convert to pointer
	return *new(keygen.ByteSource), nil
}
