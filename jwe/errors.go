package jwe

import (
	"errors"

	"github.com/lestrrat-go/jwx/v4/jwa"
)

type encryptError struct {
	error
}

func (e encryptError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (encryptError) Is(err error) bool { _ = "STUB: not implemented"; return false }

var errDefaultEncryptError = encryptError{errors.New(`encrypt error`)}

// EncryptError returns an error that can be passed to `errors.Is` to check if the error is an error returned by `jwe.Encrypt`.
func EncryptError() error { _ = "STUB: not implemented"; return nil }

func makeEncryptError(prefix string, f string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

type decryptError struct {
	error
}

func (e decryptError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (decryptError) Is(err error) bool { _ = "STUB: not implemented"; return false }

var errDefaultDecryptError = decryptError{errors.New(`decrypt error`)}

// DecryptError returns an error that can be passed to `errors.Is` to check if the error is an error returned by `jwe.Decrypt`.
func DecryptError() error { _ = "STUB: not implemented"; return nil }

func makeDecryptError(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

type recipientError struct {
	error
}

func (e recipientError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (recipientError) Is(err error) bool { _ = "STUB: not implemented"; return false }

var errDefaultRecipientError = recipientError{errors.New(`recipient error`)}

// RecipientError returns an error that can be passed to `errors.Is` to check if the error is
// an error that occurred while attempting to decrypt a JWE message for a particular recipient.
//
// For example, if the JWE message failed to parse during `jwe.Decrypt`, it will be a
// `jwe.DecryptError`, but NOT `jwe.RecipientError`. However, if the JWE message could not
// be decrypted for any of the recipients, then it will be a `jwe.RecipientError`
// (actually, it will be _multiple_ `jwe.RecipientError` errors, one for each recipient)
func RecipientError() error { _ = "STUB: not implemented"; return nil }

func makeRecipientError(err error) error { _ = "STUB: not implemented"; return nil }

type hpkeError struct {
	error
}

func (e hpkeError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (hpkeError) Is(err error) bool { _ = "STUB: not implemented"; return false }

var errDefaultHPKEError = hpkeError{errors.New(`HPKE error`)}

// HPKEError returns an error that can be passed to `errors.Is` to check
// if the error originated from an HPKE encrypt or decrypt operation.
func HPKEError() error { _ = "STUB: not implemented"; return nil }

func makeHPKEError(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

type parseError struct {
	error
}

func (e parseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (parseError) Is(err error) bool { _ = "STUB: not implemented"; return false }

var errDefaultParseError = parseError{errors.New(`parse error`)}

// ParseError returns an error that can be passed to `errors.Is` to check if the error
// is an error returned by `jwe.Parse` and related functions.
func ParseError() error { _ = "STUB: not implemented"; return nil }

func makeParseError(prefix string, f string, args ...any) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------------------
// MissingContentEncryptionError
//-------------------------------------------------------------------

// MissingContentEncryptionError is returned when jwe.Decrypt cannot
// locate the content encryption algorithm ("enc") in the protected
// headers of the JWE message.
//
// Use errors.Is with a zero-value MissingContentEncryptionError{} to
// detect this failure mode programmatically:
//
//	if errors.Is(err, jwe.MissingContentEncryptionError{}) { ... }
type MissingContentEncryptionError struct{}

func (MissingContentEncryptionError) Error() string { _ = "STUB: not implemented"; return "" }

func (MissingContentEncryptionError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// AlgorithmMismatchError
//-------------------------------------------------------------------

// AlgorithmMismatchError is returned when jwe.Decrypt detects that
// the key encryption algorithm ("alg") declared in the JWE headers
// does not match the algorithm associated with the decryption key
// supplied by the caller.
//
// Expected holds the algorithm bound to the caller's key (for example
// via jwe.WithKey). Got holds the algorithm found in the per-recipient
// or protected headers of the message.
//
// Use errors.Is with a zero-value AlgorithmMismatchError{} to detect
// this failure mode, or errors.AsType to recover the Expected and Got
// fields:
//
//	if mismatch, ok := errors.AsType[jwe.AlgorithmMismatchError](err); ok {
//	    log.Printf("alg mismatch: expected %s, got %s", mismatch.Expected, mismatch.Got)
//	}
type AlgorithmMismatchError struct {
	// Expected is the key encryption algorithm associated with the
	// decryption key supplied by the caller.
	Expected jwa.KeyEncryptionAlgorithm
	// Got is the key encryption algorithm declared in the JWE
	// per-recipient or protected headers.
	Got jwa.KeyEncryptionAlgorithm
}

func (e AlgorithmMismatchError) Error() string { _ = "STUB: not implemented"; return "" }

func (AlgorithmMismatchError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// FieldNotFoundError
//-------------------------------------------------------------------

// FieldNotFoundError is returned when jwe.Get fails to find the
// requested field on a jwe.Headers.
type FieldNotFoundError struct {
	// Name is the name of the field that was not found.
	Name string
}

func (e FieldNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e FieldNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// FieldTypeMismatchError
//-------------------------------------------------------------------

// FieldTypeMismatchError is returned when jwe.Get finds the requested
// field but the stored value cannot be converted to the requested type.
//
// Callers that need to distinguish "field missing" from "field present
// but wrong type" should use errors.Is with FieldNotFoundError{} /
// FieldTypeMismatchError{}, or errors.AsType to recover Name, Got, and
// Want fields.
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
