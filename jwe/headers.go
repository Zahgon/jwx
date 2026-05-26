package jwe

type isZeroer interface {
	isZero() bool
}

func (h *stdHeaders) Clone() (Headers, error) { _ = "STUB: not implemented"; return *new(Headers), nil }

func (h *stdHeaders) Copy(dst Headers) error { _ = "STUB: not implemented"; return nil }

// copyNoLock copies all fields from h to dst without acquiring any mutexes.
// Both h and dst must be exclusively owned by the caller (not shared).
func (h *stdHeaders) copyNoLock(dst *stdHeaders) { _ = "STUB: not implemented"; return }

func (h *stdHeaders) Merge(h2 Headers) (Headers, error) {
	_ = "STUB: not implemented"
	return *new(Headers), nil
}

// mergeIntoNoLock copies non-nil fields from h into dst without acquiring
// any mutexes. Both h and dst must be exclusively owned by the caller.
// Unlike cloneFrom, this only overwrites fields that are set in h,
// leaving existing values in dst untouched.
func (h *stdHeaders) mergeIntoNoLock(dst *stdHeaders) { _ = "STUB: not implemented"; return }

func (h *stdHeaders) Encode() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (h *stdHeaders) Decode(buf []byte) error {
	_ = "STUB: not implemented"
	// base64 json string -> json object representation of header
	return nil
}
