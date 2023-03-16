package rule

import (
	"context"
)

const NameValid Name = "valid"

type (
	Validator interface {
		Valid() bool
	}
	ValidatorRule = Rule[Validator]
)

func Valid[T Validator](ctx context.Context, v T) error {
	if !v.Valid() {
		return NewError(NameValid)
	}
	return nil
}
