package jws

import (
	"io"

	"github.com/lestrrat-go/jwx/v4/internal/pool"
)

type signContext struct {
	format        int
	detached      bool
	validateKey   bool
	payload       []byte
	payloadReader io.Reader
	encoder       Base64Encoder
	none          *signatureBuilder // special signature builder
	sigbuilders   []*signatureBuilder
}

var signContextPool = pool.New[*signContext](allocSignContext, freeSignContext)

func allocSignContext() *signContext { _ = "STUB: not implemented"; return nil }

func freeSignContext(ctx *signContext) *signContext { _ = "STUB: not implemented"; return nil }

func (sc *signContext) ProcessOptions(options []SignOption) error {
	_ = "STUB: not implemented"
	return nil
}

// No, we don't accept "none" here.

// Surface any deferred error captured while precomputing the
// fast-path header JSON (e.g. an algorithm name that would
// require JSON escaping).

// Streaming sign rejects WithInsecureNoSignature up-front so the
// caller's payload Reader is not touched. The narrower streaming
// signer used to catch this after touching the reader.

func (sc *signContext) PopulateMessage(m *Message) error { _ = "STUB: not implemented"; return nil }

// Create signature for each builders
