package jwx

import (
	"encoding/json"
)

type FormatKind int

// These constants describe the result from guessing the format
// of the incoming buffer.
const (
	// InvalidFormat is returned when the format of the incoming buffer
	// has been deemed conclusively invalid
	InvalidFormat FormatKind = iota
	// UnknownFormat is returned when GuessFormat was not able to conclusively
	// determine the format of the
	UnknownFormat
	JWE
	JWS
	JWK
	JWKS
	JWT
)

type formatHint struct {
	Payload    json.RawMessage `json:"payload"`    // Only in JWS
	Signatures json.RawMessage `json:"signatures"` // Only in JWS
	Ciphertext json.RawMessage `json:"ciphertext"` // Only in JWE
	KeyType    json.RawMessage `json:"kty"`        // Only in JWK
	Keys       json.RawMessage `json:"keys"`       // Only in JWKS
	Audience   json.RawMessage `json:"aud"`        // Only in JWT
}

// GuessFormat is used to guess the format the given payload is in
// using heuristics. See the type FormatKind for a full list of
// possible types.
//
// This may be useful in determining your next action when you may
// encounter a payload that could either be a JWE, JWS, or a plain JWT.
//
// Because JWTs are almost always JWS signed, you may be thrown off
// if you pass what you think is a JWT payload to this function.
// If the function is in the "Compact" format, it means it's a JWS
// signed message, and its payload is the JWT. Therefore this function
// will return JWS, not JWT.
//
// This function requires an extra parsing of the payload, and therefore
// may be inefficient if you call it every time before parsing.
func GuessFormat(payload []byte) FormatKind {
	_ = "STUB: not implemented"
	// The check against kty, keys, and aud are something this library
	// made up. for the distinctions between JWE and JWS, we used
	// https://datatracker.ietf.org/doc/html/rfc7516#section-9.
	//
	// The above RFC described several ways to distinguish between
	// a JWE and JWS JSON, but we're only using one of them
	return *new(FormatKind)
}

// Compact format. It's probably a JWS or JWE
// I want to const this :/

// Note: this counts the number of occurrences of the
// separator, but the RFC talks about the number of segments.
// number of tokens.Period == segments - 1, so that's why we have 2 and 4 here

// If we got here, we probably have JSON.
