package rule

import (
	"context"
)

func Union[T any](rules ...Rule[T]) Rule[T] {
	return func(ctx context.Context, v T) error {
		var e error
		for _, r := range rules {
			if err := r(ctx, v); err != nil {
				e = err // TODO only return las t
			} else {
				return nil
			}
		}
		return e
	}
}
