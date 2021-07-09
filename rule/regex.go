package rule

import (
	"errors"
	"regexp"
)

const NameRegex Name = "regex"

var ErrRegex = errors.New(NameRegex.String())

// NewRegexError constructor
func NewRegexError() *Error {
	return NewError(ErrRegex, NameRegex.String())
}

// Regex validation using go standard package
func Regex(regexp *regexp.Regexp) StringRule {
	return func(v string) Rule {
		return func() *Error {
			if !regexp.MatchString(v) {
				return NewRegexError()
			}
			return nil
		}
	}
}
