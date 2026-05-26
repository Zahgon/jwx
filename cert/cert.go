package cert

import (
	"crypto/x509"
	"io"
)

// Create is a wrapper around x509.CreateCertificate, but it additionally
// encodes it in base64 so that it can be easily added to `x5c` fields
func Create(rand io.Reader, template, parent *x509.Certificate, pub, priv any) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// EncodeBase64 is a utility function to encode ASN.1 DER certificates
// using base64 encoding. This operation is normally done by `pem.Encode`
// but since PEM would include the markers (`-----BEGIN`, and the like)
// while `x5c` fields do not need this, this function can be used to
// shave off a few lines
func EncodeBase64(der []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Parse decodes a base64-encoded ASN.1 DER certificate and validates that it
// parses as X.509.
//
// The certificate must be in PKIX format and it must not contain PEM markers.
// The maximum decoded certificate size is controlled by `cert.Settings()`.
func Parse(src []byte) (*x509.Certificate, error) { _ = "STUB: not implemented"; return nil, nil }

func validateDERCertificate(der []byte) (*x509.Certificate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func validateEncodedCertificateSize(src []byte) error { _ = "STUB: not implemented"; return nil }

func decodedCertificateSize(src []byte) int { _ = "STUB: not implemented"; return 0 }

func normalizeAndValidateChainCertificate(src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
