package jwsbb

import (
	"crypto"
	"crypto/ecdsa"
	"io"
	"math/big"
)

// ecdsaHashToDsigAlgorithm maps ECDSA hash functions to dsig algorithm constants
func ecdsaHashToDsigAlgorithm(h crypto.Hash) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// UnpackASN1ECDSASignature unpacks an ASN.1 encoded ECDSA signature into r and s values.
// This is typically used when working with crypto.Signer interfaces that return ASN.1 encoded signatures.
func UnpackASN1ECDSASignature(signed []byte, r, s *big.Int) error {
	_ = "STUB: not implemented"
	// Okay, this is silly, but hear me out. When we use the
	// crypto.Signer interface, the PrivateKey is hidden.
	// But we need some information about the key (its bit size).
	//
	// So while silly, we're going to have to make another call
	// here and fetch the Public key.
	// (This probably means that this information should be cached somewhere)
	return nil
}

// TODO: get this from a pool?

// UnpackECDSASignature unpacks a JWS-format ECDSA signature into r and s values.
// The signature should be in the format specified by RFC 7515 (r||s as fixed-length byte arrays).
func UnpackECDSASignature(signature []byte, pubkey *ecdsa.PublicKey, r, s *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// PackECDSASignature packs the r and s values from an ECDSA signature into a JWS-format byte slice.
// The output format follows RFC 7515: r||s as fixed-length byte arrays.
func PackECDSASignature(r *big.Int, sbig *big.Int, curveBits int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Serialize r and s into fixed-length bytes

// Output as r||s

// SignECDSA generates an ECDSA signature for the given payload using the specified private key and hash.
// The raw parameter should be the pre-computed signing input (typically header.payload).
//
// rr is an io.Reader that provides randomness for signing. if rr is nil, it defaults to rand.Reader.
//
// This function is now a thin wrapper around dsig.SignECDSA. For new projects, you should
// consider using dsig instead of this function.
func SignECDSA(key *ecdsa.PrivateKey, payload []byte, h crypto.Hash, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SignECDSACryptoSigner generates an ECDSA signature using a crypto.Signer interface.
// This function works with hardware security modules and other crypto.Signer implementations.
// The signature is converted from ASN.1 format to JWS format (r||s).
//
// rr is an io.Reader that provides randomness for signing. If rr is nil, it defaults to rand.Reader.
func SignECDSACryptoSigner(signer crypto.Signer, raw []byte, h crypto.Hash, rr io.Reader) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func signECDSACryptoSigner(signer crypto.Signer, signed []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ecdsaVerify(key *ecdsa.PublicKey, buf []byte, h crypto.Hash, r, s *big.Int) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyECDSA verifies an ECDSA signature for the given payload.
// This function verifies the signature using the specified public key and hash algorithm.
// The payload parameter should be the pre-computed signing input (typically header.payload).
//
// This function is now a thin wrapper around dsig.VerifyECDSA. For new projects, you should
// consider using dsig instead of this function.
func VerifyECDSA(key *ecdsa.PublicKey, payload, signature []byte, h crypto.Hash) error {
	_ = "STUB: not implemented"
	return nil
}

// VerifyECDSACryptoSigner verifies an ECDSA signature for crypto.Signer implementations.
// This function is useful for verifying signatures created by hardware security modules
// or other implementations of the crypto.Signer interface.
// The payload parameter should be the pre-computed signing input (typically header.payload).
func VerifyECDSACryptoSigner(signer crypto.Signer, payload, signature []byte, h crypto.Hash) error {
	_ = "STUB: not implemented"
	return nil
}
