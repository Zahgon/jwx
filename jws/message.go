package jws

import (
	"github.com/lestrrat-go/jwx/v4/internal/json"
)

func NewSignature() *Signature { _ = "STUB: not implemented"; return nil }

func (s *Signature) DecodeCtx() DecodeCtx { _ = "STUB: not implemented"; return *new(DecodeCtx) }

func (s *Signature) SetDecodeCtx(dc DecodeCtx) { _ = "STUB: not implemented"; return }

func (s Signature) PublicHeaders() Headers { _ = "STUB: not implemented"; return *new(Headers) }

func (s *Signature) SetPublicHeaders(v Headers) *Signature { _ = "STUB: not implemented"; return nil }

func (s Signature) ProtectedHeaders() Headers { _ = "STUB: not implemented"; return *new(Headers) }

func (s *Signature) SetProtectedHeaders(v Headers) *Signature {
	_ = "STUB: not implemented"
	return nil
}

func (s Signature) Signature() []byte { _ = "STUB: not implemented"; return nil }

func (s *Signature) SetSignature(v []byte) *Signature { _ = "STUB: not implemented"; return nil }

type signatureUnmarshalProbe struct {
	Header    Headers `json:"header,omitempty"`
	Protected *string `json:"protected,omitempty"`
	Signature *string `json:"signature,omitempty"`
}

func (s *Signature) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// RFC 7515 §3 mandates that "protected" be base64url-encoded.
// Earlier code carried a relaxed probe that accepted a literal-
// JSON form (a JSON string whose content begins with "{") and
// skipped base64 decoding — that was asymmetric with the
// flattened branch (which only base64-decodes) and gave callers
// a non-conforming wire form useful for evading byte-exact JWS
// dedup / replay caches.

//nolint:forcetypeassert

//nolint:forcetypeassert

func NewMessage() *Message {
	_ = "STUB: not implemented"

	// Clears the internal raw buffer that was accumulated during
	// the verify phase
	return nil
}

func (m *Message) clearRaw() { _ = "STUB: not implemented"; return }

func (m *Message) SetDecodeCtx(dc DecodeCtx) { _ = "STUB: not implemented"; return }

func (m *Message) DecodeCtx() DecodeCtx {
	_ = "STUB: not implemented"

	// Payload returns the decoded payload
	return *new(DecodeCtx)
}

func (m Message) Payload() []byte { _ = "STUB: not implemented"; return nil }

func (m *Message) SetPayload(v []byte) *Message { _ = "STUB: not implemented"; return nil }

func (m Message) Signatures() []*Signature { _ = "STUB: not implemented"; return nil }

func (m *Message) AppendSignature(v *Signature) *Message { _ = "STUB: not implemented"; return nil }

func (m *Message) ClearSignatures() *Message { _ = "STUB: not implemented"; return nil }

// LookupSignature looks up a particular signature entry using
// the `kid` value
func (m Message) LookupSignature(kid string) []*Signature { _ = "STUB: not implemented"; return nil }

// This struct is used to first probe for the structure of the
// incoming JSON object. We then decide how to parse it
// from the fields that are populated.
type messageUnmarshalProbe struct {
	Payload    *string           `json:"payload"`
	Signatures []json.RawMessage `json:"signatures,omitempty"`
	Header     json.RawMessage   `json:"header,omitempty"`
	Protected  *string           `json:"protected,omitempty"`
	Signature  *string           `json:"signature,omitempty"`
}

func (m *Message) UnmarshalJSON(buf []byte) error { _ = "STUB: not implemented"; return nil }

// Enforce the signature cap before we decode any signature entry.
// The probe above leaves mup.Signatures as []json.RawMessage, so no
// headers/base64 work has happened yet. Doing the check here prevents
// a large signatures array from allocating O(input) work before Parse
// gets a chance to reject it.

// flattened signature is NOT present

// RFC 7515 §7.2.1 places the unprotected JOSE header inside each
// signature entry for the general (multi-signature) form; a
// top-level "header" sibling of "signatures" is not defined.
// Reject rather than silently drop — silent drop hid both typos
// and an attacker-controlled trigger surface for any
// RegisterCustomDecoder side effects on the dropped contents.

// Instead of barfing on a nil protected header, use an empty header

// .signature is present, it's a flattened structure

//nolint:forcetypeassert

//nolint:forcetypeassert

// Instead of barfing on a nil protected header, use an empty header

// NOT base64 encoded

func (m Message) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m Message) marshalFlattened() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// RFC 7797 b64=false: emit the raw payload as a JSON string
// rather than re-base64-encoding it. json.Marshal handles
// any necessary escaping for byte sequences that aren't
// JSON-safe as-is.

func (m Message) marshalFull() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// RFC 7797 b64=false: emit the raw payload as a JSON string
// rather than re-base64-encoding it. The general JWS form has
// one shared payload across signatures; per RFC 7797, all
// signers must agree on the b64 flag, so we consult the first
// signature's protected header.

// If InsecureNoSignature is enabled, signature may not exist

// Compact generates a JWS message in compact serialization format from
// `*jws.Message` object. The object contain exactly one signature, or
// an error is returned.
//
// If using a detached payload, the payload must already be stored in
// the `*jws.Message` object, and the `jws.WithDetached()` option
// must be passed to the function.
func Compact(msg *Message, options ...CompactOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// XXX check if this is correct
