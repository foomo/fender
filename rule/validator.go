package rule

import (
	"context"
)

type Validator interface {
	Valid() bool
}

// NewValidError constructor
func NewValidError() *Error {
	return NewError(NameValid)
}

func Valid[T Validator](ctx context.Context, v T) error {
	if !v.Valid() {
		return NewError(NameValid)
	}
	return nil
}
