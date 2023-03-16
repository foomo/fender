package rule

import (
	"context"
)

func Dynamic[T any](fn func(ctx context.Context) error) Rule[T] {
	return func(ctx context.Context, v T) error {
		return fn(ctx)
	}
}
