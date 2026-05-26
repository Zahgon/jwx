package json

import (
	stdjson "encoding/json"
	"encoding/json/jsontext"
	jsonv2 "encoding/json/v2"
	"reflect"
	"sync"
)

// customDecoder is the internal interface for field decoders stored in the registry.
// It returns any because different fields decode to different types.
type customDecoder interface {
	Decode([]byte) (any, error)
}

// CustomDecoder is the public generic interface for custom field decoders.
type CustomDecoder[T any] interface {
	Decode([]byte) (T, error)
}

// CustomDecodeFunc is a function-based implementation of CustomDecoder[T].
type CustomDecodeFunc[T any] func([]byte) (T, error)

func (fn CustomDecodeFunc[T]) Decode(data []byte) (T, error) {
	_ = "STUB: not implemented"

	// customDecoderAdapter wraps a CustomDecoder[T] to satisfy the internal customDecoder interface.
	return *new(T), nil
}

type customDecoderAdapter[T any] struct {
	dec CustomDecoder[T]
}

func (a *customDecoderAdapter[T]) Decode(data []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// objectTypeDecoder is a reflect-based decoder used by the untyped Register path.
type objectTypeDecoder struct {
	typ  reflect.Type
	name string
}

func (dec *objectTypeDecoder) Decode(data []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// TypedDecoder is a generic decoder that unmarshals JSON into a concrete type T,
// eliminating the need for reflect.New.
type TypedDecoder[T any] struct {
	name string
}

func (dec *TypedDecoder[T]) Decode(data []byte) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// useNumberUnmarshalers is a pre-built json/v2 option that intercepts
// unmarshalling of JSON numbers into any, producing json.Number instead
// of float64.
var useNumberUnmarshalers = jsonv2.WithUnmarshalers(
	jsonv2.UnmarshalFromFunc(func(dec *jsontext.Decoder, val *any) error {
		if dec.PeekKind() != '0' {
			return jsonv2.SkipFunc
		}
		raw, err := dec.ReadValue()
		if err != nil {
			return err
		}
		*val = stdjson.Number(raw.String())
		return nil
	}),
)

type Registry struct {
	mu   *sync.RWMutex
	ctrs map[string]customDecoder
}

func NewRegistry() *Registry { _ = "STUB: not implemented"; return nil }

// RegisterTyped registers a generic TypedDecoder[T] for the given field name.
func RegisterTyped[T any](r *Registry, name string) { _ = "STUB: not implemented"; return }

// RegisterCustomDecoder registers a CustomDecoder[T] for the given field name.
func RegisterCustomDecoder[T any](r *Registry, name string, dec CustomDecoder[T]) {
	_ = "STUB: not implemented"
	return
}

// Register registers a decoder for the given field name using the untyped
// dispatch path. If object is nil, the registration is removed.
// If object implements customDecoder, it is used directly.
// Otherwise, an objectTypeDecoder is created using reflect.
//
// This is used internally by WithTypedField for per-parse local registries.
// New code should prefer RegisterTyped or RegisterCustomDecoder.
func (r *Registry) Register(name string, object any) { _ = "STUB: not implemented"; return }

// Unregister removes the decoder for the given field name.
func (r *Registry) Unregister(name string) { _ = "STUB: not implemented"; return }

// Decode decodes the raw JSON value using the registered decoder for the
// given field name. If no decoder is registered, the raw value is decoded
// into any.
func (r *Registry) Decode(name string, raw RawMessage) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
