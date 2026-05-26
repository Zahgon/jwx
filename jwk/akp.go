package jwk

import (
	"crypto"
	"encoding/json"
	"fmt"

	"github.com/lestrrat-go/jwx/v4/jwa"
)

func init() {
	// Register probe field for "priv" so the parser can distinguish
	// AKP public keys from AKP private keys (which use "priv" instead of "d").
	if err := RegisterProbeField[json.RawMessage]("Priv", "priv"); err != nil {
		panic(fmt.Errorf("jwk/akp: failed to register probe for 'priv' field: %w", err))
	}
}

var normalizedAKP KeyKind

func init() {
	normalizedAKP = KeyKind(jwa.AKP().String()).normalize()
}

const akpPrivateZKey = "z"

func akpKeyKind(algfn func() (jwa.KeyAlgorithm, bool)) KeyKind {
	_ = "STUB: not implemented"
	return *new(KeyKind)
}

func (k *akpPublicKey) KeyKind() KeyKind  { _ = "STUB: not implemented"; return *new(KeyKind) }
func (k *akpPrivateKey) KeyKind() KeyKind { _ = "STUB: not implemented"; return *new(KeyKind) }

func makeAKPPublicKey(src Key) (Key, error) { _ = "STUB: not implemented"; return *new(Key), nil }

func (k *akpPublicKey) PublicKey() (Key, error) { _ = "STUB: not implemented"; return *new(Key), nil }

func (k *akpPrivateKey) PublicKey() (Key, error) { _ = "STUB: not implemented"; return *new(Key), nil }

// akpThumbprint hashes the canonical JSON form defined by RFC 7638 §3.2
// for AKP keys: the required members {alg, kty, pub} in lexicographic order.
// RFC 9802 makes alg a required thumbprint input for AKP because pub is
// algorithm-scoped — omitting alg would break cross-implementation kid
// lookup.
func akpThumbprint(hash crypto.Hash, alg, pub string) []byte { _ = "STUB: not implemented"; return nil }

// Thumbprint returns the RFC 7638 thumbprint of this AKP key.
//
// AKP keys hash the canonical JSON form `{alg, kty, pub}` per RFC 9802 §7
// — different from the per-kty schemas RFC 7638 §3.2 defines for RSA, EC,
// and oct. Both `alg` and `pub` are required at thumbprint time; an
// error is returned if `alg` has not been set on the key. Other key
// types tolerate a missing `alg` because their canonical thumbprint
// input doesn't include it.
func (k *akpPublicKey) Thumbprint(hash crypto.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Thumbprint returns the RFC 7638 thumbprint of this AKP key.
//
// The thumbprint is computed over the public components only, so the
// returned value is identical to that of the corresponding
// [akpPublicKey]. AKP keys hash the canonical JSON form `{alg, kty, pub}`
// per RFC 9802 §7; both `alg` and `pub` are required at thumbprint time.
func (k *akpPrivateKey) Thumbprint(hash crypto.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *akpPublicKey) Validate() error { _ = "STUB: not implemented"; return nil }

func (k *akpPrivateKey) Validate() error { _ = "STUB: not implemented"; return nil }
