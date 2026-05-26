package jwe

func uncompress(src []byte, maxBufferSize int64) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if we have a read error, and it's not EOF, then we need to stop

// if it got here, then readErr == io.EOF, we're done

func compress(plaintext []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
