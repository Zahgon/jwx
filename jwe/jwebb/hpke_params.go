package jwebb

import (
	"crypto/ecdh"
	"crypto/hpke"

	"github.com/lestrrat-go/jwx/v4/internal/tokens"
)

type hpkeCiphersuite struct {
	curve ecdh.Curve
	kdf   hpke.KDF
	aead  hpke.AEAD
}

var hpkeSuites = map[string]hpkeCiphersuite{
	tokens.HPKE_0_KE: {ecdh.P256(), hpke.HKDFSHA256(), hpke.AES128GCM()},
	tokens.HPKE_1_KE: {ecdh.P384(), hpke.HKDFSHA384(), hpke.AES256GCM()},
	tokens.HPKE_2_KE: {ecdh.P521(), hpke.HKDFSHA512(), hpke.AES256GCM()},
	tokens.HPKE_3_KE: {ecdh.X25519(), hpke.HKDFSHA256(), hpke.AES128GCM()},
	tokens.HPKE_4_KE: {ecdh.X25519(), hpke.HKDFSHA256(), hpke.ChaCha20Poly1305()},
	tokens.HPKE_7_KE: {ecdh.P256(), hpke.HKDFSHA256(), hpke.AES256GCM()},
}

func hpkeSuite(alg string) (hpke.KDF, hpke.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(hpke.KDF), *new(hpke.AEAD), nil
}

// hpkeKEInfo builds the HPKE info parameter for Key Encryption mode
// per draft-ietf-jose-hpke-encrypt-16
// (https://datatracker.ietf.org/doc/draft-ietf-jose-hpke-encrypt/16/):
//
//	"JOSE-HPKE rcpt" || 0xFF || enc_value || 0xFF
func hpkeKEInfo(calg string) []byte { _ = "STUB: not implemented"; return nil }

// hpkePublicKey converts a raw key to hpke.PublicKey.
// Accepts *ecdh.PublicKey and *ecdsa.PublicKey (converted via keyconv).
func hpkePublicKey(alg string, key any) (hpke.PublicKey, error) {
	_ = "STUB: not implemented"
	return *new(hpke.PublicKey), nil
}

// hpkePrivateKey converts a raw key to hpke.PrivateKey.
// Accepts *ecdh.PrivateKey and *ecdsa.PrivateKey (converted via keyconv).
func hpkePrivateKey(alg string, key any) (hpke.PrivateKey, error) {
	_ = "STUB: not implemented"
	return *new(hpke.PrivateKey), nil
}

func toECDHPublicKey(alg string, key any) (*ecdh.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func toECDHPrivateKey(alg string, key any) (*ecdh.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
