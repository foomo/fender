package fender

import (
	"regexp"
)

const RuleRegex Rule = "regex"

// NewFendRegexError constructor
func NewFendRegexError() *RuleError {
	return NewFendError(RuleRegex, RuleRegex.String())
}

// Regex validation using go standard package
func Regex(v string, regexp *regexp.Regexp) error {
	if !regexp.MatchString(v) {
		return NewFendRegexError()
	}
	return nil
}
