package jws

import (
	"context"
	"hash"
	"io"

	"github.com/lestrrat-go/dsig"

	"github.com/lestrrat-go/jwx/v4/internal/base64"
	"github.com/lestrrat-go/jwx/v4/internal/json"
	"github.com/lestrrat-go/jwx/v4/jwa"
)

// ctxReader wraps an io.Reader so that ctx.Err() is checked between
// reads. Any cancellation or deadline expiry propagates out of Read
// as the ctx.Err() value, short-circuiting the outer io.Copy loop.
// A zero-length Read is still delivered to the underlying reader so
// that an empty-payload verify completes normally.
//
//nolint:containedctx // narrow, request-scoped wrapper that the caller owns for the lifetime of a single Verify call; storing the ctx is the point.
type ctxReader struct {
	ctx context.Context
	r   io.Reader
}

func (c *ctxReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// wrapReaderWithContext returns a Reader that surfaces ctx.Err()
// between underlying Reads. If ctx is nil or context.Background(),
// r is returned unchanged to avoid the indirection.
func wrapReaderWithContext(ctx context.Context, r io.Reader) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

// This file implements the streaming detached-payload variant of jws.Sign()
// and jws.Verify(), reached via the jws.WithDetachedPayloadReader() option.
// It deliberately bypasses the jws.Signer / jws.Verifier dispatch path and
// talks to dsig directly, because it needs incremental hashing through a
// hash.Hash — the Signer interface takes a fully materialized []byte payload.
//
// Consequences: algorithms registered via jws.RegisterSigner() /
// jws.RegisterVerifier() are unreachable here, as are algorithm families
// that cannot be driven from a digest (EdDSA, custom).

// streamingSigner carries per-signature state through signStreaming:
// the header prefix already written into hasher, the hasher itself
// (fed with the prefix, ready to consume payload bytes), the resolved
// dsig alg info, and the raw key / unprotected header needed for final
// signing and JSON assembly.
type streamingSigner struct {
	dsigInfo   streamingAlgorithmInfo
	rawKey     any
	hasher     hash.Hash
	hdrEncoded string // base64(protected header JSON)
	public     Headers
}

// signStreaming is invoked from Sign() when sc.payloadReader is set. It
// assembles the signing input for each configured signer by feeding
// base64(header) "." base64(payload) (or raw payload when b64=false)
// into a hash.Hash, then calls dsig.SignDigest once per signer. When
// more than one signer is registered the payload is streamed once and
// fanned out to each hasher via [io.MultiWriter]; the general JSON
// serialization is used for the output.
func (sc *signContext) signStreaming() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// For compact serialization RFC 7515 requires the unprotected
// header to be merged into the protected header because there
// is no separate slot for it on the wire. JSON serializations
// keep them separate.

// A multi-signature JWS has a single payload segment on the
// wire, so every signer must agree on the RFC 7797 "b64" flag
// or the produced JWS is internally inconsistent.

// Upstream validation in jws.Sign guarantees compact implies
// exactly one signer, but guard against future drift.

// verifyStreaming is invoked from VerifyMessage() when vc.payloadReader is
// set. It parses the JWS envelope through jws.Parse, then re-builds the
// signing input by feeding base64(header) "." base64(payload) into a
// hash.Hash fed from the supplied io.Reader and calls dsig.VerifyDigest.
func (vc *verifyContext) verifyStreaming(buf []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// When the caller supplied a context via WithContext, surface
// cancellation and deadline expiry between Reads on the payload
// reader. Without this the streaming verify path would ignore ctx
// entirely and keep draining an attacker-controlled reader long
// after the caller cancelled.

// Non-nil zero-length slice is the sentinel: the payload was streamed
// from the caller so there are no bytes to hand back, but returning
// nil would be indistinguishable from "ignored return" and invite
// `len(payload) == 0` silent-logic bugs in callers.

// streamingAlgorithmInfo carries the resolved dsig metadata plus the dsig
// algorithm name, since dsig.AlgorithmInfo itself does not include it.
type streamingAlgorithmInfo struct {
	dsig.AlgorithmInfo

	Name string
}

// resolveStreamingAlgorithm maps a JWS algorithm to its dsig metadata and
// enforces the family restrictions for the streaming path. It routes through
// jwsbb.GetDsigAlgorithm so algorithms registered by extension modules work
// just like algorithms built in to jws.
func resolveStreamingAlgorithm(alg jwa.SignatureAlgorithm) (streamingAlgorithmInfo, error) {
	_ = "STUB: not implemented"
	return *new(streamingAlgorithmInfo), nil
}

// For custom dsig algorithms registered directly with dsig the JWS
// name may equal the dsig name.

// newStreamingHasher returns a hash.Hash preloaded with the key material
// the family needs (HMAC is keyed; RSA/ECDSA just hash).
func newStreamingHasher(info streamingAlgorithmInfo, key any) (hash.Hash, error) {
	_ = "STUB: not implemented"
	return *new(hash.Hash), nil
}

// Mirror dispatchHMACSign in jws/jwsbb/sign.go: route through
// keyconv.KeyAs[[]byte] so common footguns (e.g. passing a
// string secret) surface the same actionable message as the
// non-streaming path instead of a bare type mismatch.

// streamPayloadIntoHashers copies the payload once and fans it out to
// every hasher via [io.MultiWriter]. When b64=true the bytes are
// routed through per-hasher [io.WriteCloser]s returned by the
// [base64.StreamEncoder] (each encoder keeps an unflushed 3-byte tail,
// so the wrappers cannot be shared). When b64=false the hashers receive
// the payload bytes directly.
func streamPayloadIntoHashers(hashers []hash.Hash, payload io.Reader, encodePayload bool, enc base64.StreamEncoder) error {
	_ = "STUB: not implemented"
	return nil
}

type streamingDetachedJSONEntry struct {
	Header    json.RawMessage `json:"header,omitempty"`
	Protected string          `json:"protected"`
	Signature string          `json:"signature"`
}

type streamingDetachedJSONGeneral struct {
	Signatures []streamingDetachedJSONEntry `json:"signatures"`
}

// assembleStreamingDetachedJSON emits the flattened form for a single
// signature and the general form for multiple. The "payload" member is
// omitted in both cases per RFC 7515 Appendix F.
func assembleStreamingDetachedJSON(signers []streamingSigner, rawSignatures [][]byte, enc base64.StreamEncoder, pretty bool) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// cloneOrNewHeaders returns a defensive copy of hdr, or a fresh empty
// Headers if hdr is nil. The streaming path mutates the protected headers
// to set "alg" / "kid", so we never mutate a caller-supplied value.
func cloneOrNewHeaders(hdr Headers) (Headers, error) {
	_ = "STUB: not implemented"
	return *new(Headers), nil
}

func convertStreamingSignKey(key any, family dsig.Family) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

func convertStreamingVerifyKey(key any, family dsig.Family) (any, error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}
