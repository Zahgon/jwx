package jwxcodegen

import "github.com/lestrrat-go/codegen"

// GenerateKeysMethod emits a complete Keys() []string method.
// It iterates fields with nil-check, appends to a keys slice,
// then appends private params.
func GenerateKeysMethod(o *codegen.Output, cfg KeysConfig) { _ = "STUB: not implemented"; return }

// Always-present keys (e.g., KeyTypeKey for genjwk)
