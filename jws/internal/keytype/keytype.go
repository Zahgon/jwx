package keytype

// Because the keys defined in github.com/lestrrat-go/jwx/jwk may also implement
// crypto.Signer, it would be possible for to mix up key types when signing/verifying
// for example, when we specify jws.WithKey(jwa.RSA256, cryptoSigner), the cryptoSigner
// can be for RSA, or any other type that implements crypto.Signer... even if it's for the
// wrong algorithm.
//
// These functions are there to differentiate between the valid KNOWN key types.
// For any other key type that is outside of the Go std library and our own code,
// we must rely on the user to be vigilant.
//
// Notes: symmetric keys are obviously not part of this. for v2 OKP keys,
// x25519 does not implement Sign()
func IsValidRSAKey(key any) bool { _ = "STUB: not implemented"; return false }

// these are NOT ok

func IsValidECDSAKey(key any) bool { _ = "STUB: not implemented"; return false }

// these are NOT ok

func IsValidEDDSAKey(key any) bool { _ = "STUB: not implemented"; return false }

// these are NOT ok
