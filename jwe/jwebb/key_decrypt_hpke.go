package jwebb

// KeyDecryptHPKEKE performs HPKE key decryption per
// draft-ietf-jose-hpke-encrypt-16 (https://datatracker.ietf.org/doc/draft-ietf-jose-hpke-encrypt/16/).
// encryptedCEK is the HPKE-sealed CEK from the JWE Encrypted Key field.
// ek is the HPKE encapsulated key from the "ek" header field.
func KeyDecryptHPKEKE(encryptedCEK []byte, alg, calg string, privkey any, ek []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	// Try custom HPKE key decrypter (e.g., X448 from external modules)
	return nil, nil
}

// AAD is empty for key encryption mode per spec
