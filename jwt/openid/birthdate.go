package openid

import (
	"io"
	"math"
	"regexp"
)

// https://openid.net/specs/openid-connect-core-1_0.html
//
// End-User's birthday, represented as an ISO 8601:2004 [ISO8601‑2004] YYYY-MM-DD format.
// The year MAY be 0000, indicating that it is omitted. To represent only the year, YYYY
// format is allowed. Note that depending on the underlying platform's date related function,
// providing just year can result in varying month and day, so the implementers need to
// take this factor into account to correctly process the dates.

type BirthdateClaim struct {
	year  *int
	month *int
	day   *int
}

func (b BirthdateClaim) Year() int { _ = "STUB: not implemented"; return 0 }

func (b BirthdateClaim) Month() int { _ = "STUB: not implemented"; return 0 }

func (b BirthdateClaim) Day() int { _ = "STUB: not implemented"; return 0 }

func (b *BirthdateClaim) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

var intSize int

func init() {
	intSize = 64
	if math.MaxInt == math.MaxInt32 {
		intSize = 32
	}
}

func parseBirthdayInt(s string) int { _ = "STUB: not implemented"; return 0 }

var birthdateRx = regexp.MustCompile(`^(\d{4})-(\d{2})-(\d{2})$`)

// Accept accepts a value read from JSON, and converts it to a BirthdateClaim.
// This method DOES NOT verify the correctness of a date.
// Consumers should check for validity of dates such as Apr 31 et al
func (b *BirthdateClaim) Accept(v any) error { _ = "STUB: not implemented"; return nil }

// yeah, regexp is slow. PR's welcome

// Okay, this really isn't kosher, but we're doing this for
// the coverage game... Because birthdateRx already checked that
// the string contains 3 strings with consecutive decimal values
// we can assume that strconv.ParseInt always succeeds.
// strconv.ParseInt (and strconv.ParseUint that it uses internally)
// only returns range errors, so we should be safe.

// year == 0 (i.e. "0000") means omitted per OIDC spec; leave tmp.year as nil

func (b BirthdateClaim) encode(dst io.Writer) { _ = "STUB: not implemented"; return }

func (b BirthdateClaim) String() string { _ = "STUB: not implemented"; return "" }

func (b BirthdateClaim) MarshalText() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }
