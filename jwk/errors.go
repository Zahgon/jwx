package jwk

import (
	"errors"
	"reflect"
)

var cpe = &continueError{}

// ContinueError returns an opaque error that can be returned
// when a `KeyParser`, `KeyImporter`, or `KeyExporter` cannot handle the given payload,
// but would like the process to continue with the next handler.
func ContinueError() error { _ = "STUB: not implemented"; return nil }

type continueError struct{}

func (e *continueError) Error() string { _ = "STUB: not implemented"; return "" }

type importError struct {
	error
}

func (e importError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (importError) Is(err error) bool { _ = "STUB: not implemented"; return false }

func importerr(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

var errDefaultImportError = importError{errors.New(`import error`)}

func ImportError() error { _ = "STUB: not implemented"; return nil }

type parseError struct {
	error
}

func (e parseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (parseError) Is(err error) bool { _ = "STUB: not implemented"; return false }

func bparseerr(prefix string, f string, args ...any) error { _ = "STUB: not implemented"; return nil }

func parseerr(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

func rparseerr(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

func sparseerr(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

func kparseerr(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

func kasparseerr(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

var errDefaultParseError = parseError{errors.New(`parse error`)}

func ParseError() error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// KeyTypeMismatchError
//-------------------------------------------------------------------

// KeyTypeMismatchError is returned by [Import] / [ParseKeyAs] /
// [Export] / [ExportAll] when the value the function produced does
// not match the generic type parameter supplied by the caller.
//
// Got is the runtime type of the value the library produced; Want is
// the type the caller asked for via the type parameter. Callers that
// need to distinguish "wrong generic type parameter" from other
// failures should use [errors.Is] with KeyTypeMismatchError{}, or
// [errors.AsType] to recover the Got and Want fields.
type KeyTypeMismatchError struct {
	// Got is the runtime type of the value the library produced.
	Got reflect.Type
	// Want is the type requested via the function's type parameter.
	Want reflect.Type
}

func (e KeyTypeMismatchError) Error() string { _ = "STUB: not implemented"; return "" }

func (e KeyTypeMismatchError) Is(target error) bool { _ = "STUB: not implemented"; return false }

// typeName renders a reflect.Type similarly to the %T verb so that
// KeyTypeMismatchError's message remains recognizable when either
// field is nil.
func typeName(t reflect.Type) string { _ = "STUB: not implemented"; return "" }

//-------------------------------------------------------------------
// UnknownKeyTypeError
//-------------------------------------------------------------------

// UnknownKeyTypeError is returned by [Parse] / [ParseKey] / [ParseKeyAs]
// when the input's "kty" hint cannot be resolved to a known key
// family.
//
// KeyType is empty when the input had no "kty" field at all, or when
// "kty" was present but not a JSON string (the probe could not extract
// a usable identifier). KeyType is populated when the input carried a
// string "kty" that didn't match any registered key family — useful
// for callers that want to suggest installing an extension module.
//
// Use [errors.Is] with UnknownKeyTypeError{} to recognize the
// condition, or [errors.AsType] to recover the KeyType field. The
// error chain also satisfies [errors.Is] with [ParseError].
type UnknownKeyTypeError struct {
	// KeyType is the raw "kty" value the input carried, or "" when
	// "kty" was missing or non-string.
	KeyType string
}

func (e UnknownKeyTypeError) Error() string { _ = "STUB: not implemented"; return "" }

func (UnknownKeyTypeError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// FieldNotFoundError
//-------------------------------------------------------------------

// FieldNotFoundError is returned when jwk.Get fails to find the
// requested field on a jwk.Key.
type FieldNotFoundError struct {
	// Name is the name of the field that was not found.
	Name string
}

func (e FieldNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e FieldNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// FieldTypeMismatchError
//-------------------------------------------------------------------

// FieldTypeMismatchError is returned when jwk.Get finds the
// requested field but the stored value cannot be converted to the
// requested type.
//
// Callers that need to distinguish "field missing" from "field present
// but wrong type" should use errors.Is with FieldNotFoundError{} /
// FieldTypeMismatchError{}, or errors.AsType to recover Name, Got,
// and Want fields.
type FieldTypeMismatchError struct {
	// Name is the name of the field whose value could not be converted.
	Name string
	// Got is the value currently stored under the field. Use %T to
	// inspect its concrete type.
	Got any
	// Want is a zero value of the requested type T. Use %T to inspect
	// its concrete type.
	Want any
}

func (e FieldTypeMismatchError) Error() string { _ = "STUB: not implemented"; return "" }

func (e FieldTypeMismatchError) Is(target error) bool { _ = "STUB: not implemented"; return false }
