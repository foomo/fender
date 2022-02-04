package rule

import (
	"regexp"
)

const NameRegex = "regex"

var ErrRegex = &Error{Rule: NameRegex}

// NewRegexError constructor
func NewRegexError() *Error {
	return NewError(NameRegex)
}

// Regex validation using go standard package
func Regex(regexp *regexp.Regexp) StringRule {
	return func(v string) Rule {
		return func() (*Error, error) {
			if !regexp.MatchString(v) {
				return NewRegexError(), nil
			}
			return nil, nil
		}
	}
}
