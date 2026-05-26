package jwsbb

import (
	"crypto"
	"crypto/rsa"
	"io"
)

// rsaHashToDsigAlgorithm maps RSA hash functions to dsig algorithm constants
func rsaHashToDsigAlgorithm(h crypto.Hash, pss bool) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// SignRSA generates an RSA signature for the given payload using the specified private key and options.
// The raw parameter should be the pre-computed signing input (typically header.payload).
// If pss is true, RSA-PSS is used; otherwise, PKCS#1 v1.5 is used.
//
// The rr parameter is an optional io.Reader that can be used to provide randomness for signing.
// If rr is nil, it defaults to rand.Reader.
//
// This function is now a thin wrapper around dsig.SignRSA. For new projects, you should
// consider using dsig instead of this function.
func SignRSA(key *rsa.PrivateKey, payload []byte, h crypto.Hash, pss bool, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// VerifyRSA verifies an RSA signature for the given payload and header.
// This function constructs the signing input by encoding the header and payload according to JWS specification,
// then verifies the signature using the specified public key and hash algorithm.
// If pss is true, RSA-PSS verification is used; otherwise, PKCS#1 v1.5 verification is used.
//
// This function is now a thin wrapper around dsig.VerifyRSA. For new projects, you should
// consider using dsig instead of this function.
func VerifyRSA(key *rsa.PublicKey, payload, signature []byte, h crypto.Hash, pss bool) error {
	_ = "STUB: not implemented"
	return nil
}
