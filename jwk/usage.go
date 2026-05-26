package jwk

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var strictKeyUsage = atomic.Bool{}
var keyUsageNames = map[string]struct{}{}
var muKeyUsageName sync.RWMutex

// RegisterKeyUsage registers a possible value that can be used for KeyUsageType.
// Normally, key usage (or the "use" field in a JWK) is either "sig" or "enc",
// but other values may be used.
//
// While this module only works with "sig" and "enc", it is possible that
// systems choose to use other values. This function allows users to register
// new values to be accepted as valid key usage types. Values are case sensitive.
//
// Furthermore, the check against registered values can be completely turned off
// by setting the global option `jwk.WithStrictKeyUsage(false)`.
//
// The error return is reserved for future validation. The current
// implementation always returns nil, but callers — especially extension
// modules calling this from init() — must check the return value and panic
// on failure to stay forward-compatible.
func RegisterKeyUsage(v string) error { _ = "STUB: not implemented"; return nil }

// UnregisterKeyUsage removes v from the allowlist maintained by
// [RegisterKeyUsage]. The error return is reserved for future
// validation (for example, refusing to unregister a built-in usage
// value like "sig" or "enc") and is always nil today. Callers
// scripting Register/Unregister cycles should check the returned
// value and propagate on failure to stay forward-compatible,
// matching the convention on [RegisterKeyUsage].
func UnregisterKeyUsage(v string) error { _ = "STUB: not implemented"; return nil }

func init() {
	strictKeyUsage.Store(true)
	if err := RegisterKeyUsage("sig"); err != nil {
		panic(fmt.Sprintf("jwk: failed to register builtin KeyUsage: %s", err))
	}
	if err := RegisterKeyUsage("enc"); err != nil {
		panic(fmt.Sprintf("jwk: failed to register builtin KeyUsage: %s", err))
	}
}

func isValidUsage(v string) bool {
	_ = "STUB: not implemented"
	// This function can return true if strictKeyUsage is false
	return false
}

func (k KeyUsageType) String() string { _ = "STUB: not implemented"; return "" }

func (k *KeyUsageType) Accept(v any) error { _ = "STUB: not implemented"; return nil }
