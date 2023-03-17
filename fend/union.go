package fend

import (
	"context"

	"github.com/foomo/fender/rule"
	"go.uber.org/multierr"
)

func Union[T any](rules ...Rules[T]) rule.Rule[T] {
	return func(ctx context.Context, v T) error {
		var e error
		for _, r := range rules {
			var e2 error
			for _, r2 := range r {
				if err := r2(ctx, v); err != nil {
					e2 = multierr.Append(e2, err)
				}
			}
			if e2 != nil {
				e = e2
			} else {
				return nil
			}
		}
		return e
	}
}
