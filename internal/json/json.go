package json

import (
	"encoding/json/jsontext"
	"io"

	"sync/atomic"
)

var globalUseNumber atomic.Bool

// SetUseNumber controls whether JSON numbers in private/custom fields
// should be decoded as json.Number instead of float64.
func SetUseNumber(v bool) { _ = "STUB: not implemented"; return }

// GetUseNumber returns the current UseNumber setting.
func GetUseNumber() bool { _ = "STUB: not implemented"; return false }

type (
	Decoder    = jsontext.Decoder
	Encoder    = jsontext.Encoder
	RawMessage = jsontext.Value
)

func Engine() string { _ = "STUB: not implemented"; return "" }

func NewDecoder(r io.Reader) *jsontext.Decoder { _ = "STUB: not implemented"; return nil }

func NewEncoder(w io.Writer) *jsontext.Encoder { _ = "STUB: not implemented"; return nil }

func Marshal(v any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Unmarshal(b []byte, v any) error { _ = "STUB: not implemented"; return nil }

func MarshalEncode(enc *jsontext.Encoder, v any) error { _ = "STUB: not implemented"; return nil }

func UnmarshalDecode(dec *jsontext.Decoder, v any) error { _ = "STUB: not implemented"; return nil }

func AssignNextBytesToken(dst *[]byte, dec *Decoder) error { _ = "STUB: not implemented"; return nil }

func shouldRejectNullStrings(dc DecodeCtx) bool { _ = "STUB: not implemented"; return false }

// ReadNextStringToken reads the next JSON token from the decoder and
// returns it as a string. By default, JSON null is silently accepted as "".
// When the given DecodeCtx implements StrictStringDecodeCtx and StrictStrings()
// returns true, null values are rejected.
func ReadNextStringToken(dec *Decoder, dc DecodeCtx) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func AssignNextStringToken(dst **string, dec *Decoder, dc DecodeCtx) error {
	_ = "STUB: not implemented"
	return nil
}

// FlattenAudience is a flag to specify if we should flatten the "aud"
// entry to a string when there's only one entry.
// In jwx < 1.1.8 we just dumped everything as an array of strings,
// but apparently AWS Cognito doesn't handle this well.
//
// So now we have the ability to dump "aud" as a string if there's
// only one entry, but we need to retain the old behavior so that
// we don't accidentally break somebody else's code. (e.g. messing
// up how signatures are calculated)
var FlattenAudience uint32

func MarshalAudience(aud []string, flatten bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func EncodeAudience(enc *Encoder, aud []string, flatten bool) error {
	_ = "STUB: not implemented"
	return nil
}

// DecodeCtx is an interface for objects that needs that extra something
// when decoding JSON into an object.
type DecodeCtx interface {
	Registry() *Registry
}

// DecodeCtxContainer is used to differentiate objects that can carry extra
// decoding hints and those who can't.
type DecodeCtxContainer interface {
	DecodeCtx() DecodeCtx
	SetDecodeCtx(DecodeCtx)
}

// StrictStringDecodeCtx is an optional interface that DecodeCtx implementations
// can satisfy to control per-call null string rejection.
type StrictStringDecodeCtx interface {
	StrictStrings() bool
}

// stock decodeCtx. should cover 80% of the cases
type decodeCtx struct {
	registry      *Registry
	strictStrings bool
}

// NewDecodeCtx creates a new DecodeCtx with the given registry.
func NewDecodeCtx(r *Registry) DecodeCtx { _ = "STUB: not implemented"; return *new(DecodeCtx) }

// NewDecodeCtxStrictStrings creates a new DecodeCtx with the given registry
// and strict string rejection flag.
func NewDecodeCtxStrictStrings(r *Registry, strict bool) DecodeCtx {
	_ = "STUB: not implemented"
	return *new(DecodeCtx)
}

func (dc *decodeCtx) Registry() *Registry { _ = "STUB: not implemented"; return nil }

func (dc *decodeCtx) StrictStrings() bool { _ = "STUB: not implemented"; return false }
