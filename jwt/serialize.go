package jwt

import (
	"github.com/lestrrat-go/jwx/v4/jwe"
	"github.com/lestrrat-go/jwx/v4/jws"
)

type SerializeCtx interface {
	Step() int
	Nested() bool
}

type serializeCtx struct {
	step   int
	nested bool
}

func (ctx *serializeCtx) Step() int { _ = "STUB: not implemented"; return 0 }

func (ctx *serializeCtx) Nested() bool { _ = "STUB: not implemented"; return false }

type SerializeStep interface {
	Serialize(SerializeCtx, any) (any, error)
}

// errStep is always an error. used to indicate that a method like
// serializer.Sign or Encrypt already errored out on configuration
type errStep struct {
	err error
}

func (e errStep) Serialize(_ SerializeCtx, _ any) (any, error) {
	_ = "STUB: not implemented"

	// Serializer is a generic serializer for JWTs. Whereas other convenience
	// functions can only do one thing (such as generate a JWS signed JWT),
	// Using this construct you can serialize the token however you want.
	//
	// By default, the serializer only marshals the token into a JSON payload.
	// You must set up the rest of the steps that should be taken by the
	// serializer.
	//
	// For example, to marshal the token into JSON, then apply JWS and JWE
	// in that order, you would do:
	//
	//	serialized, err := jwt.NewSerializer().
	//	   Sign(jwa.RS256, key).
	//	   Encrypt(jwe.WithEncryptOption(jwe.WithKey(jwa.RSA_OAEP(), publicKey))).
	//	   Serialize(token)
	//
	// The `jwt.Sign()` function is equivalent to
	//
	//	serialized, err := jwt.NewSerializer().
	//	   Sign(...args...).
	//	   Serialize(token)
	return *new(any), nil
}

type Serializer struct {
	steps []SerializeStep
}

// NewSerializer creates a new empty serializer.
func NewSerializer() *Serializer { _ = "STUB: not implemented"; return nil }

// Reset clears all of the registered steps.
func (s *Serializer) Reset() *Serializer { _ = "STUB: not implemented"; return nil }

// Step adds a new Step to the serialization process
func (s *Serializer) Step(step SerializeStep) *Serializer { _ = "STUB: not implemented"; return nil }

type jsonSerializer struct{}

func (jsonSerializer) Serialize(_ SerializeCtx, v any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

type genericHeader interface {
	Field(string) (any, bool)
	Set(string, any) error
	Has(string) bool
}

func setTypeOrCty(ctx SerializeCtx, hdrs genericHeader) error {
	_ = "STUB: not implemented"
	// cty and typ are common between JWE/JWS, so we don't use
	// the constants in jws/jwe package here
	return nil
}

// We are executed immediately after json marshaling

// If this is part of a nested sequence, we should set cty = 'JWT'
// https://datatracker.ietf.org/doc/html/rfc7519#section-5.2

type jwsSerializer struct {
	options []jws.SignOption
}

func (s *jwsSerializer) Serialize(ctx SerializeCtx, v any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// this is already wrapped

// JWTs MUST NOT use b64 = false
// https://datatracker.ietf.org/doc/html/rfc7797#section-7

func (s *Serializer) Sign(options ...SignOption) *Serializer { _ = "STUB: not implemented"; return nil }

// we need to from SignOption to Option because ... reasons
// (todo: when go1.18 prevails, use type parameters

func (s *Serializer) sign(options ...jws.SignOption) *Serializer {
	_ = "STUB: not implemented"
	return nil
}

type jweSerializer struct {
	options []jwe.EncryptOption
}

func (s *jweSerializer) Serialize(ctx SerializeCtx, v any) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// this is already wrapped

// Encrypt specifies the JWT to be serialized as an encrypted payload.
//
// One notable difference between this method and `jwe.Encrypt()` is that
// while `jwe.Encrypt()` OVERWRITES the previous headers when `jwe.WithProtectedHeaders()`
// is provided, this method MERGES them. This is due to the fact that we
// MUST add some extra headers to construct a proper JWE message.
// Be careful when you pass multiple `jwe.EncryptOption`s.
func (s *Serializer) Encrypt(options ...EncryptOption) *Serializer {
	_ = "STUB: not implemented"
	return nil
}

// we need to from SignOption to Option because ... reasons
// (todo: when go1.18 prevails, use type parameters

func (s *Serializer) encrypt(options ...jwe.EncryptOption) *Serializer {
	_ = "STUB: not implemented"
	return nil
}

func (s *Serializer) Serialize(t Token) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
