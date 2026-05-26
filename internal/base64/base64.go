package base64

import (
	"encoding/base64"
	"io"
	"sync/atomic"
)

type Decoder interface {
	Decode([]byte) ([]byte, error)
}

type Encoder interface {
	Encode([]byte, []byte)
	EncodedLen(int) int
	EncodeToString([]byte) string
	AppendEncode([]byte, []byte) []byte
}

// StreamEncoder is an [Encoder] that can also produce an incremental
// [io.WriteCloser] for encoding a byte stream directly into a downstream
// writer. This is the shape the jws streaming detached-payload path needs
// to avoid materializing the payload in memory.
//
// The stdlib *[encoding/base64.Encoding] satisfies this interface
// automatically, so the default jwx encoder does. Extension modules
// providing custom encoders should implement [io.WriteCloser]-returning
// NewEncoder if they want their encoder honored by the streaming path.
type StreamEncoder interface {
	Encoder
	// NewEncoder returns a new [io.WriteCloser] that encodes bytes
	// written to it and forwards the encoded output to w. Close must
	// be called to flush any partial final block.
	NewEncoder(w io.Writer) io.WriteCloser
}

// AsStreamEncoder reports whether e can be used as a [StreamEncoder]
// and returns the stream-capable view. Callers should error out when
// the second return value is false rather than silently falling back to
// a different encoder, to avoid mixing encodings within a single
// signing operation.
//
// The stdlib [*encoding/base64.Encoding] is supported as a special case
// (its streaming form is a top-level function rather than a method, so
// it does not directly satisfy the interface).
func AsStreamEncoder(e Encoder) (StreamEncoder, bool) {
	_ = "STUB: not implemented"
	return *new(StreamEncoder), false
}

// stdStreamEncoder wraps the stdlib [*base64.Encoding] so it satisfies
// [StreamEncoder]. It is used for the default encoder and as the
// fallback in [AsStreamEncoder] when a caller passes a raw
// [*base64.Encoding].
type stdStreamEncoder struct {
	*base64.Encoding
}

func (e stdStreamEncoder) NewEncoder(w io.Writer) io.WriteCloser {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser)
}

// encoderHolder and decoderHolder are fixed concrete types so that
// atomic.Value.Store never sees a type change (which would panic).
type encoderHolder struct{ enc Encoder }
type decoderHolder struct{ dec Decoder }

var atomicEncoder atomic.Value
var atomicDecoder atomic.Value

func init() {
	atomicEncoder.Store(encoderHolder{base64.RawURLEncoding})
	atomicDecoder.Store(decoderHolder{defaultDecoder{}})
}

func SetEncoder(enc Encoder) { _ = "STUB: not implemented"; return }

func getEncoder() Encoder {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return *new(Encoder)
}

func DefaultEncoder() Encoder { _ = "STUB: not implemented"; return *new(Encoder) }

func SetDecoder(dec Decoder) { _ = "STUB: not implemented"; return }

func getDecoder() Decoder {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return *new(Decoder)
}

func Encode(src []byte) []byte { _ = "STUB: not implemented"; return nil }

func AppendEncode(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func EncodedLen(n int) int { _ = "STUB: not implemented"; return 0 }

func EncodeToString(src []byte) string { _ = "STUB: not implemented"; return "" }

func EncodeUint64ToString(v uint64) string { _ = "STUB: not implemented"; return "" }

const (
	InvalidEncoding = iota
	Std
	URL
	RawStd
	RawURL
)

func Guess(src []byte) int { _ = "STUB: not implemented"; return 0 }

// defaultDecoder is a Decoder that detects the encoding of the source and
// decodes it accordingly. This shouldn't really be required per the spec, but
// it exist because we have seen in the wild JWTs that are encoded using
// various versions of the base64 encoding.
type defaultDecoder struct{}

func (defaultDecoder) Decode(src []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func Decode(src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func DecodeString(src string) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// DecodeStrict decodes base64url-encoded data (RFC 7515 / RFC 4648 §5, no padding)
// directly using base64.RawURLEncoding. It writes into the provided dst buffer
// and returns the number of bytes written.
//
// Unlike Decode, this function does not auto-detect the encoding variant,
// does not acquire any mutex, and does not allocate. The caller must ensure
// dst is large enough (use DecodedStrictLen).
func DecodeStrict(dst, src []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// DecodedStrictLen returns the maximum decoded length for a base64url-encoded
// input of length n (no padding).
func DecodedStrictLen(n int) int { _ = "STUB: not implemented"; return 0 }
