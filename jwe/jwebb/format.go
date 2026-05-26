package jwebb

import (
	"github.com/lestrrat-go/jwx/v4/internal/base64"
)

// JoinCompact builds JWE compact serialization directly from raw parts.
// The protected header must already be base64url-encoded. The remaining
// parts (encryptedKey, iv, ciphertext, tag) are raw bytes that will be
// base64url-encoded into a single pre-sized output buffer.
//
// The result format is: base64(protected).base64(encryptedKey).base64(iv).base64(ciphertext).base64(tag)
func JoinCompact(encoder base64.Encoder, protected, encryptedKey, iv, ciphertext, tag []byte) []byte {
	_ = "STUB: not implemented"
	return nil
}
