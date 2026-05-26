package jwe

import (
	"github.com/lestrrat-go/jwx/v4/internal/json"
)

// NewRecipient creates a Recipient object
func NewRecipient() Recipient { _ = "STUB: not implemented"; return *new(Recipient) }

func (r *stdRecipient) SetHeaders(h Headers) error { _ = "STUB: not implemented"; return nil }

func (r *stdRecipient) SetEncryptedKey(v []byte) error { _ = "STUB: not implemented"; return nil }

func (r *stdRecipient) Headers() Headers { _ = "STUB: not implemented"; return *new(Headers) }

func (r *stdRecipient) EncryptedKey() []byte { _ = "STUB: not implemented"; return nil }

type recipientMarshalProxy struct {
	Headers      Headers `json:"header"`
	EncryptedKey string  `json:"encrypted_key"`
}

func (r *stdRecipient) UnmarshalJSON(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (r *stdRecipient) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// NewMessage creates a new message
func NewMessage() *Message { _ = "STUB: not implemented"; return nil }

func (m *Message) AuthenticatedData() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) CipherText() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) InitializationVector() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) Tag() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) ProtectedHeaders() Headers { _ = "STUB: not implemented"; return *new(Headers) }

func (m *Message) Recipients() []Recipient { _ = "STUB: not implemented"; return nil }

func (m *Message) UnprotectedHeaders() Headers { _ = "STUB: not implemented"; return *new(Headers) }

const (
	AuthenticatedDataKey    = "aad"
	CipherTextKey           = "ciphertext"
	CountKey                = "p2c"
	InitializationVectorKey = "iv"
	ProtectedHeadersKey     = "protected"
	RecipientsKey           = "recipients"
	SaltKey                 = "p2s"
	TagKey                  = "tag"
	UnprotectedHeadersKey   = "unprotected"
	HeadersKey              = "header"
	EncryptedKeyKey         = "encrypted_key"
)

func (m *Message) Set(k string, v any) error { _ = "STUB: not implemented"; return nil }

type messageMarshalProxy struct {
	AuthenticatedData    string            `json:"aad,omitempty"`
	CipherText           string            `json:"ciphertext"`
	InitializationVector string            `json:"iv,omitempty"`
	ProtectedHeaders     json.RawMessage   `json:"protected"`
	Recipients           []json.RawMessage `json:"recipients,omitempty"`
	Tag                  string            `json:"tag,omitempty"`
	UnprotectedHeaders   Headers           `json:"unprotected,omitempty"`

	// For flattened structure. Headers is NOT a Headers type,
	// so that we can detect its presence by checking proxy.Headers != nil
	Headers      json.RawMessage `json:"header,omitempty"`
	EncryptedKey string          `json:"encrypted_key,omitempty"`
}

type jsonKV struct {
	Key   string
	Value string
}

func marshalField(v any) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (m *Message) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// This is slightly convoluted, but we need to encode the
	// protected headers, so we do it by hand
	return nil, nil
}

// '{}'

// Use flattened format

func (m *Message) UnmarshalJSON(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Get the string value

// It's now in _quoted_ base64 string. Decode it

// if this were a flattened message, we would see a "header" and "ciphertext"
// field. TODO: do both of these conditions need to meet, or just one?

// `"heders"` could be empty. If that's the case, just skip the
// following unmarshaling step

// RFC 7516 §7.2: "ciphertext", "iv", and "tag" MUST be present and
// non-empty for any AEAD-protected JWE. Reject missing/empty values
// here so that a zero-length authentication tag cannot reach the
// AEAD verification code path.

// this is later used for decryption

func (m *Message) makeDummyRecipient(enckeybuf string, protected Headers) error {
	_ = "STUB: not implemented"
	// Recipients in this case should not contain the content encryption key,
	// so move that out
	return nil
}

// Compact generates a JWE message in compact serialization format from a
// `*jwe.Message` object. The object contain exactly one recipient, or
// an error is returned.
//
// This function currently does not take any options, but the function
// signature contains `options` for possible future expansion of the API
func Compact(m *Message, _ ...CompactOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// The protected header must be a merge between the message-wide
// protected header AND the recipient header

// There's something wrong if m.protectedHeaders is nil, but
// it could happen
