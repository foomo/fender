package rule

import (
	"context"
)

type Validator interface {
	Valid() bool
}

func Valid[T Validator](ctx context.Context, v T) error {
	if !v.Valid() {
		return NewError(NameValid)
	}
	return nil
}
