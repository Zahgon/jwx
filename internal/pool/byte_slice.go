package pool

var byteSlicePool = SlicePool[byte]{
	pool: New[[]byte](allocByteSlice, freeByteSlice),
}

func allocByteSlice() []byte { _ = "STUB: not implemented"; return nil }

// Default capacity of 64 bytes

func freeByteSlice(b []byte) []byte {
	_ = "STUB: not implemented"
	// Defensive: scrub the entire backing array, not just b[:len(b)]. No
	// current caller is known to reslice past len(b) and observe stale
	// bytes, but a defer Put(buf) that captures buf at len=0 (before a
	// subsequent buf = buf[:n]) would otherwise leave plaintext resident
	// in the pool's backing storage.
	return nil
}

func ByteSlice() SlicePool[byte] { _ = "STUB: not implemented"; return nil }
