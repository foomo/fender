package rule

import (
	"fmt"
)

const NameBool = "bool"

var ErrBool = &Error{Rule: NameBool}

// NewBoolRuleError constructor
func NewBoolRuleError(v bool) *Error {
	return NewError(NameBool, fmt.Sprintf("%t", v))
}

func Bool(expected bool) BoolRule {
	return func(v bool) Rule {
		return func() (*Error, error) {
			if v != expected {
				return NewBoolRuleError(expected), nil
			}
			return nil, nil
		}
	}
}

func True(v bool) Rule {
	return func() (*Error, error) {
		if !v {
			return NewBoolRuleError(true), nil
		}
		return nil, nil
	}
}

func False(v bool) Rule {
	return func() (*Error, error) {
		if v {
			return NewBoolRuleError(false), nil
		}
		return nil, nil
	}
}
