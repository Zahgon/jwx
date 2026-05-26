package jwxcodegen

import (
	"github.com/lestrrat-go/codegen"
)

// GenerateUnmarshalCases emits case clauses for the streaming JSON decoder
// switch inside UnmarshalJSON. Each field type uses a different decode strategy:
//   - string: json.AssignNextStringToken
//   - []byte: json.AssignNextBytesToken
//   - jwk.Key: dec.ReadValue() then jwk.ParseKey
//   - slice types ([]...): json.UnmarshalDecode into slice, assign directly
//   - noDeref pointer types or IsPointer types: json.UnmarshalDecode into PointerElem, assign &decoded
//   - other types: json.UnmarshalDecode into type, assign &decoded
//
// The decodeCtxExpr parameter is the expression for the decode context passed to
// AssignNextStringToken (e.g., "nil" for genheaders, "t.dc" for genjwt, "h.dc" for genjwk).
//
// Fields in the Skip set are not emitted; the caller handles them separately.
func GenerateUnmarshalCases(o *codegen.Output, cfg CaseConfig, decodeCtxExpr string) {
	_ = "STUB: not implemented"
	return
}
