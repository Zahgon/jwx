package jwxcodegen

import (
	"github.com/lestrrat-go/codegen"
)

// GenerateSetCases emits case clauses for a setNoLock(name string, value any) error switch.
//
// The fieldStorageTypeIsIndirect parameter controls whether a field's type is stored
// behind a pointer. Each generator has slightly different rules for this.
//
// Fields in the Skip set are not emitted; the caller handles them (e.g., algorithm
// has custom handling in genheaders and genjwk).
func GenerateSetCases(o *codegen.Output, cfg CaseConfig, fieldStorageTypeIsIndirect func(codegen.Field) bool) {
	_ = "STUB: not implemented"
	return
}
