package jwebb

import (
	"github.com/lestrrat-go/jwx/v4/jwe/internal/keygen"
)

// KeyEncryptAESKW encrypts the CEK using AES key wrap
func KeyEncryptAESKW(cek []byte, _ string, sharedkey []byte) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}

// KeyEncryptDirect returns the CEK directly for DIRECT algorithm
func KeyEncryptDirect(_ []byte, _ string, sharedkey []byte) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}

// KeyEncryptPBES2 encrypts the CEK using PBES2 password-based encryption.
// count is the PBKDF2 iteration count. If count <= 0, the OWASP 2023
// per-variant default is used as a safety fallback; public callers go
// through jwe.Encrypt / jwe.Settings and can override via the
// WithPBES2Count option.
func KeyEncryptPBES2(cek []byte, alg string, password []byte, count int) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}

// Derive key using PBKDF2

// Use the derived key for AES key wrap

// KeyEncryptAESGCMKW encrypts the CEK using AES GCM key wrap
func KeyEncryptAESGCMKW(cek []byte, _ string, sharedkey []byte) (keygen.ByteSource, error) {
	_ = "STUB: not implemented"
	return *new(keygen.ByteSource), nil
}
