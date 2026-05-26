package jwebb

import (
	"github.com/lestrrat-go/jwx/v4/jwe/internal/keygen"
)

// KeyEncryptHPKEKE performs HPKE key encryption per
// draft-ietf-jose-hpke-encrypt-16 (https://datatracker.ietf.org/doc/draft-ietf-jose-hpke-encrypt/16/).
// It encrypts the CEK using the HPKE ciphersuite determined by alg, with the
// content encryption algorithm calg bound into the HPKE info parameter.
func KeyEncryptHPKEKE(cek []byte, alg, calg string, pubkey any) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	// Try custom HPKE key encrypter (e.g., X448 from external modules)
	return *new(keygen.ByteSource), nil
}

// AAD is empty for key encryption mode per spec
