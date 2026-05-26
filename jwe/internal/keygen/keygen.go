package keygen

// Bytes returns the byte from this ByteKey
func (k ByteKey) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func Random(n int) (ByteSource, error) { _ = "STUB: not implemented"; return *new(ByteSource), nil }

// HeaderPopulate populates the header with the required EC-DSA public key
// information ('epk' key)
func (k ByteWithECPublicKey) Populate(h Setter) error { _ = "STUB: not implemented"; return nil }

// HeaderPopulate populates the header with the required AES GCM
// parameters ('iv' and 'tag')
func (k ByteWithIVAndTag) Populate(h Setter) error { _ = "STUB: not implemented"; return nil }

// Populate populates the header with the KEM ciphertext ('ek')
func (k ByteWithEncapsulatedKey) Populate(h Setter) error { _ = "STUB: not implemented"; return nil }

// HeaderPopulate populates the header with the required PBES2
// parameters ('p2s' and 'p2c')
func (k ByteWithSaltAndCount) Populate(h Setter) error { _ = "STUB: not implemented"; return nil }
