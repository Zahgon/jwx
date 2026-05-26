package main

import (
	"github.com/lestrrat-go/codegen"
	"github.com/lestrrat-go/jwx/v4/internal/jwxcodegen"
)

func runJWT(args []string) error { _ = "STUB: not implemented"; return nil }

func computePkgPrefix(obj *codegen.Object) string { _ = "STUB: not implemented"; return "" }

func tokenKeyName(f codegen.Field) string { _ = "STUB: not implemented"; return "" }

func makeTokenCaseConfig(obj *codegen.Object) jwxcodegen.CaseConfig {
	_ = "STUB: not implemented"
	return *new(jwxcodegen.CaseConfig)
}

func generateTokenConstants(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	return
}

// end const

func generateTokenInterface(o *codegen.Output, obj *codegen.Object, pkgPrefix string) {
	_ = "STUB: not implemented"
	return
}

func generateTokenStruct(o *codegen.Output, obj *codegen.Object, pkgPrefix string) {
	_ = "STUB: not implemented"
	return
}

// end type Token

func generateTokenConstructor(o *codegen.Output, obj *codegen.Object, pkgPrefix string) {
	_ = "STUB: not implemented"
	return
}

func generateTokenHas(o *codegen.Output, obj *codegen.Object) { _ = "STUB: not implemented"; return }

func generateTokenFieldAndGet(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	return
}

// end switch name
// end of Field

func generateTokenRemove(o *codegen.Output, obj *codegen.Object) { _ = "STUB: not implemented"; return }

// currently unused, but who knows

func generateTokenSet(o *codegen.Output, obj *codegen.Object) { _ = "STUB: not implemented"; return }

// Build skip set for fields with special handling in setNoLock

// Emit algorithm special case inline if present

// end if t.privateClaims == nil

// end switch name

// end func (t *%s) Set(name string, value any)

func generateTokenGetters(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	return
}

// func (h *stdHeaders) %s() %s

func generateTokenPrivateClaims(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	return
}

func generateTokenUnmarshalJSON(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	return
}

// This looks like bad code, but we're unrolling things for maximum
// runtime efficiency

func generateTokenKeys(o *codegen.Output, obj *codegen.Object) { _ = "STUB: not implemented"; return }

func generateTokenClaims(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	// Generate Claims() iter.Seq2[string, any] method.
	// Snapshots claims under the read lock and releases it before yielding,
	// so callers can safely access other token methods inside the yield closure
	// without deadlocking on the token's RWMutex.
	return
}

func generateTokenCloneFrom(o *codegen.Output, obj *codegen.Object) {
	_ = "STUB: not implemented"
	return
}

func generateTokenMakePairsAndMarshal(o *codegen.Output, obj *codegen.Object, pkgPrefix string) {
	_ = "STUB: not implemented"
	return
}

// Need to handle audience flattening specially

func generateToken(obj *codegen.Object) error { _ = "STUB: not implemented"; return nil }

func genBuilder(obj *codegen.Object) error { _ = "STUB: not implemented"; return nil }
