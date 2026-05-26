package jws

func (h *stdHeaders) Copy(dst Headers) error { _ = "STUB: not implemented"; return nil }

// mergeHeaders merges two headers, and works even if the first Header
// object is nil. This is not exported because ATM it felt like this
// function is not frequently used, and MergeHeaders seemed a clunky name
func mergeHeaders(h1, h2 Headers) (Headers, error) {
	_ = "STUB: not implemented"
	return *new(Headers), nil
}

func (h *stdHeaders) Merge(h2 Headers) (Headers, error) {
	_ = "STUB: not implemented"
	return *new(Headers), nil
}

// Clone creates a deep copy of the header
func (h *stdHeaders) Clone() (Headers, error) { _ = "STUB: not implemented"; return *new(Headers), nil }
