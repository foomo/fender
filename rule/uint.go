package rule

import (
	"context"
)

func UIntMin(expected uint) UIntRule {
	return func(ctx context.Context, v uint) error {
		if v < expected {
			return NewErrorf(NameMin, 'd', expected)
		}
		return nil
	}
}

func UIntMax(expected uint) UIntRule {
	return func(ctx context.Context, v uint) error {
		if v > expected {
			return NewErrorf(NameMax, 'd', expected)
		}
		return nil
	}
}

func UIntRequired(ctx context.Context, v uint) error {
	if v == 0 {
		return NewError(NameRequired)
	}
	return nil
}

func UIntNotRequired(ctx context.Context, v uint) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}

func UIntSize(expected uint) UIntRule {
	return func(ctx context.Context, v uint) error {
		if v != expected {
			return NewErrorf(NameSize, 'd', expected)
		}
		return nil
	}
}

func UIntOmitEmpty(ctx context.Context, v uint) error {
	if v == 0 {
		return ErrBreak
	}
	return nil
}
