package jwk

import (
	"encoding/json/jsontext"
	"iter"
)

const keysKey = `keys` // appease linter

func newSet() *set { _ = "STUB: not implemented"; return nil }

// NewSet creates and empty `jwk.Set` object
func NewSet() Set { _ = "STUB: not implemented"; return *new(Set) }

func (s *set) Set(n string, v any) error { _ = "STUB: not implemented"; return nil }

func (s *set) Field(name string) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func (s *set) Key(idx int) (Key, bool) { _ = "STUB: not implemented"; return *new(Key), false }

func (s *set) Len() int { _ = "STUB: not implemented"; return 0 }

// indexNL is Index(), but without the locking
func (s *set) indexNL(key Key) int { _ = "STUB: not implemented"; return 0 }

func (s *set) Index(key Key) int { _ = "STUB: not implemented"; return 0 }

func (s *set) AddKey(key Key) error { _ = "STUB: not implemented"; return nil }

func (s *set) Remove(name string) error { _ = "STUB: not implemented"; return nil }

func (s *set) RemoveKey(key Key) error { _ = "STUB: not implemented"; return nil }

func (s *set) Clear() error { _ = "STUB: not implemented"; return nil }

func (s *set) Keys() []string { _ = "STUB: not implemented"; return nil }

func (s *set) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (s *set) setMaxKeys(n int) { _ = "STUB: not implemented"; return }

func (s *set) setRejectDuplicateKID(v bool) { _ = "STUB: not implemented"; return }

// UnmarshalJSON delegates to UnmarshalJSONFrom so the streaming /
// cap-before-allocate behavior is identical for stdlib v1 callers
// (encoding/json) and jsonv2 callers. This entry point requires JWKS
// shape — bare JWK input is rejected here. Callers that don't know
// the shape ahead of time should use [Parse], which dispatches.
func (s *set) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

// UnmarshalJSONFrom streams the JWKS document via the supplied decoder.
// The "keys" array is read element-by-element with the per-Set cap
// (or the global default) enforced BEFORE the (cap+1)-th element is
// tokenized — an attacker-controlled input length cannot force
// allocation past the cap.
func (s *set) UnmarshalJSONFrom(dec *jsontext.Decoder) error { _ = "STUB: not implemented"; return nil }

func (s *set) LookupKeyID(kid string) (Key, bool) {
	_ = "STUB: not implemented"
	return *new(Key), false
}

func (s *set) All() iter.Seq2[int, Key] { _ = "STUB: not implemented"; return nil }

func (s *set) Fields() iter.Seq2[string, any] { _ = "STUB: not implemented"; return nil }

func (s *set) DecodeCtx() DecodeCtx { _ = "STUB: not implemented"; return *new(DecodeCtx) }

func (s *set) SetDecodeCtx(dc DecodeCtx) { _ = "STUB: not implemented"; return }

func (s *set) Clone() (Set, error) { _ = "STUB: not implemented"; return *new(Set), nil }
