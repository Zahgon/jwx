package jwk

import (
	"crypto"
	"reflect"

	"github.com/lestrrat-go/jwx/v4/jwa"
)

func init() {
	panicOnRegistrationError(RegisterKeyExporter(KeyKind(jwa.OctetSeq().String()), KeyExportFunc(octetSeqToRaw)))
}

func (k *symmetricKey) Import(rawKey []byte) error { _ = "STUB: not implemented"; return nil }

var symmetricConvertibleKeys = []reflect.Type{
	reflect.TypeFor[SymmetricKey](),
}

func octetSeqToRaw(keyif Key, _ any) (any, error) {
	_ = "STUB: not implemented"
	// Fast path: built-in concrete types need no reflection
	return *new(any), nil
}

// already a concrete type, skip extractEmbeddedKey

//nolint:forcetypeassert // rlocker is unexported; only our concrete types implement it

// External implementation — use self-locking interface getters.

// Thumbprint returns the JWK thumbprint using the indicated
// hashing algorithm, according to RFC 7638
func (k *symmetricKey) Thumbprint(hash crypto.Hash) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (k *symmetricKey) PublicKey() (Key, error) { _ = "STUB: not implemented"; return *new(Key), nil }

func (k *symmetricKey) Validate() error { _ = "STUB: not implemented"; return nil }
