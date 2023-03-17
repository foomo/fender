package fend

import (
	"context"

	"github.com/foomo/fender/rule"
)

type Rules[T any] []rule.Rule[T]

func NewRules[T any](rules ...rule.Rule[T]) Rules[T] {
	return rules
}

func (r Rules[T]) Field(path string, value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, path, value, rule.Rules[T](rules).Append(r...)...)
	}
}

func (r Rules[T]) Var(value T, rules ...rule.Rule[T]) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fend(ctx, mode, "", value, rule.Rules[T](rules).Append(r...)...)
	}
}

func (r Rules[T]) DynamicField(path string, rules ...rule.DynamicRule) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendDynamic(ctx, mode, path, rules...)
	}
}

func (r Rules[T]) DynamicVar(rules ...rule.DynamicRule) Fend {
	return func(ctx context.Context, mode Mode) error {
		return fendDynamic(ctx, mode, "", rules...)
	}
}
