package config

import "github.com/go-playground/validator/v10"

const (
	DefaultRuleDelimiter  = ";"
	DefaultFieldDelimiter = ":"
	DefaultMetaDelimiter  = "="
	DefaultErrorDelimiter = ","
)

var (
	FieldDelimiter = DefaultFieldDelimiter
	RuleDelimiter  = DefaultRuleDelimiter
	ErrorDelimiter = DefaultErrorDelimiter
	MetaDelimiter  = DefaultMetaDelimiter
)

// EmailRegexWeak as https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
var EmailRegexWeak = ".+@.+\\..+"

var Validator = validator.New()
