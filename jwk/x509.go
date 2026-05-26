package jwk

// decodeX509 decodes a single PEM block from src. It returns the
// decoded value, the remaining bytes after the consumed block, and any
// error. Callers iterate by calling this repeatedly with the returned
// rest until empty.
//
// Dispatch is delegated to [jwkbb.DecodeX509], which routes by
// block.Type to a decoder registered via [jwkbb.RegisterX509Decoder].
func decodeX509(src []byte) (any, []byte, error) {
	_ = "STUB: not implemented"
	return *new(any), nil, nil
}
