package cipher

import (
	"crypto/cipher"
)

var gcm = &gcmFetcher{}
var cbc = &cbcFetcher{}

func (f gcmFetcher) Fetch(key []byte, size int) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

func (f cbcFetcher) Fetch(key []byte, size int) (cipher.AEAD, error) {
	_ = "STUB: not implemented"
	return *new(cipher.AEAD), nil
}

func (c AesContentCipher) KeySize() int { _ = "STUB: not implemented"; return 0 }

func (c AesContentCipher) TagSize() int { _ = "STUB: not implemented"; return 0 }

func NewAES(alg string) (*AesContentCipher, error) { _ = "STUB: not implemented"; return nil, nil }

func (c AesContentCipher) Encrypt(cek, plaintext, aad []byte) (iv, ciphertxt, tag []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil, nil
}

// CBC+HMAC's Seal may panic (buffer size limits, auth tag errors),
// so we must recover. GCM's Seal from the stdlib does not panic
// with valid inputs, but we protect uniformly for safety.

func (c AesContentCipher) Decrypt(cek, iv, ciphertxt, tag, aad []byte) (plaintext []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CBC+HMAC's Open may panic (buffer size limits), so we must recover.
