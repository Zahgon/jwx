package pool

import "bytes"

var bytesBufferPool = New[*bytes.Buffer](allocBytesBuffer, freeBytesBuffer)

func allocBytesBuffer() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

func freeBytesBuffer(b *bytes.Buffer) *bytes.Buffer {
	_ = "STUB: not implemented"
	// Zero the backing array before returning to pool — the buffer
	// may hold private-key material, plaintext, or HMAC input.
	// b.Bytes() shares the internal slice (offset is always 0 in
	// our write-only usage); reslicing to cap reaches all residual bytes.
	return nil
}

func BytesBuffer() *Pool[*bytes.Buffer] { _ = "STUB: not implemented"; return nil }
