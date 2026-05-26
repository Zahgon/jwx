package jwxcodegen

import "github.com/lestrrat-go/codegen"

// GenerateMarshalJSON emits a complete MarshalJSON method. The pattern builds a
// pooled []fieldPair slice under RLock, sorts, then writes JSON directly to a
// pooled buffer using fmt.Fprintf and json.Marshal. []byte values are base64-encoded.
//
// The caller must ensure the target package contains the fieldPair type,
// fieldPairPool, getFieldPairList, putFieldPairList, and fieldPairLess.
//
// Used by genheaders and genjwk, but NOT genjwt (which uses claimPair pool).
func GenerateMarshalJSON(o *codegen.Output, cfg MarshalConfig) { _ = "STUB: not implemented"; return }

// Always-present entries (e.g., KeyTypeKey for genjwk)
