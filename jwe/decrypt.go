package jwe

import (
	"github.com/lestrrat-go/jwx/v4/jwa"
	"github.com/lestrrat-go/jwx/v4/jwe/internal/content_crypt"
	"github.com/lestrrat-go/jwx/v4/jwe/jwebb"
)

// decryptCEKContext holds algorithm-agnostic context needed during CEK decryption.
type decryptCEKContext struct {
	maxPBES2Count int
	minPBES2Count int
	ctalg         jwa.ContentEncryptionAlgorithm
	contentCipher content_crypt.Cipher
}

// decryptCEK dispatches key decryption to the appropriate per-family
// function based on the algorithm. Each function extracts its own
// algorithm-specific parameters from the merged headers.
func decryptCEK(alg jwa.KeyEncryptionAlgorithm, key any, msg *Message, recipient Recipient, headers Headers, ctx *decryptCEKContext) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKeyDirect(recipientKey []byte, alg string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKeyPBES2(recipientKey []byte, alg string, key any, headers Headers, maxCount, minCount int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse p2c into int64 directly. Float64 cannot represent integers
// above 2^53 exactly; comparing a parsed value against a high
// MaxPBES2Count cap in float-space and then casting via int(...) lets
// out-of-range values silently round into the accepted range, which
// would defeat the cap when callers raise it past 2^53. int64 keeps
// the bound check exact.

// Reject values outside int64 range before casting; the cast
// of an out-of-range float to int is implementation-defined.
// Use explicit float-domain bounds (2^63 / -2^63) instead of
// math.MaxInt64 / MinInt64 so the comparison is independent
// of the platform's int width and the constants do not need
// implicit conversion.

// 2^63, the smallest float > MaxInt64
// -2^63, exact float = MinInt64

func decryptKeyAESGCMKW(recipientKey []byte, alg string, key any, headers Headers) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKeyECDHES(recipientKey []byte, alg string, ctalg jwa.ContentEncryptionAlgorithm, key any, headers Headers) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Extract ephemeral public key from headers

func decryptKeyHPKE(recipientKey []byte, alg string, ctalg jwa.ContentEncryptionAlgorithm, key any, headers Headers) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKeyMLKEM(recipientKey []byte, alg string, ctalg jwa.ContentEncryptionAlgorithm, key any, headers Headers) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// mlkemDecrypterFromKey is the decrypt-side counterpart to
// mlkemEncrypterFromKey. See its doc for the conversion strategy.
func mlkemDecrypterFromKey(key any) (jwebb.MLKEMKeyDecrypter, error) {
	_ = "STUB: not implemented"
	return *new(jwebb.MLKEMKeyDecrypter), nil
}

func decryptKeyRSA15(recipientKey []byte, _ string, key any, contentCipher content_crypt.Cipher) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKeyRSAOAEP(recipientKey []byte, alg string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func decryptKeyAESKW(recipientKey []byte, alg string, key any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
