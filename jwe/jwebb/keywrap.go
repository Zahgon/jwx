package jwebb

import (
	"crypto/cipher"
)

var keywrapDefaultIV = []byte{0xa6, 0xa6, 0xa6, 0xa6, 0xa6, 0xa6, 0xa6, 0xa6}

func Wrap(kek cipher.Block, cek []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Single flat buffer for register values instead of [][]byte

func Unwrap(block cipher.Block, ciphertxt []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Single flat buffer for register values instead of [][]byte
