package jwebb

// AES key wrap decryption functions

// Use constants from tokens package
// No need to redefine them here

func KeyDecryptAESKW(_, enckey []byte, _ string, sharedkey []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyDecryptDirect(_, _ []byte, _ string, cek []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func KeyDecryptPBES2(_, enckey []byte, alg string, password []byte, salt []byte, count int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Derive key using PBKDF2

// Use the derived key for AES key wrap

func KeyDecryptAESGCMKW(recipientKey, _ []byte, _ string, sharedkey []byte, iv []byte, tag []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Combine recipient key and tag for GCM decryption. Allocate a fresh
// buffer so we never alias into recipientKey's backing array, which
// is owned by the parsed message.
