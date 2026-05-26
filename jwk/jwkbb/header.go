package jwkbb

import (
	"github.com/valyala/fastjson"
)

type headerNotFoundError struct {
	key string
}

func (e headerNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e headerNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// ErrHeaderNotFound returns an error that can be passed to errors.Is
// to check if the error is the result of a field not being found in
// the parsed JSON.
func ErrHeaderNotFound() error { _ = "STUB: not implemented"; return nil }

// Header is an opaque handle to a parsed JWK or JWKS JSON object.
// It exists for fast, allocation-light field probing without paying
// the cost of a full encoding/json or jsonv2 unmarshal.
//
// Header instances are NOT safe for concurrent use. Create a new one
// per goroutine. Values returned by HeaderGet* helpers may alias
// memory owned by the Header; do not retain them past the Header's
// lifetime unless the helper explicitly copies (HeaderGetString does;
// HeaderGetStringBytes does not).
//
// This type is experimental and may change or be removed in the future.
type Header interface {
	// Sealed so callers can't depend on the underlying fastjson type
	// or substitute their own implementation.
	jwkbbHeader()
}

type header struct {
	v   *fastjson.Value
	err error
}

func (h *header) jwkbbHeader() {
	_ = "STUB: not implemented"

	// HeaderParse parses a JSON byte slice and returns a Header for fast
	// field access. Parse errors are deferred to the first HeaderGet* /
	// HeaderHas call.
	//
	// This function is experimental and may change or be removed in the future.
	return
}

func HeaderParse(buf []byte) Header { _ = "STUB: not implemented"; return *new(Header) }

func headerGet(h Header, key string) (*fastjson.Value, error) {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return nil, nil
}

// we _know_ this can't be another type

// HeaderHas reports whether the given key exists in the parsed JSON object.
// Returns false on parse errors.
//
// This function is experimental and may change or be removed in the future.
func HeaderHas(h Header, key string) bool { _ = "STUB: not implemented"; return false }

// HeaderGetString returns the string value for the given key as a
// freshly-allocated Go string. The returned value remains valid after
// the Header is garbage collected.
//
// This function is experimental and may change or be removed in the future.
func HeaderGetString(h Header, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// HeaderGetStringBytes returns the JSON string bytes for the given key
// without copying.
//
// WARNING: the returned slice aliases memory owned by h. It becomes
// invalid as soon as h is reused, re-parsed, or goes out of scope and
// is garbage collected. Do not retain the slice, share it across
// goroutines, or use it after any further call on h. If you need a
// value that outlives h, use [HeaderGetString].
//
// This function is experimental and may change or be removed in the future.
func HeaderGetStringBytes(h Header, key string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
