package rule

import (
	"context"

	"golang.org/x/exp/constraints"
)

func IntegerMin[T constraints.Integer](expected T) IntegerRule[T] {
	return func(ctx context.Context, v T) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func IntegerMax[T constraints.Integer](expected T) IntegerRule[T] {
	return func(ctx context.Context, v T) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func IntegerRequired[T constraints.Integer](ctx context.Context, v T) error {
	if v == T(0) {
		return NewError(NameRequired)
	}
	return nil
}

func IntegerSize[T constraints.Integer](expected T) IntegerRule[T] {
	return func(ctx context.Context, v T) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func IntegerOmitEmpty[T constraints.Integer](ctx context.Context, v T) error {
	if v == T(0) {
		return ErrBreak
	}
	return nil
}
