package cert

import (
	"sync/atomic"
)

const (
	defaultMaxChainLength     = 10
	defaultMaxCertificateSize = 256 * 1024
)

var maxChainLength atomic.Int64
var maxCertificateSize atomic.Int64

func init() {
	maxChainLength.Store(defaultMaxChainLength)
	maxCertificateSize.Store(defaultMaxCertificateSize)
}

// Settings configures process-global validation limits for `cert.Parse()` and
// `cert.Chain` ingestion.
//
// These settings are read atomically, so changing them at runtime is race-free.
// However, concurrent parses may observe a mix of old and new values. Configure
// them once at program startup when possible.
//
// Returns a non-nil error and applies no changes if any option fails
// validation (for example, a negative [WithMaxChainLength] or
// [WithMaxCertificateSize]).
func Settings(options ...GlobalOption) error {
	_ = "STUB: not implemented"
	// Validate first so the call is all-or-nothing on error.
	return nil
}

func currentMaxChainLength() int64 { _ = "STUB: not implemented"; return 0 }

func currentMaxCertificateSize() int64 { _ = "STUB: not implemented"; return 0 }

func validateChainLength(n int) error { _ = "STUB: not implemented"; return nil }

func validateCertificateSize(n int) error { _ = "STUB: not implemented"; return nil }
