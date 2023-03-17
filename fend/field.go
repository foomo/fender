package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func Field[T any](path string, value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, path, value, rules...)
	}
}
