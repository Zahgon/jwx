//go:generate ../scripts/jwxcodegen.sh generate-headers -objects=objects.yml

// Package jws implements the digital signature on JSON based data
// structures as described in https://tools.ietf.org/html/rfc7515
//
// If you do not care about the details, the only things that you
// would need to use are the following functions:
//
//	jws.Sign(payload, jws.WithKey(algorithm, key))
//	jws.Verify(serialized, jws.WithKey(algorithm, key))
//
// To sign, simply use `jws.Sign`. `payload` is a []byte buffer that
// contains whatever data you want to sign. `alg` is one of the
// jwa.SignatureAlgorithm constants from package jwa. For RSA and
// ECDSA family of algorithms, you will need to prepare a private key.
// For HMAC family, you just need a []byte value. The `jws.Sign`
// function will return the encoded JWS message on success.
//
// To verify, use `jws.Verify`. It will parse the `encodedjws` buffer
// and verify the result using `algorithm` and `key`. Upon successful
// verification, the original payload is returned, so you can work on it.
//
// `jws.Sign()` and `jws.Verify()` are the default general-purpose entry
// points. For detached payloads already available as `[]byte`, pass
// `jws.WithDetachedPayload()`. For detached payloads that should be
// streamed from an `io.Reader` without materializing them in memory, pass
// `jws.WithDetachedPayloadReader()`. The streaming path is intentionally
// narrower — single key, detached only, HMAC/RSA/ECDSA only.
//
// As a sidenote, consider using github.com/lestrrat-go/htmsig if you
// looking for HTTP Message Signatures (RFC9421) -- it uses the same
// underlying signing/verification mechanisms as this module.
package jws

import (
	"io"
	"sync"
	"sync/atomic"

	"github.com/lestrrat-go/jwx/v4/internal/json"
	"github.com/lestrrat-go/jwx/v4/jwa"
)

var registry = json.NewRegistry()

var maxSignatures atomic.Int64

func init() {
	maxSignatures.Store(100)
}

type defaultSigner struct {
	alg jwa.SignatureAlgorithm
}

func (s defaultSigner) Sign(key any, payload []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

const (
	fmtInvalid = 1 << iota
	fmtCompact
	fmtJSON
	fmtJSONPretty
	fmtMax
)

func validateKeyBeforeUse(key any) error { _ = "STUB: not implemented"; return nil }

// Sign generates a JWS message for the given payload and returns
// it in serialized form, which can be in either compact or
// JSON format. Default is compact.
//
// You must pass at least one key to `jws.Sign()` by using `jws.WithKey()`
// option.
//
//	jws.Sign(payload, jws.WithKey(alg, key))
//	jws.Sign(payload, jws.WithJSON(), jws.WithKey(alg1, key1), jws.WithKey(alg2, key2))
//
// Note that in the second example the `jws.WithJSON()` option is
// specified as well. This is because the compact serialization
// format does not support multiple signatures, and users must
// specifically ask for the JSON serialization format.
//
// Read the documentation for `jws.WithKey()` to learn more about the
// possible values that can be used for `alg` and `key`.
//
// You may create JWS messages with the "none" (jwa.NoSignature) algorithm
// if you use the `jws.WithInsecureNoSignature()` option. This option
// can be combined with one or more signature keys, as well as the
// `jws.WithJSON()` option to generate multiple signatures (though
// the usefulness of such constructs is highly debatable)
//
// Note that this library does not allow you to successfully call `jws.Verify()` on
// signatures with the "none" algorithm. To parse these, use `jws.Parse()` instead.
//
// If you want to use a detached payload, use `jws.WithDetachedPayload()` as
// one of the options. When you use this option, you must always set the
// first parameter (`payload`) to `nil`, or the function will return an error
//
// You may also want to look at how to pass protected headers to the
// signing process, as you will likely be required to set the `b64` field
// when using detached payload.
//
// RFC 7797 note: producing an in-band compact JWS with `b64=false`
// (i.e. setting the `b64` protected header to `false` without also
// passing [WithDetachedPayload]) is "NOT RECOMMENDED" per §5.2; strict
// peers commonly reject such messages. The canonical pairing for
// `b64=false` is [WithDetachedPayload] (or [WithDetachedPayloadReader]
// for streaming), which keeps the unencoded payload out of the wire
// format. Sign auto-declares `"b64"` in `crit` whenever `b64=false`
// is set, so the produced JWS is at least RFC 7797 §3 conformant on
// the producer side.
//
// Look for options that return `jws.SignOption` or `jws.SignVerifyOption`
// for a complete list of options that can be passed to this function.
//
// You can use `errors.Is` with `jws.SignError()` to check if an error is from this function.
func Sign(payload []byte, options ...SignOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Design note: while we could have easily set format = fmtJSON when
// lsigner > 1, I believe the decision to change serialization formats
// must be explicitly stated by the caller. Otherwise, I'm pretty sure
// there would be people filing issues saying "I get JSON when I expected
// compact serialization".
//
// Therefore, instead of making implicit format conversions, we force the
// user to spell it out as `jws.Sign(..., jws.WithJSON(), jws.WithKey(...), jws.WithKey(...))`

// For compact single-signature (the overwhelmingly common case),
// bypass Message construction and Compact() entirely.
// Build() returns the signing input buffer (base64(hdr).base64(payload))
// so we can append the signature directly without re-encoding.

// Detached: output is base64(hdr)..base64(sig) (empty payload segment).
// The combined buffer is base64(hdr).base64(payload), so slice
// up to and including the period to get base64(hdr). and append
// the signature after that.

// Non-detached: append .base64(sig) to the existing signing buffer.

// JSON serialization path - needs full Message construction

// Verify checks if the given JWS message is verifiable using `alg` and `key`.
// `key` may be a "raw" key (e.g. rsa.PublicKey) or a jwk.Key
//
// If the verification is successful, `err` is nil, and the content of the
// payload that was signed is returned. If you need more fine-grained
// control of the verification process, manually generate a
// `Verifier` in `verify` subpackage, and call `Verify` method on it.
// If you need to access signatures and JOSE headers in a JWS message,
// use `Parse` function to get `Message` object.
//
// Because the use of "none" (jwa.NoSignature) algorithm is strongly discouraged,
// this function DOES NOT consider it a success when `{"alg":"none"}` is
// encountered in the message (it would also be counterintuitive when the code says
// it _verified_ something when in fact it did no such thing). If you want to
// accept messages with "none" signature algorithm, use `jws.Parse` to get the
// raw JWS message.
//
// The error returned by this function is of type can be checked against
// `jws.VerifyError()` and `jws.VerificationError()`. The latter is returned
// when the verification process itself fails (e.g. invalid signature, wrong key),
// while the former is returned when any other part of the `jws.Verify()`
// function fails.
//
// When `jws.WithDetachedPayloadReader()` is used, the payload is streamed
// from the caller's `io.Reader` and is not extracted from the JWS envelope.
// In that case, the returned `[]byte` is a non-nil zero-length slice on
// success; the verified bytes are whatever the caller read from the Reader.
// Do not treat the returned slice as "the payload is empty" — callers that
// need the payload bytes must retain their own copy.
//
// Context cancellation is governed by [WithContext]. The slow-path verify
// loop checks ctx.Err() between each signature, each key provider, and
// each (alg, key) attempt; jkuProvider passes ctx to its underlying
// jwk.Fetcher; the streaming path checks ctx between payload Reads.
// staticKeyProvider and keySetProvider do not consult ctx inside
// FetchKeys themselves (their backing data is already in memory) — see
// the [WithContext] godoc for the full per-layer breakdown.
func Verify(buf []byte, options ...VerifyOption) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// getB64Value reads the typed "b64" header field and returns its value,
// or RFC 7797's default of true when the field is unset. The field is
// declared in jws/objects.yml as a typed bool, so Set rejects non-bool
// values at the API boundary; this helper exists so callers do not have
// to write the same nil-default check at every read site.
func getB64Value(hdr Headers) bool { _ = "STUB: not implemented"; return false }

// RFC 7797 default

// detectParseFormat inspects the first non-whitespace rune in src to
// classify the input as compact or JSON serialization. Returns 0 on
// empty or whitespace-only input so callers can distinguish "no usable
// bytes" from a successful classification.
func detectParseFormat(src []byte) int { _ = "STUB: not implemented"; return 0 }

// Parse parses contents from the given source and creates a jws.Message
// struct. By default the input can be in either compact or full JSON serialization.
//
// You may pass `jws.WithJSON()` and/or `jws.WithCompact()` to specify
// explicitly which format to use. If neither or both is specified, the function
// will attempt to autodetect the format. If one or the other is specified,
// only the specified format will be attempted.
//
// Bounding the input size is the caller's responsibility; this function
// trusts the caller-provided src. See docs/13-input-size.md.
//
// On error, returns a jws.ParseError.
func Parse(src []byte, options ...ParseOption) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if format is 0 or both JSON/Compact, auto detect

// ParseString parses contents from the given source and creates a jws.Message
// struct. The input can be in either compact or full JSON serialization.
//
// On error, returns a jws.ParseError.
func ParseString(src string, options ...ParseOption) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ParseReader parses contents from the given source and creates a jws.Message
// struct. The input can be in either compact or full JSON serialization.
//
// Bounding the input size is the caller's responsibility: wrap src with
// [io.LimitReader] or [net/http.MaxBytesReader] before passing it in. See
// docs/13-input-size.md for the rationale.
//
// On error, returns a jws.ParseError.
func ParseReader(src io.Reader, options ...ParseOption) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseJSON(data []byte, maxSigs int) (result *Message, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseCompact(data []byte) (m *Message, err error) { _ = "STUB: not implemented"; return nil, nil }

func parse(protected, payload, signature []byte) (*Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// CustomDecoder is a generic interface for custom field decoders.
type CustomDecoder[T any] = json.CustomDecoder[T]

// CustomDecodeFunc is a function-based implementation of CustomDecoder[T].
type CustomDecodeFunc[T any] = json.CustomDecodeFunc[T]

// RegisterCustomField registers a private field to be decoded as type T
// using json.Unmarshal. This option has a global effect.
//
//	jws.RegisterCustomField[time.Time](`x-birthday`)
//
// For more fine-tuned control over the decoding process,
// use RegisterCustomDecoder instead.
//
// The error return is reserved for future validation. The current
// implementation always returns nil, but callers — especially extension
// modules calling this from init() — must check the return value and panic
// on failure to stay forward-compatible.
func RegisterCustomField[T any](name string) error { _ = "STUB: not implemented"; return nil }

// RegisterCustomDecoder registers a private field with a custom decoder
// function. This option has a global effect.
//
//	jws.RegisterCustomDecoder(`x-birthday`, jws.CustomDecodeFunc[time.Time](func(data []byte) (time.Time, error) {
//	  var s string
//	  if err := json.Unmarshal(data, &s); err != nil {
//	    return time.Time{}, err
//	  }
//	  return time.Parse(time.RFC1123, s)
//	}))
//
// The error return is reserved for future validation. The current
// implementation always returns nil, but callers — especially extension
// modules calling this from init() — must check the return value and panic
// on failure to stay forward-compatible.
func RegisterCustomDecoder[T any](name string, dec CustomDecodeFunc[T]) error {
	_ = "STUB: not implemented"
	return nil
}

// UnregisterCustomField removes the registration for a custom field.
//
// The error return is reserved for future validation (for example,
// refusing to unregister a built-in field) and is always nil today.
// Callers — especially extension modules scripting Register/Unregister
// cycles from init() — should check the returned value and propagate
// on failure to stay forward-compatible, matching the convention on
// [RegisterCustomField] / [RegisterCustomDecoder].
func UnregisterCustomField(name string) error { _ = "STUB: not implemented"; return nil }

// Helpers for signature verification
var muAlgorithmMaps sync.RWMutex
var keyTypeToAlgorithms = make(map[jwa.KeyType][]jwa.SignatureAlgorithm)
var algorithmToKeyTypes = make(map[jwa.SignatureAlgorithm][]jwa.KeyType)
var curveToAlgorithms = make(map[jwa.EllipticCurveAlgorithm][]jwa.SignatureAlgorithm)

func init() {
	mustRegisterAlgorithmForKeyType(jwa.OKP(), jwa.EdDSA())
	mustRegisterAlgorithmForCurve(jwa.Ed25519(), jwa.EdDSAEd25519())
	for _, alg := range []jwa.SignatureAlgorithm{jwa.HS256(), jwa.HS384(), jwa.HS512()} {
		mustRegisterAlgorithmForKeyType(jwa.OctetSeq(), alg)
	}
	for _, alg := range []jwa.SignatureAlgorithm{jwa.RS256(), jwa.RS384(), jwa.RS512(), jwa.PS256(), jwa.PS384(), jwa.PS512()} {
		mustRegisterAlgorithmForKeyType(jwa.RSA(), alg)
	}
	for _, alg := range []jwa.SignatureAlgorithm{jwa.ES256(), jwa.ES384(), jwa.ES512()} {
		mustRegisterAlgorithmForKeyType(jwa.EC(), alg)
	}
}

func mustRegisterAlgorithmForKeyType(kty jwa.KeyType, alg jwa.SignatureAlgorithm) {
	_ = "STUB: not implemented"
	return
}

func mustRegisterAlgorithmForCurve(crv jwa.EllipticCurveAlgorithm, alg jwa.SignatureAlgorithm) {
	_ = "STUB: not implemented"
	return
}

// RegisterAlgorithmForKeyType registers an additional algorithm as valid for
// the given key type. This is used internally by init() and can also be called
// from external modules that provide support for additional algorithms (e.g. Ed448).
//
// The error return is reserved for future validation. The current
// implementation always returns nil, but callers — especially extension
// modules calling this from init() — must check the return value and panic
// on failure to stay forward-compatible.
func RegisterAlgorithmForKeyType(kty jwa.KeyType, alg jwa.SignatureAlgorithm) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterAlgorithmForCurve registers an algorithm as valid for the given
// elliptic curve. When [AlgorithmsForKey] can determine the curve of a key,
// it returns the union of key-type-level algorithms and curve-specific
// algorithms instead of all algorithms for the key type.
//
// This function is append-only and deduplicates entries, so builtin
// registrations cannot be overwritten by external modules.
//
// The error return is reserved for future validation. The current
// implementation always returns nil, but callers — especially extension
// modules calling this from init() — must check the return value and panic
// on failure to stay forward-compatible.
func RegisterAlgorithmForCurve(crv jwa.EllipticCurveAlgorithm, alg jwa.SignatureAlgorithm) error {
	_ = "STUB: not implemented"
	return nil
}

// AlgorithmsForKey returns the possible signature algorithms that can
// be used for a given key. It only takes in consideration keys/algorithms
// for verification purposes, as this is the only usage where one may need
// dynamically figure out which method to use.
//
// When the key's curve can be determined (via [jwk.Key] Crv() method or
// inferred from the raw Go type), curve-specific algorithms registered via
// [RegisterAlgorithmForCurve] are combined with key-type-level algorithms
// to produce a more precise result.
//
// Accepted key shapes (resolved in order):
//
//  1. [jwk.Key] — kty is read directly; if the implementation also exposes
//     Crv(), the curve refines the result.
//  2. Stdlib crypto types: [rsa.PublicKey] / [rsa.PrivateKey] (and pointer
//     forms), [ecdsa.PublicKey] / [ecdsa.PrivateKey] (and pointer forms),
//     [ed25519.PublicKey], [ed25519.PrivateKey], and [byte] slices for
//     symmetric keys.
//  3. [crypto/ecdh.PublicKey] / [crypto/ecdh.PrivateKey] (and pointer
//     forms) — explicitly rejected; ECDH keys are key-agreement only.
//     Returns an error wrapping [ErrUnclassifiableKey].
//  4. [crypto.Signer] (e.g. KMS-backed adapters) — resolved once via
//     .Public(); the public key is then re-classified through tiers 1–2
//     or the [jwk.Import] fallback below. To prevent infinite recursion,
//     a Signer whose .Public() is itself a Signer is left for the
//     downstream dispatcher to handle.
//  5. [jwk.Import] fallback — anything else is offered to the import
//     registry, allowing extension modules to register their own raw key
//     types.
//
// All "we cannot classify this key" failures wrap [ErrUnclassifiableKey],
// so callers can branch with errors.Is rather than pattern-matching error
// strings. The wrapping error keeps the concrete %T or %q diagnostic in
// its message for human readers.
func AlgorithmsForKey(key any) ([]jwa.SignatureAlgorithm, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ecdh keys are for key agreement (X25519/X448), not signing.
// Reject at the API boundary instead of returning a misleading
// algorithm list that would fail deeper in the signing stack.

// For crypto.Signer from external packages (e.g. KMS-backed signers),
// extract the underlying public key type via .Public().
// Standard library types (*rsa.PrivateKey, etc.) are already handled
// by the concrete cases above.

// Guard: only recurse if the public key is not itself a crypto.Signer,
// to prevent infinite recursion from pathological implementations.

// Save the inner classification error so a
// downstream Import-fallback failure can surface
// both diagnostics. A successful Import discards
// signerPubErr — only the eventual failure path
// joins them.

// If we know the curve and there are curve-specific registrations,
// return only key-type-level algorithms (those not registered under
// any curve) plus curve-specific algorithms for this curve.

// filterAlgorithmsForCurve returns the subset of ktyAlgs that are not
// registered under any curve (i.e., generic for the key type) plus the
// curve-specific algorithms from crvAlgs.
func filterAlgorithmsForCurve(ktyAlgs, crvAlgs []jwa.SignatureAlgorithm) []jwa.SignatureAlgorithm {
	_ = "STUB: not implemented"
	return nil
}

// Add key-type-level algorithms that are not claimed by any curve

// Add curve-specific algorithms

func isRegisteredUnderAnyCurve(alg jwa.SignatureAlgorithm) bool {
	_ = "STUB: not implemented"
	return false
}

// validateAlgorithmForKey checks that alg is compatible with key.
// Three classification failures are intentionally allowed through:
//
// (a) a nil key, used by keyless algorithms (see GH910);
// (b) any key handed to an algorithm with a user-registered custom
// [Signer] or [Verifier] — custom implementations may accept arbitrary
// key types that AlgorithmsForKey cannot classify;
// (c) an opaque crypto.Signer whose .Public() is itself a crypto.Signer,
// the one case AlgorithmsForKey refuses to recurse into.
//
// Every other classification failure is surfaced so callers get a crisp
// option-boundary rejection instead of a deep-stack error.
//
// Carve-out (b) is OR-symmetric across the two registries: hasCustomSigVerifier
// checks both signerDB and verifierDB, so registering EITHER a custom Signer
// OR a custom Verifier loosens the gate for that alg on BOTH the sign and
// verify paths. Importing a verifier-only extension therefore also affects
// jws.Sign-path validation. This is intentional: downstream dispatchers in
// jws/jwsbb re-gate the key shape (via keyconv.KeyAs[T] or
// signer.Public().(*concrete)) before any cryptographic call, so a loose
// validateAlgorithmForKey verdict can only produce a deeper-stack
// type-mismatch error, never a forged signature.
func validateAlgorithmForKey(alg jwa.SignatureAlgorithm, key any) error {
	_ = "STUB: not implemented"
	return nil
}

// hasCustomSigVerifier reports whether a non-default Signer or
// Verifier has been registered for alg. When this is true, key-type
// validation must be skipped: the custom implementation decides what
// key types it accepts.
func hasCustomSigVerifier(alg jwa.SignatureAlgorithm) bool { _ = "STUB: not implemented"; return false }

// Settings allows you to set global settings for JWS operations.
//
// Returns a non-nil error and applies no changes if any option fails
// validation (for example, a non-positive [WithMaxSignatures]).
func Settings(options ...GlobalOption) error { _ = "STUB: not implemented"; return nil }

// VerifyCompactFast is a fast path verification function for JWS messages
// in compact serialization format.
//
// This function is considered experimental, and may change or be removed
// in the future.
//
// VerifyCompactFast performs signature verification on a JWS compact
// serialization without fully parsing the message into a jws.Message object.
// This makes it more efficient for cases where you only need to verify
// the signature and extract the payload, without needing access to headers
// or other JWS metadata.
//
// Returns the original payload that was signed if verification succeeds.
//
// Unlike jws.Verify(), this function requires you to specify the
// algorithm explicitly rather than extracting it from the JWS headers.
// This can be useful for performance-critical applications where the
// algorithm is known in advance.
//
// This function uses strict base64url encoding without padding (RFC 4648 §5)
// for decoding the signature and payload. It does not auto-detect other
// base64 variants. If your JWS uses non-standard encoding (e.g. padded
// base64url), use jws.Verify() instead, which auto-detects the encoding.
//
// Since this function avoids doing many checks that jws.Verify would perform,
// you must ensure to perform the necessary checks including ensuring that algorithm is safe to use for your payload yourself.
//
// VerifyCompactFast cross-checks the protected header's "alg" against
// the caller-supplied alg: if the header omits "alg" (required by
// RFC 7515 §4.1.1) or advertises a different value, it returns a
// verification error. This prevents silently verifying a message
// under a different discipline than the one its header advertises.
//
// VerifyCompactFast refuses messages whose protected header carries a
// "crit" list. RFC 7515 §4.1.11 requires every critical extension to be
// understood by the recipient, and the fast path has no WithCritExtension
// allowlist to consult. On crit-present input it returns a sentinel error
// that callers can detect with errors.Is(err, jws.ErrCritPresent()) and
// retry through jws.Verify, which enforces the full validateCritical rule
// set. Applications that may legitimately receive "crit" headers should
// call jws.Verify directly.
//
// VerifyCompactFast assumes the JWS uses the default "b64":true
// (base64url-encoded) payload encoding. Any protected header carrying
// a "b64" entry is refused with jws.ErrB64Present(), regardless of
// whether "crit" also lists it: the fast path's signing-input
// reconstruction and post-verify base64 decode both depend on the
// default encoding, and a non-conformant b64=false producer (one that
// omits "b64" from "crit") would otherwise verify cryptographically
// while returning bytes that differ from the producer's intent.
// Detached-payload callers must use jws.Verify with jws.WithDetachedPayload
// regardless, since VerifyCompactFast has no way to accept a detached
// payload.
func VerifyCompactFast(key any, compact []byte, alg jwa.SignatureAlgorithm) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Split the serialized JWS into its components

// Refuse crit-bearing messages: the fast path has no WithCritExtension
// allowlist, so accepting them would silently violate RFC 7515 §4.1.11.
// Callers that wrap VerifyCompactFast can detect this via
// errors.Is(err, jws.ErrCritPresent()) and fall through to jws.Verify.
// The sentinel is wrapped in verifyError so the same error also matches
// errors.Is(err, jws.VerifyError()) — fast-path refusals are a verify
// error, just one with a more specific classification available.

// Refuse "b64"-bearing messages, regardless of whether "crit" also
// lists it. The signing-input reconstruction and the post-verify
// base64 decode both assume the default b64=true encoding; a
// b64=false JWS that the fast path "verified" would either fail the
// post-verify base64 decode with a misleading error, or — worse —
// return base64-decoded garbage as the payload while the producer's
// raw bytes silently disagree. jws.Verify has the WithDetachedPayload
// / WithCritExtension machinery to handle b64=false correctly. As with
// the crit refusal above, the sentinel is wrapped in verifyError so the
// same error matches both jws.ErrB64Present() and jws.VerifyError().

// Cross-check the protected header "alg" against the caller-supplied
// alg. RFC 7515 §4.1.1 makes "alg" mandatory in the protected header
// for compact serialization, and a mismatch between what the message
// advertises and the discipline under which we verify is the sort of
// silent divergence that downstream code (e.g. JWT consumers) should
// not be asked to re-discover on its own.

// Decode signature into pooled buffer (strict base64url, no padding per RFC 7515)

// Instead of appending, copy the data from hdr/payload

// Verify the signature

// Decode payload (strict base64url, no padding per RFC 7515)
