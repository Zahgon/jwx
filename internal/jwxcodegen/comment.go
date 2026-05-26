package jwxcodegen

import (
	"regexp"

	"github.com/lestrrat-go/codegen"
)

var reLooksLikeCodeBlock = regexp.MustCompile(`^\s+`)

// WriteComment formats a multi-line YAML comment string as Go doc
// comment lines on the given output. Returns true if any comment
// was written, false if the comment was empty.
func WriteComment(o *codegen.Output, comment string) bool { _ = "STUB: not implemented"; return false }
