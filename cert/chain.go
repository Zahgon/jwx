package cert

// Chain represents a certificate chain as used in the `x5c` field of
// various objects within JOSE.
//
// It stores the certificates as a list of base64-encoded byte sequences. Every
// certificate added to or decoded into the chain must parse as X.509 and is
// subject to the global limits configured by `cert.Settings()`.
type Chain struct {
	certificates [][]byte
}

func (cc Chain) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// UnmarshalJSON decodes an `x5c` JSON array and validates each entry as a
// base64-encoded X.509 certificate.
func (cc *Chain) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// Get returns the n-th ASN.1 DER + base64 encoded certificate
// stored. `false` will be returned in the second argument if
// the corresponding index is out of range.
func (cc *Chain) Get(index int) ([]byte, bool) { _ = "STUB: not implemented"; return nil, false }

// Len returns the number of certificates stored in this Chain
func (cc *Chain) Len() int { _ = "STUB: not implemented"; return 0 }

func (cc *Chain) AddString(der string) error { _ = "STUB: not implemented"; return nil }

// Add appends a certificate to the chain.
//
// Input may be either a PEM `CERTIFICATE` block or a base64-encoded DER value
// as stored in JOSE `x5c` fields. The certificate is validated as X.509 and is
// subject to the global limits configured by `cert.Settings()`.
func (cc *Chain) Add(der []byte) error { _ = "STUB: not implemented"; return nil }

// Accept a PEM-encoded CERTIFICATE block and convert it to the
// base64(DER) form that x5c requires.

// Non-PEM input must be base64(DER). Strip any internal whitespace
// (callers commonly pass multi-line base64 literals) and validate.

func stripASCIIWhitespace(src []byte) []byte { _ = "STUB: not implemented"; return nil }
