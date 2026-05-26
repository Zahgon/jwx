package jwt

import "sync/atomic"

// TokenOptionSet is a bit flag containing per-token options.
type TokenOptionSet uint64

var defaultOptions atomic.Uint64

// TokenOption describes a single token option that can be set on
// the per-token option set (TokenOptionSet)
type TokenOption uint64

const (
	// FlattenAudience option controls whether the "aud" claim should be flattened
	// to a single string upon the token being serialized to JSON.
	//
	// This is sometimes important when a JWT consumer does not understand that
	// the "aud" claim can actually take the form of an array of strings.
	// (We have been notified by users that AWS Cognito has manifested this behavior
	// at some point)
	//
	// Unless the global option is set using `jwt.Settings()`, the default value is
	// `disabled`, which means that "aud" claims are always rendered as a arrays of
	// strings when serialized to JSON.
	FlattenAudience TokenOption = 1 << iota

	// MaxPerTokenOption is a marker to denote the last value that an option can take.
	// This value has no meaning other than to be used as a marker.
	MaxPerTokenOption
)

// Value returns the uint64 value of a single option
func (o TokenOption) Value() uint64 {
	_ = "STUB: not implemented"

	// Value returns the uint64 bit flag value of an option set
	return 0
}

func (o TokenOptionSet) Value() uint64 {
	_ = "STUB: not implemented"

	// DefaultOptionSet creates a new TokenOptionSet using the default
	// option set. This may differ depending on if/when functions that
	// change the global state has been called, such as `jwt.Settings`
	return 0
}

func DefaultOptionSet() TokenOptionSet { _ = "STUB: not implemented"; return *new(TokenOptionSet) }

// Clear sets all bits to zero, effectively disabling all options
func (o *TokenOptionSet) Clear() { _ = "STUB: not implemented"; return }

// Set sets the value of this option set, effectively *replacing*
// the entire option set with the new value. This is NOT the same
// as Enable/Disable.
func (o *TokenOptionSet) Set(s TokenOptionSet) {
	_ = "STUB: not implemented"

	// Enable sets the appropriate value to enable the option in the
	// option set
	return
}

func (o *TokenOptionSet) Enable(flag TokenOption) { _ = "STUB: not implemented"; return }

// Disable sets the appropriate value to disable the option in the
// option set
func (o *TokenOptionSet) Disable(flag TokenOption) { _ = "STUB: not implemented"; return }

// IsEnabled returns true if the given bit on the option set is enabled.
func (o TokenOptionSet) IsEnabled(flag TokenOption) bool { _ = "STUB: not implemented"; return false }
