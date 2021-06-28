package fender

import (
	"regexp"
)

const RuleRegex Rule = "regex"

// NewFendRegexError constructor
func NewFendRegexError() *RuleError {
	return NewRuleError(RuleRegex)
}

// Regex validation using go standard package
func Regex(v string, regexp *regexp.Regexp) Fend {
	return func() *RuleError {
		if !regexp.MatchString(v) {
			return NewFendRegexError()
		}
		return nil
	}
}
