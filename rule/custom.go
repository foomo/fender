package rule

import (
	"errors"
)

const NameCustom Name = "custom"

var ErrCustom = errors.New(NameCustom.String())

// NewCustomRuleError constructor
func NewCustomRuleError(rule string, meta ...string) *Error {
	return NewError(ErrCustom, rule, meta...)
}
