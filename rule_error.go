package fender

import "strings"

type RuleError struct {
	Rule Rule
	Meta string
}

// NewRuleError constructor
func NewRuleError(fend Rule, meta ...string) *RuleError {
	return &RuleError{
		Rule: fend,
		Meta: strings.Join(meta, MetaDelimiter),
	}
}

// RuleError interface
func (e *RuleError) Error() string {
	s := string(e.Rule)
	if e.Meta != "" {
		s += MetaDelimiter + e.Meta
	}
	return s
}
