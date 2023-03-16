package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

type Rules[T any] []rule.Rule[T]

func NewRules[T any](rules ...rule.Rule[T]) Rules[T] {
	return rules
}

func (r Rules[T]) Fend(path string, value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, "", value, rule.Rules[T](rules).Append(r...)...)
	}
}

func (r Rules[T]) FendVar(value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, "", value, rule.Rules[T](rules).Append(r...)...)
	}
}
