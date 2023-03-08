package rule

import (
	"context"
	"regexp"
)

const NameRegex Name = "regex"

// NewRegexError constructor
func NewRegexError() *Error {
	return NewError(NameRegex)
}

// Regex validation using go standard package
func Regex(regexp *regexp.Regexp) StringRule {
	return func(ctx context.Context, v string) error {
		if !regexp.MatchString(v) {
			return NewRegexError()
		}
		return nil
	}
}
