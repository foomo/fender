package config

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

const (
	DefaultRuleDelimiter  = ";"
	DefaultFieldDelimiter = ":"
	DefaultMetaDelimiter  = "="
)

var (
	FieldDelimiter = DefaultFieldDelimiter
	RuleDelimiter  = DefaultRuleDelimiter
	MetaDelimiter  = DefaultMetaDelimiter
)

// EmailRegexWeak as https://davidcel.is/posts/stop-validating-email-addresses-with-regex/
var EmailRegexWeak = ".+@.+\\..+"

var Validator = validator.New()

func init() {
	Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		return strings.ToLower(fld.Name[:1]) + fld.Name[1:]
	})
}
