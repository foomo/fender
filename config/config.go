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

var (
	// EmailRegexWeak as https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
	EmailRegexWeak = ".+@.+\\..+"
	// EmailRegexRFC5322 as https://stackoverflow.com/questions/201323/how-to-validate-an-email-address-using-a-regular-expression/201378#201378
	EmailRegexRFC5322 = "^(?i)(?:[a-z0-9!#$%&'*+/=?^_`{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_`{|}~-]+)*|\"(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21\\x23-\\x5b\\x5d-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])*\")@(?:(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?|\\[(?:(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9]))\\.){3}(?:(2(5[0-5]|[0-4][0-9])|1[0-9][0-9]|[1-9]?[0-9])|[a-z0-9-]*[a-z0-9]:(?:[\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x21-\\x5a\\x53-\\x7f]|\\\\[\\x01-\\x09\\x0b\\x0c\\x0e-\\x7f])+)\\])*$"
)

var Validator = validator.New()
