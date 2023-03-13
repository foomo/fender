package rule

import (
	"context"
)

func FloatMin[T Float](expected T) FloatRule[T] {
	return func(ctx context.Context, v T) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func FloatMax[T Float](expected T) FloatRule[T] {
	return func(ctx context.Context, v T) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func FloatRequired[T Float](ctx context.Context, v T) error {
	if v == T(0) {
		return NewError(NameRequired)
	}
	return nil
}

func FloatNotRequired[T Float](ctx context.Context, v T) error {
	if v == T(0) {
		return ErrBreak
	}
	return nil
}

func FloatSize[T Float](expected T) FloatRule[T] {
	return func(ctx context.Context, v T) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func FloatOmitEmpty[T Float](ctx context.Context, v T) error {
	if v == T(0) {
		return ErrBreak
	}
	return nil
}
