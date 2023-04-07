package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func Var[T any](value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, "", value, rules...)
	}
}
