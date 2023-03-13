package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

func Field[T any](name string, value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, name, value, rules...)
	}
}

func FieldFn(name string, rules ...rule.Fn) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendFn(ctx, mode, name, rules...)
	}
}
