package rule

import (
	"context"
)

func IntegerMin[T Integer](expected T) IntegerRule[T] {
	return func(ctx context.Context, v T) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func IntegerMax[T Integer](expected T) IntegerRule[T] {
	return func(ctx context.Context, v T) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func IntegerRequired[T Integer](ctx context.Context, v T) error {
	if v == T(0) {
		return NewError(NameRequired)
	}
	return nil
}

func IntegerNotRequired[T Integer](ctx context.Context, v T) error {
	if v == T(0) {
		return ErrBreak
	}
	return nil
}

func IntegerSize[T Integer](expected T) IntegerRule[T] {
	return func(ctx context.Context, v T) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func IntegerOmitEmpty[T Integer](ctx context.Context, v T) error {
	if v == T(0) {
		return ErrBreak
	}
	return nil
}
