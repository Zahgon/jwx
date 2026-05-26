package keyconv

import (
	"crypto/ecdh"
	"crypto/ecdsa"
	"crypto/ed25519"
	"crypto/elliptic"
	"crypto/rsa"
)

// KeyAs converts src to type T. src may be a jwk.Key or a raw crypto key.
// If src is a jwk.Key, it is exported via jwk.Export[T]. Otherwise, a direct
// type assertion is attempted, with a reflect fallback for value→pointer
// conversion (e.g. rsa.PrivateKey when T = *rsa.PrivateKey).
func KeyAs[T any](src any) (T, error) { _ = "STUB: not implemented"; return *new(T), nil }

// value → pointer: e.g. rsa.PrivateKey when T = *rsa.PrivateKey

// RSAPublicKey extracts an *rsa.PublicKey from src.
// src may be rsa.PublicKey, *rsa.PublicKey, rsa.PrivateKey, *rsa.PrivateKey, or jwk.Key.
func RSAPublicKey(src any) (*rsa.PublicKey, error) { _ = "STUB: not implemented"; return nil, nil }

// ECDSAPublicKey extracts an *ecdsa.PublicKey from src.
// src may be ecdsa.PublicKey, *ecdsa.PublicKey, ecdsa.PrivateKey, *ecdsa.PrivateKey, or jwk.Key.
func ECDSAPublicKey(src any) (*ecdsa.PublicKey, error) { _ = "STUB: not implemented"; return nil, nil }

func Ed25519PrivateKey(src any) (*ed25519.PrivateKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Export may return ed25519.PrivateKey (not pointer)

func Ed25519PublicKey(src any) (*ed25519.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type privECDHer interface {
	ECDH() (*ecdh.PrivateKey, error)
}

// ECDHPrivateKey extracts an *ecdh.PrivateKey from src.
// In addition to jwk.Key and direct type matches, it also handles
// types that implement the ECDH() method (e.g. *ecdsa.PrivateKey).
func ECDHPrivateKey(src any) (*ecdh.PrivateKey, error) { _ = "STUB: not implemented"; return nil, nil }

type pubECDHer interface {
	ECDH() (*ecdh.PublicKey, error)
}

// ECDHPublicKey extracts an *ecdh.PublicKey from src.
// In addition to jwk.Key and direct type matches, it also handles
// types that implement the ECDH() method (e.g. *ecdsa.PublicKey).
func ECDHPublicKey(src any) (*ecdh.PublicKey, error) { _ = "STUB: not implemented"; return nil, nil }

// ecdhCurveToElliptic maps ECDH curves to elliptic curves
func ecdhCurveToElliptic(ecdhCurve ecdh.Curve) (elliptic.Curve, error) {
	_ = "STUB: not implemented"
	return *new(elliptic.Curve), nil
}

// ecdhPublicKeyToECDSA converts an ECDH public key to an ECDSA public key
func ecdhPublicKeyToECDSA(ecdhPubKey *ecdh.PublicKey) (*ecdsa.PublicKey, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Parse the uncompressed point format (0x04 prefix + X + Y coordinates)

// ECDHToECDSA converts an ECDH key to an ECDSA key.
// Returns *ecdsa.PublicKey for public keys, *ecdsa.PrivateKey for private keys.
func ECDHToECDSA(src any) (any, error) {
	_ = "STUB: not implemented"
	// First, handle value types by converting to pointers
	return *new(any), nil
}

// convert the public key
