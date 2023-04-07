package rule

import (
	"context"
)

const NameEqual Name = "equal"

func Equal[T comparable](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v != expected {
			return NewError(NameEqual, Meta('s', expected))
		}
		return nil
	}
}

func NotEqual[T comparable](expected T) Rule[T] {
	return func(ctx context.Context, v T) error {
		if v == expected {
			return NewError(NameEqual, Meta('s', expected))
		}
		return nil
	}
}
