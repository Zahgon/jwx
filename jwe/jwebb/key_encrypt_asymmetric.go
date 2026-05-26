package jwebb

import (
	"crypto/rsa"

	"github.com/lestrrat-go/jwx/v4/jwe/internal/keygen"
)

// KeyEncryptRSA15 encrypts the CEK using RSA PKCS#1 v1.5
func KeyEncryptRSA15(cek []byte, _ string, pubkey *rsa.PublicKey) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}

// KeyEncryptRSAOAEP encrypts the CEK using RSA OAEP
func KeyEncryptRSAOAEP(cek []byte, alg string, pubkey *rsa.PublicKey) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}
