package content_crypt //nolint:golint

import (
	"sync"

	"github.com/lestrrat-go/jwx/v4/jwa"
)

var genericCache sync.Map // map[string]*Generic

func (c Generic) Algorithm() jwa.ContentEncryptionAlgorithm {
	_ = "STUB: not implemented"
	return *new(jwa.ContentEncryptionAlgorithm)
}

func (c Generic) Encrypt(cek, plaintext, aad []byte) ([]byte, []byte, []byte, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

func (c Generic) Decrypt(cek, iv, ciphertext, tag, aad []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewGeneric(alg jwa.ContentEncryptionAlgorithm) (*Generic, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:forcetypeassert

func (c Generic) KeySize() int { _ = "STUB: not implemented"; return 0 }
