package jwebb

import (
	"crypto/ecdh"
	"crypto/ecdsa"

	"github.com/lestrrat-go/jwx/v4/jwe/internal/keygen"
)

// ECDHESKeyGenerator is implemented by raw public key types that can
// perform ECDH-ES key generation for JWE encryption. This allows
// external modules to provide ECDH-ES support for key types not in
// Go's standard library (e.g., X448 from cloudflare/circl).
//
// When jwe.Encrypt encounters a raw key implementing this interface
// in the ECDH-ES path, it delegates key generation to the key itself.
type ECDHESKeyGenerator interface {
	// GenerateECDHES generates an ephemeral key pair, performs the ECDH
	// operation with this public key, and derives the key encryption key
	// via Concat KDF.
	//
	// alg is the derived algorithm label used in the KDF (the content
	// encryption algorithm for bare ECDH-ES, or the key wrapping
	// algorithm for ECDH-ES+AxxxKW).
	//
	// Returns the derived key bytes and the ephemeral public key. The
	// ephemeral public key must be importable by jwk.Import so it can
	// be stored as the 'epk' JWE header.
	GenerateECDHES(alg string, keysize int, apu, apv []byte) (derivedKey []byte, ephemeralPubKey any, err error)
}

// ECDHESKeyDeriver is implemented by raw private key types that can
// perform ECDH-ES key derivation for JWE decryption. This allows
// external modules to provide ECDH-ES support for key types not in
// Go's standard library (e.g., X448 from cloudflare/circl).
//
// When jwe.Decrypt encounters a raw key implementing this interface
// in the ECDH-ES path, it delegates key derivation to the key itself.
type ECDHESKeyDeriver interface {
	// DeriveECDHES performs the ECDH operation using this private key and
	// the given ephemeral public key, then derives the key via Concat KDF.
	//
	// The ephemeralPubKey is the raw key exported from the 'epk' JWE header.
	DeriveECDHES(alg string, keysize int, ephemeralPubKey any, apu, apv []byte) ([]byte, error)
}

// DeriveECDHESRaw performs the Concat KDF key derivation used in ECDH-ES,
// given a pre-computed ECDH shared secret (Z). This is a low-level helper
// for ECDHESKeyGenerator/ECDHESKeyDeriver implementations that handle the
// ECDH computation themselves.
func DeriveECDHESRaw(alg string, zBytes, apu, apv []byte, keysize int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// KeyEncryptECDHESCustom encrypts using ECDH-ES with a custom key type
// that implements ECDHESKeyGenerator.
func KeyEncryptECDHESCustom(cek []byte, alg string, apu, apv []byte, gen ECDHESKeyGenerator, keysize uint32, ctalg string, keywrap bool) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}

// KeyDecryptECDHESCustom decrypts using ECDH-ES with a custom key type
// that implements ECDHESKeyDeriver.
func KeyDecryptECDHESCustom(recipientKey []byte, alg string, apu, apv []byte, deriver ECDHESKeyDeriver, pubkey any, keysize uint32, keywrap bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewECDHESKeyGenerator normalizes a raw key into an ECDHESKeyGenerator.
// If the key already implements ECDHESKeyGenerator, it is returned as-is.
// Otherwise, stdlib key types (*ecdh.PublicKey, *ecdsa.PublicKey, and their
// private key counterparts) are wrapped in an adapter.
func NewECDHESKeyGenerator(key any) (ECDHESKeyGenerator, error) {
	_ = "STUB: not implemented"
	return *new(ECDHESKeyGenerator), nil
}

// ecdhGeneratorFromECDSAPublic converts an *ecdsa.PublicKey into an
// *ecdh.PublicKey via stdlib (*ecdsa.PublicKey).ECDH() and wraps it in
// ecdhGenerator. This routes every ecdsa-input ECDH-ES path through
// crypto/ecdh, which uses identity matching on named NIST curves and
// refuses anything else — including the generic elliptic.CurveParams
// big-int path. That closes the invalid-curve attack surface that a
// caller-controlled ecdsa.PublicKey.Curve field would otherwise expose
// via the deprecated crypto/elliptic.Curve.ScalarMult.
func ecdhGeneratorFromECDSAPublic(pub *ecdsa.PublicKey) (ECDHESKeyGenerator, error) {
	_ = "STUB: not implemented"
	return *new(ECDHESKeyGenerator), nil
}

// NewECDHESKeyDeriver normalizes a raw key into an ECDHESKeyDeriver.
// If the key already implements ECDHESKeyDeriver, it is returned as-is.
// Otherwise, stdlib private key types (*ecdh.PrivateKey, *ecdsa.PrivateKey)
// are wrapped in an adapter.
func NewECDHESKeyDeriver(key any) (ECDHESKeyDeriver, error) {
	_ = "STUB: not implemented"
	return *new(ECDHESKeyDeriver), nil
}

// ecdhGenerator wraps *ecdh.PublicKey to implement ECDHESKeyGenerator.
// Handles both NIST curves (P-256, P-384, P-521) and X25519.
type ecdhGenerator struct {
	key *ecdh.PublicKey
}

func (g *ecdhGenerator) GenerateECDHES(alg string, keysize int, apu, apv []byte) ([]byte, any, error) {
	_ = "STUB: not implemented"
	return nil, *new(any), nil
}

// ecdhDeriver wraps *ecdh.PrivateKey to implement ECDHESKeyDeriver.
type ecdhDeriver struct {
	key *ecdh.PrivateKey
}

func (d *ecdhDeriver) DeriveECDHES(alg string, keysize int, ephemeralPubKey any, apu, apv []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ecdsaDeriver wraps *ecdsa.PrivateKey to implement ECDHESKeyDeriver.
type ecdsaDeriver struct {
	key *ecdsa.PrivateKey
}

func (d *ecdsaDeriver) DeriveECDHES(alg string, keysize int, ephemeralPubKey any, apu, apv []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
