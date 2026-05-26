package concatkdf

import (
	"crypto"
)

type KDF struct {
	buf       []byte
	otherinfo []byte
	z         []byte
	hash      crypto.Hash
}

func New(hash crypto.Hash, alg, Z, apu, apv, pubinfo, privinfo []byte) *KDF {
	_ = "STUB: not implemented"
	// Write length-prefixed fields directly into a single buffer,
	// avoiding intermediate allocations from ndata().
	return nil
}

func (k *KDF) Read(out []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
