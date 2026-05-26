package jwxcodegen

import "github.com/lestrrat-go/codegen"

// GenerateFieldCases emits case clauses for a Field(name string) (any, bool) switch.
// Each case does a nil check, then returns the value (dereferenced for indirect types).
//
// The fieldStorageTypeIsIndirect parameter controls whether a field's type is stored
// behind a pointer. Each generator has slightly different rules for this.
func GenerateFieldCases(o *codegen.Output, cfg CaseConfig, fieldStorageTypeIsIndirect func(codegen.Field) bool) {
	_ = "STUB: not implemented"
	return
}
