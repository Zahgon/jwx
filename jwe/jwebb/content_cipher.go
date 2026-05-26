package jwebb

import (
	"sync"

	"github.com/lestrrat-go/jwx/v4/jwe/internal/content_crypt"
)

var contentCipherCache sync.Map // map[string]content_crypt.Cipher

// ContentEncryptionIsSupported checks if the content encryption algorithm is supported
func ContentEncryptionIsSupported(alg string) bool { _ = "STUB: not implemented"; return false }

// CreateContentCipher creates a content encryption cipher for the given algorithm string
func CreateContentCipher(alg string) (content_crypt.Cipher, error) {
	_ = "STUB: not implemented"
	return *new(content_crypt.Cipher), nil
}

//nolint:forcetypeassert
