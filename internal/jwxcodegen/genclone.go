package jwxcodegen

import (
	"github.com/lestrrat-go/codegen"
)

// GenerateCloneFrom emits a complete cloneFrom(src *StructName) method.
// Per-field nil-check with type-appropriate copy:
//   - Slices ([] prefix or List suffix): slices.Clone
//   - noDeref pointers: direct copy
//   - Other pointers: shallow copy
//   - Indirect types: tmp := *(src.field); dst.field = &tmp
//   - Direct types: direct copy
//
// Also copies the private params map and any extra fields.
func GenerateCloneFrom(o *codegen.Output, cfg CloneConfig) { _ = "STUB: not implemented"; return }

// Extra fields that are direct-copied before field iteration (e.g., options)

// Pointer type without noDeref — shallow copy

// Private params

// Extra fields that are not direct-copied (e.g., "raw" with slices.Clone)
