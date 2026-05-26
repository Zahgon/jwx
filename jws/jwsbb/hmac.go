package jwsbb

import (
	"hash"
)

// hmacHashToDsigAlgorithm maps HMAC hash function sizes to dsig algorithm constants
func hmacHashToDsigAlgorithm(hfunc func() hash.Hash) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SHA256

// SHA384

// SHA512

// SignHMAC generates an HMAC signature for the given payload using the specified hash function and key.
// The raw parameter should be the pre-computed signing input (typically header.payload).
//
// This function is now a thin wrapper around dsig.SignHMAC. For new projects, you should
// consider using dsig instead of this function.
func SignHMAC(key, payload []byte, hfunc func() hash.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyHMAC verifies an HMAC signature for the given payload.
// This function verifies the signature using the specified key and hash function.
// The payload parameter should be the pre-computed signing input (typically header.payload).
//
// This function is now a thin wrapper around dsig.VerifyHMAC. For new projects, you should
// consider using dsig instead of this function.
func VerifyHMAC(key, payload, signature []byte, hfunc func() hash.Hash) error {
	_ = "STUB: not implemented"
	return nil
}
