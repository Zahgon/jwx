package openid

const (
	AddressFormattedKey     = "formatted"
	AddressStreetAddressKey = "street_address"
	AddressLocalityKey      = "locality"
	AddressRegionKey        = "region"
	AddressPostalCodeKey    = "postal_code"
	AddressCountryKey       = "country"
)

// AddressClaim is the address claim as described in https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
type AddressClaim struct {
	formatted     *string // https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
	streetAddress *string // https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
	locality      *string // https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
	region        *string // https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
	postalCode    *string // https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
	country       *string // https://openid.net/specs/openid-connect-core-1_0.html#AddressClaim
}

type addressClaimMarshalProxy struct {
	Xformatted     *string `json:"formatted,omitempty"`
	XstreetAddress *string `json:"street_address,omitempty"`
	Xlocality      *string `json:"locality,omitempty"`
	Xregion        *string `json:"region,omitempty"`
	XpostalCode    *string `json:"postal_code,omitempty"`
	Xcountry       *string `json:"country,omitempty"`
}

func NewAddress() *AddressClaim { _ = "STUB: not implemented"; return nil }

// Formatted is a convenience function to retrieve the corresponding value store in the token
// if there is a problem retrieving the value, the zero value is returned. If you need to differentiate between existing/non-existing values, use `Get` instead
func (t AddressClaim) Formatted() string { _ = "STUB: not implemented"; return "" }

// StreetAddress is a convenience function to retrieve the corresponding value store in the token
// if there is a problem retrieving the value, the zero value is returned. If you need to differentiate between existing/non-existing values, use `Get` instead
func (t AddressClaim) StreetAddress() string { _ = "STUB: not implemented"; return "" }

// Locality is a convenience function to retrieve the corresponding value store in the token
// if there is a problem retrieving the value, the zero value is returned. If you need to differentiate between existing/non-existing values, use `Get` instead
func (t AddressClaim) Locality() string { _ = "STUB: not implemented"; return "" }

// Region is a convenience function to retrieve the corresponding value store in the token
// if there is a problem retrieving the value, the zero value is returned. If you need to differentiate between existing/non-existing values, use `Get` instead
func (t AddressClaim) Region() string { _ = "STUB: not implemented"; return "" }

// PostalCode is a convenience function to retrieve the corresponding value store in the token
// if there is a problem retrieving the value, the zero value is returned. If you need to differentiate between existing/non-existing values, use `Get` instead
func (t AddressClaim) PostalCode() string { _ = "STUB: not implemented"; return "" }

// Country is a convenience function to retrieve the corresponding value store in the token
// if there is a problem retrieving the value, the zero value is returned. If you need to differentiate between existing/non-existing values, use `Get` instead
func (t AddressClaim) Country() string { _ = "STUB: not implemented"; return "" }

func (t *AddressClaim) Get(s string) (any, bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func (t *AddressClaim) Set(key string, value any) error { _ = "STUB: not implemented"; return nil }

func (t *AddressClaim) Accept(v any) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON serializes the token in JSON format.
func (t AddressClaim) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// String values go through json.Marshal, not strconv.Quote.
// strconv.Quote produces Go source-form escaping (\xNN, invalid
// for JSON; control bytes 0x00–0x1F and 0x7F must be \u00NN).
// Routing through json.Marshal also handles invalid UTF-8 the
// way the rest of the package's encoders do.

// Field order preserved from the historical implementation so the
// MarshalJSON output is byte-stable for callers that compare it.

// UnmarshalJSON deserializes data from a JSON data buffer into a AddressClaim
func (t *AddressClaim) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
