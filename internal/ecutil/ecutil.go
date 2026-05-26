// Package ecutil defines tools that help with elliptic curve related
// computation
package ecutil

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

const (
	// size of buffer that needs to be allocated for EC521 curve
	ec521BufferSize = 66 // (521 / 8) + 1
)

var ecpointBufferPool = sync.Pool{
	New: func() any {
		// In most cases the curve bit size will be less than this length
		// so allocate the maximum, and keep reusing
		buf := make([]byte, 0, ec521BufferSize)
		return &buf
	},
}

func getCrvFixedBuffer(size int) []byte {
	_ = "STUB: not implemented"
	//nolint:forcetypeassert
	return nil
}

// ReleaseECPointBuffer releases the []byte buffer allocated.
func ReleaseECPointBuffer(buf []byte) { _ = "STUB: not implemented"; return }

func CalculateKeySize(crv elliptic.Curve) int {
	_ = "STUB: not implemented"
	// We need to create a buffer that fits the entire curve.
	// If the curve size is 66, that fits in 9 bytes. If the curve
	// size is 64, it fits in 8 bytes.
	return 0
}

// For most common cases we know before hand what the byte length
// is going to be. optimize

// TODO: use constant?

// AllocECPointBuffer allocates a buffer for the given point in the given
// curve. This buffer should be released using the ReleaseECPointBuffer
// function.
func AllocECPointBuffer(v *big.Int, crv elliptic.Curve) []byte {
	_ = "STUB: not implemented"
	return nil
}
