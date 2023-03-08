package rule

import (
	"context"
)

const NameValid Name = "valid"

type Validator interface {
	Valid() bool
}

// NewValidError constructor
func NewValidError() *Error {
	return NewError(NameValid)
}

func Valid() ValidatorRule {
	return func(ctx context.Context, v Validator) error {
		if !v.Valid() {
			return NewValidError()
		}
		return nil
	}
}
