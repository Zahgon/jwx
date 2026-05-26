package jwxcodegen

import "github.com/lestrrat-go/codegen"

// DecodeCtxConfig holds configuration for GenerateDecodeCtx.
type DecodeCtxConfig struct {
	// ReceiverName is the receiver variable name (e.g., "h", "t", "k").
	ReceiverName string

	// StructName is the receiver struct name.
	StructName string

	// DecodeCtxType is the type of the decode context (e.g., "DecodeCtx",
	// "json.DecodeCtx").
	DecodeCtxType string
}

// GenerateDecodeCtx emits two complete methods: DecodeCtx() getter and
// SetDecodeCtx() setter. These are simple RLock/Lock wrappers around a `dc` field.
func GenerateDecodeCtx(o *codegen.Output, cfg DecodeCtxConfig) { _ = "STUB: not implemented"; return }
