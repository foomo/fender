package rule

import (
	"context"
	"fmt"
)

const NameBool Name = "bool"

// NewBoolRuleError constructor
func NewBoolRuleError(v bool) *Error {
	return NewError(NameBool, fmt.Sprintf("%t", v))
}

func Bool(expected bool) BoolRule {
	return func(ctx context.Context, v bool) error {
		if v != expected {
			return NewBoolRuleError(expected)
		}
		return nil
	}
}

func True(ctx context.Context, v bool) error {
	if !v {
		return NewBoolRuleError(true)
	}
	return nil
}

func False(ctx context.Context, v bool) error {
	if v {
		return NewBoolRuleError(false)
	}
	return nil
}
