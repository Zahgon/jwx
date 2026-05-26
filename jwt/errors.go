package jwt

import (
	"errors"
	"time"
)

// errUnknownPayloadType is the sentinel for unknown payload type errors.
var errUnknownPayloadType = errors.New(`unknown payload type (payload is not JWT?)`)

// UnknownPayloadTypeError returns the opaque error value that is returned when
// jwt.Parse fails due to not being able to deduce the format of
// the incoming buffer.
//
// This value should only be used for comparison using errors.Is().
func UnknownPayloadTypeError() error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// ClaimNotFoundError
//-------------------------------------------------------------------

// ClaimNotFoundError is returned when jwt.Get fails to find the requested claim.
type ClaimNotFoundError struct {
	// Name is the name of the claim that was not found.
	Name string
}

func (e ClaimNotFoundError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ClaimNotFoundError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// ClaimTypeMismatchError
//-------------------------------------------------------------------

// ClaimTypeMismatchError is returned when jwt.Get finds the requested
// claim but the stored value cannot be converted to the requested type.
//
// Callers that need to distinguish "claim missing" from "claim present
// but wrong type" should use errors.Is with ClaimNotFoundError{} /
// ClaimTypeMismatchError{}, or errors.AsType to recover the Name, Got,
// and Want fields.
type ClaimTypeMismatchError struct {
	// Name is the name of the claim whose value could not be converted.
	Name string
	// Got is the value currently stored under the claim. Use %T to
	// inspect its concrete type.
	Got any
	// Want is a zero value of the requested type T. Use %T to inspect
	// its concrete type.
	Want any
}

func (e ClaimTypeMismatchError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ClaimTypeMismatchError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// ClaimAssignmentFailedError
//-------------------------------------------------------------------

// ClaimAssignmentFailedError is returned when jwt.Get fails to assign
// the value to the destination.
type ClaimAssignmentFailedError struct {
	// Err is the underlying error.
	Err error
}

func (e ClaimAssignmentFailedError) Error() string { _ = "STUB: not implemented"; return "" }

func (e ClaimAssignmentFailedError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (e ClaimAssignmentFailedError) Is(target error) bool { _ = "STUB: not implemented"; return false }

//-------------------------------------------------------------------
// ParseError
//-------------------------------------------------------------------

// ParseError is returned when jwt.Parse fails.
type ParseError struct {
	error
}

func (e ParseError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (ParseError) Is(err error) bool { _ = "STUB: not implemented"; return false }

func parseErrorf(prefix, f string, args ...any) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// ValidationError
//-------------------------------------------------------------------

// ValidationError is the blanket error returned by jwt.Validate.
// It wraps the specific validation failure(s).
type ValidationError struct {
	error
}

func (ValidationError) Is(err error) bool { _ = "STUB: not implemented"; return false }

func (err ValidationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func validateErrorf(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

func validateErrorJoin(errs ...error) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// InvalidIssuerError
//-------------------------------------------------------------------

// InvalidIssuerError is returned when the iss claim is not satisfied.
type InvalidIssuerError struct {
	error
}

func (err InvalidIssuerError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err InvalidIssuerError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func issuerErrorf(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// TokenExpiredError
//-------------------------------------------------------------------

// TokenExpiredError is returned when the exp claim is not satisfied.
// The structured fields allow callers to inspect what values were
// compared without re-parsing the token.
type TokenExpiredError struct {
	error

	// Expiration is the token's exp claim value (after truncation).
	Expiration time.Time
	// Now is the time used for comparison (after truncation).
	Now time.Time
	// Skew is the acceptable skew duration.
	Skew time.Duration
}

func (err TokenExpiredError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err TokenExpiredError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newTokenExpiredError(expiration, now time.Time, skew time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------------------
// InvalidIssuedAtError
//-------------------------------------------------------------------

// InvalidIssuedAtError is returned when the iat claim is not satisfied.
type InvalidIssuedAtError struct {
	error

	// IssuedAt is the token's iat claim value (after truncation).
	IssuedAt time.Time
	// Now is the time used for comparison (after truncation).
	Now time.Time
	// Skew is the acceptable skew duration.
	Skew time.Duration
}

func (err InvalidIssuedAtError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err InvalidIssuedAtError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newInvalidIssuedAtError(issuedAt, now time.Time, skew time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------------------
// TokenNotYetValidError
//-------------------------------------------------------------------

// TokenNotYetValidError is returned when the nbf claim is not satisfied.
type TokenNotYetValidError struct {
	error

	// NotBefore is the token's nbf claim value (after truncation).
	NotBefore time.Time
	// Now is the time used for comparison (after truncation).
	Now time.Time
	// Skew is the acceptable skew duration.
	Skew time.Duration
}

func (err TokenNotYetValidError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err TokenNotYetValidError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newTokenNotYetValidError(notBefore, now time.Time, skew time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------------------
// InvalidAudienceError
//-------------------------------------------------------------------

// InvalidAudienceError is returned when the aud claim is not satisfied.
type InvalidAudienceError struct {
	error
}

func (err InvalidAudienceError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err InvalidAudienceError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func audienceErrorf(f string, args ...any) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// MissingRequiredClaimError
//-------------------------------------------------------------------

// MissingRequiredClaimError is returned when a required claim is missing.
type MissingRequiredClaimError struct {
	error

	// Claim is the name of the missing required claim.
	Claim string
}

func (err MissingRequiredClaimError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err MissingRequiredClaimError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func missingRequiredClaimErrorf(name string) error { _ = "STUB: not implemented"; return nil }

//-------------------------------------------------------------------
// ClaimValidationError
//-------------------------------------------------------------------

// ClaimValidationError is returned when a generic claim validator
// (ClaimValueIs, ClaimContainsString) detects a mismatch.
//
// The Error() message intentionally omits the Expected and Actual
// values so that arbitrary claim payloads do not end up in log output
// via "%v"/"%s" formatting. Callers that need to inspect the mismatched
// values can access the fields directly, but should treat them as
// potentially sensitive (PII, secrets, etc.) and avoid logging them
// verbatim.
type ClaimValidationError struct {
	error

	// Claim is the name of the claim that failed validation.
	Claim string
	// Expected is the value the validator expected. May contain
	// sensitive data; see the type-level godoc.
	Expected any
	// Actual is the value found in the token. May contain sensitive
	// data (PII, secrets, etc.); see the type-level godoc.
	Actual any
}

func (err ClaimValidationError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err ClaimValidationError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newClaimValidationError(claim string, expected, actual any, msg string) error {
	_ = "STUB: not implemented"
	return nil
}

//-------------------------------------------------------------------
// TimeDeltaError
//-------------------------------------------------------------------

// TimeDeltaError is returned when a time delta validator
// (MaxDeltaIs, MinDeltaIs) detects that the delta between two
// claims is out of range.
type TimeDeltaError struct {
	error

	// Claim1 is the name of the first claim (or "" for current time).
	Claim1 string
	// Claim2 is the name of the second claim (or "" for current time).
	Claim2 string
	// Value1 is the resolved time value for Claim1.
	Value1 time.Time
	// Value2 is the resolved time value for Claim2.
	Value2 time.Time
	// Delta is the actual duration between Value1 and Value2.
	Delta time.Duration
	// Limit is the threshold duration that was exceeded or not met.
	Limit time.Duration
	// Skew is the acceptable skew duration.
	Skew time.Duration
}

func (err TimeDeltaError) Is(target error) bool { _ = "STUB: not implemented"; return false }

func (err TimeDeltaError) Unwrap() error { _ = "STUB: not implemented"; return nil }

func newTimeDeltaError(msg, c1, c2 string, v1, v2 time.Time, delta, limit, skew time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}
