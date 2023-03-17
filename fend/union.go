package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func Union[T any](rules ...rule.Rule[T]) rule.Rule[T] {
	return func(ctx context.Context, v T) error {
		var e error
		for _, r := range rules {
			if err := r(ctx, v); err != nil {
				e = err // TODO only return last error?
			} else {
				return nil
			}
		}
		return e
	}
}
